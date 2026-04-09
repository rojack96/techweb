package dto

type SightingDTO struct {
	ID        uint64     `json:"id"`
	AnimalID  uint64     `json:"animalId"`
	BreedID   *uint64    `json:"breedId"`
	Position  [2]float64 `json:"position"` // [latitude, longitude]
	SpottedAt *int64     `json:"spottedAt"`
}
