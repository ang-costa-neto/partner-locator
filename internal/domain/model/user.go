package model

import "gorm.io/gorm"

type UserType string

const (
	Admin        UserType = "Admin"
	SuperAdmin   UserType = "SuperAdmin"
	Collaborator UserType = "Collaborator"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	UserType  UserType
}
