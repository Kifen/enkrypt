package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type EnkryptServer struct {
	port           int
	key            string
	ef             string
	encryptedFiles []string
	nonce          int
	router         *mux.Router
}

type EncryptedFiles struct {
	Files []string `json:"files"`
}

type Response struct {
	Message string         `json:"message"`
	Data    EncryptedFiles `json:"data"`
}

func (e *EnkryptServer) UpdateEncryptedFolder(f string) {
	e.ef = f
}

func (e *EnkryptServer) UpdateNOnce(n int) {
	e.nonce = n
}

func NewServer(key string, source string, target string) *EnkryptServer {
	return &EnkryptServer{
		port:           5000,
		key:            key,
		nonce:          0,
		encryptedFiles: make([]string, 0),
		router:         mux.NewRouter().StrictSlash(true),
	}
}

func (e *EnkryptServer) ListEncryptedFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)

	var files = &EncryptedFiles{}
	if e.nonce != 1 {
		err := DecryptFolder(e.ef, e.key)
		e.UpdateNOnce(1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if MetaFile == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{Message: "No encryped file(s)"})
		return
	}

	file, err := os.Open(MetaFile.Name())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files.Files = append(files.Files, scanner.Text())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(files)
}

func (e *EnkryptServer) DownloadFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	isEncrypted := func(file string) bool {
		for _, v := range e.encryptedFiles {
			if file == v {
				return true
			}
		}
		return false
	}

	param := r.URL.Query().Get("file")

	b := isEncrypted(param)
	if !b {
		file, err := Download(param, e.key)
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		e.encryptedFiles = append(e.encryptedFiles, param)
	}

	http.ServeFile(w, r, param)
	return
}

func (e *EnkryptServer) Serve() {
	router := mux.NewRouter().StrictSlash(true)
	corsObj := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/listencryptedfiles", e.ListEncryptedFiles)
	router.HandleFunc("/downloadfile/", e.DownloadFile)

	log.Printf("Server serving on port %d", e.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", e.port), handlers.CORS(corsObj)(router)))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
