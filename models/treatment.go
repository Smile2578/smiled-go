package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Treatment struct {
	ID            primitive.ObjectID `bson:"_id"`
	PatientID     primitive.ObjectID `bson:"patientId"`
	Type          string             `bson:"type"`      // Description of the act
	Specialty     string             `bson:"specialty"` // e.g., Soins, Endodontie, etc.
	Price         float64            `bson:"price"`
	CostPrice     float64            `bson:"costPrice"` // Price of procurement or cost to the clinic
	Quantity      int                `bson:"quantity"`
	CCAMCode      string             `bson:"ccamCode"`      // CCAM code for the treatment
	AssignedTeeth []string           `bson:"assignedTeeth"` // Array of teeth numbers or regions
	Status        string             `bson:"status"`
	StartDate     time.Time          `bson:"startDate"`
	EndDate       time.Time          `bson:"endDate,omitempty"`
	Details       string             `bson:"details"`
}
