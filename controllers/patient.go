package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"yourapp/models" // Adjust the import path to your models' location

	"go.mongodb.org/mongo-driver/bson"
)

// GetPatientDetails handles fetching patient details
func GetPatientDetails(w http.ResponseWriter, r *http.Request) {
	// Extract patient ID from request, for example, using URL parameters
	patientID := r.URL.Query().Get("id")

	var patient models.Patient
	collection := client.Database("your_database").Collection("patients")
	if err := collection.FindOne(context.TODO(), bson.M{"_id": patientID}).Decode(&patient); err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}
