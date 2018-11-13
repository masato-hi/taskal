package main

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) Execute() error {
	ret := m.Called().Get(0)

	if v, ok := ret.(error); ok {
		return v
	} else {
		return nil
	}
}

func TestExecutorImpl_Execute(t *testing.T) {
	t.Run("When an error occurred.", func(t *testing.T) {
		doExecCommand = func(name string, args ...string) error {
			return fmt.Errorf("error message")
		}

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args: []string{
				"bar",
				"buz",
			},
		}

		actual := executor.Execute()

		assert.Error(actual)
	})

	t.Run("When no error occurred.", func(t *testing.T) {
		doExecCommand = func(name string, args ...string) error {
			return nil
		}

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args: []string{
				"bar",
				"buz",
			},
		}

		actual := executor.Execute()

		assert.NoError(actual)
	})
}

func TestExecutorImpl_execOnWindows(t *testing.T) {
	var execName string
	var execArgs []string
	doExecCommand = func(name string, args ...string) error {
		execName = name
		execArgs = args
		return fmt.Errorf("error message")
	}

	t.Run("When has not sub command arguments.", func(t *testing.T) {
		iobuffer.Reset()
		execName = ""
		execArgs = nil

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args:    []string{},
		}

		actual := executor.execOnWindows()

		assert.Error(actual)

		expected := "[INFO][15:04:05] exec \"echo foo\"\n"
		assert.Equal(expected, iobuffer.String())

		expected2 := "exec"
		assert.Equal(expected2, execName)

		expected3 := []string{
			"echo foo",
		}
		assert.Equal(expected3, execArgs)
	})

	t.Run("When has sub command arguments.", func(t *testing.T) {
		iobuffer.Reset()
		execName = ""
		execArgs = nil

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args: []string{
				"bar",
				"baz",
			},
		}

		actual := executor.execOnWindows()

		assert.Error(actual)

		expected := "[INFO][15:04:05] exec \"echo foo\"\n"
		assert.Equal(expected, iobuffer.String())

		expected2 := "exec"
		assert.Equal(expected2, execName)

		expected3 := []string{
			"echo foo",
		}
		assert.Equal(expected3, execArgs)
	})
}

func TestExecutorImpl_execOnUnix(t *testing.T) {
	var execName string
	var execArgs []string
	doExecCommand = func(name string, args ...string) error {
		execName = name
		execArgs = args
		return fmt.Errorf("error message")
	}

	t.Run("When has not sub command arguments.", func(t *testing.T) {
		iobuffer.Reset()
		execName = ""
		execArgs = nil

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args:    []string{},
		}

		actual := executor.execOnUnix()

		assert.Error(actual)

		expected := "[INFO][15:04:05] sh -c \"echo foo\"\n"
		assert.Equal(expected, iobuffer.String())

		expected2 := "sh"
		assert.Equal(expected2, execName)

		expected3 := []string{
			"-c",
			"echo foo",
		}
		assert.Equal(expected3, execArgs)
	})

	t.Run("When has sub command arguments.", func(t *testing.T) {
		iobuffer.Reset()
		execName = ""
		execArgs = nil

		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun:  false,
			command: "echo foo",
			args: []string{
				"bar",
				"baz",
			},
		}

		actual := executor.execOnUnix()

		assert.Error(actual)

		expected := "[INFO][15:04:05] sh -c \"echo foo\" -- bar baz\n"
		assert.Equal(expected, iobuffer.String())

		expected2 := "sh"
		assert.Equal(expected2, execName)

		expected3 := []string{
			"-c",
			"echo foo",
			"--",
			"bar",
			"baz",
		}
		assert.Equal(expected3, execArgs)
	})
}

func TestExecutorImpl_execCommand(t *testing.T) {
	doExecCommand = func(name string, args ...string) error {
		return fmt.Errorf("error message")
	}

	t.Run("When disabled dry run flag.", func(t *testing.T) {
		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun: false,
		}
		name := "foo"
		args := []string{
			"bar",
			"baz",
		}

		actual := executor.execCommand(name, args...)

		assert.Error(actual)
	})

	t.Run("When enable dry run flag.", func(t *testing.T) {
		assert := assert2.New(t)

		executor := ExecutorImpl{
			dryRun: true,
		}
		name := "foo"
		args := []string{
			"bar",
			"baz",
		}

		actual := executor.execCommand(name, args...)

		assert.NoError(actual)
	})
}
