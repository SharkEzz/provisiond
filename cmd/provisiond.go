package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/SharkEzz/provisiond/pkg/api"
	"github.com/SharkEzz/provisiond/pkg/executor"
	"github.com/SharkEzz/provisiond/pkg/loader"
	"github.com/SharkEzz/provisiond/pkg/logging"
)

var (
	buildTime   = time.Now()
	version     = "dev"
	file        = flag.String("file", "", "The path to the configuration file to execute")
	enableAPI   = flag.Bool("api", false, "Set to true to enable the integrated REST API")
	apiPassword = flag.String("apiPassword", "", "The REST API password")
)

func main() {
	logging.LogOut(fmt.Sprintf("provisiond %s - Compiled on %s\n", version, buildTime.Format("2006-01-02 15:04:05")))

	flag.Parse()

	if *enableAPI {
		if *apiPassword == "" {
			fmt.Println(logging.Log("apiPassword cannot be blank"))
		}
		api.StartAPI("0.0.0.0", 7655, *apiPassword)
		return
	}

	if *file == "" {
		panic(fmt.Errorf("file cannot be null"))
	}

	cfg, err := loader.GetLoader(*file).Load()
	if err != nil {
		panic(err)
	}

	err = executor.NewExecutor(cfg).ExecuteJobs()
	if err != nil {
		panic(err)
	}
}
