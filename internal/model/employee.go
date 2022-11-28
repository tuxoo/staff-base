package model

import "time"

type Gender string

const (
	Male   = "Male"
	Female = "Female"
)

type NewEmployeeDto struct {
	FullName string `json:"fullName" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Gender   Gender `json:"gender" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type Employee struct {
	Id           int       `json:"id"`
	FullName     string    `json:"fullName"`
	Phone        string    `json:"phone"`
	Gender       Gender    `json:"gender"`
	Age          int       `json:"age"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
}
