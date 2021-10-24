package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck Contém à lógica para os endpoint de verificação de integridade,
// retorna se está saudavel em uma solicitação bem sucedida.
func (a App) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
