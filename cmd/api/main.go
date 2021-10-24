package main

import "goRestApi/internal/api"

func main() {
	application := api.New()
	application.Start()
}
