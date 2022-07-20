package executor

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/SharkEzz/provisiond/internal/remote"
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/logging"
	"github.com/SharkEzz/provisiond/pkg/plugin"
	"github.com/google/uuid"
)

type Executor struct {
	Deployment *deployment.Deployment
	Config     *Config
	logChannel chan string
	UUID       string
	logFile    *os.File
}

func NewExecutor(dpl *deployment.Deployment, cfg *Config, logChannel chan string) (*Executor, error) {
	if cfg == nil {
		cfg = &Config{
			JobTimeout:        3600,  // 1 hour
			DeploymentTimeout: 86400, // 1 day
			AllowFailure:      false,
		}
	}

	randId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	filePath := fmt.Sprintf("logs/deployments/%s.log", randId.String())
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		return nil, err
	}

	return &Executor{
		dpl,
		cfg,
		logChannel,
		randId.String(),
		file,
	}, nil
}

func (e *Executor) ExecuteJobs() error {
	defer e.logFile.Close()
	type deploymentContextKey string

	e.Log(fmt.Sprintf("Starting execution of deployment '%s'", e.Deployment.Name))

	deploymentContext, stopDeployment := context.WithCancel(context.Background())
	deploymentContext = context.WithValue(deploymentContext, deploymentContextKey("errorChannel"), make(chan error))
	defer stopDeployment()

	go func() {
		defer stopDeployment()
		errorChannel := deploymentContext.Value(deploymentContextKey("errorChannel")).(chan error)
		clients := map[string]*remote.Client{}

		for name, config := range e.Deployment.Config.SSH {
			client, err := remote.ConnectToHost(name, config.Host, config.Port, config.Type, config.Username, config.Password, config.KeyFile, config.KeyContent, config.KeyPass, e.Deployment.Variables)
			if err != nil {
				errorChannel <- err
				return
			}
			clients[name] = client
		}
		defer remote.CloseAllClients(clients)

		for _, job := range e.Deployment.Jobs {
			select {
			// Prevent the deployment goroutine from running other jobs if deployment is cancelled
			case <-deploymentContext.Done():
				return
			default:
			}
			jobName := job["name"].(string)
			jobHosts := job["hosts"].([]any)
			for _, host := range jobHosts {
				var client *remote.Client

				if host == "localhost" {
					client = remote.ConnectToLocalhost(e.Deployment.Variables)
				} else {
					c, ok := clients[host.(string)]
					if !ok {
						errorChannel <- fmt.Errorf("error: host %#v does not exist", host)
						return
					}
					client = c
				}

				jobContext, stopJob := deployment.NewJobContext(jobName, client, e.Log)
				defer stopJob()

				go func(host string) {
					defer stopJob()
					e.Log(fmt.Sprintf("Executing job '%s' on host '%s'", jobName, host))

					err := e.ExecuteJob(job, jobContext)
					if err != nil {
						errorChannel <- err
						return
					}

				}(host.(string))

				select {
				case <-jobContext.Done():
					continue
				case <-time.After(time.Duration(e.Config.JobTimeout) * time.Second):
					errorChannel <- fmt.Errorf("error: job '%s' on host '%s' timed out after %d seconds", jobName, host, e.Config.JobTimeout)
				}
			}
		}
	}()

	select {
	case err := <-deploymentContext.Value(deploymentContextKey("errorChannel")).(chan error):
		return err
	case <-deploymentContext.Done():
		e.Log(fmt.Sprintf("Deployment '%s' done, executed %d jobs", e.Deployment.Name, len(e.Deployment.Jobs)))
		return nil
	case <-time.After(time.Duration(e.Config.DeploymentTimeout) * time.Second):
		return fmt.Errorf("error: deployment '%s' timed out after %d seconds", e.Deployment.Name, e.Config.DeploymentTimeout)
	}
}

func (e *Executor) ExecuteJob(job map[string]any, ctx *deployment.JobContext) error {
	for key, value := range job {
		// Skip keys that are not plugins
		if key == "hosts" ||
			key == "name" ||
			key == "allow_failure" {
			continue
		}

		plg, exist := plugin.Plugins[key]
		if !exist {
			return fmt.Errorf("error: plugin '%s' does not exist", key)
		}

		output, err := plg.Execute(ctx, value)
		if err != nil {
			if allowedToFail, ok := job["allow_failure"].(bool); ok && allowedToFail || e.Config.AllowFailure {
				e.Log(fmt.Sprintf("Job %s failed but failure allowed: %s", job["name"], err))
			} else {
				return err
			}
		}
		if output != "" {
			e.Log(fmt.Sprintf("Job output: %s", output))
		}
	}

	return nil
}

func (e *Executor) Log(data string) {
	logStr := logging.Log(data)

	fmt.Fprintln(e.logFile, logStr)

	if e.logChannel != nil {
		e.logChannel <- logStr
	}

	logging.LogOut(data)
}
