package main

import (
	"github.com/Kifen/enkrypt/pkg"
	"log"
)

const (
	source = "/home/kifen/Desktop/source"
	target = "/home/kifen/Desktop/target"
)

func main()  {
	dstInfo, err := pkg.ValidatePath(target)
	if err != nil {
		log.Fatal(err)
	}

	if !dstInfo.IsDir() {
		log.Fatalf("Target %s is not a directory", target)
	}

	log.Printf("Target %s is valid", target)

	err = pkg.CopyDir(source, target)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Finsished copying source dir <%s> to target <%s>.", source, target)
	}
}
