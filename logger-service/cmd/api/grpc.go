package main

import (
	"context"
	"fmt"
	"github.com/charichu/logger-service/data"
	"github.com/charichu/logger-service/logs"
	"google.golang.org/grpc"
	"log"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return resp
	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	server := grpc.NewServer()
	logs.RegisterLogServiceServer(server, &LogServer{Models: app.Models})
	log.Printf("gRPC server started on port: %s", gRpcPort)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
