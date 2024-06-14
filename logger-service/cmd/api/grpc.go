package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/younesious/logger-service/log/data"
	"github.com/younesious/logger-service/log/logs"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	log.Println("Log:", input.Name, input.Data)

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed:"}
		return res, err
	}

	res := &logs.LogResponse{Result: "logged!"}

	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC server started at port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve for gRPC: %v", err)
	}
}
