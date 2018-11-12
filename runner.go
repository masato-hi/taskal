package main

import "fmt"

type Runner interface {
	Run() error
}

type RunnerImpl struct {
	Option Option
	Config Config
}

var NewRunner = func(option Option, config Config) Runner {
	return &RunnerImpl{
		Option: option,
		Config: config,
	}
}

func (r *RunnerImpl) Run() error {
	if !r.Option.HasSpecifiedTasks() {
		Error("Task is not specified")
		return fmt.Errorf("task is not specified")
	}

	tasks := r.specifiedDefinedTasks()
	for _, task := range tasks {
		if err := r.runOnce(task); err != nil {
			return err
		}
	}
	return nil
}

func (r *RunnerImpl) specifiedDefinedTasks() []DefinedTask {
	var tasks []DefinedTask
	for _, specifiedTask := range r.Option.SpecifiedTasks() {
		for _, definedTask := range r.Config.DefinedTasks() {
			if definedTask.Name() == specifiedTask {
				tasks = append(tasks, definedTask)
			}
		}
	}
	return tasks
}

func (r *RunnerImpl) runOnce(task DefinedTask) error {
	dryRun := r.Option.BeDryRun()
	args := r.Option.TaskArgs()
	task.Run(dryRun, args)
	return nil
}
