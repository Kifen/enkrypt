package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateSymLink(src, dst string) error {
	err := os.Symlink(src, dst)
	if err != nil {
		return fmt.Errorf("Failed in creating symlink: %v", err)
	}

	return nil
}

func ResolveSymlink(file string) (string, error) {
	var o string
	fileInfo, err := os.Lstat(file)
	if err != nil {
		return "", err
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		originFile, err := os.Readlink(file)
		if err != nil {
			return "", fmt.Errorf("Failed to resolve symlink: %v", err)
		}
		o = originFile
	}

	return o, nil
}

func EncryptFile(file, key string) error {
	overwriteFile := func(f string) (*os.File, error) {
		file, err := os.Create(f)
		if err != nil {
			return nil, err
		}

		return file, err
	}

	block, err := aes.NewCipher([]byte(createHash(key)))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return  fmt.Errorf("Failed to read file: %v", err)
	}

	f, err := overwriteFile(file)
	if err != nil {
		return fmt.Errorf("Failed tp overwrite file: %s", err)
	}
	defer f.Close()

	_, err = f.Write(gcm.Seal(nonce, nonce, data, nil))
	if err != nil {
		return fmt.Errorf("Failed to write encrypted data to file: %v", err)
	}

	return nil
}


func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

type WriteCounter struct {
	Total uint64
}

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

func DownloadFile(file, downloadPath string) (string, error) {
	in, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer in.Close()

	token := strings.Split(file, "/")
	fileName := token[len(token)-1]
	newFile := filepath.Join(downloadPath, fileName)

	out, err := os.Create(newFile+".tmp")
	if err != nil {
		return "", err
	}
	defer  out.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(in, counter))
	if err != nil {
		return "", err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	// Rename the tmp file back to the original file
	err = os.Rename(newFile+".tmp", newFile)
	if err != nil {
		return "", err
	}

	return fileName, nil
}