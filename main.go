package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(""))
}

func convertToPDF(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	file, header, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	log.Printf("Converting file %s", header.Filename)

	if _, err := io.Copy(&buf, file); err != nil {
		panic(err)
	}

	parentDir := os.TempDir()
	dirName := uuid.New().String()
	tmpDir, err := ioutil.TempDir(parentDir, dirName)

	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s/%s", tmpDir, uuid.New().String())
	if ioutil.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", fileName, "--outdir", tmpDir)

	if stdout, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		log.Println(string(stdout))
	}

	pdfFile, err := ioutil.ReadFile(fmt.Sprintf("%s.pdf", fileName))

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/pdf")
	w.Write(pdfFile)
}

func main() {
	httpPort, found := os.LookupEnv("PORT")

	if !found {
		httpPort = "3000"
	}

	httpAddress := "0.0.0.0:" + httpPort

	router := mux.NewRouter()
	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/", convertToPDF).Methods("POST")

	log.Printf("Starting HTTP server at address %s\n", httpAddress)

	if err := http.ListenAndServe(httpAddress, router); err != nil {
		log.Fatalf("Cannot start http server, cause: %+v", err)
	}
}
