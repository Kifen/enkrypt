package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func overwriteFile(f string, b []byte) error {
	file, err := os.Create(f)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func EncryptFile(file, key string) error {
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
		return fmt.Errorf("Failed to read file: %v", err)
	}

	b := gcm.Seal(nonce, nonce, data, nil)
	err = overwriteFile(file, b)
	if err != nil {
		return fmt.Errorf("Failed tp overwrite file: %s", err)
	}

	return nil
}

func DecryptFile(file string, passPhrase string) error {
	key := []byte(createHash(passPhrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("NA HEREEE")
		return err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return err
	}

	err = overwriteFile(file, plainText)
	if err != nil {
		return err
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

func ZipFolder(folder, zipfile string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("tar", "czf", zipfile, folder)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func UnzipFile(zipFile string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("tar", "xzf", zipFile)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func EncryptFolder(f, k string) (string, error) {
	done := make(chan struct{})
	errCh := make(chan error)

	t := strings.Split(f, "/")
	p := t[len(t)-1]
	h := t[:len(t)-1]
	home := strings.Join(h, "/")
	filename := fmt.Sprintf("%s.tar.gz", p)

	go func() {
		p = filepath.Clean(p)
		home = filepath.Clean(home)

		err := os.Chdir(home)
		if err != nil {
			errCh <- err
		}

		err, _, _ = ZipFolder(filepath.Join(p), filename)
		if err != nil {
			errCh <- err
		}

		err = EncryptFile(filepath.Join(home, filename), k)
		if err != nil {
			errCh <- err
		}

		done <- struct{}{}
	}()

	select {
	case err := <-errCh:
		return "", err
	case <-done:
		err, _, _ := deleteDir(f)
		if err != nil {
			return "", err
		}
		return filepath.Join(home, filename), nil
	}
}

func DecryptFolder(f, k string) error {
	t := strings.Split(f, "/")
	h := t[:len(t)-1]
	home := strings.Join(h, "/")
	home = filepath.Clean(home)

	err := DecryptFile(f, k)
	if err != nil {
		return err
	}

	err = os.Chdir(home)
	if err != nil {
		return err
	}

	err, _, _ = UnzipFile(f)
	if err != nil {
		return err
	}

	return nil
}

func deleteDir(dir string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("rm", "-rf", dir)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return err, stdout.String(), stderr.String()
}
