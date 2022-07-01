package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Encode a map of map[string]any to JSON
func jsonEncode(data map[string]any) (string, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(res), err
}

func ReturnJson(data map[string]any, w http.ResponseWriter) error {
	str, err := jsonEncode(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, str)

	return nil
}
