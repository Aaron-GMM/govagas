package router

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/Aaron-GMM/govagas/handler/openingHandler"
	"github.com/Aaron-GMM/govagas/repository"
	"github.com/Aaron-GMM/govagas/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	// 1. Pegar a instância do banco já iniciada no main
	db := config.GetDB()

	// 2. Montar o Repository
	repo := repository.NewOpeningRepository(db)

	// 3. Montar o Service
	svc := service.NewOpeningService(repo)

	// 4. Montar o Handler com o Service
	opHandler := openingHandler.NewOpeningHandler(svc)

	// 5. Inicializar o Gin
	router := gin.Default()

	// 6. Configurar as rotas passando o Handler montado
	initializeRouter(router, opHandler)

	// Executar a aplicação
	router.Run(":8080")
}
