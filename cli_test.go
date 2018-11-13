package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewCLI(t *testing.T) {
	assert := assert.New(t)

	t.Run("Will expected to returns CLI implementation.", func(t *testing.T) {
		actual := NewCLI()
		expected := (*CLI)(nil)
		assert.Implements(expected, actual)
	})
}

type MockOption struct {
	mock.Mock
}

func (m *MockOption) WillBeShowTasks() bool {
	return m.Called().Bool(0)
}
func (m *MockOption) BeDryRun() bool {
	return m.Called().Bool(0)
}
func (m *MockOption) HasSpecifiedTasks() bool {
	return m.Called().Bool(0)
}
func (m *MockOption) SpecifiedTasks() []string {
	return []string{m.Called().String(0)}
}
func (m *MockOption) ConfigPath() string {
	return m.Called().String(0)
}
func (m *MockOption) TaskArgs() []string {
	return []string{m.Called().String(0)}
}

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) ShowAllDefinedTasks() {
	m.Called()
}
func (m *MockConfig) AddDefinedTask(task DefinedTask) {
	m.Called(task)
}
func (m *MockConfig) DefinedTasks() []DefinedTask {
	m.Called()
	return []DefinedTask{}
}

type MockRunner struct {
	mock.Mock
}

func (m *MockRunner) Run() error {
	return m.Called().Error(0)
}

func TestCLIImpl_Run(t *testing.T) {
	target := NewCLI()
	args := []string{"taskal"}

	t.Run("When the option parsing fails.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			return nil, fmt.Errorf("invalid option")
		}

		actual := target.Run(args)
		expected := InvalidOption
		assert.Equal(expected, actual)
	})

	t.Run("When reading config file failed.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			option := new(MockOption)
			option.On("ConfigPath").Return("")
			return option, nil
		}
		ReadConfig = func(path string) (string, error) {
			return "", fmt.Errorf("config file not exists")
		}

		actual := target.Run(args)
		expected := UnreadConfig
		assert.Equal(expected, actual)
	})

	t.Run("When parsing of the config file fails.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			option := new(MockOption)
			option.On("ConfigPath").Return("")
			return option, nil
		}
		ReadConfig = func(path string) (string, error) {
			return "", nil
		}
		ParseConfig = func(buf string) (Config, error) {
			return nil, fmt.Errorf("invalid config")
		}

		actual := target.Run(args)
		expected := InvalidConfig
		assert.Equal(expected, actual)
	})

	t.Run("When will be show tasks was specified.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			option := new(MockOption)
			option.On("ConfigPath").Return("")
			option.On("WillBeShowTasks").Return(true)
			return option, nil
		}
		ReadConfig = func(path string) (string, error) {
			return "", nil
		}
		ParseConfig = func(buf string) (Config, error) {
			config := new(MockConfig)
			config.On("ShowAllDefinedTasks")
			return config, nil
		}

		actual := target.Run(args)
		expected := Succeeded
		assert.Equal(expected, actual)
	})

	t.Run("When the failed to run Runner.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			option := new(MockOption)
			option.On("ConfigPath").Return("")
			option.On("WillBeShowTasks").Return(false)
			return option, nil
		}
		ReadConfig = func(path string) (string, error) {
			return "", nil
		}
		ParseConfig = func(buf string) (Config, error) {
			config := new(MockConfig)
			return config, nil
		}
		NewRunner = func(option Option, config Config) Runner {
			runner := new(MockRunner)
			runner.On("Run").Return(fmt.Errorf("failed to run"))
			return runner
		}

		actual := target.Run(args)
		expected := FailedExecute
		assert.Equal(expected, actual)
	})

	t.Run("When the succeeded to run Runner.", func(t *testing.T) {
		assert := assert.New(t)
		ParseOption = func(args []string) (Option, error) {
			option := new(MockOption)
			option.On("ConfigPath").Return("")
			option.On("WillBeShowTasks").Return(false)
			return option, nil
		}
		ReadConfig = func(path string) (string, error) {
			return "", nil
		}
		ParseConfig = func(buf string) (Config, error) {
			config := new(MockConfig)
			return config, nil
		}
		NewRunner = func(option Option, config Config) Runner {
			runner := new(MockRunner)
			runner.On("Run").Return(nil)
			return runner
		}

		actual := target.Run(args)
		expected := Succeeded
		assert.Equal(expected, actual)
	})
}
