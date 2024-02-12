package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "yourapp/models" // Adjust to your actual models' import path
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    "github.com/go-playground/validator/v10"
)

var (
    client *mongo.Client // Assume initialized elsewhere
    validate = validator.New()
    collection = client.Database("yourDatabaseName").Collection("users")
)

func init() {
    // Register custom validation for the Role field
    validate.RegisterValidation("role", validateRole)
}

func validateRole(fl validator.FieldLevel) bool {
    role := models.Role(fl.Field().String())
    switch role {
    case models.Superadmin, models.Admin, models.Dentist, models.Assistant, models.Patient:
        return true
    }
    return false
}

// Custom error response to enhance error handling
type ErrorResponse struct {
    ErrorCode    int    `json:"errorCode"`
    ErrorMessage string `json:"errorMessage"`
}

func writeError(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(ErrorResponse{ErrorCode: code, ErrorMessage: message})
}

// UserInput struct now includes a validator for Role
type UserInput struct {
    FirstName string        `json:"firstName" validate:"required,alpha"`
    LastName  string        `json:"lastName" validate:"required,alpha"`
    Email     string        `json:"email" validate:"required,email"`
    Phone     string        `json:"phone" validate:"required"`
    Password  string        `json:"password" validate:"required,min=8"`
    Role      models.Role   `json:"role" validate:"required,role"`
}

// CreateUser - Enhanced to include better validation and error handling
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var input UserInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        writeError(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    if err := validate.Struct(input); err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "Failed to hash password")
        return
    }

    newUser := models.User{
        ID:           primitive.NewObjectID(),
        FirstName:    input.FirstName,
        LastName:     input.LastName,
        Email:        input.Email,
        Phone:        input.Phone,
        PasswordHash: string(hashedPassword),
        Role:         input.Role,
    }

    _, err = collection.InsertOne(context.Background(), newUser)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "Failed to create user")
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created successfully", "userId": newUser.ID.Hex()})
}

// DeleteUser - Function to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    userID, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	   if err != nil {
l {		        ht.Error(w, "Invalid user I ID", http.StatusBaequest)
equest)	      return
	   }

	   result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": userID})
c   if err != nil {
son.M{"_id":.Error(w, "Failed to delete user",http.Error(w, "Failed to rError)
le      return
"   }

t   if result.DeletedCount == 0 {
		writeError.Er http.StatusNotF      with given ID", http.StatusNotFound)
Co      return
0 {	   }

    w.WriteHeader(http.StatusOK)
o   json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// Custom validation rules and other CRUD operations should be added below
// Custom validation rules
validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	return role == "admin" || role == "user"
})

// UpdateUser - Function to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		writeError(w, "Invalid user ID", ht, "Invalid user ID")
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		t)iteError(w, http.StatusBadRequest,, http.StatusBadRequest)
		return
	}

	// Validate input
	if validationErr := validate.Struct(updatedUser); validationErr != nil {
		.Error(w, (w,valid request body", ht http.StatusBadRequest)
		return
	}

	// Update user in the database
	update := bson.M{
		"$set": bson.M{
			"firstName": updatedUser.FirstName,
			"lastName":  updatedUser.LastName,
			"email":     updatedUser.Email,
			"phone":     updatedUser.Phone,
			"role":      updatedUser.Role,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": userID}, update)
	if err != nil {
		http.Error(w, "Failed to update user", "Failed to update user", http.)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(w,given ID", http.StatusNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// GetUser - Function to get a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", "Invalid user ID", ht)
		return
	}

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		http.Error(w, "Failed		http.Error(w "Failed to get user", http.Sta)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// ListUsers - Function to list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to list users", "Failed to list users", http.S)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			ailed to l(w, users, http.StatusInter "Failed to decode user", http.)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}