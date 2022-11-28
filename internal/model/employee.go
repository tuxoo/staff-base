package model

import "time"

type Gender string

type NewEmployeeDto struct {
	FullName string `json:"fullName" xml:"fullName" binding:"required"`
	Phone    string `json:"phone" xml:"phone" binding:"required"`
	Gender   Gender `json:"gender" xml:"gender" binding:"required"`
	Age      int    `json:"age" xml:"age" binding:"required"`
	Email    string `json:"email" xml:"email" binding:"required"`
	Address  string `json:"address" xml:"address" binding:"required"`
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
