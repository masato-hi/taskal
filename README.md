# taskal
taskal is a simple and easy task runner.

[![Build Status](https://travis-ci.org/masato-hi/taskal.svg?branch=master)](https://travis-ci.org/masato-hi/taskal)

## Usage
```
$ taskal [options...] [tasks ...] -- [optional args ...]
```

### Options
```
  -T	Show all tasks.
  -c string
    	taskal -c [CONFIGFILE] (default "taskal.yml")
  -n	Do a dry run without executing actions.
```

### Example
```
$ cat taskal.yml
example: echo example
example2: echo example2

$ taskal example
[INFO][15:04:05] Execute task: example
[INFO][15:04:05] sh -c "echo example"
example
```

## Why use taskal?

It is because you only need to know YAML.


## Install
### go get
If you are a Golang developper/user; then execute go get.

```
$ go get -u github.com/masato-hi/taskal
```

## Config file format
Create taskal.yml in YAML format.

#### Define task
This defines the task whose task name is "build" and whose command is "echo build".
```
build: echo build
```

#### Define many tasks
This is add define the task whose task name is "test" and whose command is "echo test".
```
build: echo build
test: echo test
```

#### Define multi-line task
```
build: |
  if [ ! -e "/tmp/build.lock" ]; then
    touch "/tmp/build.lock"
    echo build
    rm "/tmp/build.lock"
  else
    exit 2
  fi
test: echo test
```

#### Define tasks to execute multiple commands
```
build: 
  - echo prepare build
  - echo build
  - echo clean build
```


#### Define dependent tasks
```
test: &test echo test
build:
  - *test
  - echo build
```

Of course, you can depend on tasks that execute multiple commands.

```
test: &test echo test
prepare: &prepare
  - echo do linter
  - *test
build: 
  - *prepare
  - echo build
  - echo clean build
```

#### Definition of hidden tasks
Tasks with an underscore at the starting of the name are not displayed in the task list.

```
test: &test echo test
_prepare: &prepare
  - echo do linter
  - *test
build: 
  - *prepare
  - echo build
  - echo clean build
```

```
$ taskal -T
All defined tasks:

build
test
```

#### Pass arguments to task (Only UNIX like OS)
Pass arguments after double-dash(`--`) and refer to `$@`.
```
build: echo build $@
```
```
$ taskal build -- additional arguments
[INFO][15:04:05] Execute task: build
[INFO][15:04:05] sh -c "echo build $@" -- additional arguments
build additional arguments
```

