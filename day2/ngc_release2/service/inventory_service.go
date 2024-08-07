package service

import (
	"database/sql"
	"errors"
	"ngc_release2/entity"
)

type InventoryService interface {
	CreateInventory(db *sql.DB, inv *entity.Inventory) error
	GetInventory(db *sql.DB, id int) (*entity.Inventory, error)
	GetInventories(db *sql.DB) ([]*entity.Inventory, error)
	UpdateInventory(db *sql.DB, inv *entity.Inventory) error
	DeleteInventory(db *sql.DB, id int) error
}

type inventoryService struct{}

func NewInventoryService() InventoryService {
	return &inventoryService{}
}

func (s *inventoryService) CreateInventory(db *sql.DB, inv *entity.Inventory) error {
	query := "INSERT INTO inventories (name, code, stock, description, status) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, inv.Name, inv.Code, inv.Stock, inv.Description, inv.Status)
	return err
}

func (s *inventoryService) GetInventory(db *sql.DB, id int) (*entity.Inventory, error) {
	query := "SELECT id, name, code, stock, description, status FROM inventories WHERE id = ?"
	row := db.QueryRow(query, id)

	inv := &entity.Inventory{}
	err := row.Scan(&inv.ID, &inv.Name, &inv.Code, &inv.Stock, &inv.Description, &inv.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return inv, nil
}

func (s *inventoryService) GetInventories(db *sql.DB) ([]*entity.Inventory, error) {
	query := "SELECT id, name, code, stock, description, status FROM inventories"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []*entity.Inventory
	for rows.Next() {
		inv := &entity.Inventory{}
		err := rows.Scan(&inv.ID, &inv.Name, &inv.Code, &inv.Stock, &inv.Description, &inv.Status)
		if err != nil {
			return nil, err
		}
		inventories = append(inventories, inv)
	}
	return inventories, nil
}

func (s *inventoryService) UpdateInventory(db *sql.DB, inv *entity.Inventory) error {
	query := "UPDATE inventories SET name = ?, code = ?, stock = ?, description = ?, status = ? WHERE id = ?"
	_, err := db.Exec(query, inv.Name, inv.Code, inv.Stock, inv.Description, inv.Status, inv.ID)
	return err
}

func (s *inventoryService) DeleteInventory(db *sql.DB, id int) error {
	query := "DELETE FROM inventories WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}
