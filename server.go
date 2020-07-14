package main

import (
	"context"
	"fmt"
	"github.com/Kifen/enkrypt/pkg"
	pb "github.com/Kifen/enkrypt/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

type EnkryptServer struct {
	port int
}

func (e *EnkryptServer) ListEncryptedFiles(ctx context.Context, dir *pb.Directory) (*pb.EncryptedFiles, error) {
	f, err := pkg.ValidatePath(dir.Path)
	if err != nil {
		return nil, err
	}

	if !f.IsDir() {
		return nil, fmt.Errorf("Path %s is not a directory.", dir.Path)
	}

	return nil, nil
}

func (e *EnkryptServer) serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEnkryptServer(grpcServer, &EnkryptServer{})
	grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed serving server: %v", err)
	}

	return nil
}
