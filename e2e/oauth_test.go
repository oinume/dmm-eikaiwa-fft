package e2e

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/oinume/lekcije/server/model"
	"github.com/stretchr/testify/assert"
)

var _ = time.UTC
var _ = fmt.Print

func TestOAuthGoogle(t *testing.T) {
	if os.Getenv("CIRCLECI") != "" {
		t.Skipf("Skip because it can't render Google log in page.")
	}
	a := assert.New(t)
	driver := newWebDriver()
	err := driver.Start()
	a.Nil(err)
	defer driver.Stop()

	page, err := driver.NewPage()
	a.Nil(err)
	a.Nil(page.Navigate(server.URL))
	//time.Sleep(10 * time.Second)
	link := page.FindByXPath("//div[@class='starter-template']/a")
	u, err := link.Attribute("href")
	a.Nil(err)
	fmt.Printf("u = %v, err = %v\n", u, err)
	link.Click()
	//time.Sleep(10 * time.Second)

	googleAccount := os.Getenv("E2E_GOOGLE_ACCOUNT")
	err = page.Find("#Email").Fill(googleAccount)
	a.Nil(err)
	page.Find("#gaia_loginform").Submit()
	a.Nil(err)

	time.Sleep(time.Second * 1)
	page.Find("#Passwd").Fill(os.Getenv("E2E_GOOGLE_PASSWORD"))
	a.Nil(err)
	page.Find("#gaia_loginform").Submit()
	a.Nil(err)

	time.Sleep(time.Second * 3)
	err = page.Find("#submit_approve_access").Click()
	a.Nil(err)
	//time.Sleep(time.Second * 10)
	// TODO: Check HTML content

	cookies, err := page.GetCookies()
	a.Nil(err)
	apiToken := getAPIToken(cookies)
	a.NotEmpty(apiToken)

	user, err := model.NewUserService(db).FindByUserAPIToken(apiToken)
	a.Nil(err)
	a.Equal(googleAccount, user.Email.Raw())
}

func TestOAuthGoogleLogout(t *testing.T) {
	// TODO: user_api_token will be deleted after logout
}

func getAPIToken(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == "apiToken" {
			return cookie.Value
		}
	}
	return ""
}
