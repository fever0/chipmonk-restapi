package main

import "time"

// User is a sruct
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username" gorm:"unique_index"`
	Password  string    `json:"password"`
	Session   uint      `json:"-"`
	CreatedAt time.Time `json:"time"`
}

// ActiveUser is a sruct
type ActiveUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
