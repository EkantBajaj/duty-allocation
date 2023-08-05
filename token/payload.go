package token

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Payload is the payload of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p Payload) Valid() error {
	//check if the token is expired or not
	if time.Now().After(p.ExpiredAt) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

// NewPayload creates a new payload for the token with a specific username and duration and return payload and error
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	// Create a new payload
	payload := &Payload{
		ID:        uuid.New(),
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
