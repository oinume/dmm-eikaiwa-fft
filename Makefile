E2E_TEST_ARGS=-v
GO_TEST_ARGS=-v
GO_TEST_PACKAGES=$(shell glide novendor | grep -v e2e)
#DB_URL=lekcije:lekcije@tcp(192.168.99.100:13306)/lekcije

all: install

.PHONY: install

setup:
	go get github.com/Masterminds/glide
	go get golang.org/x/tools/cmd/goimports
	glide install
	go install ./vendor/bitbucket.org/liamstask/goose/cmd/goose
	go install ./vendor/github.com/cespare/reflex

serve:
	go run server/cmd/lekcije/main.go

install:
	go install github.com/oinume/lekcije/server/cmd/lekcije

e2e_test:
	go test $(E2E_TEST_ARGS) github.com/oinume/lekcije/e2e

go_test:
	go test $(GO_TEST_ARGS) $(GO_TEST_PACKAGES)

minify_static:
	MINIFY=true VERSION_HASH=$(shell git rev-parse HEAD) npm run build

reset_db:
	mysql -h 192.168.99.100 -P 13306 -uroot -proot -e "DROP DATABASE IF EXISTS lekcije"
	mysql -h 192.168.99.100 -P 13306 -uroot -proot -e "DROP DATABASE IF EXISTS lekcije_test"
	mysql -h 192.168.99.100 -P 13306 -uroot -proot < db/create_database.sql
