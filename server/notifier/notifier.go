package notifier

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oinume/lekcije/server/config"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/fetcher"
	"github.com/oinume/lekcije/server/logger"
	"github.com/oinume/lekcije/server/model"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var lessonFetcher *fetcher.TeacherLessonFetcher

func init() {
	http.DefaultClient.Timeout = 5 * time.Second
	lessonFetcher = fetcher.NewTeacherLessonFetcher(http.DefaultClient, nil)
}

type Notifier struct {
	db            *gorm.DB
	dryRun        bool
	lessonService *model.LessonService
	teachers      map[uint32]*model.Teacher
}

func NewNotifier(db *gorm.DB, dryRun bool) *Notifier {
	return &Notifier{
		db:       db,
		dryRun:   dryRun,
		teachers: make(map[uint32]*model.Teacher, 1000),
	}
}

func (n *Notifier) SendNotification(user *model.User) error {
	followingTeacherService := model.NewFollowingTeacherService(n.db)
	n.lessonService = model.NewLessonService(n.db)

	teacherIDs, err := followingTeacherService.FindTeacherIDsByUserID(user.ID)
	if err != nil {
		return errors.Wrapperf(err, "Failed to FindTeacherIDsByUserID(): userID=%v", user.ID)
	}

	availableLessonsPerTeacher := make(map[uint32][]*model.Lesson, 1000)
	allFetchedLessons := make([]*model.Lesson, 0, 5000)
	for _, teacherID := range teacherIDs {
		teacher, fetchedLessons, newAvailableLessons, err := n.fetchAndExtractNewAvailableLessons(teacherID)
		if err != nil {
			return err
		}

		allFetchedLessons = append(allFetchedLessons, fetchedLessons...)
		n.teachers[teacherID] = teacher
		if len(newAvailableLessons) > 0 {
			availableLessonsPerTeacher[teacherID] = newAvailableLessons
		}
	}

	if err := n.sendNotificationToUser(user, availableLessonsPerTeacher); err != nil {
		return err
	}

	if !n.dryRun {
		n.lessonService.UpdateLessons(allFetchedLessons)
	}

	return nil
}

// Returns teacher, fetchedLessons, newAvailableLessons, error
func (n *Notifier) fetchAndExtractNewAvailableLessons(teacherId uint32) (
	*model.Teacher, []*model.Lesson, []*model.Lesson, error,
) {
	teacher, fetchedLessons, err := lessonFetcher.Fetch(teacherId)
	if err != nil {
		return nil, nil, nil, err
	}
	//fmt.Printf("fetchedLessons ---\n")
	//for _, l := range fetchedLessons {
	//	fmt.Printf("teacherId=%v, datetime=%v, status=%v\n", l.TeacherId, l.Datetime, l.Status)
	//}

	now := time.Now()
	fromDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, config.LocalTimezone())
	toDate := fromDate.Add(24 * 6 * time.Hour)
	lastFetchedLessons, err := n.lessonService.FindLessons(teacher.ID, fromDate, toDate)
	if err != nil {
		return nil, nil, nil, err
	}
	//fmt.Printf("lastFetchedLessons ---\n")
	//for _, l := range lastFetchedLessons {
	//	fmt.Printf("teacherId=%v, datetime=%v, status=%v\n", l.TeacherId, l.Datetime, l.Status)
	//}

	newAvailableLessons := n.lessonService.GetNewAvailableLessons(lastFetchedLessons, fetchedLessons)
	//fmt.Printf("newAvailableLessons ---\n")
	//for _, l := range newAvailableLessons {
	//	fmt.Printf("teacherId=%v, datetime=%v, status=%v\n", l.TeacherId, l.Datetime, l.Status)
	//}
	return teacher, fetchedLessons, newAvailableLessons, nil
}

func (n *Notifier) sendNotificationToUser(
	user *model.User,
	lessonsPerTeacher map[uint32][]*model.Lesson,
) error {
	lessonsCount := 0
	var teacherIDs []int
	for teacherID, lessons := range lessonsPerTeacher {
		teacherIDs = append(teacherIDs, int(teacherID))
		lessonsCount += len(lessons)
	}
	if lessonsCount == 0 {
		// Don't send notification
		return nil
	}

	sort.Ints(teacherIDs)
	var teacherIDs2 []uint32
	var teacherNames []string
	for _, id := range teacherIDs {
		teacherIDs2 = append(teacherIDs2, uint32(id))
		teacherNames = append(teacherNames, n.teachers[uint32(id)].Name)
	}

	t := template.New("email")
	t = template.Must(t.Parse(getEmailTemplate()))
	type TemplateData struct {
		TeacherIDs        []uint32
		Teachers          map[uint32]*model.Teacher
		LessonsPerTeacher map[uint32][]*model.Lesson
	}
	data := &TemplateData{
		TeacherIDs:        teacherIDs2,
		Teachers:          n.teachers,
		LessonsPerTeacher: lessonsPerTeacher,
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return errors.InternalWrapf(err, "Failed to execute template.")
	}
	//fmt.Printf("--- mail ---\n%s", body.String())

	subject := "[lekcije] Schedule of teacher " + strings.Join(teacherNames, ", ")
	sender := &EmailNotificationSender{}
	return sender.Send(user, subject, body.String())
}

func getEmailTemplate() string {
	return strings.TrimSpace(`
{{- range $teacherID := .TeacherIDs }}
{{- $teacher := index $.Teachers $teacherID -}}
--- Available lessons of {{ $teacher.Name }} ---
PC: http://eikaiwa.dmm.com/teacher/index/{{ $teacherID }}/
Mobile: http://eikaiwa.dmm.com/teacher/schedule/{{ $teacherID }}/

  {{ $lessons := index $.LessonsPerTeacher $teacherID -}}
  {{- range $lesson := $lessons }}
{{ $lesson.Datetime.Format "2006-01-02 15:04" }}
  {{- end }}

{{ end }}
	`)
}

type NotificationSender interface {
	Send(user *model.User, subject, body string) error
}

type EmailNotificationSender struct{}

func (s *EmailNotificationSender) Send(user *model.User, subject, body string) error {
	from := mail.NewEmail("noreply", "noreply@lampetty.net") // TODO: noreply@lekcije.com
	to := mail.NewEmail(user.Name, user.Email.Raw())
	content := mail.NewContent("text/html", strings.Replace(body, "\n", "<br>", -1))
	m := mail.NewV3MailInit(from, subject, to, content)

	req := sendgrid.GetRequest(
		os.Getenv("SENDGRID_API_KEY"),
		"/v3/mail/send",
		"https://api.sendgrid.com",
	)
	req.Method = "POST"
	req.Body = mail.GetRequestBody(m)
	resp, err := sendgrid.API(req)
	if err != nil {
		return errors.InternalWrapf(err, "Failed to send email by sendgrid")
	}
	if resp.StatusCode >= 300 {
		message := fmt.Sprintf(
			"Failed to send email by sendgrid: statusCode=%v, body=%v",
			resp.StatusCode, strings.Replace(resp.Body, "\n", "\\n", -1),
		)
		logger.AppLogger.Error(message)
		return errors.InternalWrapf(
			err,
			"Failed to send email by sendgrid: statusCode=%v",
			resp.StatusCode,
		)
	}

	return nil
}
