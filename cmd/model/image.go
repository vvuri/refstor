package model

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ImageID     uuid.UUID  `json:"uuid"` // UUIDv4
	Description string     `json:"description"`
	SmallImg    []byte     `json:"small_img"` // save as "image/png" in base64
	Date        *time.Time `json:"date"`      // DateOnly "2006-01-02"
}
