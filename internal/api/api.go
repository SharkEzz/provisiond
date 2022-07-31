package api

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/SharkEzz/provisiond/internal/api/handlers/v1"
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/logging"
)

type API struct {
	host        string
	port        uint16
	password    string
	v1_handlers *v1.API
}

// Create a new API instance.
func NewAPI(host string, port uint16, password string, config *executor.Config) *API {
	v1_handlers := &v1.API{
		Config: config,
	}

	return &API{
		host,
		port,
		password,
		v1_handlers,
	}
}

// Middleware to check if the password for accessing the API is provided.
func (a *API) checkPasswordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("password") != a.password {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid password")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *API) StartAPI() {
	healthcheckHandler := http.HandlerFunc(a.v1_handlers.HandleGetHealthcheck)
	deployHandler := http.HandlerFunc(a.v1_handlers.HandlePostDeploy)
	deployStatusHandler := http.HandlerFunc(a.v1_handlers.HandleGetDeploymentStatus)

	http.Handle("/v1/healthcheck", healthcheckHandler)
	http.Handle("/v1/deploy/log", a.checkPasswordMiddleware(deployStatusHandler))
	http.Handle("/v1/deploy", a.checkPasswordMiddleware(deployHandler))

	logging.LogOut("Started API server", logging.SUCCESS)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", a.host, a.port), nil))
}
