package utils

import "encoding/json"

func ConvertToType[T any](data any) (*T, error) {
	jsonStr, err := json.Marshal(data.(map[string]any))
	if err != nil {
		return nil, err
	}

	var newType T

	if err := json.Unmarshal(jsonStr, &newType); err != nil {
		return nil, err
	}

	return &newType, nil
}
