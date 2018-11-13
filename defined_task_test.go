package main

import "github.com/stretchr/testify/mock"

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
