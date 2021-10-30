package models

//Error representa e armazena informações para quaisquer erros que tenhamos como parte
//de nossas regras de negócios.
type Error struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"` //Ingnora a propriedade quando retornamos a struct como um objeto JSON.
	Validation []string `json:"validation,omitempty"`
}

//BindError erros comuns que podem acontecer em rotas
func BindError() *Error {
	return &Error{Code: 400, Message: "Error processing request.", Name: "BIND_ERROR"}
}

//ValidationError
func ValidationError(errors []string) *Error {
	return &Error{Code: 400, Name: "VALIDATION", Message: "A validation error occurred.", Validation: errors}
}
