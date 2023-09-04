package dtos

// Response dto: respuesta a consultas.
type ResponseDto[T any] struct {
	Error bool   `json:"error" validate:"required" default:"false"`
	Msg   string `json:"msg" validate:"required"`
	Count int    `json:"count"`
	Data  []T    `json:"data" `
}

type ResponseDto404 struct {
	Error bool   `json:"error" validate:"required" default:"true"`
	Msg   string `json:"msg" validate:"required" default:"Not found"`
	Count int    `json:"count" default:"0"`
	Data  []int  `json:"data"`
}

type ResponseDto500 struct {
	Error bool   `json:"error" validate:"required" default:"true"`
	Msg   string `json:"msg" validate:"required" default:"Internal server error"`
	Count int    `json:"count" default:"0"`
	Data  []int  `json:"data"`
}
