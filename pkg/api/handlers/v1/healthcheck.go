package v1

import (
	"net/http"

	"github.com/SharkEzz/provisiond/pkg/api/utils"
)

func (a *API) HandleGetHealthcheck(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	utils.ReturnJson(map[string]any{
		"result": "ok",
	}, w)
}
