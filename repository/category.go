package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz3-rizqyep/database"
	"quiz3-rizqyep/entities"
	"time"
)

type CategoryRepository interface {
	GetAll() (error, []entities.Category)
	Insert(category entities.Category) error
	GetById(id int) (error, entities.Category)
	Update(id int, category entities.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: database.GetDBConnection()}
}

func (repository *categoryRepository) GetAll() (err error, results []entities.Category) {
	sql := "SELECT * FROM categories"

	rows, err := repository.db.Query(sql)

	if err != nil {
		return err, []entities.Category{}
	}

	defer rows.Close()

	for rows.Next() {

		var category = entities.Category{}
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		fmt.Println(rows)
		if err != nil {
			return err, []entities.Category{}
		}

		results = append(results, category)
	}

	return nil, results
}

func (repository *categoryRepository) GetById(id int) (err error, result entities.Category) {
	sqlStatement := `SELECT * FROM categories WHERE id=$1;`

	row := repository.db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return errors.New(fmt.Sprintf("no record found for id %v", id)), entities.Category{}
	case nil:
		return nil, result
	default:
		return err, entities.Category{}
	}
}

func (repository *categoryRepository) Insert(category entities.Category) (err error) {
	sql := "INSERT INTO categories (name) VALUES ($1)"

	errs := repository.db.QueryRow(sql, category.Name)

	return errs.Err()
}

func (repository *categoryRepository) Update(id int, category entities.Category) (err error) {
	sql := "UPDATE categories SET name = $1, updated_at = $2 WHERE id=$3"

	errs := repository.db.QueryRow(sql, category.Name, time.Now(), id)

	return errs.Err()
}

func (repository *categoryRepository) Delete(id int) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"
	errs := repository.db.QueryRow(sql, id)

	return errs.Err()
}
