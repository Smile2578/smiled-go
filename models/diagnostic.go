type Diagnostic struct {
	ID        primitive.ObjectID `bson:"_id"`
	PatientID primitive.ObjectID `bson:"patientId"`
	Date      time.Time          `bson:"date"`
	Details   DiagnosticDetails  `bson:"details"`
}

type DiagnosticDetails struct {
	// Add fields based on the "Intelligence Artificielle" document
	// Example: Carie, Absence, Orthodontie, etc.
}
