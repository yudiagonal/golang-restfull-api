package response

import (
	"encoding/json"
	"net/http"
)

// struct for category response
type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductResponse struct {
	Id          uint64 `json:"id"`
	NamaBarang  string `json:"nama_barang"`
	Harga       string `json:"harga"`
	Jenis       string `json:"jenis"`
	MetaKeyword string `json:"meta_keyword"`
}

// struct for standar web response
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseWriter struct {
	http.ResponseWriter
	C, S    int
	B       []byte
	ErrCode string
	Err     error
}

func (r *ResponseWriter) WriteHeader(code int) {
	if r.Code() == 0 {
		r.C = code
		r.ResponseWriter.WriteHeader(code)
	}
}

func (r *ResponseWriter) Write(body []byte) (int, error) {
	r.B = body
	if r.Code() == 0 {
		r.WriteHeader(http.StatusOK)
	}

	return r.ResponseWriter.Write(body)
}

func (r *ResponseWriter) Code() int {
	return r.C
}

func (r *ResponseWriter) Size() int {
	return r.S
}

func (r *ResponseWriter) Body() interface{} {
	var body interface{}
	json.Unmarshal(r.B, &body)

	return body
}

func (r *ResponseWriter) ErrorCode() string {
	return r.ErrCode
}

func (r *ResponseWriter) Error() error {
	return r.Err
}

func New(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
	}
}
