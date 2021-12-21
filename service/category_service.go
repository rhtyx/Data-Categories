package service

import (
	"Data-Category/exception"
	"Data-Category/helper"
	"Data-Category/model/domain"
	"Data-Category/model/web"
	"Data-Category/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	DeleteAll(ctx context.Context)
	UpdateById(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	DeleteById(ctx context.Context, categoryId int)
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (cs *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := cs.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = cs.CategoryRepository.Save(ctx, tx, category)

	return (web.CategoryResponse)(category)
}

func (cs *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := cs.CategoryRepository.FindAll(ctx, tx)

	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, (web.CategoryResponse)(category))
	}
	return categoriesResponse
}

func (cs *CategoryServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cs.CategoryRepository.DeleteAll(ctx, tx)
}

func (cs *CategoryServiceImpl) UpdateById(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := cs.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := cs.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = cs.CategoryRepository.UpdateById(ctx, tx, category)

	return (web.CategoryResponse)(category)
}

func (cs *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := cs.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return (web.CategoryResponse)(category)
}

func (cs *CategoryServiceImpl) DeleteById(ctx context.Context, categoryId int) {
	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := cs.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cs.CategoryRepository.DeleteById(ctx, tx, category)
}
