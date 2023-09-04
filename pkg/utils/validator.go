package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}
	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

// establecemos las columnas requeridas
type CSVValidator struct {
	RequiredColumns []string
}

// retornamos el validador css con las columnas instanciadas
func NewCSVValidator(columns []string) *CSVValidator {
	return &CSVValidator{RequiredColumns: columns}
}

// metodo validate
func (v *CSVValidator) Validate(file multipart.File) error {
	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return err
	}
	for _, col := range v.RequiredColumns {
		found := false
		for _, h := range header {
			if h == col {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("column '%s' is missing", col)
		}
	}
	return nil
}
