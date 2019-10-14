package main

import (
	"github.com/threefoldtech/rivine/pkg/client"
)

// CommandLineClient extend for commands
type CommandLineClient struct {
	*client.CommandLineClient
}

// NewCommandLineClient creates a new eurochain commandline client
func NewCommandLineClient(address, name, userAgent string) (*CommandLineClient, error) {
	client, err := client.NewCommandLineClient(address, name, userAgent, nil)
	if err != nil {
		return nil, err
	}
	return &CommandLineClient{
		CommandLineClient: client,
	}, nil
}
