package context

import (
	"github.com/SharkEzz/provisiond/pkg/remote"
)

type PluginContext struct {
	client *remote.Client
}

func (p *PluginContext) ExecuteCommand(command string) error {
	return p.client.ExecuteCommand(command)
}

func NewPluginContext(client *remote.Client) *PluginContext {
	return &PluginContext{
		client,
	}
}
