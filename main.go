package main

import (
	"log"
	"net/http"

	"github.com/ubunifupay/db"

	"fmt"

	balance "github.com/ubunifupay/balance/pb"
	transaction "github.com/ubunifupay/transaction/pb"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

// TransactionJSON : Transaction
type TransactionJSON struct {
	ID          int64  `json:"id"`
	AccountID   string `json:"accountid"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Notes       string `json:"notes"`
}

var transactionClient transaction.TransactionServiceClient
var balanceClient balance.BalanceServiceClient

func getBalance(c *gin.Context) {

}

func credit(c *gin.Context) {

}

func debit(c *gin.Context) {

}

func storeTransaction(c *gin.Context) {
	transactionJSON := TransactionJSON{}
	req := &transaction.TransactionRequest{}
	err := c.BindJSON(&transactionJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not parse transaction data.")
	}
	storedMethod := "stored"
	if transactionJSON.ID != 0 {
		storedMethod = "updated"
	}
	req.ID = transactionJSON.ID
	req.AccountID = transactionJSON.AccountID
	req.Amount = transactionJSON.Amount
	req.Currency = transactionJSON.Currency
	req.Description = transactionJSON.Description
	req.Notes = transactionJSON.Notes
	req.CreatedAt = ptypes.TimestampNow()
	res, err := transactionClient.StoreTransaction(c, req)
	if err == nil && res.Completed {
		c.JSON(http.StatusOK, fmt.Sprintf("Transaction correctly %s.", storedMethod))
	} else {
		c.JSON(http.StatusInternalServerError, "Could not process the transaction. "+err.Error())
	}
}

func main() {
	// Database initialization
	session := db.CassandraSession
	defer session.Close()

	// Routes initialization
	r := gin.Default()
	r.POST("/createtransaction", storeTransaction)
	r.POST("/updatetransaction", storeTransaction)
	r.GET("/balance", getBalance)
	r.POST("/credit", credit)
	r.POST("/debit", debit)

	// Setup dial with transaction service
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	transactionClient = transaction.NewTransactionServiceClient(conn)

	// Setup dial with balance service
	conn, err = grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	balanceClient = balance.NewBalanceServiceClient(conn)

	// Listening
	r.Run(":8080")
}
