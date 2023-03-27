package viewmodels

import "github.com/google/uuid"

type UserImage struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	ImageId  string    `json:"image_id"`
	ImageUrl string    `json:"image_url"`
}
