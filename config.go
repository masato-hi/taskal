package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"strings"
)

type Config interface {
	AddDefinedTask(DefinedTask)
	DefinedTasks() []DefinedTask
	ShowAllDefinedTasks()
}

type ConfigImpl struct {
	path         string
	definedTasks []DefinedTask
}

type Node interface{}
type Document map[string]Node

func (c *ConfigImpl) AddDefinedTask(task DefinedTask) {
	c.definedTasks = append(c.definedTasks, task)
}

func (c *ConfigImpl) DefinedTasks() []DefinedTask {
	return c.definedTasks
}

func (c *ConfigImpl) ShowAllDefinedTasks() {
	Printf("All defined tasks:\n")
	for _, task := range c.DefinedTasks() {
		Printf("%s", task.Name())
	}
}

func (c *ConfigImpl) sortDefinedTasks() {
	sort.Slice(c.definedTasks, func(i int, j int) bool {
		return c.definedTasks[i].Name() < c.definedTasks[j].Name()
	})
}

var ReadConfig = func(path string) (string, error) {
	if buf, err := ioutil.ReadFile(path); err != nil {
		Error("Config file read error. path: %s", path)
		return "", err
	} else {
		return string(buf), nil
	}
}

var ParseConfig = func(buf string) (Config, error) {
	config := &ConfigImpl{}

	var document = make(Document)
	if err := yaml.Unmarshal([]byte(buf), &document); err != nil {
		Error(err.Error())
		return nil, err
	}

	var task DefinedTask
	for taskName, rootNode := range document {
		if strings.HasPrefix(taskName, "_") {
			continue
		}

		if command, ok := rootNode.(string); ok {
			task = NewDefinedTask(taskName)
			task.AddCommand(command)
			config.AddDefinedTask(task)
		} else if node, ok := rootNode.(Node); ok {
			task = NewDefinedTask(taskName)
			parseNode(task, node)
			config.AddDefinedTask(task)
		}
	}

	config.sortDefinedTasks()

	return config, nil
}

func parseNode(task DefinedTask, node Node) {
	if command, ok := node.(string); ok {
		task.AddCommand(command)
	} else if list, ok := node.([]interface{}); ok {
		for _, childNode := range list {
			parseNode(task, childNode)
		}
	}
}
