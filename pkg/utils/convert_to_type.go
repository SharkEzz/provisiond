package utils

import "encoding/json"

// ConvertToType is a generic function, it take any data then convert them to a pointer of type T.
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
