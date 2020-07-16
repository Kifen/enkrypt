package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	doOnce   sync.Once
	MetaFile *os.File
	Done = make(chan struct{})
)

func CopyDir(source, target string) error {
	srcInfo, err := ValidatePath(source)
	if err != nil  && os.IsNotExist(err) {
		return err
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("Source %s is not a directory", source)
	}

	log.Printf("Source %s is valid", source)

	src := filepath.Clean(source)
	dst := filepath.Clean(target)

	doOnce.Do(func() {
		go WatchDir(source, target)
	})

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	fds, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, fd := range fds {
		time.Sleep(2 * time.Second)
		sourcePath := filepath.Join(src, fd.Name())
		targetPath := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			err = CopyDir(sourcePath, targetPath)
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(sourcePath, targetPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(source, target string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	in, err := os.Open(source)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer  out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	err = os.Chmod(target, srcInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

// ValidatePath checks if directory exists
func ValidatePath(path string) (os.FileInfo, error) {
	log.Printf("Validating %s.", path)
	f, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return nil, err
	}

	return f, nil
}

func WatchDir(source, target string){
	var cOnce sync.Once

	token := func(path string) string {
		t := strings.Split(path, "/")
		return t[len(t)-1]
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add(source)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Watching directory: %s", source)

	go func() {
		for {
			select {
			case event, ok := <- watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					dst := filepath.Join(target, token(event.Name))
					go func() {
						err := CopyFile(event.Name, dst)
						if err != nil {
							log.Fatal(err)
						}
						err = EncryptFile(dst, "key")
						if err != nil {
							log.Fatalf("Failed to encrypt file: %v", err)
						}
						log.Printf("Encrypted file <%s>.", dst)
						cOnce.Do(func() {
							MetaFile, err = os.Create(filepath.Join(target, ".meta.txt"))
							if err != nil {
								log.Fatalf("Failed creating file: %s", err)
							}
						})
						_, err = fmt.Fprintln(MetaFile, dst)
						if err != nil {
							log.Fatalf("failed writing to file: %s", err)
						}
					}()
				}
			case err, ok := <- watcher.Errors:
				if !ok {
					return
				}
				log.Fatal(err)
			}
		}
	}()
	<-Done
	log.Println("Closing watcher...")
}
