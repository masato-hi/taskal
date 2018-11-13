package main

import "github.com/stretchr/testify/mock"

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
	var ret []DefinedTask
	args := m.Called()
	for _, arg := range args {
		if v, ok := arg.(DefinedTask); ok {
			ret = append(ret, v)
		}
	}
	return ret
}
