package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name" bson:"name"`
	Address       string               `json:"address" bson:"address"`
	City          string               `json:"city" bson:"city"`
	Phone         string               `json:"phone" bson:"phone"`
	Description   string               `json:"description" bson:"description"`
	HabitacionIds []primitive.ObjectID `json:"habitacionId,omitempty" bson:"habitacionId,omitempty"`
	Registros     []primitive.ObjectID `json:"registros,omitempty" bson:"registros,omitempty"`
	CreatedAt     time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at" bson:"updated_at"`
}
