package main

import (
	"cat-api/breed/gateway"
	breedRepo "cat-api/breed/repository"
	breedHandler "cat-api/breed/server/http"
	breedService "cat-api/breed/service"
	userRepo "cat-api/user/repository"
	userHandler "cat-api/user/server/http"
	userService "cat-api/user/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

const (
	driver         = "mysql"
	dataSourceName = "root:root@tcp(172.28.1.1)/cats?parseTime=true"
	secret = "@#$RF@!718"
)

func main() {
	fmt.Println("Starting API...")
	fmt.Println("Connecting to Database...")
	var db *sql.DB

	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		db, err = sql.Open(driver, dataSourceName)
		if err != nil {
			fmt.Println("Error Connecting to Database: ", err)
			os.Exit(1)
		}
		return
	})

	breedRepository, err := breedRepo.NewRepository(db)

	if err != nil {
		fmt.Println("Could not connect with the database")
		panic(err)
	}

	userRepository, err := userRepo.NewRepository(db)

	if err != nil {
		fmt.Println("Could not connect with the database")
		panic(err)
	}

	fmt.Println("Database Connected!")

	fmt.Println("Opening the Gateway...")
	gate := gateway.NewGateway()
	fmt.Println("Gateway Opened!")

	fmt.Println("Starting the Services...")
	breedServ := breedService.NewService(breedRepository, gate)
	userServ := userService.NewService(userRepository, []byte(secret))
	fmt.Println("Service Started!")

	fmt.Println("Starting the Router...")
	r := gin.New()
	userHandler.NewHandler(userServ, r)
	breedHandler.NewHandler(breedServ, r, []byte(secret))
	fmt.Println("Router Started!")
	port := ":8081"
	fmt.Println(fmt.Sprintf("API Started! Listening on Port: %s", port))
	log.Fatal(r.Run(port))
}
