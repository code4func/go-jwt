package main 

import (
	"fmt"
	"net/http"
	"go-jwt/handler"
	"go-jwt/config"
	"go-jwt/driver"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)
	
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}