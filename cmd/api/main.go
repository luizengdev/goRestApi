package main

import "goRestApi/internal/api"

// Criando nova instância api, e chamando a função start para configurar
// as rotas e executar o aplicativo.
func main() {
	application := api.New()
	application.Start()
}
