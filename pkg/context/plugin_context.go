package context

import (
	"github.com/SharkEzz/provisiond/pkg/remote"
	"github.com/sirupsen/logrus"
)

type PluginContext struct {
	client *remote.Client
	logger *logrus.Logger
}

func (p *PluginContext) ExecuteCommand(command string) error {
	return p.client.ExecuteCommand(command)
}

func (p *PluginContext) Log(level logrus.Level, data any) {
	p.logger.Log(level, data)
}

func NewPluginContext(client *remote.Client, logger *logrus.Logger) *PluginContext {
	return &PluginContext{
		client,
		logger,
	}
}
