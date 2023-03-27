package viewmodels

import "github.com/google/uuid"

type Profile struct {
	ID      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"user_id"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
}
