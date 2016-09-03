package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/model"
	"github.com/oinume/lekcije/server/notifier"
)

var (
	dryRun   = flag.Bool("dry-run", false, "Dry run")
	logLevel = flag.String("log-level", "info", "Log level")
)

// TODO: move somewhere proper and make it be struct
var definedEnvs = map[string]string{
	"DB_DSN":         "",
	"NODE_ENV":       "",
	"ENCRYPTION_KEY": "",
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("err = %v", err) // TODO: Error handling
		os.Exit(1)
	}
}

func run() error {
	// Check env
	for key := range definedEnvs {
		if value := os.Getenv(key); value != "" {
			definedEnvs[key] = value
		} else {
			log.Fatalf("Env '%v' must be defined.", key)
		}
	}

	db, _, err := model.OpenAndSetToContext(context.Background(), definedEnvs["DB_DSN"])
	if err != nil {
		return err
	}

	var users []*model.User
	// TODO: email_verified
	userSql := `SELECT * FROM user /* WHERE email_verified = 1 */`
	result := db.Raw(userSql).Scan(&users)
	if result.Error != nil && !result.RecordNotFound() {
		return errors.InternalWrapf(result.Error, "")
	}

	notifier := notifier.NewNotifier()
	for _, user := range users {
		if err := notifier.SendNotification(user); err != nil {
			panic(err)
			return err
		}
	}

	return nil
}
