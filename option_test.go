package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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

func TestOptionImpl_WillBeShowTasks(t *testing.T) {
	t.Run("When show tasks flag filed was true.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			willBeShowTasks: true,
		}

		actual := option.WillBeShowTasks()

		assert.True(actual)
	})

	t.Run("When show tasks flag filed was false.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			willBeShowTasks: false,
		}

		actual := option.WillBeShowTasks()

		assert.False(actual)
	})
}

func TestOptionImpl_BeDryRun(t *testing.T) {
	t.Run("When dry run flag filed was true.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			beDryRun: true,
		}

		actual := option.BeDryRun()

		assert.True(actual)
	})

	t.Run("When dry run flag filed was false.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			beDryRun: false,
		}

		actual := option.BeDryRun()

		assert.False(actual)
	})
}

func TestOptionImpl_HasSpecifiedTasks(t *testing.T) {
	t.Run("When has specified tasks.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			specifiedTasks: []string{
				"foo",
			},
		}

		actual := option.HasSpecifiedTasks()

		assert.True(actual)
	})

	t.Run("When has not specified tasks.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			specifiedTasks: []string{},
		}

		actual := option.HasSpecifiedTasks()

		assert.False(actual)
	})
}

func TestOptionImpl_ConfigPath(t *testing.T) {
	assert := assert2.New(t)

	option := OptionImpl{
		configPath: "./fixtures/dummy.yml",
	}

	actual := option.ConfigPath()

	expected := "./fixtures/dummy.yml"
	assert.Equal(expected, actual)
}

func TestOptionImpl_SpecifiedTasks(t *testing.T) {
	t.Run("When has once specified task.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			specifiedTasks: []string{
				"foo",
			},
		}

		actual := option.SpecifiedTasks()

		expected := 1
		assert.Len(actual, expected)

		expected2 := "foo"
		assert.Equal(expected2, actual[0])
	})

	t.Run("When has twice specified tasks.", func(t *testing.T) {
		assert := assert2.New(t)

		option := OptionImpl{
			specifiedTasks: []string{
				"foo",
				"bar",
			},
		}

		actual := option.SpecifiedTasks()

		expected := 2
		assert.Len(actual, expected)

		expected2 := "foo"
		assert.Equal(expected2, actual[0])

		expected3 := "bar"
		assert.Equal(expected3, actual[1])
	})
}

func TestParseOption(t *testing.T) {
	t.Run("When passing invalid flag.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"-invalid",
		}

		option, err := ParseOption(args)

		assert.Nil(option)

		assert.Error(err)

		expected := "flag provided but not defined: -invalid"
		assert.Contains(iobuffer.String(), expected)
	})

	t.Run("When passing help flag.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"-h",
		}

		option, err := ParseOption(args)

		assert.Nil(option)

		assert.Error(err)

		expected := "Usage: taskal"
		assert.Contains(iobuffer.String(), expected)
	})

	t.Run("When passing show tasks flags.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"-T",
		}

		option, err := ParseOption(args)

		expected := (*Option)(nil)
		assert.Implements(expected, option)

		assert.NoError(err)

		assert.True(option.WillBeShowTasks())
		assert.False(option.BeDryRun())

		expected2 := "taskal.yml"
		assert.Equal(expected2, option.ConfigPath())
	})

	t.Run("When passing dry run flags.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"-n",
		}

		option, err := ParseOption(args)

		expected := (*Option)(nil)
		assert.Implements(expected, option)

		assert.NoError(err)

		assert.False(option.WillBeShowTasks())
		assert.True(option.BeDryRun())

		expected2 := "taskal.yml"
		assert.Equal(expected2, option.ConfigPath())
	})

	t.Run("When passing config path flag.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"-c",
			"./fixtures/dummy.yml",
		}

		option, err := ParseOption(args)

		expected := (*Option)(nil)
		assert.Implements(expected, option)

		assert.NoError(err)

		assert.False(option.WillBeShowTasks())
		assert.False(option.BeDryRun())

		expected2 := "./fixtures/dummy.yml"
		assert.Equal(expected2, option.ConfigPath())
	})

	t.Run("When passing specified tasks.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"foo",
			"bar",
		}

		option, err := ParseOption(args)

		expected := (*Option)(nil)
		assert.Implements(expected, option)

		assert.NoError(err)

		expected2 := 2
		assert.Len(option.SpecifiedTasks(), expected2)

		expected3 := "foo"
		assert.Equal(expected3, option.SpecifiedTasks()[0])

		expected4 := "bar"
		assert.Equal(expected4, option.SpecifiedTasks()[1])

		expected5 := 0
		assert.Len(option.TaskArgs(), expected5)
	})

	t.Run("When passing specified tasks and task args.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		args := []string{
			"taskal",
			"foo",
			"bar",
			"--",
			"baz",
			"-buzz",
		}

		option, err := ParseOption(args)

		expected := (*Option)(nil)
		assert.Implements(expected, option)

		assert.NoError(err)

		expected2 := 2
		assert.Len(option.SpecifiedTasks(), expected2)

		expected3 := "foo"
		assert.Equal(expected3, option.SpecifiedTasks()[0])

		expected4 := "bar"
		assert.Equal(expected4, option.SpecifiedTasks()[1])

		expected5 := 2
		assert.Len(option.TaskArgs(), expected5)

		expected6 := "baz"
		assert.Equal(expected6, option.TaskArgs()[0])

		expected7 := "-buzz"
		assert.Equal(expected7, option.TaskArgs()[1])
	})
}
