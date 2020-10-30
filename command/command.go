package command

import (
	"os"
	"path/filepath"
)

type Command func(source, destination string) error

func MkDirCommand(_, destination string) error {
	dir, _ := filepath.Split(destination)
	return os.MkdirAll(dir, os.ModePerm)
}

type Executor struct {
	source      string
	destination string
}

func NewExecutor(source, destination string) *Executor {
	return &Executor{
		source:      source,
		destination: destination,
	}
}

func (e *Executor) Execute(commands ...Command) error {
	for _, command := range commands {
		if err := command(e.source, e.destination); err != nil {
			return err
		}
	}

	return nil
}
