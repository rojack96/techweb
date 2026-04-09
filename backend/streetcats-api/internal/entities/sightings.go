package entities

type Animal struct {
	ID        uint64 `db:"id"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
}

type AnimalEntities struct {
	ID          uint64  `db:"id"`
	AnimalID    uint64  `db:"animal_id"`
	BreedID     *uint64 `db:"breed_id"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	SpottedAt   *int64  `db:"spotted_at"`
	CreatedAt   int64   `db:"created_at"`
}

type AnimalEntitiesView struct {
	ID          uint64  `db:"id"`
	AnimalID    uint64  `db:"animal_id"`
	Breed       *string `db:"breed"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	SpottedAt   *int64  `db:"spotted_at"`
	CreatedAt   int64   `db:"created_at"`
}

type Breed struct {
	ID        uint64 `db:"id"`
	AnimalID  uint64 `db:"animal_id"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
}

type Sighting struct {
	AnimalEntityID uint64  `db:"animal_entity_id"`
	Latitude       float64 `db:"latitude"`
	Longitude      float64 `db:"longitude"`
	SpottedAt      *int64  `db:"spotted_at"`
	CreatedAt      int64   `db:"created_at"`
}
