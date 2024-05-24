package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"my-reading-app/internal/handler"
	"my-reading-app/internal/repository"
	"my-reading-app/internal/service"
	"my-reading-app/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENVIRONMENT")

	var envPath string
	if env != "" {
		envPath = filepath.Join("/app", ".env."+env)
	} else {
		envPath = filepath.Join("/app", ".env")
	}

	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env %s file from path %s: %v", env, envPath, err)
	}

	router := gin.Default()

	// Configurar CORS dinamicamente
	config := cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	// Conectar ao MongoDB
	ctx := context.Background()
	client, err := db.ConnectMongo(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Dependências
	readingRepo := repository.NewMongoRepository(client)
	readingService := service.NewReadingService(readingRepo)
	readingHandler := handler.NewReadingHandler(readingService)

	userRepo := repository.NewMongoUserRepository(client)
	authService := service.NewAuthService(userRepo, "your_secret_key")
	authHandler := handler.NewAuthHandler(authService)

	// Rotas de autenticação
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Rotas de leitura
	router.GET("/readings/:day", readingHandler.GetReading)
	router.POST("/readings/:day/next", readingHandler.NextReading)
	router.POST("/readings/:day/previous", readingHandler.PreviousReading)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
