type Appointment struct {
	ID        primitive.ObjectID `bson:"_id"`
	PatientID primitive.ObjectID `bson:"patientId"`
	Date      time.Time          `bson:"date"`
	Purpose   string             `bson:"purpose"`
	Notes     string             `bson:"notes"`
}
