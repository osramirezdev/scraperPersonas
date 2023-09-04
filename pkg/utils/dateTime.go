package utils

import (
	"fmt"
	"time"
)

func ObtenerHora(format string) string {
	location, err := time.LoadLocation("America/Asuncion")
	if err != nil {
		fmt.Println("Error al cargar la ubicación:", err)
		return ""
	}

	currentTime := time.Now().In(location)
	formattedTime := currentTime.Format(format)
	fmt.Println("Hora en America/Asuncion:", formattedTime)
	return formattedTime
}
