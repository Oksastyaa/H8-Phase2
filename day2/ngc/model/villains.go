package model

import (
	"database/sql"
)

type Villains struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageUrl string `json:"image_url"`
}

// Create creates a new villain
func (v *Villains) Create(db *sql.DB) error {
	query := "INSERT INTO villains (name, universe, image_url) VALUES (?, ?, ?)"
	result, err := db.Exec(query, v.Name, v.Universe, v.ImageUrl)
	if err != nil {
		return err
	}
	//retrieve the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	v.Id = int(id)
	return nil
}

// GetByID retrieves a villain by its ID
func (v *Villains) GetByID(db *sql.DB, id int) error {
	query := "SELECT id, name, universe, image_url FROM villains WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&v.Id, &v.Name, &v.Universe, &v.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all villains
func (v *Villains) GetAll(db *sql.DB) ([]Villains, error) {
	query := "SELECT id, name, universe, image_url FROM villains"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var villains []Villains
	for rows.Next() {
		var villain Villains
		err := rows.Scan(&villain.Id, &villain.Name, &villain.Universe, &villain.ImageUrl)
		if err != nil {
			return nil, err
		}
		villains = append(villains, villain)
	}
	return villains, nil
}

// Update updates a villain
func (v *Villains) Update(db *sql.DB) error {
	query := "UPDATE villains SET name = ?, universe = ?, image_url = ? WHERE id = ?"
	_, err := db.Exec(query, v.Name, v.Universe, v.ImageUrl, v.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a villain
func (v *Villains) Delete(db *sql.DB) error {
	query := "DELETE FROM villains WHERE id = ?"
	_, err := db.Exec(query, v.Id)
	if err != nil {
		return err
	}
	return nil
}
