package context

import (
	"github.com/SharkEzz/provisiond/pkg/remote"
	"github.com/sirupsen/logrus"
)

type JobContext struct {
	jobName string
	client  *remote.Client
	logger  *logrus.Logger
}

func (p *JobContext) ExecuteCommand(command string) error {
	return p.client.ExecuteCommand(command)
}

func (p *JobContext) Log(level logrus.Level, data any) {
	p.logger.Log(level, data)
}

func NewPluginContext(jobName string, client *remote.Client, logger *logrus.Logger) *JobContext {
	return &JobContext{
		jobName,
		client,
		logger,
	}
}
