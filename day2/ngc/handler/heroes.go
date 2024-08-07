package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ngc/model"
	"strconv"
)

// CreateHeroHandler creates a new hero
func CreateHeroHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var hero model.Heroes

		err := json.NewDecoder(r.Body).Decode(&hero)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = hero.Create(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a map to hold the response data without Id
		// Create a response struct with a success message and hero data
		response := struct {
			Message  string `json:"message"`
			Name     string `json:"name"`
			Universe string `json:"universe"`
			Skill    string `json:"skill"`
			ImageUrl string `json:"image_url"`
		}{
			Message:  "Hero created successfully",
			Name:     hero.Name,
			Universe: hero.Universe,
			Skill:    hero.Skill,
			ImageUrl: hero.ImageUrl,
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed Create Hero", http.StatusInternalServerError)
			return
		}
	}
}

// GetAllHeroesHandler retrieves all heroes
func GetAllHeroesHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hero model.Heroes
		heroes, err := hero.GetAll(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(heroes)
	}
}

// GetHeroByIDHandler retrieves a hero by its ID
func GetHeroByIDHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hero model.Heroes
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			// error response for invalid ID
			errorResponse := map[string]string{"message": "Invalid Hero ID"}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encondeErr := json.NewEncoder(w).Encode(errorResponse); encondeErr != nil {
				http.Error(w, encondeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		err = hero.GetByID(db, id)
		if err != nil {
			// error response for invalid ID
			errorResponse := map[string]string{"message": "Hero Not Found"}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encondeErr := json.NewEncoder(w).Encode(errorResponse); encondeErr != nil {
				http.Error(w, encondeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Success response
		response := struct {
			Hero    model.Heroes `json:"hero"`
			Message string       `json:"message"`
		}{
			Hero:    hero,
			Message: "Hero retrieved successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}
	}
}

// UpdateHeroHandler updates a hero
func UpdateHeroHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hero model.Heroes
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			// error response for invalid ID
			errorResponse := map[string]string{"message": "Invalid ID"}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encondeErr := json.NewEncoder(w).Encode(errorResponse); encondeErr != nil {
				http.Error(w, encondeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		err = json.NewDecoder(r.Body).Decode(&hero)
		if err != nil {
			// Error response for invalid JSON body
			errorResponse := map[string]string{"message": "Invalid request body"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encodeErr := json.NewEncoder(w).Encode(errorResponse); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Set the ID from URL parameter to the heroes struct
		hero.Id = id
		err = hero.Update(db)
		if err != nil {
			// Error response if update fails
			errorResponse := map[string]string{"message": "Failed to update Hero"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if encodeErr := json.NewEncoder(w).Encode(errorResponse); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Success response
		response := struct {
			Hero    model.Heroes `json:"hero"`
			Message string       `json:"message"`
		}{
			Hero:    hero,
			Message: "Hero Updated successfully",
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}
	}
}

// DeleteHeroHandler deletes a hero
func DeleteHeroHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hero model.Heroes
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid Hero ID", http.StatusBadRequest)
			return
		}

		// Get the hero by its ID
		err = hero.GetByID(db, id)
		if err != nil {
			http.Error(w, "Hero Not Found", http.StatusNotFound)
			return
		}
		err = hero.Delete(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hero deleted successfully"})
	}
}
