package deployment

import (
	"context"

	"github.com/SharkEzz/provisiond/internal/remote"
)

// A JobContext is passed to each executed jobs.
// It it provide useful functions for executing commands, logging output, and more!
type JobContext struct {
	context.Context
	JobName string
	client  *remote.Client
	logger  func(data string)
}

// Execute a command on the remote host.
func (p *JobContext) ExecuteCommand(command string) (string, error) {
	return p.client.ExecuteCommand(command)
}

// Log a message to the logger, the output will be logged to the deployment log.
func (p *JobContext) Log(data string) {
	p.logger(data)
}

// Create a new JobContext.
func NewJobContext(jobName string, client *remote.Client, logger func(data string)) (*JobContext, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	return &JobContext{
		ctx,
		jobName,
		client,
		logger,
	}, cancel
}
