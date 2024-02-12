type Specialty struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"` // e.g., Soins, Endodontie, etc.
}
