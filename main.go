package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Print("...Server running")
	server := echo.New()

	server.GET("/status", Handler)

	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	data := time.Now().Format("02.01.2006")
	err := ctx.String(http.StatusOK, data)
	if err != nil {
		return err
	}
	return nil
}
