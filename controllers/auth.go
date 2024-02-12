package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"yourapp/models" // Adjust the import path to your models' location

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Assuming you have a setup like this for your MongoDB client
var client *mongo.Client // Initialize this in your main function or a setup function

// SignInHandler handles user authentication
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you'd hash the password provided and compare it with the one in the database
	// For demonstration, this is simplified
	var dbUser models.User
	collection := client.Database("your_database").Collection("users")
	if err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&dbUser); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Assuming you verify the password here (omitted for brevity)

	// Generate JWT token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.Claims{
		Email: dbUser.Email,
		Role:  string(dbUser.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("YourSigningKeyFromEnv")) // Use the key from your env

	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Return token and potentially user info
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
