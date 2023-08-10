package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductDetail struct {
	NamaBarang  string `json:"nama_barang"`
	Harga       string `json:"harga"`
	Jenis       string `json:"jenis"`
	MetaKeyword string `json:"meta_keyword"`
	IsPublished bool   `json:"is_published"`
}

type Product struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	NamaBarang  string         `gorm:"column:nama_barang" json:"nama_barang"`
	Harga       string         `gorm:"column:harga" json:"harga"`
	Jenis       string         `gorm:"column:jenis" json:"jenis"`
	MetaKeyword string         `gorm:"column:meta_keyword" json:"meta_keyword"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index,column:deleted_at" json:"deleted_at"`
}

// ProductDetail is a struct for ProductDetail.
type ProductList struct {
	ID          uint64    `json:"id"`
	NamaBarang  string    `json:"nama_barang"`
	Harga       string    `json:"harga"`
	Jenis       string    `json:"jenis"`
	MetaKeyword string    `gorm:"column:meta_keyword" json:"meta_keyword"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DataProduct []Product

func (data DataProduct) ParseToCMSList() []ProductList {
	var result []ProductList
	for _, item := range data {
		resultItem := ProductList{
			ID:         item.ID,
			NamaBarang: item.NamaBarang,
			Harga:      item.Harga,
			Jenis:      item.Jenis,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		result = append(result, resultItem)
	}
	return result
}
