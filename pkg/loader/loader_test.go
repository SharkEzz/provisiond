package loader_test

import (
	"testing"

	loaders "github.com/SharkEzz/provisiond/internal/loader"
	"github.com/SharkEzz/provisiond/pkg/loader"
)

func TestLoader(t *testing.T) {
	ld := loader.GetLoader(`name: Test
jobs:
  - name: Test1
  hosts: [localhost]
  shell: echo oui > test.txt`)

	if _, ok := ld.(loaders.StringLoader); !ok {
		t.Errorf("expected StringLoader")
	}

	ld = loader.GetLoader("name: Test\n")

	if _, ok := ld.(loaders.StringLoader); !ok {
		t.Errorf("expected StringLoader")
	}

	ld = loader.GetLoader("./test.yaml")
	if _, ok := ld.(loaders.FileLoader); !ok {
		t.Errorf("expected FileLoader")
	}

	ld = loader.GetLoader("should be a StringLoader")
	if _, ok := ld.(loaders.StringLoader); ok {
		t.Errorf("expected StringLoader")
	}
}
