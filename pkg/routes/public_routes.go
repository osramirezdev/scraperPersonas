package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osramirezdev/scraperPersonas/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET Persona:
	route.Get("/persona/:cedula", controllers.PersonaController.ObtenerDatosPersona)

	// Routes for Scraper method:
	route.Post("/scraper/obtener/ruc", controllers.ScrapRuc)
	route.Post("/scraper/obtener/ips", controllers.ScrapIps)
	route.Post("/scraper/obtener/funcionarios", controllers.ScrapFuncionarios)
}
