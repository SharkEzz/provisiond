package remote

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/SharkEzz/provisiond/pkg/logging"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	name        string
	sshClient   *ssh.Client
	variables   map[string]string
	isLocalhost bool
}

func (c *Client) ExecuteCommand(command string) (string, error) {
	if c.sshClient == nil && !c.isLocalhost {
		err := fmt.Errorf("error: client is not created")
		fmt.Println(logging.Log(err.Error()))
		return "", err
	}

	// If the client is remote (SSH)
	if !c.isLocalhost {
		session, err := c.sshClient.NewSession()
		if err != nil {
			fmt.Println(logging.Log(err.Error()))
			return "", err
		}
		defer session.Close()

		for name, value := range c.variables {
			err := session.Setenv(name, value)
			if err != nil {
				return "", fmt.Errorf("error while setting %s variable", name)
			}
		}

		// err can be *ssh.ExitError
		output, err := session.Output(command)
		if err != nil {
			fmt.Println(logging.Log(err.Error()))
			return "", err
		}

		return string(output), nil
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "sh"
		}
		cmd = exec.Command(shell, "-c", command)
	}

	// Set environment for the current command
	for key, value := range c.variables {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(logging.Log(err.Error()))
		return "", err
	}

	return string(output), nil
}

func CloseAllClients(clients map[string]*Client) {
	for _, client := range clients {
		if client.sshClient != nil {
			err := client.sshClient.Close()
			if err != nil {
				fmt.Println(logging.Log(fmt.Sprintf("error closing client %s: %s", client.name, err.Error())))
			}
		}
	}
}

func ConnectToLocalhost(variables map[string]string) *Client {
	return &Client{
		name:        "localhost",
		sshClient:   nil,
		isLocalhost: true,
		variables:   variables,
	}
}

func ConnectToHost(name, host string, port uint16, connectionType, username, password, keyFile, keyContent, keyPass string, variables map[string]string) (*Client, error) {
	var sshClient *ssh.Client
	var err error

	switch connectionType {
	case "password":
		sshClient, err = connectWithPassword(host, port, username, password)
	case "key":
		sshClient, err = connectWithPrivateKey(host, port, username, keyFile, keyContent, keyPass)
	default:
		return nil, fmt.Errorf("error: unknown login method '%s'", connectionType)
	}
	if err != nil {
		return nil, err
	}

	client := &Client{
		name,
		sshClient,
		variables,
		false,
	}

	return client, nil
}

func connectWithPassword(host string, port uint16, username, password string) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	// TODO: fix this
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func connectWithPrivateKey(host string, port uint16, username, keyFile, keyContent, keyPass string) (*ssh.Client, error) {
	var signer ssh.Signer
	var err error

	if keyFile == "" && keyContent == "" {
		return nil, fmt.Errorf("error: either keyFile or keyContent must be set")
	}
	if keyFile != "" && keyContent != "" {
		return nil, fmt.Errorf("error: keyFile and keyContent cannot be set at the same time")
	}

	if keyFile != "" {
		fileContent, err := os.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		keyContent = string(fileContent)
	}

	if keyPass != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(keyContent), []byte(keyPass))
		if err != nil {
			return nil, err
		}
	} else {
		signer, err = ssh.ParsePrivateKey([]byte(keyContent))
		if err != nil {
			return nil, err
		}
	}

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	// TODO: fix this
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
