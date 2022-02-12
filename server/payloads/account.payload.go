package payloads

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id       uuid.UUID  `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated"`
}
