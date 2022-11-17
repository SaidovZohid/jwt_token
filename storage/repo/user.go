package repo

import "time"

type User struct {
	ID        int64
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(user_id int64) (*User, error)
}
