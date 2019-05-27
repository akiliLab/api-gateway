package main

import (
	"context"
	"log"

	transaction "github.com/akililab/transaction/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
)

// Server : Gateway struct
type Server struct{}

var (
	ts transaction.TransactionService
)

// GetTransactions : GetTransaction struct
func (s *Server) GetTransactions(c *gin.Context) {

	log.Print("Received Server.GetTransactions API request")

	accountid := c.Param("accountid")

	response, err := ts.GetTransactions(context.TODO(), &transaction.TransactionRequest{AccountId: accountid})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)

}

func main() {

	// Create service
	service := web.NewService(
		web.Name("micro.akililab.api.gateway"),
	)

	service.Init()

	// setup Greeter Server Client
	ts = transaction.NewTransactionService("micro.akililab.transaction", client.DefaultClient)

	// Create RESTful handler (using Gin)

	server := new(Server)
	router := gin.Default()
	router.GET("/transactions/:accountid", server.GetTransactions)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
