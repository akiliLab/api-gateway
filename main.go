package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nicholasjackson/env"
	"google.golang.org/grpc"

	balance "github.com/akililab/api-gateway/proto/balance"
)

var (
	// the go.micro.srv.balance address
	endpoint = env.String("BALANCE_ADDRESS", false, "127.0.0.1:8005", "go.micro.srv.balance address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := balance.RegisterBalanceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
