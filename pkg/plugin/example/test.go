package main

import (
	"github.com/SharkEzz/provisiond/pkg/deployment"
)

type Test struct{}

func (t *Test) Execute(data any, ctx *deployment.JobContext) (string, error) {
	return "test", nil
}

func GetPlugin() (p any) {
	return &Test{}
}
