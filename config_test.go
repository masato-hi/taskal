package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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

func TestConfigImpl_AddDefinedTask(t *testing.T) {
	t.Run("When add once defined task.", func(t *testing.T) {
		assert := assert2.New(t)

		config := ConfigImpl{}
		task := &DefinedTaskImpl{
			name: "foo",
		}

		config.AddDefinedTask(task)

		expected := 1
		assert.Equal(expected, len(config.definedTasks))

		expected2 := "foo"
		assert.Equal(expected2, config.definedTasks[0].Name())
	})

	t.Run("When add twice defined task.", func(t *testing.T) {
		assert := assert2.New(t)

		config := ConfigImpl{}
		task := &DefinedTaskImpl{
			name: "foo",
		}
		task2 := &DefinedTaskImpl{
			name: "bar",
		}

		config.AddDefinedTask(task)
		config.AddDefinedTask(task2)

		expected := 2
		assert.Equal(expected, len(config.definedTasks))

		expected2 := "foo"
		assert.Equal(expected2, config.definedTasks[0].Name())

		expected3 := "bar"
		assert.Equal(expected3, config.definedTasks[1].Name())
	})
}

func TestConfigImpl_DefinedTasks(t *testing.T) {
	t.Run("When added once defined task.", func(t *testing.T) {
		assert := assert2.New(t)

		config := ConfigImpl{
			definedTasks: []DefinedTask{
				&DefinedTaskImpl{
					name: "foo",
				},
			},
		}

		expected := 1
		assert.Equal(expected, len(config.definedTasks))

		expected2 := "foo"
		assert.Equal(expected2, config.definedTasks[0].Name())
	})

	t.Run("When added twice defined task.", func(t *testing.T) {
		assert := assert2.New(t)

		config := ConfigImpl{
			definedTasks: []DefinedTask{
				&DefinedTaskImpl{
					name: "foo",
				},
				&DefinedTaskImpl{
					name: "bar",
				},
			},
		}

		expected := 2
		assert.Equal(expected, len(config.definedTasks))

		expected2 := "foo"
		assert.Equal(expected2, config.definedTasks[0].Name())

		expected3 := "bar"
		assert.Equal(expected3, config.definedTasks[1].Name())
	})
}

func TestConfigImpl_ShowAllDefinedTasks(t *testing.T) {
	t.Run("When added once defined task.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		config := ConfigImpl{
			definedTasks: []DefinedTask{
				&DefinedTaskImpl{
					name: "foo",
				},
			},
		}

		config.ShowAllDefinedTasks()

		expected := "All defined tasks:\n\nfoo\n"
		assert.Equal(expected, iobuffer.String())
	})

	t.Run("When added twice defined task.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		config := ConfigImpl{
			definedTasks: []DefinedTask{
				&DefinedTaskImpl{
					name: "foo",
				},
				&DefinedTaskImpl{
					name: "bar",
				},
			},
		}

		config.ShowAllDefinedTasks()

		expected := "All defined tasks:\n\nfoo\nbar\n"
		assert.Equal(expected, iobuffer.String())
	})
}

func TestReadConfig(t *testing.T) {
	t.Run("When config file exists.", func(t *testing.T) {
		assert := assert2.New(t)

		path := "./fixtures/dummy.yml"

		actual, err := ReadConfig(path)

		expected := "foo:\n  - echo bar\n  - echo baz\n"
		assert.Equal(expected, actual)

		assert.Nil(err)
	})

	t.Run("When config file not exists.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		path := "./fixtures/dummy-notexists.yml"

		actual, err := ReadConfig(path)

		expected := ""
		assert.Equal(expected, actual)

		assert.Error(err)

		actual2 := err.Error()
		expected2 := "open ./fixtures/dummy-notexists.yml: no such file or directory"
		assert.Equal(expected2, actual2)

		expected3 := "[ERROR][15:04:05] Config file read error. path: ./fixtures/dummy-notexists.yml\n"
		assert.Equal(expected3, iobuffer.String())
	})
}

func TestParseConfig(t *testing.T) {
	originNewDefinedTask := NewDefinedTask

	var restoreOriginFunc = func() {
		NewDefinedTask = originNewDefinedTask
	}

	defer restoreOriginFunc()

	t.Run("When passing invalid yaml.", func(t *testing.T) {
		iobuffer.Reset()

		assert := assert2.New(t)

		buf := "invalid yaml"
		actual, err := ParseConfig(buf)

		assert.Nil(actual)

		assert.Error(err)

		expected := "[ERROR][15:04:05] yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `invalid...` into main.Document\n"
		assert.Equal(expected, iobuffer.String())
	})

	t.Run("When passing valid yaml.", func(t *testing.T) {
		t.Run("Only has task name starting underscore.", func(t *testing.T) {
			assert := assert2.New(t)

			buf := "_foo: echo foo"
			actual, err := ParseConfig(buf)

			expected := (*Config)(nil)
			assert.Implements(expected, actual)

			assert.Nil(err)

			expected2 := 0
			assert.Equal(expected2, len(actual.DefinedTasks()))
		})

		t.Run("Has once task.", func(t *testing.T) {
			assert := assert2.New(t)

			buf := "_foo: echo _foo\nfoo: echo foo"
			actual, err := ParseConfig(buf)

			expected := (*Config)(nil)
			assert.Implements(expected, actual)

			assert.Nil(err)

			expected2 := 1
			assert.Equal(expected2, len(actual.DefinedTasks()))

			expected3 := "foo"
			assert.Equal(expected3, actual.DefinedTasks()[0].Name())

			expected4 := []string{
				"echo foo",
			}
			assert.Equal(expected4, actual.DefinedTasks()[0].Commands())
		})

		t.Run("Has twice task.", func(t *testing.T) {
			assert := assert2.New(t)

			buf := "foo: echo foo\nbar: echo bar"
			actual, err := ParseConfig(buf)

			expected := (*Config)(nil)
			assert.Implements(expected, actual)

			assert.Nil(err)

			expected2 := 2
			assert.Equal(expected2, len(actual.DefinedTasks()))

			expected3 := "bar"
			assert.Equal(expected3, actual.DefinedTasks()[0].Name())

			expected4 := []string{
				"echo bar",
			}
			assert.Equal(expected4, actual.DefinedTasks()[0].Commands())

			expected5 := "foo"
			assert.Equal(expected5, actual.DefinedTasks()[1].Name())

			expected6 := []string{
				"echo foo",
			}
			assert.Equal(expected6, actual.DefinedTasks()[1].Commands())
		})

		t.Run("Has nested task.", func(t *testing.T) {
			assert := assert2.New(t)

			buf := "foo: echo foo\n" +
				"bar: echo bar\n" +
				"baz:\n" +
				"  - echo foo\n" +
				"  - echo bar\n" +
				"  -\n" +
				"    - echo baz\n" +
				"    - echo buzz\n" +
				""
			actual, err := ParseConfig(buf)

			expected := (*Config)(nil)
			assert.Implements(expected, actual)

			assert.Nil(err)

			expected2 := 3
			assert.Equal(expected2, len(actual.DefinedTasks()))

			expected3 := "bar"
			assert.Equal(expected3, actual.DefinedTasks()[0].Name())

			expected4 := []string{
				"echo bar",
			}
			assert.Equal(expected4, actual.DefinedTasks()[0].Commands())

			expected5 := "baz"
			assert.Equal(expected5, actual.DefinedTasks()[1].Name())

			expected6 := []string{
				"echo foo",
				"echo bar",
				"echo baz",
				"echo buzz",
			}
			assert.Equal(expected6, actual.DefinedTasks()[1].Commands())

			expected7 := "foo"
			assert.Equal(expected7, actual.DefinedTasks()[2].Name())

			expected8 := []string{
				"echo foo",
			}
			assert.Equal(expected8, actual.DefinedTasks()[2].Commands())
		})
	})
}
