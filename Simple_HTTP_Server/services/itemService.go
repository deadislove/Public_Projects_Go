package services

import (
	"Simple_HTTP_Server/database"
	"Simple_HTTP_Server/models"
	"errors"
)

// Create an item
func CreateItem(item models.ItemDTO) error {
	_, err := database.DB.Exec("INSERT INTO items (id, name) VALUES (?, ?)", item.ID, item.Name)
	return err
}

// Read all items
func GetAllItems() ([]models.Item, error) {
	rows, err := database.DB.Query("SELECT id, name FROM items")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Update an item
func UpdateItem(item models.ItemDTO) error {
	result, err := database.DB.Exec("UPDATE items SET name = ? WHERE id = ?", item.Name, item.ID)

	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

// Delete an item
func DeleteItem(id string) error {
	result, err := database.DB.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}
