package viewmodels

import "github.com/google/uuid"

type Campaign struct {
	ID          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	IsCompleted bool      `json:"is_completed"`
	IsDeleted   bool      `json:"is_deleted"`
	IsFeatured  bool      `json:"is_featured"`
	IsSponsored bool      `json:"is_sponsored"`
	IsVerified  bool      `json:"is_verified"`
}
