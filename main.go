package main

import (
	"github.com/Kifen/enkrypt/pkg/util"
	"log"
)

const (
	source = "/home/kifen/Desktop/source"
	target = "/home/kifen/Desktop/target"
)

func main()  {
	dstInfo, err := util.ValidatePath(target)
	if err != nil {
		log.Fatal(err)
	}

	if !dstInfo.IsDir() {
		log.Fatalf("Target %s is not a directory", target)
	}

	log.Printf("Target %s is valid", target)

	err = util.CopyDir(source, target)
	if err != nil {
		log.Fatal(err)
	}

	util.Done <- struct{}{}
	log.Println("Files copied...")
	srv := NewServer(8080)
	srv.serve()
}
