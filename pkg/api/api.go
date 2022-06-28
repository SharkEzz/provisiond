package api

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/SharkEzz/provisiond/pkg/api/handlers/v1"
)

func StartAPI(host string, port uint16, password string) {
	outputChannel := make(chan map[string]string)

	handlers := v1.API{
		OutputChannel: outputChannel,
	}

	http.HandleFunc("/v1/healthcheck", handlers.HandleGetHealthcheck)
	http.HandleFunc("/v1/deploy", handlers.HandlePostDeploy)
	http.HandleFunc("/v1/stream", handlers.HandleGetStream)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
