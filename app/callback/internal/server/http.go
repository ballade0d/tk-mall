package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"mall/api/mall/service/v1"
	"net/http"
)

func NewHTTPServer() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	err := v1.RegisterCallbackServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Println("HTTP server is running on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
