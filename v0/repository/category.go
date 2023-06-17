package repository

import (
	"Go-crud-api/helper"
	modelCategory "Go-crud-api/v0/model"
	"context"
	"database/sql"
	"errors"
)

// interface for category repository implementation
type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, model modelCategory.Category) modelCategory.Category
	Update(ctx context.Context, tx *sql.Tx, model modelCategory.Category) modelCategory.Category
	Delete(ctx context.Context, tx *sql.Tx, model modelCategory.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (modelCategory.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []modelCategory.Category
}

// category repository implementation
type CategoryRepositoryImpl struct {
}

// new category repository
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category modelCategory.Category) modelCategory.Category {
	SQL := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category modelCategory.Category) modelCategory.Category {
	SQL := "UPDATE category SET name=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category modelCategory.Category) {
	SQL := "DELETE FROM category WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (modelCategory.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := modelCategory.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil

	} else {
		return category, errors.New("category is not found")

	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []modelCategory.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []modelCategory.Category
	for rows.Next() {
		category := modelCategory.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories
}
