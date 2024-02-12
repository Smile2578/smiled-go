type Patient struct {
	ID             primitive.ObjectID   `bson:"_id"`
	Name           string               `bson:"name"`
	DateOfBirth    time.Time            `bson:"dateOfBirth"`
	Gender         string               `bson:"gender"`
	ContactInfo    ContactInfo          `bson:"contactInfo"`
	MedicalHistory []MedicalRecord      `bson:"medicalHistory"`
	Appointments   []primitive.ObjectID `bson:"appointments"`
	Treatments     []primitive.ObjectID `bson:"treatments"`
}

type ContactInfo struct {
	Phone   string `bson:"phone"`
	Email   string `bson:"email"`
	Address string `bson:"address"`
}

type MedicalRecord struct {
	Date      time.Time `bson:"date"`
	Summary   string    `bson:"summary"`
	Diagnosis string    `bson:"diagnosis"`
	Treatment string    `bson:"treatment"`
}

