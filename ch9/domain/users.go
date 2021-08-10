package domain

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("user number %d, with username %s , pass : %s and email %s . Created at %s and updated at %s",
		u.Id, u.UserName, u.Password, u.Email, u.CreatedAt.String(), time.Since(u.UpdatedAt).String())
}

func (u *User) GoStringer() string {
	return fmt.Sprintf("user number %d, with username %s , pass : %s and email %s . Created at %s and updated at %s",
		u.Id, u.UserName, u.Password, u.Email, u.CreatedAt.String(), time.Since(u.UpdatedAt).String())
}
