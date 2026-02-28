package domain

import "time"

type User struct {
	ID           int64
	FirstName    string
	SecondName   string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
}
