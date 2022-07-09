package context

import (
	"context"

	"github.com/SharkEzz/provisiond/internal/remote"
)

// A JobContext is passed to each executed jobs
type JobContext struct {
	context.Context
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

func NewJobContext(jobName string, client *remote.Client, logger func(data string)) (*JobContext, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	return &JobContext{
		ctx,
		jobName,
		client,
		logger,
	}, cancel
}
