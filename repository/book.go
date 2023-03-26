package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz3-rizqyep/database"
	"quiz3-rizqyep/entities"
	"time"
)

type BookRepository interface {
	Insert(book entities.Book) error
	GetAll() ([]entities.Book, error)
	GetById(id int) (entities.Book, error)
	Update(id int, book entities.Book) error
	Delete(id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		db: database.GetDBConnection(),
	}
}

func (repository *bookRepository) Insert(book entities.Book) error {
	sql := "INSERT INTO books (title, description, image_url, release_year, total_page, price, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	errs := repository.db.QueryRow(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.TotalPage, book.Price, book.Thickness, book.CategoryId)

	return errs.Err()
}

func (repository *bookRepository) GetAll() (results []entities.Book, err error) {
	sqlStatement := `SELECT books.* , categories.* FROM books JOIN categories ON categories.id = books.category_id ;`

	rows, err := repository.db.Query(sqlStatement)

	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {

		var book = entities.Book{}
		err = rows.Scan(
			&book.ID,
			&book.CategoryId,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Category.ID,
			&book.Category.Name,
			&book.Category.CreatedAt,
			&book.Category.UpdatedAt,
		)
		fmt.Println(rows)
		if err != nil {
			return []entities.Book{}, err
		}
		results = append(results, book)
	}

	return results, nil
}

func (repository *bookRepository) GetById(id int) (result entities.Book, err error) {
	sqlStatement := `SELECT books.* , categories.* FROM books JOIN categories ON categories.id = books.category_id  WHERE books.id = $1;`

	row := repository.db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&result.ID,
		&result.CategoryId,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.ReleaseYear,
		&result.Price,
		&result.TotalPage,
		&result.Thickness,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Category.ID,
		&result.Category.Name,
		&result.Category.CreatedAt,
		&result.Category.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return entities.Book{}, errors.New(fmt.Sprintf("no record found for id %v", id))
	case nil:
		return result, nil
	default:
		return entities.Book{}, nil
	}
}
func (repository *bookRepository) Update(id int, book entities.Book) (err error) {
	sql := "UPDATE books SET title = $1, description = $2, image_url = $3, release_year=$4, total_page=$5, price=$6, thickness=$7, category_id=$8, updated_at=$9 WHERE id=$10"

	errs := repository.db.QueryRow(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.TotalPage, book.Price, book.Thickness, book.CategoryId, time.Now(), id)

	return errs.Err()
}

func (repository *bookRepository) Delete(id int) (err error) {
	sql := "DELETE FROM books WHERE id = $1"
	errs := repository.db.QueryRow(sql, id)

	return errs.Err()
}
