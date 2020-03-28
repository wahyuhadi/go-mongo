package models

import (
	"gopkg.in/mgo.v2/bson"
	// Import forms
	"framework/forms"
	"framework/helpers"
	"time"
)

// User defines user object structure
type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
	CreatedAt time.Time 	 `json:"created_at,omitempty" bson:"created_at,omitempty"`
 	UpdatedAt time.Time 	 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}


type DetailUsers struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
	CreatedAt time.Time 	 `json:"created_at,omitempty" bson:"created_at,omitempty"`
 	UpdatedAt time.Time 	 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// UserModel defines the model structure
type UserModel struct{}

// Signup handles registering a user
func (u *UserModel) Signup(data forms.SignupUserCommand) error {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
		// This will come later when adding verification
		"is_verified": false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})

	return err
}

// GetUserByEmail handles fetching user by email
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err = collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}

// GetUserByEmail handles fetching user by email
func (u *UserModel) GetAllUser() (user []User, err error) {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err = collection.Find(nil).Select(bson.M{"email": 1, "name": 1}).All(&user)
	return user, err
}


func (u *UserModel) GetUserDetails(email string) (user DetailUsers, err error) {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err = collection.Find(bson.M{"email": email}).Select(bson.M{"password": 0}).One(&user)
	return user, err
}

