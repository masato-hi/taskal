package main

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRunner struct {
	mock.Mock
}

func (m *MockRunner) Run() error {
	return m.Called().Error(0)
}

func TestRunnerImpl_Run(t *testing.T) {
	buffer := &bytes.Buffer{}
	Stdout = buffer
	Stderr = buffer
	color.NoColor = true
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("When task is not specified.", func(t *testing.T) {
		option := new(MockOption)
		config := new(MockConfig)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("HasSpecifiedTasks").Return(false)

		t.Run("And return errors.", func(t *testing.T) {
			buffer.Reset()

			assert := assert2.New(t)
			actual := runner.Run()
			expected := "task is not specified"
			assert.Error(actual, expected)
		})

		t.Run("And log error messages.", func(t *testing.T) {
			buffer.Reset()

			assert := assert2.New(t)
			runner.Run()
			expected := "[ERROR][15:04:05] Task is not specified\n"
			assert.Equal(expected, buffer.String())
		})
	})

	t.Run("When specified task is not defined.", func(t *testing.T) {
		buffer.Reset()

		option := new(MockOption)
		config := new(MockConfig)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("HasSpecifiedTasks").Return(true)
		option.On("SpecifiedTasks").Return("bar")
		config.On("DefinedTasks").Return(
			&DefinedTaskImpl{
				name:     "foo",
				commands: []string{"echo foo", "echo foobar"},
			},
		)

		assert := assert2.New(t)
		actual := runner.Run()
		expected := "task is not specified"
		assert.Error(actual, expected)
	})

	t.Run("When specified task is defined.", func(t *testing.T) {
		buffer.Reset()

		option := new(MockOption)
		config := new(MockConfig)
		task := new(MockDefinedTask)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("HasSpecifiedTasks").Return(true)
		option.On("SpecifiedTasks").Return("foo")

		option.On("BeDryRun").Once().Return(true)
		option.On("BeDryRun").Twice().Return(false)
		option.On("TaskArgs").Return("foo", "bar")
		config.On("DefinedTasks").Return(task, task)

		task.On("Name").Once().Return("bar")
		task.On("Name").Twice().Return("foo")
		task.On("Run", true, []string{"foo", "bar"}).Return(fmt.Errorf("mock return"))
		task.On("Run", false, []string{"foo", "bar"}).Return(nil)

		t.Run("When an error occurred on run tasks.", func(t *testing.T) {
			task.On("Run", true, []string{"foo", "bar"}).Return(fmt.Errorf("mock return"))

			assert := assert2.New(t)
			actual := runner.Run()
			expected := "mock return"
			assert.Error(actual, expected)
		})

		t.Run("When no error occurred on run tasks.", func(t *testing.T) {
			task.On("Run", true, []string{"foo", "bar"}).Return(nil)
			task.On("Run", false, []string{"foo", "bar"}).Return(nil)

			assert := assert2.New(t)
			actual := runner.Run()
			assert.Nil(actual)
		})
	})
}

func TestRunnerImpl_specifiedDefinedTasks(t *testing.T) {
	buffer := &bytes.Buffer{}
	Stdout = buffer
	Stderr = buffer
	color.NoColor = true
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("Found specified Tasks.", func(t *testing.T) {
		assert := assert2.New(t)
		option := new(MockOption)
		config := new(MockConfig)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("SpecifiedTasks").Return("bar", "foo")
		config.On("DefinedTasks").Return(
			&DefinedTaskImpl{
				name:     "foo",
				commands: []string{"echo foo", "echo foobar"},
			},
			&DefinedTaskImpl{
				name:     "bar",
				commands: []string{"echo bar"},
			},
			&DefinedTaskImpl{
				name:     "baz",
				commands: []string{"echo baz"},
			},
		)

		tasks, err := runner.specifiedDefinedTasks()
		assert.NoError(err)

		expected := 2
		assert.Len(tasks, expected)

		expected2 := "bar"
		assert.Equal(expected2, tasks[0].Name())

		expected3 := "foo"
		assert.Equal(expected3, tasks[1].Name())
	})

	t.Run("Not found specified Tasks.", func(t *testing.T) {
		assert := assert2.New(t)
		option := new(MockOption)
		config := new(MockConfig)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("SpecifiedTasks").Return("pii", "poo")
		config.On("DefinedTasks").Return(
			&DefinedTaskImpl{
				name:     "foo",
				commands: []string{"echo foo", "echo foobar"},
			},
			&DefinedTaskImpl{
				name:     "bar",
				commands: []string{"echo bar"},
			},
			&DefinedTaskImpl{
				name:     "baz",
				commands: []string{"echo baz"},
			},
		)

		tasks, err := runner.specifiedDefinedTasks()

		assert.Nil(tasks)

		expected := "specified task is not defined"
		assert.Errorf(err, expected)

		expected2 := "[WARN][15:04:05] Specified task is not defined. task: pii\n"
		assert.Equal(expected2, buffer.String())
	})
}

func TestRunnerImpl_runOnce(t *testing.T) {
	t.Run("When runOnce is executed", func(t *testing.T) {
		assert := assert2.New(t)
		option := new(MockOption)
		config := new(MockConfig)
		task := new(MockDefinedTask)
		runner := RunnerImpl{
			Option: option,
			Config: config,
		}

		option.On("BeDryRun").Return(true)
		option.On("TaskArgs").Return("foo", "bar")
		task.On("Run", true, []string{"foo", "bar"}).Return(fmt.Errorf("mock return"))

		actual := runner.runOnce(task)
		expected := "mock return"
		assert.Error(actual, expected)
	})
}
