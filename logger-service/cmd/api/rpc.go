package main

import (
	"context"
	"log"
	"time"

	"github.com/younesious/logger-service/log/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.Background(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error writing log in rpc.go, LogInfo", err)
		return err
	}

	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
