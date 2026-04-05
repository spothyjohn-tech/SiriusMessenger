package main

import (
	"log"
	"strconv"


	"MesEdge/internal/repository"
	"MesEdge/internal/service/user"  // переименуй папку user в service/user
	"MesEdge/pkg/database"
	"MesEdge/internal/transport/ws"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB("messanger.db")
	if err != nil {
		log.Fatal(err)
	}
	
	// Инициализация репозиториев и сервисов
	userRepo := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	
	hub := ws.NewHub()
	r := gin.Default()
	
	// Регистрация с использованием сервиса
	r.POST("/register", func(c *gin.Context) {
		var req user.RegistrationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Bad Data"})
			return
		}
		
		if err := userService.Register(req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	r.GET("/ws", func(c *gin.Context) {
		userIDstr := c.Query("id")
		userID, _ := strconv.ParseUint(userIDstr, 10, 32)
		
		if userID == 0 {
			c.JSON(401, gin.H{"error": "Wrong user"})
			return
		}
		hub.WSHandler(c.Writer, c.Request, uint(userID))
	})
	
	r.Run(":8080")
}