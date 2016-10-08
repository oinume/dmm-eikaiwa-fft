package main

import (
	"os"
)

type Table struct {
	PrimaryKey PrimaryKey
	Columns []Column
}

type Column struct {
	Name string
	Type string
	Nullable bool
	DefaultValue string
	Extra string
	Comment string
}

type PrimaryKey struct {
	Columns []Column
}

func main() {
	dbURL := os.Getenv("DB_URL")
	println(dbURL)
	// テーブル一覧を取得
	// カラム一覧を取得
}

/*
type User struct {
	ID            uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name          string
	Email         Email
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
*/
