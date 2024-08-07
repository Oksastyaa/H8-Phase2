package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ngc/model"
	"strconv"
)

// CreateVillainHandler creates a new villain
func CreateVillainHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var villain model.Villains

		err := json.NewDecoder(r.Body).Decode(&villain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = villain.Create(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//success response
		response := struct {
			Message string         `json:"message"`
			Villain model.Villains `json:"villain"`
		}{
			Message: "Villain created successfully",
			Villain: villain,
		}

		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to create Villain", http.StatusInternalServerError)
			return
		}
	}
}

// GetAllVillainsHandler retrieves all villains
func GetAllVillainsHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var villain model.Villains
		villains, err := villain.GetAll(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(villains)
	}
}

// GetVillainByIDHandler retrieves a villain by its ID
func GetVillainByIDHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var villain model.Villains
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			errorResponse := map[string]string{"message": "Invalid Villain ID"}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encondeErr := json.NewEncoder(w).Encode(errorResponse); encondeErr != nil {
				http.Error(w, encondeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		err = villain.GetByID(db, id)
		if err != nil {
			// error response for invalid ID
			errorResponse := map[string]string{"message": "Villain Not Found"}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encondeErr := json.NewEncoder(w).Encode(errorResponse); encondeErr != nil {
				http.Error(w, encondeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		// success response
		response := struct {
			Villain model.Villains `json:"villain"`
			Message string         `json:"message"`
		}{
			Villain: villain,
			Message: "Villain retrieved successfully",
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}
	}
}

// UpdateVillainHandler updates a villain by its ID
func UpdateVillainHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var villain model.Villains
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			// Error response for invalid ID
			errorResponse := map[string]string{"message": "Invalid ID"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if encodeErr := json.NewEncoder(w).Encode(errorResponse); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		err = json.NewDecoder(r.Body).Decode(&villain)
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
		// Set the ID from URL parameter to the villain struct
		villain.Id = id

		err = villain.Update(db)
		if err != nil {
			// Error response if update fails
			errorResponse := map[string]string{"message": "Failed to update villain"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if encodeErr := json.NewEncoder(w).Encode(errorResponse); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Success response
		response := struct {
			Villain model.Villains `json:"villain"`
			Message string         `json:"message"`
		}{
			Villain: villain,
			Message: "Villain updated successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}
	}
}

// DeleteVillainHandler deletes a villain by its ID
func DeleteVillainHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var villain model.Villains
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid Hero ID", http.StatusBadRequest)
			return
		}

		// Get the villain by its ID
		err = villain.GetByID(db, id)
		if err != nil {
			http.Error(w, "Villains Not Found", http.StatusNotFound)
			return
		}

		err = villain.Delete(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hero deleted successfully"})
	}
}
