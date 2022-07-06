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
	host     string
	port     uint16
	password string
	handlers *v1.API
	config   *executor.Config
}

func NewAPI(host string, port uint16, password string, config *executor.Config) *API {
	v1 := &v1.API{
		Config: config,
	}

	return &API{
		host,
		port,
		password,
		v1,
		config,
	}
}

func (a *API) checkPassword(next http.Handler) http.Handler {
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
	handlers := v1.API{}

	healthcheckHandler := http.HandlerFunc(handlers.HandleGetHealthcheck)
	deployHandler := http.HandlerFunc(handlers.HandlePostDeploy)

	http.Handle("/v1/healthcheck", healthcheckHandler)
	http.Handle("/v1/deploy", a.checkPassword(deployHandler))

	logging.LogOut("Started API server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", a.host, a.port), nil))
}
