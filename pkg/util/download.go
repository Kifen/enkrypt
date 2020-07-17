package util

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"os"
	"strings"
)

var OP int

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress prints the progress of a file write
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(file string) (string, error) {
	in, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer in.Close()

	token := strings.Split(file, "/")
	fileName := token[len(token)-1]
	//newFile := filepath.Join(downloadPath, fileName)

	out, err := os.Create(file + ".tmp")
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(in, counter))
	if err != nil {
		return "", err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	// Rename the tmp file back to the original file
	err = os.Rename(file+".tmp", file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func Download(file, key string) (*os.File, error) {
	err := DecryptFile(file, key)
	if err != nil {
		return nil, err
	}

	log.Printf("File decrypted: <%s>", file)

	in, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return in, nil
}
