package main

import (
	"flag"
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
	"github.com/sirupsen/logrus"
)

var (
	commitHash = ""
	file       = flag.String("file", "", "The path to the configuration file to execute")
)

func main() {
	flag.Parse()

	if *file == "" {
		logrus.Panic(fmt.Errorf("file cannot be null"))
	}

	cfg, err := loader.GetLoader(*file).Load()
	if err != nil {
		logrus.Panic(err)
	}

	exec := executor.NewExecutor(cfg)
	err = exec.ExecuteJobs()
	if err != nil {
		logrus.Panic(err)
	}
}
