package internal

import "github.com/SharkEzz/provisiond/pkg/context"

// The Shell plugin take one string as shell command and execute it.
type Shell struct{}

func (s *Shell) Execute(data any, ctx *context.JobContext) (string, error) {
	return ctx.ExecuteCommand(data.(string))
}
