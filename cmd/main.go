package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OkaniYoshiii/lesfraternels/internal/routes"
)

const LogsDir = "./logs"
const LogsFile = LogsDir + "/dev.log"

var address = flag.String("address", ":8000", "ip address the server will listen to (example: 127.0.0.1:8000)")

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	if _, err := os.Stat(LogsDir); os.IsNotExist(err) {
		if err := os.Mkdir(LogsDir, 0750); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.OpenFile(LogsFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.Default().Flags())

	fileServer := http.FileServer(http.Dir("./website/dist"))

	mux.Handle("GET /{$}", new(routes.HomeHandler))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer))
	mux.Handle("GET /connexion", new(routes.LoginHandler))
	mux.Handle("POST /connexion", new(routes.LoginHandler))

	server := http.Server{
		Addr:              *address,
		Handler:           mux,
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 10,
		MaxHeaderBytes:    1 << 20, // 1MB
		ErrorLog:          logger,
	}

	log.Fatal(server.ListenAndServe())
}
