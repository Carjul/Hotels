package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Registro struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Cedula          string             `json:"cedula" bson:"cedula"`
	Nombre          string             `json:"nombre" bson:"nombre"`
	FechaNacimiento string             `json:"fecha_nacimiento" bson:"fecha_nacimiento"`
	CiudadOrigen    string             `json:"ciudadOrigen" bson:"ciudadOrigen"`
	CiudadDestino   string             `json:"ciudadDestino" bson:"ciudadDestino"`
	Tiempo          string             `json:"tiempo" bson:"tiempo"`
	Celular         string             `json:"celular" bson:"celular"`
	HabitacionId    primitive.ObjectID `json:"habitacionId,omitempty" bson:"habitacionId,omitempty"`
	TipoRegistro    string             `json:"tipoRegistro" bson:"tipoRegistro"`
	Entrada         time.Time          `json:"entrada" bson:"entrada"`
	Salida          time.Time          `json:"salida" bson:"salida"`
	Total           float64            `json:"total" bson:"total"`
	Pago            bool               `json:"pago" bson:"pago"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
