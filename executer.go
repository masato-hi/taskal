package main

import (
	"github.com/fatih/color"
	"os/exec"
	"runtime"
	"strings"
)

type Executer interface {
	Execute() (string, error)
}

type ExecuterImpl struct {
	dryRun  bool
	command string
	args    []string
}

var NewExecuter = func(dryRun bool, command string, args []string) Executer {
	return &ExecuterImpl{dryRun, command, args}
}

func (e *ExecuterImpl) Execute() (string, error) {
	raw, err := e.executeInner()
	out := TrimTailingSpace(string(raw))

	if err == nil {
		return out, nil
	} else {
		return out, err
	}
}

func (e *ExecuterImpl) executeInner() ([]byte, error) {
	if runtime.GOOS == "windows" {
		return e.executeWindows()
	} else {
		return e.executeUnix()
	}
}

func (e *ExecuterImpl) executeWindows() ([]byte, error) {
	Info(color.HiBlackString(e.command))
	return e.execCommand(e.command)
}

func (e *ExecuterImpl) executeUnix() ([]byte, error) {
	execArgs := []string{
		"-c",
		e.command,
	}

	if len(e.args) > 0 {
		execArgs = append(execArgs, "--")
		execArgs = append(execArgs, e.args...)

		Info(color.HiBlackString("sh -c %s -- %s", QuoteString(e.command), strings.Join(e.args, " ")))
	} else {
		Info(color.HiBlackString("sh -c %s", QuoteString(e.command)))
	}

	return e.execCommand("sh", execArgs...)
}

func (e ExecuterImpl) execCommand(name string, args ...string) ([]byte, error) {
	if !e.dryRun {
		return doExecCommand(name, args...)
	} else {
		return []byte{}, nil
	}
}

var doExecCommand = func(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).CombinedOutput()
}
