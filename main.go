package main

import (
	"golang/handler"
	"golang/message"
	"golang/room"
	"golang/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-course?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{})

	db.AutoMigrate(&room.Room{})

	db.AutoMigrate(&message.Message{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	roomRepository := room.NewRepository(db)
	roomService := room.NewService(roomRepository)
	roomHandler := handler.NewRoomHandler(roomService)

	messageRepository := message.NewRepository(db)
	messageService := message.NewService(messageRepository)
	messageHandler := handler.NewMessageHandler(messageService)

	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/user", userHandler.StoreUser)
	api.GET("/user", userHandler.FetchAllUser)
	api.GET("/user/:id", userHandler.FetchUserById)
	api.PUT("/user/:id", userHandler.UpdateUser)
	//api.GET("/user/:id", messageHandler.GetUserByRoomID)
	api.DELETE("/user/:id", userHandler.DeleteUser)
	api.POST("/room", roomHandler.StoreRoom)
	api.GET("/room", roomHandler.FetchAllRoom)
	api.GET("/room/:id", roomHandler.FetchRoomById)
	api.PUT("/room/:id", roomHandler.UpdateRoom)
	api.DELETE("/room/:id", roomHandler.DeleteRoom)

	api.POST("/message", messageHandler.StoreMessage)
	api.GET("/message", messageHandler.FetchAllMessage)
	api.GET("/message/:id", messageHandler.FetchMessageById)
	api.PUT("/message/:id", messageHandler.UpdateMessage)
	api.DELETE("/message/:id", messageHandler.DeleteMessage)
	// api.GET("/message/:id", messageHandler.GetMessagesByUserID)
	// api.GET("/message/:id", messageHandler.GetMessagesByRoomID)

	r.Run(":3000")
}
