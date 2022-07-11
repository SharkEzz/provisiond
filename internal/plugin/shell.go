package plugin

import "github.com/SharkEzz/provisiond/pkg/deployment"

// The Shell plugin take one string as shell command and execute it.
type Shell struct{}

func (s *Shell) Execute(ctx *deployment.JobContext, data any) (string, error) {
	return ctx.ExecuteCommand(data.(string))
}
