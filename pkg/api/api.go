package api

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/SharkEzz/provisiond/pkg/api/handlers/v1"
	"github.com/SharkEzz/provisiond/pkg/logging"
)

func StartAPI(host string, port uint16, password string) {
	handlers := v1.API{}

	http.HandleFunc("/v1/healthcheck", handlers.HandleGetHealthcheck)
	http.HandleFunc("/v1/deploy", handlers.HandlePostDeploy)

	logging.LogOut("Started API server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
