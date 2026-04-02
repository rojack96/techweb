package dto

type SightingDTO struct {
	ID        uint64  `json:"id"`
	AnimalID  uint64  `json:"animal_id"`
	BreedID   *uint64 `json:"breed_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	SpottedAt *int64  `json:"spotted_at"`
}
