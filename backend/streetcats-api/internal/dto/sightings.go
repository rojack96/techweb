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

type CreateSightingDTO struct {
	ID          uint64     `json:"id,omitempty"`
	AnimalID    uint64     `json:"animalId"`
	BreedID     uint64     `json:"breedId"`
	Position    [2]float64 `json:"position"` // [latitude, longitude]
	Title       string     `json:"title"`
	Description string     `json:"description"`
	SpottedAt   *int64     `json:"spottedAt"`
}

type BreedDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}
