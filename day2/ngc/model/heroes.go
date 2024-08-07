package model

import (
	"database/sql"
)

type Heroes struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageUrl string `json:"image_url "`
}

// Create creates a new hero
func (h *Heroes) Create(db *sql.DB) error {
	query := "INSERT INTO heroes (name, universe, skill, image_url) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, h.Name, h.Universe, h.Skill, h.ImageUrl)
	if err != nil {
		return err

	}
	return nil
}

// GetByID retrieves a hero by its ID
func (h *Heroes) GetByID(db *sql.DB, id int) error {
	query := "SELECT id,name, universe, skill, image_url FROM heroes WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&h.Id, &h.Name, &h.Universe, &h.Skill, &h.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all heroes
func (h *Heroes) GetAll(db *sql.DB) ([]Heroes, error) {
	query := "SELECT id, name, universe, skill, image_url FROM heroes"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var heroes []Heroes
	for rows.Next() {
		var hero Heroes
		err := rows.Scan(&hero.Id, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageUrl)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}
	return heroes, nil
}

// Update updates a hero
func (h *Heroes) Update(db *sql.DB) error {
	query := "UPDATE heroes SET name = ?, universe = ?, skill = ?, image_url = ? WHERE id = ?"
	_, err := db.Exec(query, h.Name, h.Universe, h.Skill, h.ImageUrl, h.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a hero
func (h *Heroes) Delete(db *sql.DB) error {
	query := "DELETE FROM heroes WHERE id = ?"
	_, err := db.Exec(query, h.Id)
	if err != nil {
		return err
	}
	return nil
}
