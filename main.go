package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MGT06/shortlink-app_backend/internal/config"
	"github.com/MGT06/shortlink-app_backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			ShortLink APP
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		localhost:5080
//	@BasePath	/

//	@securityDefinitions.apikey		BearerToken
//	@in								header
//	@name							Authorization
//	@description					Bearer token used as identitiy for access resource

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	psqlConn := config.NewPsqlDb(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))

	pool, err := psqlConn.Connect()
	if err != nil {
		log.Fatalf("cannot connect to db\nreason: %s", err.Error())
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("db not response\nreason: %s", err.Error())
	}

	log.Printf(" PONG ! Connected PostgreSQL successfully")

	redis := config.NewRedisClient(os.Getenv("RDBUSER"), os.Getenv("RDBPASS"), os.Getenv("RDBHOST"), os.Getenv("RDBPORT"))

	rc := redis.ConnectRedis()
	if err := rc.Ping(context.Background()).Err(); err != nil {
		log.Printf("Close Redis : %s\n\n", err.Error())
	}

	log.Println("PONG ! Conected redis successfully")
	
	r := gin.Default()
	
	routes.MainRouter(r, pool, rc)
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("APPHOST"), os.Getenv("APPPORT")))
}
