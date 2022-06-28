package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) HandleGetStream(w http.ResponseWriter, res *http.Request) {
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		data := <-a.OutputChannel

		jsonData, err := json.Marshal(data)
		if err != nil {
			// TODO: better handling
			return
		}

		fmt.Fprintf(w, `data: %s`, jsonData)
		fmt.Fprintln(w)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}
