package domain

import "time"

type House struct {
	ID              int64      `json:"id"`
	Address         string     `json:"address"`
	Year            int16      `json:"year"`
	Developer       *string    `json:"developer,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	LastFlatAddedAt *time.Time `json:"last_flat_added_at,omitempty"`
}
