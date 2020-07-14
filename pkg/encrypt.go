package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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