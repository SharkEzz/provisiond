package context

import (
	"github.com/SharkEzz/provisiond/internal/remote"
)

// A JobContext is passed to each executed jobs
//
// TODO: use a context.Context
type JobContext struct {
	JobName string
	client  *remote.Client
	logger  func(data string)
}

func (p *JobContext) ExecuteCommand(command string) (string, error) {
	return p.client.ExecuteCommand(command)
}

func (p *JobContext) Log(data string) {
	p.logger(data)
}

func NewJobContext(jobName string, client *remote.Client, logger func(data string)) *JobContext {
	return &JobContext{
		jobName,
		client,
		logger,
	}
}
