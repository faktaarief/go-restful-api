package helper

import (
	"faktaarief/go-restful-api/model/domain"
	"faktaarief/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
