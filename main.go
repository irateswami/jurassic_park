package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	PARK          JurassicPark
	INFOLOGGER    *log.Logger
	WARNINGLOGGER *log.Logger
	ERRORLOGGER   *log.Logger
)

func isCarnivore(species string) bool {
	switch species {
	case "Tyrannosaurus":
		return true
	case "Velociraptor":
		return true
	case "Spinosaurus":
		return true
	case "Megalosaurus":
		return true

	case "Brachiosaurus":
		return false
	case "Stegosaurus":
		return false
	case "Ankylosaurus":
		return false
	case "Triceratops":
		return false

	default:
		return false
	}
}

type JurassicPark struct {
	cages map[string]Cage
	mut   *sync.RWMutex
}

type Cage struct {
	Id          string              `json:"id"`
	Species     Species             `json:"species"`
	MaxCapacity uint64              `json:"max_capacity"`
	Dinosaurs   map[string]Dinosaur `json:"dinosaurs"`
	Active      bool                `json:"active"`
}

type Species struct {
	Species   string `json:"species"`
	Carnivore bool   `json:"carnivore"`
}

type Dinosaur struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
}

type dinosPutRequest struct {
	Dinosaur Dinosaur `json:"dinosaur"`
	CageID   string   `json:"cage_id"`
}

func init() {
	var mut sync.RWMutex
	cages := make(map[string]Cage)

	PARK = JurassicPark{
		mut:   &mut,
		cages: cages,
	}

	INFOLOGGER = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARNINGLOGGER = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERRORLOGGER = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func cageHandler(w http.ResponseWriter, req *http.Request) {

	// always going to be this
	w.Header().Set("Content-Type", "application/json")

	// What type of request are we handling?
	switch req.Method {

	// this is the hard one, rest is cake
	case "PUT": //---------------------------------------------------------------------------------

		var cagePutRequestBody []Cage

		// decode the array of requests
		err := json.NewDecoder(req.Body).Decode(&cagePutRequestBody)
		if err != nil {
			ERRORLOGGER.Println("parsing cage PUT request body went bad")
			ERRORLOGGER.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// now lock until we're done
		PARK.mut.Lock()
		defer PARK.mut.Unlock()

		var badRequest bool // don't quit because of one thing
		for _, v := range cagePutRequestBody {

			INFOLOGGER.Printf("%+v\n", v)

			// does the cage exist?
			if _, present := PARK.cages[v.Id]; present {
				badRequest = true
				ERRORLOGGER.Printf("%s cageID already exists\n", v.Id)
				continue // move on
			}

			PARK.cages[v.Id] = v
		}

		if badRequest == true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func dinoHandler(w http.ResponseWriter, req *http.Request) {

	// always going to be this
	w.Header().Set("Content-Type", "application/json")

	// What type of request are we handling?
	switch req.Method {

	// this is the hard one, rest is cake
	case "PUT": //---------------------------------------------------------------------------------

		var dinosPutRequestBody []dinosPutRequest

		// decode the array of requests
		err := json.NewDecoder(req.Body).Decode(&dinosPutRequestBody)
		if err != nil {
			ERRORLOGGER.Println("parsing dino request body went bad")
			ERRORLOGGER.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// now lock until we're done
		PARK.mut.Lock()
		defer PARK.mut.Unlock()

		var badRequest bool // don't quit because of one thing
		for _, v := range dinosPutRequestBody {

			// does the cage exist?
			tempCage, present := PARK.cages[v.CageID]
			if !present {
				badRequest = true
				ERRORLOGGER.Printf("#%s cage doesn't exist\n", v.CageID)
				continue // move on
			}

			// is it active?
			if tempCage.Active != true {
				badRequest = true
				ERRORLOGGER.Printf("#%s cage isn't active\n", v.CageID)
				continue // move on
			}

			// is it empty?
			if len(tempCage.Dinosaurs) == 0 {

				tempCage.Dinosaurs = make(map[string]Dinosaur)

				// set everything
				tempCage.Species.Carnivore = isCarnivore(v.Dinosaur.Species)
				tempCage.Dinosaurs[v.Dinosaur.ID] = v.Dinosaur
				continue // move on
			}

			// what species are allowed
			if tempCage.Species.Carnivore == true && tempCage.Species.Species != v.Dinosaur.Species {
				badRequest = true
				ERRORLOGGER.Printf("%s species isn't compatible with this cage\n", v.CageID)
				continue // move on
			}

			// no special cases
			tempCage.Dinosaurs[v.Dinosaur.ID] = v.Dinosaur

		}

		if badRequest == true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case "GET": //---------------------------------------------------------------------------------

		//		objectID := names[3]
		//
		//		// only getting so read lock is appropriate
		//		PARK.mut.RLock()
		//		defer PARK.mut.RUnlock()
		//
		//		resp := PARK.get(repositoryName, objectID)
		//		w.WriteHeader(resp.Code)
		//		w.Write(resp.Data)

	case "DELETE": //------------------------------------------------------------------------------
		//		objectID := names[3]
		//
		//		// only getting so read lock is appropriate
		//		PARK.mut.RLock()
		//		defer PARK.mut.RUnlock()
		//		w.WriteHeader(PARK.delete(repositoryName, objectID))
	default: //------------------------------------------------------------------------------------
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/dinosaurs/", dinoHandler)
	http.HandleFunc("/cages/", cageHandler)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8282", nil))
}
