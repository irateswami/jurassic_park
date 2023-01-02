package objects

type Cage struct {
	Id          string `json:"id" db:"id"`
	Carnivore   bool   `json:"carnivore" db:"carnivore"`
	Active      bool   `json:"active" db:"active"`
	MaxCapacity uint64 `json:"max_capacity" db:"max_capacity"`
}

type Dinosaur struct {
	Id      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Species string `json:"species" db:"species"`
	Cage    string `json:"cage" db:"cage"`
}
