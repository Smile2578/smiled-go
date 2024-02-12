package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Superadmin Role = "superadmin"
	Admin      Role = "admin"
	Dentist    Role = "dentist"
	Assistant  Role = "assistant"
	Patient    Role = "patient"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id"`
	FirstName         string             `bson:"firstName"`
	LastName          string             `bson:"lastName"`
	Email             string             `bson:"email"`
	Phone             string             `bson:"phone"`
	PasswordHash      string             `bson:"passwordHash"`
	Role              Role               `bson:"role"`
	RPPSNumber        string             `bson:"rppsNumber,omitempty"`
	LinkedAssistantID primitive.ObjectID `bson:"linkedAssistantId,omitempty"`
	LinkedDentistID   primitive.ObjectID `bson:"linkedDentistId,omitempty"`
	RefreshToken      string             `bson:"refreshToken,omitempty"`
}
