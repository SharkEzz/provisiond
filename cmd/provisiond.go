package main

import (
	"flag"
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/api"
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
	"github.com/sirupsen/logrus"
)

var (
	commitHash  = ""
	file        = flag.String("file", "", "The path to the configuration file to execute")
	enableAPI   = flag.Bool("api", false, "Set to true to enable the integrated REST API")
	apiPassword = flag.String("apiPassword", "", "The REST API password")
)

func main() {
	flag.Parse()

	if *enableAPI {
		if *apiPassword == "" {
			logrus.Panic("apiPassword cannot be blank")
		}
		api.NewAPI("0.0.0.0", 7655, *apiPassword).Start()
		return
	}

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
