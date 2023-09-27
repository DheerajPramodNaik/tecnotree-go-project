package entities

import "time"

type Register struct {
	Id              int       `json:"id" bson:"_id"`
	FirstName       string    `json:"FirstName" bson:"FirstName,required"`
	LastName        string    `json:"LastName" bson:"LastName,required"`
	Age             int       `json:"Age" bson:"Age,required"`
	Email           string    `json:"Email" bson:"Email,required"`
	Password        string    `json:"Password" bson:"Password,required"`
	ConfirmPassword string    `json:"ConfirmPassword" bson:"ConfirmPassword"`
	CreatedAt       time.Time `json:"CreatedAt" bson:"CreatedAt,required"`
	UpdatedAt       time.Time `json:"UpdatedAt" bson:"UpdatedAt,required"`
}

type Login struct {
	Email    string `json:"Email" bson:"Email,required"`
	Password string `json:"Password" bson:"Password,required"`
}
