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

func (m *MockExecutor) Execute() (string, error) {
	str := m.Called().String(0)
	ret := m.Called().Get(1)

	if v, ok := ret.(error); ok {
		return str, v
	} else {
		return str, nil
	}
}

func TestExecutorImpl_Execute(t *testing.T) {
	t.Run("When an error occurred in execInner.", func(t *testing.T) {
		doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
			return ([]byte)("  output message  "), fmt.Errorf("  error message  ")
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

		actual, err := executor.Execute()

		expected := "  output message"
		assert.Equal(expected, actual)

		expected2 := "  error message"
		assert.Error(err, expected2)
	})

	t.Run("When no error occurred in execInner.", func(t *testing.T) {
		doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
			return ([]byte)("  output message  "), nil
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

		actual, err := executor.Execute()

		expected := "  output message"
		assert.Equal(expected, actual)

		assert.Nil(err)
	})
}

func TestExecutorImpl_execInner(t *testing.T) {
	doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
		return ([]byte)("output message"), fmt.Errorf("error message")
	}

	assert := assert2.New(t)

	executor := ExecutorImpl{
		dryRun:  false,
		command: "echo foo",
		args: []string{
			"bar",
			"baz",
		},
	}

	actual, err := executor.execInner()

	expected := ([]byte)("output message")
	assert.Equal(expected, actual)

	expected2 := "error message"
	assert.Error(err, expected2)
}

func TestExecutorImpl_execOnWindows(t *testing.T) {
	var execName string
	var execArgs []string
	doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
		execName = name
		execArgs = args
		return ([]byte)("output message"), fmt.Errorf("error message")
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

		actual, err := executor.execOnWindows()

		expected := ([]byte)("output message")
		assert.Equal(expected, actual)

		expected2 := "error message"
		assert.Error(err, expected2)

		expected3 := "[INFO][15:04:05] exec \"echo foo\"\n"
		assert.Equal(expected3, iobuffer.String())

		expected4 := "exec"
		assert.Equal(expected4, execName)

		expected5 := []string{
			"echo foo",
		}
		assert.Equal(expected5, execArgs)
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

		actual, err := executor.execOnWindows()

		expected := ([]byte)("output message")
		assert.Equal(expected, actual)

		expected2 := "error message"
		assert.Error(err, expected2)

		expected3 := "[INFO][15:04:05] exec \"echo foo\"\n"
		assert.Equal(expected3, iobuffer.String())

		expected4 := "exec"
		assert.Equal(expected4, execName)

		expected5 := []string{
			"echo foo",
		}
		assert.Equal(expected5, execArgs)
	})
}

func TestExecutorImpl_execOnUnix(t *testing.T) {
	var execName string
	var execArgs []string
	doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
		execName = name
		execArgs = args
		return ([]byte)("output message"), fmt.Errorf("error message")
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

		actual, err := executor.execOnUnix()

		expected := ([]byte)("output message")
		assert.Equal(expected, actual)

		expected2 := "error message"
		assert.Error(err, expected2)

		expected3 := "[INFO][15:04:05] sh -c \"echo foo\"\n"
		assert.Equal(expected3, iobuffer.String())

		expected4 := "sh"
		assert.Equal(expected4, execName)

		expected5 := []string{
			"-c",
			"echo foo",
		}
		assert.Equal(expected5, execArgs)
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

		actual, err := executor.execOnUnix()

		expected := ([]byte)("output message")
		assert.Equal(expected, actual)

		expected2 := "error message"
		assert.Error(err, expected2)

		expected3 := "[INFO][15:04:05] sh -c \"echo foo\" -- bar baz\n"
		assert.Equal(expected3, iobuffer.String())

		expected4 := "sh"
		assert.Equal(expected4, execName)

		expected5 := []string{
			"-c",
			"echo foo",
			"--",
			"bar",
			"baz",
		}
		assert.Equal(expected5, execArgs)
	})
}

func TestExecutorImpl_execCommand(t *testing.T) {
	doExecCommand = func(name string, args ...string) (bytes []byte, e error) {
		return ([]byte)("output message"), fmt.Errorf("error message")
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

		actual, err := executor.execCommand(name, args...)

		expected := ([]byte)("output message")
		assert.Equal(expected, actual)

		expected2 := "error message"
		assert.Error(err, expected2)
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

		actual, err := executor.execCommand(name, args...)

		expected := make([]byte, 0)
		assert.Equal(expected, actual)

		assert.Nil(err)
	})
}
