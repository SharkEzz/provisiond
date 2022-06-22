package executor

import (
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/plugin"
	"github.com/SharkEzz/provisiond/pkg/remote"
	"github.com/SharkEzz/provisiond/pkg/types"
	"github.com/sirupsen/logrus"
)

type Executor struct {
	Deployment *types.Deployment
}

func NewExecutor(deployment *types.Deployment) *Executor {
	return &Executor{
		deployment,
	}
}

func (e *Executor) ExecuteJobs() error {
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
			ctx := context.NewPluginContext(client, logrus.New())

			logrus.Info(fmt.Sprintf("executing job '%s' on host '%s'", name, host))

			err := e.ExecuteJob(name, job, ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Executor) ExecuteJob(jobName string, job map[string]any, ctx *context.PluginContext) error {
	for key, value := range job {
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
