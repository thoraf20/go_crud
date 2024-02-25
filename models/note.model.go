package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID            			primitive.ObjectID 							`bson:'_id'`
	Text								*string 												`json:'_text'`
	Title								*string 												`json:'_title'`
	Created_at    			*time.Time 											`json:"created_at"`
	Update_at     			*time.Time 											`json:"updated_at"`
	Note_id							*string													`json:"note_id"`
}