package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type EnkryptServer struct {
	port   int
	key    string
	srcdir string
	dstdir string
	ef     string
	nonce int
	router *mux.Router
}

type EncryptedFiles struct {
	Files []string `json:"files"`
}

type Response struct {
	Message string `json:"message"`
	Data EncryptedFiles `json:"data"`
}

func (e *EnkryptServer) UpdateEncryptedFolder(f string) {
	e.ef = f
}

func (e *EnkryptServer) UpdateNOnce(n int) {
	e.nonce = n
}

func NewServer(port int, key string, source string, target string) *EnkryptServer {
	return &EnkryptServer{
		port:   port,
		key:    key,
		srcdir: source,
		dstdir: target,
		nonce: 0,
		router: mux.NewRouter().StrictSlash(true),
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

func (e *EnkryptServer) GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["file"]
	file, err := DownloadFile(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return	}

	log.Printf("File %s downloaded.", file)
	w.WriteHeader(http.StatusOK)
}

func (e *EnkryptServer) Serve() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/listencryptedfiles", e.ListEncryptedFiles)
	router.HandleFunc("/downloadfile", e.GetFile)

	log.Printf("Server serving on port %d", e.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", e.port), router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}