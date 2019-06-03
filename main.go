package main

import (
	"context"
	"log"

	transaction "github.com/akililab/transaction/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
)

// Server : Gateway struct
type Server struct{}

var (
	ts transaction.TransactionService
)

// GetTransactions : GetTransaction struct
func (s *Server) GetTransactions(c echo.Context) error {

	accountid := c.Param("accountid")

	response, err := ts.GetTransactions(context.TODO(), &transaction.TransactionRequest{AccountId: accountid})

	if err != nil {
		c.JSON(500, err)
	}

	return c.JSON(200, response)

}

func main() {

	// Create service
	service := web.NewService(
		web.Name("go.micro.api.transactions"),
	)

	service.Init()

	// setup Greeter Server Client
	ts = transaction.NewTransactionService("go.micro.srv.transactions", client.DefaultClient)

	// Create RESTful handler (using Gin)

	server := new(Server)
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/transactions/:accountid", server.GetTransactions)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
