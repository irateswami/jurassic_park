package objects

const (
	// Carnivore dinosaurs like Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus.
	CARNIVORE_TYRANNOSAURUS = "Tyrannosaurus"
	CARNIVORE_VELOCIRAPTOR  = "Velociraptor"
	CARNIVORE_SPINOSAURUS   = "Spinosaurus"
	CARNIVORE_MEGALOSAURUS  = "Megalosaurus"

	//Herbivores like Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops.
	HERBIVORES_BRACHIOSAURUS = "Brachiosaurus"
	HERBIVORES_STEGOSAURUS   = "Stegosaurus"
	HERBIVORES_ANKYLOSAURUS  = "Ankylosaurus"
	HERBIVORES_TRICERATOPS   = "Triceratops"
)

const (
	CARNIVORE = 0
	HERBIVORE = 1
)

type Cage struct {
	Id          uint8
	HerbOrCarn  uint8
	MaxCapacity uint64
}

type Dinosaur struct {
	Id         uint8  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Species    string `json:"species" db:"species"`
	HerbOrCarn uint8  `json:"herb_or_carn" db:"herb_or_carn"`
}
