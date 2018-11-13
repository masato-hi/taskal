package main

import "github.com/stretchr/testify/mock"

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
