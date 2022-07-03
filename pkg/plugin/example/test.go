package main

import (
	"github.com/SharkEzz/provisiond/pkg/context"
)

type Test struct{}

func (t *Test) Execute(data any, ctx *context.JobContext) (string, error) {
	return "test", nil
}

func GetPlugin() (p any) {
	return &Test{}
}
