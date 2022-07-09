package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/SharkEzz/provisiond/internal/api"
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"gopkg.in/yaml.v3"
)

var (
	buildTime   = time.Now()
	version     = "dev"
	file        = flag.String("file", "", "The path to the configuration file to execute")
	enableAPI   = flag.Bool("api", false, "Set to true to enable the integrated REST API")
	apiPassword = flag.String("apiPassword", "", "The REST API password")
	apiPort     = flag.Uint("apiPort", 7655, "The port to listen on for the REST API")
)

func main() {
	logging.LogOut(fmt.Sprintf("provisiond %s - Compiled on %s\n", version, buildTime.Format("2006-01-02 15:04:05")))

	flag.Parse()

	var config *executor.Config

	if _, err := os.Stat("./config.yaml"); err == nil {
		content, err := os.ReadFile("./config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			panic(err)
		}
		logging.LogOut("Loaded config from ./config.yaml")
	}

	if *enableAPI {
		if *apiPassword == "" {
			panic(fmt.Errorf("apiPassword is required when enabling the API"))
		}
		api.NewAPI("0.0.0.0", uint16(*apiPort), *apiPassword, config).StartAPI()
		return
	}

	if *file == "" {
		panic(fmt.Errorf("file cannot be null"))
	}

	deployment, err := loader.GetLoader(*file).Load()
	if err != nil {
		panic(err)
	}

	exec, err := executor.NewExecutor(deployment, config, nil)
	if err != nil {
		panic(err)
	}

	err = exec.ExecuteJobs()
	if err != nil {
		panic(err)
	}
}
