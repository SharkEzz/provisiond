package v1

import (
	"io"
	"net/http"

	utils "github.com/SharkEzz/provisiond/pkg/api"
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
)

func (a *API) HandlePostDeploy(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body := req.Body
	data, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(map[string]any{
			"error": err,
		}, w)
		return
	}

	cfg, err := loader.GetLoader(string(data)).Load()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ReturnJson(map[string]any{
			"error": err,
		}, w)
		return
	}

	exec, err := executor.NewExecutor(cfg, a.Config, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(map[string]any{
			"error": err,
		}, w)
		return
	}

	go func() {
		exec.ExecuteJobs()
	}()

	utils.ReturnJson(map[string]any{
		"status": "started",
		"id":     exec.UUID,
	}, w)
}
