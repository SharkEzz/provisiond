package internal

import "context"

// The Shell plugin take one or more strings as shell commands and execute them.
type Shell struct{}

func (s *Shell) Execute(ctx context.Context, data any) error {
	return nil
}
