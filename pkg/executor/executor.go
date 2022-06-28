package executor

import (
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"github.com/SharkEzz/provisiond/pkg/plugin"
	"github.com/SharkEzz/provisiond/pkg/remote"
	"github.com/google/uuid"
)

type Executor struct {
	Deployment    *deployment.Deployment
	UUID          string
	outputChannel chan map[string]string
}

func NewExecutor(dployment *deployment.Deployment, outputChannel chan map[string]string) *Executor {
	uuid := uuid.NewString()

	return &Executor{
		dployment,
		uuid,
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

	for _, job := range e.Deployment.Jobs {
		jobName := job["name"].(string)
		jobHosts := job["hosts"].([]any)
		for _, host := range jobHosts {
			client, ok := clients[host.(string)]
			if !ok {
				return fmt.Errorf("host '%s' does not exist", host)
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
		if key == "hosts" || key == "name" {
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
		logData := map[string]string{
			"log":  logging.Log(data),
			"uuid": e.UUID,
		}
		select {
		case e.outputChannel <- logData:
			break
		default:
			break
		}
	}
}
