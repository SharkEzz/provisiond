package api

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/SharkEzz/provisiond/pkg/api/handlers/v1"
	"github.com/SharkEzz/provisiond/pkg/logging"
)

var apiPassword string

func checkPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("password") != apiPassword {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid password")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func StartAPI(host string, port uint16, password string) {
	apiPassword = password
	handlers := v1.API{}

	healthcheckHandler := http.HandlerFunc(handlers.HandleGetHealthcheck)
	deployHandler := http.HandlerFunc(handlers.HandlePostDeploy)

	http.Handle("/v1/healthcheck", checkPassword(healthcheckHandler))
	http.Handle("/v1/deploy", checkPassword(deployHandler))

	logging.LogOut("Started API server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
