package helper

import (
	"Go-crud-api/log/response"
	"Go-crud-api/v0/model"
)

func ToCategoryResponse(category model.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:          product.ID,
		NamaBarang:  product.NamaBarang,
		Harga:       product.Harga,
		Jenis:       product.Jenis,
		MetaKeyword: product.MetaKeyword,
	}
}
