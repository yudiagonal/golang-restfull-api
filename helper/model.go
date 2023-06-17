package helper

import (
	"Go-crud-api/log/response"
	modelCategory "Go-crud-api/v0/model"
)

func ToCategoryResponse(category modelCategory.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
