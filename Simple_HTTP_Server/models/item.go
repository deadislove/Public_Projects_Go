package models

// Item represents the item structure in the database
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ItemDTO is used for transferring data between services and HTTP handlers
type ItemDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
