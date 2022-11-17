package models

import "time"

type RegisterUser struct {
	Name     string `json:"name" binding:"required,min=2,max=30"`
	Username string `json:"username" binding:"required,min=2,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type RegisterResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	AccessToken string    `json:"access_token"`
}
