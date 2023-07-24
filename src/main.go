// main.go
package main

import (
	"tasks_api/src/api"
	"tasks_api/src/config"
)

func main() {
	cfg := config.GetConfig() // Recupera a configuração do projeto

	api.InitServer(cfg) // Inicializa o servidor da API
}
