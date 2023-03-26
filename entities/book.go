package entities

import "time"

type Book struct {
	ID          int       `json:"id"`
	CategoryId  int       `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	TotalPage   int       `json:"total_page"`
	Price       string    `json:"price"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    Category  `json:"category"`
}
