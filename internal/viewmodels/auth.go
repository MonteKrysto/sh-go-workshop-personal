package viewmodels

import "github.com/google/uuid"

type AuthedUser struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
