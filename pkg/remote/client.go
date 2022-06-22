package remote

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	name      string
	sshClient *ssh.Client
}

func (c *Client) ExecuteCommand(command string) error {
	if c.sshClient == nil {
		err := fmt.Errorf("error: client is not created")
		logrus.Error(err)
		return err
	}

	session, err := c.sshClient.NewSession()
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer session.Close()

	err = session.Run(command)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// TODO: get the stdout / stderr output ?

	return nil
}

func CloseAllClients(clients map[string]*Client) {
	for _, client := range clients {
		if client.sshClient != nil {
			err := client.sshClient.Close()
			if err != nil {
				logrus.Error(fmt.Errorf("cannot close client '%s'", client.name))
			}
		}
	}
}

func ConnectToHost(name, host string, port uint16, connectionType, username, password, keyFile string) (*Client, error) {
	var sshClient *ssh.Client
	var err error

	switch connectionType {
	case "password":
		sshClient, err = connectWithPassword(host, port, username, password)
	case "key":
		sshClient, err = connectWithPrivateKey(host, port, username, keyFile)
	default:
		return nil, fmt.Errorf("error: unknown login method '%s'", connectionType)
	}

	if err != nil {
		return nil, err
	}

	client := &Client{
		name,
		sshClient,
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

func connectWithPrivateKey(host string, port uint16, username, keyFile string) (*ssh.Client, error) {
	keyFileContent, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(keyFileContent)
	if err != nil {
		return nil, err
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
