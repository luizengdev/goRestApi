package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Estrutura que conterá infomações partilhavel para serem utilizadas em outras funções
type App struct {
	server *echo.Echo
}

// New inicializa o servidor Echo e qualquer lógica adicional como middleware global e dependencias
func New() *App {
	// Echo instance
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())

	return &App{
		server: server,
	}
}

// ConfigureRoutes cria e configura os endpoint
func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
}

// Start configura à rota e depois starta o servidor na porta.
func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
