package v1

import (
	"net/http"

	"github.com/SharkEzz/provisiond/pkg/utils"
)

func (a *API) HandleGetHealthcheck(w http.ResponseWriter, req *http.Request) {
	utils.ReturnJson(map[string]any{
		"result": "ok",
	}, w)
}
