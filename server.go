package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/Kifen/enkrypt/pkg/proto"
	"github.com/Kifen/enkrypt/pkg/util"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type EnkryptServer struct {
	port int
}

func NewServer(port int) *EnkryptServer {
	return &EnkryptServer{
		port: port,
	}
}

func (e *EnkryptServer) ListEncryptedFiles(ctx context.Context, d *pb.Empty) (*pb.EncryptedFiles, error) {
	encryptedFiles := &pb.EncryptedFiles{
		Files: make([]string, 0),
	}

	file, err := os.Open(util.MetaFile.Name())
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		encryptedFiles.Files = append(encryptedFiles.Files, scanner.Text())
	}

	return encryptedFiles, nil
}

func (e *EnkryptServer) GetFile(ctx context.Context, f *pb.File) (*pb.File, error) {
	file, err := util.DownloadFile(f.File)
	if err != nil {
		return nil, err
	}

	log.Println("Download finished")
	return &pb.File{File: file}, nil
}

func (e *EnkryptServer) serve(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEnkryptServer(grpcServer, &EnkryptServer{})

	log.Printf("Grpc server listening on port %d", e.port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed serving server: %v", err)
	}
}
