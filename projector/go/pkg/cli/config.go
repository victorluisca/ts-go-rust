package projector

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Delete
)

type Config struct {
	Arguments []string
	Operation Operation
	Pwd       string
	Config    string
}

func getPwd(opts *ProjectorOptions) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getConfig(opts *ProjectorOptions) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}

	config, err := os.UserConfigDir()
	if err != nil {
		return "", nil
	}

	return path.Join(config, "projector", "projector.json"), nil
}

func getOperation(opts *ProjectorOptions) Operation {
	if len(opts.Arguments) == 0 {
		return Print
	}

	if opts.Arguments[0] == "add" {
		return Add
	}

	if opts.Arguments[0] == "del" {
		return Delete
	}

	return Print
}

func getArgs(opts *ProjectorOptions) ([]string, error) {
	if len(opts.Arguments) == 0 {
		return []string{}, nil
	}

	operation := getOperation(opts)

	if operation == Add {
		if len(opts.Arguments) != 3 {
			return nil, fmt.Errorf("add requires 2 arguments but received %v", len(opts.Arguments)-1)
		}
		return opts.Arguments[1:], nil
	}

	if operation == Delete {
		if len(opts.Arguments) != 2 {
			return nil, fmt.Errorf("delete requires 1 argument but received %v", len(opts.Arguments)-1)
		}
		return opts.Arguments[1:], nil
	}

	if len(opts.Arguments) > 1 {
		return nil, fmt.Errorf("print requires 0 or 1 arguments but received %v", len(opts.Arguments)-1)
	}

	return opts.Arguments, nil
}

func NewConfig(opts *ProjectorOptions) (*Config, error) {
	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}

	config, err := getConfig(opts)
	if err != nil {
		return nil, err
	}

	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Arguments: args,
		Operation: getOperation(opts),
		Pwd:       pwd,
		Config:    config,
	}, nil
}
