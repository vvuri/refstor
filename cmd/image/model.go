package image

import (
	"time"
)

type ImageLink struct {
	ImageID     string     `json:"uuid"` // UUIDv4
	Description string     `json:"description"`
	SmallImg    []byte     `json:"small_img"` // save as "image/png" in base64
	Date        *time.Time `json:"date"`      // DateOnly "2006-01-02"
	URL         string     `json:"url"`
}
