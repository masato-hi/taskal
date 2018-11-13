package main

import "github.com/stretchr/testify/mock"

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
	var ret []string
	args := m.Called()
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			ret = append(ret, v)
		}
	}
	return ret
}
func (m *MockOption) ConfigPath() string {
	return m.Called().String(0)
}
func (m *MockOption) TaskArgs() []string {
	var ret []string
	args := m.Called()
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			ret = append(ret, v)
		}
	}
	return ret
}
