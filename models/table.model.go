package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID 									primitive.ObjectID 				`bson:'_id'`
	Number_of_guests		*int                    	`json:'number_of_guests'`
	Table_number        *int                    	`json:'table_number`
	Created_at          *time.Time                `json:"created_at"`
	Update_at           *time.Time                `json:"created_at"`
	Table_id            string                    `json:"table_id" validate:"required"`
}
