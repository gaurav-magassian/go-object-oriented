package models

// Order ...
type Order struct {
	ID    int
	Items []Item
}

// Item ...
type Item struct {
	Name  string
	Price float32
}
