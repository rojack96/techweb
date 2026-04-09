package dto

type SightingDTO struct {
	ID          string     `json:"id"`
	SightingID  uint64     `json:"sightingId"`
	Breed       *string    `json:"breed,omitempty"`
	Position    [2]float64 `json:"position"` // [latitude, longitude]
	Title       string     `json:"title"`
	Description string     `json:"description"`
	SpottedAt   *int64     `json:"spottedAt"`
}
