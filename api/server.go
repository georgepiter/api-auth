package api

import (
	"api-auth/api/models"
	"api-auth/api/routes"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	db := models.Connect()
	err := db.AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}()

	listen(8080)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Escutando na porta %s...\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}
