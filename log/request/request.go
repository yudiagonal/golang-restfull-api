package request

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type ProductCreateRequest struct {
	NamaBarang  string `validate:"required,min=1,max=100" json:"nama_barang"`
	Harga       string `validate:"required" json:"harga"`
	Jenis       string `validate:"required" json:"jenis"`
	MetaKeyword string `validate:"required" json:"meta_keyword"`
}

type ProductUpdateRequest struct {
	Id          uint64 `validate:"required"`
	NamaBarang  string `validate:"required,max=200,min=1" json:"nama_barang"`
	Harga       string `validate:"required" json:"harga"`
	Jenis       string `validate:"required" json:"jenis"`
	MetaKeyword string `validate:"required" json:"meta_keyword"`
}
