package model

// DefaultPayload godoc
// @Summary Default payload structure
// @Description Wraps the API response data
// @Model DefaultPayload
type DefaultPayload[T any] struct {
	Data T `json:"data"`
}

type PagingPayload struct {
	Paging `json:"paging"`
	Data   interface{} `json:"data"`
}
