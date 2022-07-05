package executor

import (
	"fmt"

	"github.com/SharkEzz/provisiond/internal/remote"
	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"github.com/SharkEzz/provisiond/pkg/plugin"
)

type Executor struct {
	Deployment *deployment.Deployment
}

func NewExecutor(dployment *deployment.Deployment) *Executor {
	return &Executor{
		dployment,
	}
}

func (e *Executor) ExecuteJobs() error {
	e.Log(fmt.Sprintf("Starting execution of deployment '%s'", e.Deployment.Name))

	clients := map[string]*remote.Client{}

	for name, config := range e.Deployment.Config.SSH {
		client, err := remote.ConnectToHost(name, config.Host, config.Port, config.Type, config.Username, config.Password, config.KeyFile, config.KeyPass, e.Deployment.Variables)
		if err != nil {
			return err
		}
		clients[name] = client
	}
	defer remote.CloseAllClients(clients)

	for _, job := range e.Deployment.Jobs {
		jobName := job["name"].(string)
		jobHosts := job["hosts"].([]any)
		for _, host := range jobHosts {
			var client *remote.Client

			if host == "localhost" {
				client = remote.ConnectToLocalhost()
			} else {
				c, ok := clients[host.(string)]
				if !ok {
					return fmt.Errorf("host '%s' does not exist", host)
				}
				client = c
			}

			ctx := context.NewPluginContext(jobName, client, e.Log)

			e.Log(fmt.Sprintf("Executing job '%s' on host '%s'", jobName, host))

			err := e.ExecuteJob(job, ctx)
			if err != nil {
				return err
			}
		}
	}

	e.Log(fmt.Sprintf("Deployment '%s' done, executed %d jobs", e.Deployment.Name, len(e.Deployment.Jobs)))
	return nil
}

func (e *Executor) ExecuteJob(job map[string]any, ctx *context.JobContext) error {
	for key, value := range job {
		// Skip keys that are not plugins
		if key == "hosts" ||
			key == "name" ||
			key == "allow_failure" {
			continue
		}

		plg, exist := plugin.Plugins[key]
		if !exist {
			return fmt.Errorf("error: plugin '%s' does not exist", key)
		}

		output, err := plg.Execute(value, ctx)
		if err != nil {
			if allowedToFail, ok := job["allow_failure"].(bool); ok && allowedToFail {
				e.Log(fmt.Sprintf("Job %s failed but failure allowed: %s", job["name"], err))
			} else {
				return err
			}
		}
		if output != "" {
			e.Log(fmt.Sprintf("Job output: %s", output))
		}
	}

	return nil
}

func (e *Executor) Log(data string) {
	// TODO: re-enable file logging
	// logStr := logging.Log(data)

	// os.Mkdir(".deployments", 0750)
	// filePath := fmt.Sprintf(".deployments/%s.txt", e.UUID)

	// file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	// if err == nil {
	// 	defer file.Close()
	// }

	// if !e.logFileCreated {
	// 	fmt.Fprintln(file, e.Deployment.Name)
	// 	e.logFileCreated = true
	// }
	// fmt.Fprintln(file, logStr)

	logging.LogOut(data)
}
