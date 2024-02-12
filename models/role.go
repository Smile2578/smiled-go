// models/role.go

package models

type Role string

const (
	Superadmin Role = "superadmin"
	Admin      Role = "admin"
	Dentist    Role = "dentist"
	Assistant  Role = "assistant"
	Patient    Role = "patient"
)
