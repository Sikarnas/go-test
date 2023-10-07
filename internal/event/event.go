package event

import (
	"time"
)

type CreateRequest struct {
	Name         string    `json:"name" validate:"required"`
	Date         time.Time `json:"date" validate:"required"`
	Languages    []string  `json:"languages" validate:"required"`
	VideoQuality []string  `json:"VideoQuality" validate:"required"`
	AudioQuality []string  `json:"AudioQuality" validate:"required"`
	Invitees     []string  `json:"invitees" validate:"required"`
	Description  *string   `json:"description"`
}

type Entity struct {
	ID string `json:"id" `
	CreateRequest
}
