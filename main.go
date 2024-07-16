package main

import (
	"auth_git/databases/mongodb"
	"auth_git/databases/postgresql"
	"auth_git/servers"
	"auth_git/utilts/addtionalQueryAndEncryptDecrypt"

	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := postgresql.Postgresql()
	defer db.Close()
	// # check how many tables available
	addtionalQueryAndEncryptDecrypt.CountTables(db)
	// # connect postgresql #
	conn := mongodb.MongoDB()
	defer conn.Client().Disconnect(context.Background())
	// # connect mongoDB #
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	servers.SetupRoutesAuth(router, db)

	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
