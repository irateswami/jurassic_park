# About
This is a repo that I created to build an api, end to end, for a video to give an example of many different GoLang design paradigms.

# Reference
TODO : here is where the api reference docs will go.

## Cages
### PUT
1. You're required to input an id, everything else will either be empty or 0, false, etc.

# Specs
## Problem Statement
It's 1993 and you're the lead software developer for the new Jurassic Park! Park
operations needs a system to keep track of the different cages around the park and the
different dinosaurs in each one. You'll need to develop a JSON formatted RESTful API
to allow the builders to create new cages. It will also allow doctors and scientists the
ability to edit/retrieve the statuses of dinosaurs and cages.

## Requirements
### Business Requirements
- All requests should respond with the correct HTTP status codes and a response, if necessary, representing either the success or error conditions.
- Each dinosaur must have a name.
- Each dinosaur is considered an herbivore or a carnivore, depending on its species.
- Carnivores can only be in a cage with other dinosaurs of the same species.
- Each dinosaur must have a species (See enumerated list below, feel free to add others).
- Herbivores cannot be in the same cage as carnivores.
- Use Carnivore dinosaurs like Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus.
- Use Herbivores like Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops.

### Technical Requirements
- You are allowed to use scaffolding technology.
- This project should be done with version Golang 1.19 or newer.
- This project can use the Gin API framework or similar technology.
- This should be done using version control, preferably git.
- The project should include a README that addresses anything you may not have completed. It should also address what additional changes you might need to make if the application were intended to run in a concurrent environment. Any other comments or thoughts about the project are also welcome.

### Bonus
- Cages have a maximum capacity for how many dinosaurs it can hold.
- Cages know how many dinosaurs are contained.
- Cages have a power status of ACTIVE or DOWN.
- Cages cannot be powered off if they contain dinosaurs.
- Dinosaurs cannot be moved into a cage that is powered down.
- Must be able to query a listing of dinosaurs in a specific cage.
- When querying dinosaurs or cages they should be filterable on their attributes (Cages on their power status and dinosaurs on species).
- Automated tests that ensure the business logic implemented is correct.
