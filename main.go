package main

import (
	"github.com/Kifen/enkrypt/pkg/util"
	"log"
	"time"
)

const (
	Source = "/home/kifen/Desktop/Source"
	Target = "/home/kifen/Desktop/Target"
)

func main()  {
	dstInfo, err := util.ValidatePath(Target)
	if err != nil {
		log.Fatal(err)
	}

	if !dstInfo.IsDir() {
		log.Fatalf("Target %s is not a directory", Target)
	}

	log.Printf("Target %s is valid", Target)

	err = util.CopyDir(Source, Target)
	if err != nil {
		log.Fatal(err)
	}

	util.Done <- struct{}{}
	log.Println("Files copied...")

	err = util.EncryptFolder(Target, "key")
	if err != nil {
		log.Fatalf("failed to encrypt folder: %v", err)
	}

	log.Println("Successfully encrypted folder...")

	time.Sleep(10 *time.Second)

	err = util.DecryptFolder("/home/kifen/Desktop/Target.tar.gz","key")
	if err != nil {
		log.Fatalf("failed to decrypt folder: %v", err)
	}

	log.Println("Successfully decrypted folder...")

	srv := util.NewServer(8080)
	srv.serve()
}

