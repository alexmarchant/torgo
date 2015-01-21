package commands

import (
	"errors"
	"os"
)

var (
	BadCommand = errors.New("Bad command")
)

type Command struct {
	Name    string
	Options []string
}

func NewCommand() (cmd *Command, err error) {
	if len(os.Args) < 2 {
		return nil, BadCommand
	}

	cmd = &Command{
		Name:    os.Args[1],
		Options: os.Args[2:],
	}

	return
}

func Execute() error {
	cmd, err := NewCommand()

	if err != nil {
		return err
	}

	switch cmd.Name {
	case "download":
		return Download(cmd)
	case "status":
		return Status(cmd)
	}

	return BadCommand
}
