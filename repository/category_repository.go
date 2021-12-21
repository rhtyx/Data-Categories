package repository

import (
	"Data-Category/helper"
	"Data-Category/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	DeleteAll(ctx context.Context, tx *sql.Tx)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	UpdateById(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	DeleteById(ctx context.Context, tx *sql.Tx, category domain.Category)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	querySQL := "SELECT id, name FROM data_category"
	rows, err := tx.QueryContext(ctx, querySQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySQL := "INSERT INTO data_category(name) VALUES ($1) RETURNING id"
	var id int
	rows, err := tx.QueryContext(ctx, querySQL, category.Name)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		helper.PanicIfError(err)
	}
	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	querySQL := "DELETE FROM data_category"
	_, err := tx.ExecContext(ctx, querySQL)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	querySQL := "SELECT id, name FROM data_category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, querySQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var category domain.Category
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (c *CategoryRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySQL := "UPDATE data_category SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, querySQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (c *CategoryRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, category domain.Category) {
	querySQL := "DELETE FROM data_category WHERE id = $1"
	_, err := tx.ExecContext(ctx, querySQL, category.Id)
	helper.PanicIfError(err)
}
