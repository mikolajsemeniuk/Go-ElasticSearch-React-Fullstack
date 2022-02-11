package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id       uuid.UUID  `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated"`
}

func NewAccount() Account {
	return Account{
		Id:      uuid.New(),
		Created: time.Now(),
	}
}
