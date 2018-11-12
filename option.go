package main

import (
	"flag"
	"fmt"
)

type Option interface {
	WillBeShowTasks() bool
	BeDryRun() bool
	HasSpecifiedTasks() bool
	SpecifiedTasks() []string
	ConfigPath() string
	TaskArgs() []string
}

type OptionImpl struct {
	willBeShowTasks bool
	beDryRun        bool
	specifiedTasks  []string
	configPath      string
	taskArgs        []string
}

var ParseOption = func(args []string) (Option, error) {
	option := &OptionImpl{}

	f := flag.NewFlagSet("taskal", flag.ContinueOnError)
	f.SetOutput(Stderr)
	f.Usage = func() {
		fmt.Fprintln(f.Output(), fmt.Sprintf("taskal %s", VERSION))
		fmt.Fprintln(f.Output(), "Usage: taskal [options...] [tasks ...] -- [args...]")
		f.PrintDefaults()
	}
	f.BoolVar(&option.willBeShowTasks, "T", false, "Show all tasks.")
	f.BoolVar(&option.beDryRun, "n", false, "Do a dry run without executing actions.")
	f.StringVar(&option.configPath, "c", "taskal.yml", "taskal -c [CONFIGFILE]")

	if err := f.Parse(args[1:]); err != nil {
		return nil, err
	}

	option.specifiedTasks, option.taskArgs = parseTaskAndArgs(f.Args())

	Debug("%v, %v", option.specifiedTasks, option.taskArgs)

	return option, nil
}

var parseTaskAndArgs = func(args []string) ([]string, []string) {
	if len(args) == 0 {
		return []string{}, []string{}
	}

	for i := 0; i < len(args); i++ {
		if args[i] == "--" {
			return args[:i], args[i+1:]
		}
	}

	return args, []string{}
}

func (o *OptionImpl) WillBeShowTasks() bool {
	return o.willBeShowTasks
}

func (o *OptionImpl) BeDryRun() bool {
	return o.beDryRun
}

func (o *OptionImpl) HasSpecifiedTasks() bool {
	return len(o.specifiedTasks) > 0
}

func (o *OptionImpl) SpecifiedTasks() []string {
	return o.specifiedTasks
}

func (o *OptionImpl) ConfigPath() string {
	return o.configPath
}

func (o *OptionImpl) TaskArgs() []string {
	return o.taskArgs
}
