package main

import (
	"github.com/fatih/color"
	"strings"
)

type DefinedTask interface {
	Name() string
	AddCommand(string)
	Commands() []string
	Run(bool, []string) error
}

type DefinedTaskImpl struct {
	name     string
	commands []string
}

var NewDefinedTask = func(name string) DefinedTask {
	Debug("Define Task: %s", name)
	return &DefinedTaskImpl{
		name: name,
	}
}

func (d *DefinedTaskImpl) Name() string {
	return d.name
}

func (d *DefinedTaskImpl) AddCommand(command string) {
	Debug("  Add Command: %s", command)
	d.commands = append(d.commands, strings.TrimSpace(command))
}

func (d *DefinedTaskImpl) Commands() []string {
	return d.commands
}

func (d *DefinedTaskImpl) Run(dryRun bool, args []string) error {
	Info(color.HiYellowString("Execute task: %s", d.name))

	commands := d.Commands()
	for _, command := range commands {
		if err := d.runOnce(dryRun, command, args); err != nil {
			return err
		}
	}
	return nil
}

func (d *DefinedTaskImpl) runOnce(dryRun bool, command string, args []string) error {
	executor := NewExecutor(dryRun, command, args)
	if out, err := executor.Execute(); err != nil {
		Error(err.Error())

		if len(out) > 0 {
			Error(out)
		}

		return err
	} else {
		if len(out) > 0 {
			Printf(out)
		}

		return nil
	}
}
