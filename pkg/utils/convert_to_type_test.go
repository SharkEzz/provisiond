package utils_test

import (
	"testing"

	"github.com/SharkEzz/provisiond/pkg/utils"
)

func TestConvertToType(t *testing.T) {
	var testData = map[string]any{
		"name": "John",
		"age":  30,
	}

	type Test struct {
		Name string
		Age  int
	}

	newType, err := utils.ConvertToType[Test](testData)
	if err != nil {
		t.Error(err)
	}

	if newType.Name != "John" || newType.Age != 30 {
		t.Errorf("Expected %v, got %v", testData, newType)
	}
}
