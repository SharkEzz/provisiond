package v1

import (
	"fmt"
	"net/http"
	"os"

	utils "github.com/SharkEzz/provisiond/pkg/utils"
)

func (a *API) HandleGetDeploymentStatus(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := req.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.ReturnJson(map[string]any{
			"error": "id is required",
		}, w)
		return
	}

	file, err := os.ReadFile(fmt.Sprintf("logs/deployments/%s.log", id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(map[string]any{
			"error": "failed to read log file",
		}, w)
		return
	}

	utils.ReturnJson(map[string]any{
		"content": string(file),
	}, w)
}
