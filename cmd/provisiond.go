package main

import (
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
)

var (
	commitHash = ""
)

func main() {
	cfg, err := loader.GetLoader("./test.yaml").Load()
	if err != nil {
		panic(err)
	}

	exec := executor.NewExecutor(cfg)
	err = exec.ExecuteJobs()
	if err != nil {
		panic(err)
	}
}
