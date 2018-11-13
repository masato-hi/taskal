package main

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockDefinedTask struct {
	mock.Mock
}

func (m *MockDefinedTask) Name() string {
	return m.Called().String(0)
}

func (m *MockDefinedTask) AddCommand(command string) {
	m.Called()
}

func (m *MockDefinedTask) Commands() []string {
	var ret []string
	args := m.Called()
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (m *MockDefinedTask) Run(dryRun bool, args []string) error {
	ret := m.Called(dryRun, args).Get(0)
	if v, ok := ret.(error); ok {
		return v
	} else {
		return nil
	}
}

func TestNewDefinedTask(t *testing.T) {
	t.Run("When called this func.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		name := "foo"
		actual := NewDefinedTask(name)
		expected := (*DefinedTask)(nil)
		assert.Implements(expected, actual)

		expected2 := "foo"
		assert.Equal(expected2, actual.Name())

		expected3 := "[DEBUG][15:04:05] Define Task: foo\n"
		assert.Equal(expected3, iobuffer.String())
	})
}

func TestDefinedTaskImpl_Name(t *testing.T) {
	t.Run("When test name is foo.", func(t *testing.T) {
		assert := assert2.New(t)

		task := DefinedTaskImpl{
			name: "foo",
		}

		actual := task.Name()

		expected := "foo"
		assert.Equal(expected, actual)
	})

	t.Run("When test name is bar.", func(t *testing.T) {
		assert := assert2.New(t)

		task := DefinedTaskImpl{
			name: "bar",
		}

		actual := task.Name()

		expected := "bar"
		assert.Equal(expected, actual)
	})
}

func TestDefinedTaskImpl_AddCommand(t *testing.T) {
	t.Run("When called this func at once.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		task := DefinedTaskImpl{}

		task.AddCommand("echo foo")

		expected := 1
		assert.Len(task.commands, expected)

		expected2 := "echo foo"
		assert.Equal(expected2, task.commands[0])

		expected3 := "[DEBUG][15:04:05]   Add Command: echo foo\n"
		assert.Equal(expected3, iobuffer.String())
	})

	t.Run("When called this func at twice.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		task := DefinedTaskImpl{}

		task.AddCommand("echo foo")
		task.AddCommand("echo bar")

		expected := 2
		assert.Len(task.commands, expected)

		expected2 := "echo foo"
		assert.Equal(expected2, task.commands[0])

		expected3 := "echo bar"
		assert.Equal(expected3, task.commands[1])

		expected4 := "[DEBUG][15:04:05]   Add Command: echo foo\n[DEBUG][15:04:05]   Add Command: echo bar\n"
		assert.Equal(expected4, iobuffer.String())
	})
}

func TestDefinedTaskImpl_Commands(t *testing.T) {
	t.Run("When added once command.", func(t *testing.T) {
		assert := assert2.New(t)

		task := DefinedTaskImpl{
			commands: []string{
				"echo foo",
			},
		}

		expected := 1
		assert.Len(task.commands, expected)

		expected2 := "echo foo"
		assert.Equal(expected2, task.commands[0])
	})

	t.Run("When added twice commands.", func(t *testing.T) {
		assert := assert2.New(t)

		task := DefinedTaskImpl{
			commands: []string{
				"echo foo",
				"echo bar",
			},
		}

		expected := 2
		assert.Len(task.commands, expected)

		expected2 := "echo foo"
		assert.Equal(expected2, task.commands[0])

		expected3 := "echo bar"
		assert.Equal(expected3, task.commands[1])
	})
}

func TestDefinedTaskImpl_Run(t *testing.T) {
	var executor *MockExecutor
	NewExecutor = func(dryRun bool, command string, args []string) Executor {
		return executor
	}

	t.Run("When at error occurred in runOnce.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		task := DefinedTaskImpl{
			name: "foo",
			commands: []string{
				"echo foo",
				"echo bar",
			},
		}
		dryRun := true
		args := []string{"bar"}
		executor = new(MockExecutor)

		executor.On("Execute").Return("error message", fmt.Errorf("mock return"))

		actual := task.Run(dryRun, args)

		expect := "mock return"
		assert.Error(actual, expect)

		expect2 := "[INFO][15:04:05] Execute task: foo\n[ERROR][15:04:05] mock return\n[ERROR][15:04:05] error message\n"
		assert.Equal(expect2, iobuffer.String())
	})

	t.Run("When no error occured in runOnce.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		task := DefinedTaskImpl{
			name: "foo",
			commands: []string{
				"echo foo",
				"echo bar",
			},
		}
		dryRun := true
		args := []string{"bar"}
		executor = new(MockExecutor)

		executor.On("Execute").Return("output message", nil)

		actual := task.Run(dryRun, args)

		assert.Nil(actual)

		expect2 := "[INFO][15:04:05] Execute task: foo\noutput message\noutput message\n"
		assert.Equal(expect2, iobuffer.String())
	})
}

func TestDefinedTaskImpl_runOnce(t *testing.T) {
	var executor *MockExecutor
	NewExecutor = func(dryRun bool, command string, args []string) Executor {
		return executor
	}

	t.Run("When at error occurred in executor.", func(t *testing.T) {
		t.Run("And not return output form executor.", func(t *testing.T) {
			iobuffer.Reset()

			assert := assert2.New(t)

			task := DefinedTaskImpl{}
			dryRun := true
			command := "echo foo"
			args := []string{"bar"}
			executor = new(MockExecutor)

			executor.On("Execute").Return("", fmt.Errorf("mock return"))

			actual := task.runOnce(dryRun, command, args)

			expect := "mock return"
			assert.Error(actual, expect)

			expect2 := "[ERROR][15:04:05] mock return\n"
			assert.Equal(expect2, iobuffer.String())
		})

		t.Run("And return output form executor.", func(t *testing.T) {
			iobuffer.Reset()

			assert := assert2.New(t)

			task := DefinedTaskImpl{}
			dryRun := true
			command := "echo foo"
			args := []string{"bar"}
			executor = new(MockExecutor)

			executor.On("Execute").Return("error message", fmt.Errorf("mock return"))

			actual := task.runOnce(dryRun, command, args)

			expect := "mock return"
			assert.Error(actual, expect)

			expect2 := "[ERROR][15:04:05] mock return\n[ERROR][15:04:05] error message\n"
			assert.Equal(expect2, iobuffer.String())
		})
	})

	t.Run("When no error occured in executor.", func(t *testing.T) {
		t.Run("And not return output form executor.", func(t *testing.T) {
			iobuffer.Reset()

			assert := assert2.New(t)

			task := DefinedTaskImpl{}
			dryRun := true
			command := "echo foo"
			args := []string{"bar"}
			executor = new(MockExecutor)

			executor.On("Execute").Return("", nil)

			actual := task.runOnce(dryRun, command, args)

			assert.Nil(actual)

			expected := ""
			assert.Equal(expected, iobuffer.String())
		})

		t.Run("And return output form executor.", func(t *testing.T) {
			iobuffer.Reset()

			assert := assert2.New(t)

			task := DefinedTaskImpl{}
			dryRun := true
			command := "echo foo"
			args := []string{"bar"}
			executor = new(MockExecutor)

			executor.On("Execute").Return("output message", nil)

			actual := task.runOnce(dryRun, command, args)

			assert.Nil(actual)

			expect2 := "output message\n"
			assert.Equal(expect2, iobuffer.String())
		})
	})
}
