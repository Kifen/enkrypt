package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
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
	nonce  int
	router *mux.Router
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

func NewServer(port int, key string, source string, target string) *EnkryptServer {
	return &EnkryptServer{
		port:   port,
		key:    key,
		srcdir: source,
		dstdir: target,
		nonce:  0,
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	f := vars["file"]
	file, err := DownloadFile(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("File %s downloaded.", file)
	w.WriteHeader(http.StatusOK)
}

func (e *EnkryptServer) DownloadFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	param := r.URL.Query().Get("file")
	//w.Header().Add("Content-Disposition", "Attachment")
	//w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	//w.Header().Set("Content-Length", w.Header.("Content-Length"))

	file, err := Download(param, e.key)
	if err != nil {
		log.Println("FAILS HERE: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	log.Println("FIles is downloaded: ", file.Name())
	http.ServeFile(w, r, param)

	log.Println("SUCCESSFULLY DOWNLOADED FILE...")
	return
}

func (e *EnkryptServer) Serve() {
	router := mux.NewRouter().StrictSlash(true)
	corsObj := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/listencryptedfiles", e.ListEncryptedFiles)
	router.HandleFunc("/downloadfile/", e.DownloadFile)

	log.Printf("Server serving on port %d", e.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", e.port),  handlers.CORS(corsObj)(router)))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

