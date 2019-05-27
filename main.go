package main

import (
	transaction "github.com/akililab/transaction/proto"
)

// Server : Gateway struct
type Server struct{}

var (
	cl transaction.TransactionService
)

// TransactionRequest : Transaction Request
type TransactionRequest struct {
	AccountID string `json:"account_id"`
}

func main() {
	
}
