package main

import (
	"context"
	"fmt"
	"log"
	"tecnotree-go-project/config"
	"tecnotree-go-project/controllers"
	"tecnotree-go-project/routes"
	"tecnotree-go-project/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	server      *gin.Engine
)

func main() {
	server = gin.Default()
	InitializeDatabase()
	InitializeUsers()
	InitializeProducts()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)

	server.Run(":4800")
}

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to the database : %v", err)
	} else {
		fmt.Println("Connected to the Database")
	}
}

func InitializeUsers() {
	usersCollection := config.GetCollection(mongoClient, "Dheeraj", "users")
	userSvc := services.InitUserService(usersCollection)
	userCtrl := controllers.InitUserController(userSvc)
	routes.UserRoutes(server, *userCtrl)
}

func InitializeProducts() {
	productsCollection := config.GetCollection(mongoClient, "Dheeraj", "products")
	productSvc := services.InitProductService(productsCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)

}
