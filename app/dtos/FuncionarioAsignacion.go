package dtos

import "github.com/google/uuid"

type FuncionarioAsignacion struct {
	ID                   uuid.UUID `json:"id"`
	Anho                 string    `json:"anho"`
	Mes                  string    `json:"mes"`
	Nivel                string    `json:"nivel"`
	DescripcionNivel     string    `json:"descripcion_nivel"`
	Entidad              string    `json:"entidad"`
	DescripcionEntidad   string    `json:"descripcion_entidad"`
	Oee                  string    `json:"oee"`
	DescripcionOee       string    `json:"descripcion_oee"`
	Documento            string    `json:"documento"`
	Nombres              string    `json:"nombres"`
	Apellidos            string    `json:"apellidos"`
	Funcion              string    `json:"funcion"`
	Estado               string    `json:"estado"`
	CargaHoraria         string    `json:"carga_horaria"`
	AnhoIngreso          string    `json:"anho_ingreso"`
	Sexo                 string    `json:"sexo"`
	Discapacidad         string    `json:"discapacidad"`
	TipoDiscapacidad     string    `json:"tipo_discapacidad"`
	FuenteFinanciamiento string    `json:"fuente_financiamiento"`
	ObjetoGasto          string    `json:"objeto_gasto"`
	Concepto             string    `json:"concepto"`
	Linea                string    `json:"linea"`
	Categoria            string    `json:"categoria"`
	Cargo                string    `json:"cargo"`
	Presupuestado        string    `json:"presupuestado"`
	Devengado            string    `json:"devengado"`
	Movimiento           string    `json:"movimiento"`
	Lugar                string    `json:"lugar"`
	FechaNacimiento      string    `json:"fecha_nacimiento"`
	FecUltModif          string    `json:"fec_ult_modif"`
	URI                  string    `json:"uri"`
	FechaActo            string    `json:"fecha_acto"`
	Correo               string    `json:"correo"`
	Profesion            string    `json:"profesion"`
	MotivoMovimiento     string    `json:"motivo_movimiento"`
}
