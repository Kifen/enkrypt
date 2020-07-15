package main

import (
	"context"
	pb "github.com/Kifen/enkrypt/pkg/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	log.Printf("Client on...")
	client := pb.NewEnkryptClient(conn)
	files, err := client.ListEncryptedFiles(context.Background(), &pb.E{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Encrypted files: %s", files)

	f, err := client.GetFile(context.Background(), &pb.Info{FilePath: "/home/kifen/Desktop/source/source.txt", DownloadPath: "/home/kifen/Desktop/"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File downloaded: %s", f.File)
}
