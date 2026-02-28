package dto

import "time"

type CreateUserRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}

type CreateUserResponse struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"createdAt"`
}

type UpdateUserRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}

type GetUserListResponse struct {
	Users []GetUserResponse `json:"users"`
}

type GetUserResponse struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"createdAt"`
}
