package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id        primitive.ObjectID `bson:"_id"`
	ProductId int                `json:"ProductId" bson:"ProductId,required"`
	Name      string             `json:"Name" bson:"Name,required"`
	Category  string             `json:"Category" bson:"Category,required"`
	Quantity  int                `json:"Quantity" bson:"Quantity,required"`
	CreatedAt time.Time          `json:"CreatedAt" bson:"CreatedAt,required"`
	UpdatedAt time.Time          `json:"UpdatedAt" bson:"UpdatedAt,required"`
}

type SearchById struct {
	ProductId int `json:"ProductId" bson:"ProductId,required"`
}
