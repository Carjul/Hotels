package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType struct {
	ID   int    `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type Users struct {
	ID         primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Nombre     string               `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Status     bool                 `json:"status,omitempty" bson:"status,omitempty"`
	Imagen     string               `json:"imagen,omitempty" bson:"imagen,omitempty"`
	Correo     string               `json:"correo,omitempty" bson:"correo,omitempty"`
	Contraseña string               `json:"contraseña,omitempty" bson:"contraseña,omitempty"`
	Rol        UserType             `json:"rol,omitempty" bson:"rol,omitempty"`
	HotelIds   []primitive.ObjectID `json:"hotelId,omitempty" bson:"hotelId,omitempty"`
	CreatedAt  time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at" bson:"updated_at"`
}

var Roles = []UserType{
	{ID: 1, Name: "superAdmin"},
	{ID: 2, Name: "administrador"},
	{ID: 3, Name: "receptionista"},
	{ID: 4, Name: "cliente"},
}
