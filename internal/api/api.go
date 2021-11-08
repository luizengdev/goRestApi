package api

import (
	"GoRestApi/pkg/config"
	"GoRestApi/pkg/data"
	"GoRestApi/pkg/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// Estrutura que conterá infomações partilhavel para serem utilizadas em outras funções
type App struct {
	server  *echo.Echo
	userSvc services.IUserService
	cfg     *config.Settings
}

// New inicializa o servidor Echo e qualquer lógica adicional como middleware global e dependencias
func New(cfg *config.Settings, client *mongo.Client) *App {
	// Echo instance
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, userProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		cfg:     cfg,
	}
}

// ConfigureRoutes cria e configura os endpoint
func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
	a.server.POST("/v1/public/account/login", a.Login)

	protected := a.server.Group("v1/api")

	middleware := Middleware{config: a.cfg}

	protected.Use(middleware.Auth)
	protected.GET("/secret", func(c echo.Context) error {
		userId := c.Get("user").(string)
		return c.String(200, userId)
	})
}

// Start configura à rota e depois starta o servidor na porta.
func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
