package remote

import (
	"fmt"
	"os"
	"os/exec"

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

		// err can be *exec.ExitError
		output, err := session.Output(command)
		if err != nil {
			fmt.Println(logging.Log(err.Error()))
			return "", err
		}

		return string(output), nil
	}

	cmd := exec.Command("sh", "-c", command)
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

func ConnectToLocalhost() *Client {
	return &Client{
		name:        "localhost",
		sshClient:   nil,
		isLocalhost: true,
	}
}

func ConnectToHost(name, host string, port uint16, connectionType, username, password, keyFile, keyPass string, variables map[string]string) (*Client, error) {
	var sshClient *ssh.Client
	var err error

	switch connectionType {
	case "password":
		sshClient, err = connectWithPassword(host, port, username, password)
	case "key":
		sshClient, err = connectWithPrivateKey(host, port, username, keyFile, keyPass)
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

func connectWithPrivateKey(host string, port uint16, username, keyFile, keyPass string) (*ssh.Client, error) {
	keyFileContent, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	var signer ssh.Signer

	if keyPass != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(keyFileContent, []byte(keyPass))
		if err != nil {
			return nil, err
		}
	} else {
		signer, err = ssh.ParsePrivateKey(keyFileContent)
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
