package executor

import (
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"github.com/SharkEzz/provisiond/pkg/plugin"
	"github.com/SharkEzz/provisiond/pkg/remote"
)

type Executor struct {
	Deployment    *deployment.Deployment
	outputChannel chan string
}

func NewExecutor(dployment *deployment.Deployment, outputChannel chan string) *Executor {
	return &Executor{
		dployment,
		outputChannel,
	}
}

func (e *Executor) ExecuteJobs() error {
	e.Log(fmt.Sprintf("Starting execution of deployment '%s'", e.Deployment.Name))

	clients := map[string]*remote.Client{}

	for name, config := range e.Deployment.Config.SSH {
		client, err := remote.ConnectToHost(name, config.Host, config.Port, config.Type, config.Username, config.Password, config.KeyFile)
		if err != nil {
			return err
		}
		clients[name] = client
	}
	defer remote.CloseAllClients(clients)

	for name, job := range e.Deployment.Jobs {
		jobHosts := job["hosts"].([]any)
		for _, host := range jobHosts {
			client, ok := clients[host.(string)]
			if !ok {
				return fmt.Errorf("host '%s' does not exist", host)
			}
			ctx := context.NewPluginContext(name, client, e.Log)

			e.Log(fmt.Sprintf("Executing job '%s' on host '%s'", name, host))

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
		if key == "hosts" {
			continue
		}

		plg, exist := plugin.Plugins[key]
		if !exist {
			return fmt.Errorf("error: plugin '%s' does not exist", key)
		}

		plg.Execute(value.(string), ctx)
	}

	return nil
}

func (e *Executor) Log(data string) {
	logging.LogOut(data)

	if e.outputChannel != nil {
		select {
		case e.outputChannel <- logging.Log(data):
			break
		default:
			break
		}
	}
}
