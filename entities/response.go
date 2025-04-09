package entities

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Errors []error     `json:"errors,omitempty"`
}

type ResponsePaging struct {
	Total uint32 `json:"total"`
}
