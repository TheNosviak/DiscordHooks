package discordhooks

import (
	"encoding/json"
	"time"
)

// Hook will help produce the sender required information for sending a webhook
type Hook struct {
	ID	 	 int
	Key      string
}

// New will attempt to create a new webhooking instance
func New(id int, key string) *Hook {
	return &Hook{
		ID: id,
		Key: key,
	}
}

// NewWebhook attempts to create a brand new webhook instance
func (H *Hook) NewWebhook(m *webhook) error {
	contents, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return H.forward(contents, 2 * time.Second)
}