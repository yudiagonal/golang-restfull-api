package controller

import (
	"Go-crud-api/helper"
	"Go-crud-api/log/request"
	"Go-crud-api/log/response"
	"Go-crud-api/v0/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// interface for product controller
type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByIdProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAllProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

// product for route in main
func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// struct for category product impl
type ProductControllerImpl struct {
	ProductService service.ProductService
}

// DeleteProduct implements ProductController.
func (controller *ProductControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindAllProduct implements ProductController.
func (controller *ProductControllerImpl) FindAllProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindByIdProduct implements ProductController.
func (controller *ProductControllerImpl) FindByIdProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// UpdateProduct implements ProductController.
func (controller *ProductControllerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *ProductControllerImpl) CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	productCreateRequest := request.ProductCreateRequest{}
	err := decoder.Decode(&productCreateRequest)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.CreateProduct(r.Context(), productCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
