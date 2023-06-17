package response

// struct for category response
type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// struct for standar web response
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
