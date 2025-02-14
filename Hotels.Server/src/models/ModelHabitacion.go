package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Habitacion struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nombre      string             `json:"nombre" bson:"nombre"`
	Tiempo      string             `json:"tiempo" bson:"tiempo"`
	Valor       float64            `json:"valor" bson:"valor"`
	Status      string             `json:"status" bson:"status"`
	Descripcion string             `json:"descripcion" bson:"descripcion"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
