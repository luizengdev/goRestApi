package models

import (
	"GoRestApi/pkg/domain"

	"github.com/labstack/echo/v4"
)

//RegisterRequest representa os tipos de solicitação que API aceita, registro.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//LoginRequest representa os tipos de solicitação que API aceita, login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//ValidateRegisterRequest verificará o comprimento do nome de usuário e da senha registrada,
//Se a validação falhar, retornamos a função ValidationError da pasta Erros.
func ValidateRegisterRequest(c echo.Context) (*domain.User, *Error) {
	RegisterRequest := new(RegisterRequest)
	if err := c.Bind(RegisterRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(RegisterRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characterers")
	}

	if len(RegisterRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characterers")
	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: RegisterRequest.Username,
		Password: RegisterRequest.Password,
	}, nil
}

//ValidateLoginRequest verificará o comprimento do nome de usuário e da senha logada,
//Se a validação falhar, retornamos a função ValidationError da pasta Erros.
func ValidateLoginRequest(c echo.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(loginRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}, nil

}
