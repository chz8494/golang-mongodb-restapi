package models

import (
	"github.com/chz8494/golang-mongodb-restapi/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// connect db
var collection = config.ConnectDB()

//Article fields
type Coin struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp    int64              `json:"timestam,omitempty" bson:"timestam,omitempty" validate:"required"`
	Price        float64            `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Initialprice string             `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	High         string             `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Low          string             `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
/*
	Title   string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" bson:"content,omitempty" validate:"required"`
	Author  string             `json:"author,omitempty" bson:"author,omitempty" validate:"required"`
*/
}

//Articles of article
type Coins []Article
