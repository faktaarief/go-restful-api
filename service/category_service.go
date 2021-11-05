package service

import (
	"context"
	"faktaarief/go-restful-api/model/web"
)

type CategoryService interface {
	Create(cx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(cx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(cx context.Context, categoryId int)
	FindById(cx context.Context, categoryId int) web.CategoryResponse
	FindAll(cx context.Context) []web.CategoryResponse
}
