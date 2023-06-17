package request

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}
