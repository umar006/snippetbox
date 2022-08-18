package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Create         time.Time
}

type UserModel struct {
	DB *sql.DB
}
