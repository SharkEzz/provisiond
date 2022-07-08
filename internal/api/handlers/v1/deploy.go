package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")

	running := true
	ended := make(chan bool, 1)
	logChannel := make(chan string)

	exec := executor.NewExecutor(cfg, a.Config, logChannel)

	go func() {
		exec.ExecuteJobs()
		ended <- true
	}()

	for running {
		select {
		case log := <-logChannel:
			var buf bytes.Buffer
			json.NewEncoder(&buf).Encode(log)
			fmt.Fprintf(w, "data: %s\n", buf.String())
		case <-ended:
			running = false
		}

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}
