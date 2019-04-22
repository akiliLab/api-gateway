package main

import (
	"log"
	"net/http"

	"fmt"

	balance "github.com/ubunifupay/balance/pb"
	mastercard "github.com/ubunifupay/mastercard/pb"
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

// SearchMerchant type
type SearchMerchant struct {
	MerchantID string `json:"merchantid"`
	Search     int64  `json:"search"`
}

var transactionClient transaction.TransactionServiceClient
var balanceClient balance.BalanceServiceClient
var mastercardClient mastercard.MastercardServiceClient

func getBalance(c *gin.Context) {

}

func credit(c *gin.Context) {

}

func debit(c *gin.Context) {

}

func queryMerchants(c *gin.Context) {
	searchJSON := SearchMerchant{}
	req := &mastercard.MastercardRequest{}
	err := c.BindJSON(&searchJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not parse Mastercard data.")
	}
	req.MerchantID = searchJSON.MerchantID
	req.Search = searchJSON.Search
	res, err := mastercardClient.GetMerchantIdentifiers(c, req)

	if err == nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, "Could not process the Mastercard. "+err.Error())
	}

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

	// Routes initialization
	r := gin.Default()
	r.POST("/createtransaction", storeTransaction)
	r.POST("/updatetransaction", storeTransaction)
	r.GET("/balance", getBalance)
	r.POST("/credit", credit)
	r.POST("/debit", debit)
	r.POST("/queryMerchants", queryMerchants)

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

	// Setup dial with mastercard service
	conn, err = grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	mastercardClient = mastercard.NewMastercardServiceClient(conn)

	// Listening
	r.Run(":8080")
}
