package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cha-casa-nova-backend/controllers"
	"cha-casa-nova-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Senha hardcoded para simplificar
		const adminPassword = "admin123"

		// Verifica se a senha está no header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Remove o prefixo "Bearer " se existir
		password := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			password = authHeader[7:]
		}

		if password != adminPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar ao banco de dados
	database.ConnectDatabase()

	// Criar router Gin
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Middleware para adicionar headers de resposta
	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	// Rota de saúde
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "API do Chá de Casa Nova funcionando!",
			"time":    time.Now(),
		})
	})

	// Grupo de rotas v1
	v1 := r.Group("/api/v1")
	{
		// Rotas públicas
		v1.GET("/items", controllers.GetItems)
		v1.GET("/items/:id", controllers.GetItem)
		v1.GET("/stats", controllers.GetStats)
		v1.GET("/evento", controllers.GetEvento)
		v1.POST("/mensagens", controllers.CreateMensagem)
		v1.GET("/mensagens", controllers.GetMensagensAprovadas)
		v1.POST("/confirmacoes", controllers.CreateConfirmacao)
		v1.POST("/items/:id/resgate", controllers.ResgateItem)
		v1.POST("/items/:id/cancela-resgate", controllers.CancelaResgate)

		// Rotas protegidas (administrativas)
		admin := v1.Group("/admin")
		admin.Use(authMiddleware())
		{
			// Itens
			admin.GET("/items", controllers.GetAdminItems)
			admin.POST("/items", controllers.CreateItem)
			admin.PUT("/items/:id", controllers.UpdateItem)
			admin.DELETE("/items/:id", controllers.DeleteItem)

			// Mensagens
			admin.GET("/mensagens", controllers.GetMensagens)
			admin.DELETE("/mensagens/:id", controllers.DeleteMensagem)
			admin.POST("/mensagens/:id/aprovar", controllers.AprovarMensagem)

			// Confirmações
			admin.GET("/confirmacoes", controllers.GetConfirmacoes)
			admin.PUT("/confirmacoes/:id", controllers.UpdateConfirmacao)
			admin.DELETE("/confirmacoes/:id", controllers.DeleteConfirmacao)

			// Estatísticas
			admin.GET("/stats/detalhadas", controllers.GetEstatisticasDetalhadas)

			// Evento
			admin.PUT("/evento", controllers.UpdateEvento)
		}
	}

	// Middleware para log de requisições
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Middleware para recuperação de panic
	r.Use(gin.Recovery())

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Falha ao iniciar servidor:", err)
	}
}
