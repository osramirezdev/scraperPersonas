package dtos

type FuncionarioDto struct {
	Search []Search `json:"search" validate:"required,dive"`
}

type Search struct {
	Documento string `json:"documento" validate:"required" example:"900999"`
	Tipo      string `json:"tipo" example:"ruc"`
}
