package main


import (
	routes "github.com/jayaramsivaramannair/joys_of_GO/Project_JWT-Authentication/routes"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port =="" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)


	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access granted for api-1"})
	})


	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access granted for api-2"})
	})


	router.Run(":" + port)
}