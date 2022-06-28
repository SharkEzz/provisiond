package context

import (
	"github.com/SharkEzz/provisiond/pkg/remote"
)

type JobContext struct {
	jobName string
	client  *remote.Client
	logger  func(data string)
}

func (p *JobContext) ExecuteCommand(command string) error {
	return p.client.ExecuteCommand(command)
}

func (p *JobContext) Log(data string) {
	p.logger(data)
}

func NewPluginContext(jobName string, client *remote.Client, logger func(data string)) *JobContext {
	return &JobContext{
		jobName,
		client,
		logger,
	}
}