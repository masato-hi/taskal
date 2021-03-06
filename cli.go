package main

const (
	Succeeded = 0 + iota
	InvalidOption
	InvalidConfig
	UnreadConfig
	FailedExecute
)

type CLI interface {
	Run([]string) int
}

type CLIImpl struct {
}

var NewCLI = func() CLI {
	return &CLIImpl{}
}

func (c *CLIImpl) Run(args []string) int {
	option, err := ParseOption(args)
	if err != nil {
		return InvalidOption
	}

	buf, err := ReadConfig(option.ConfigPath())
	if err != nil {
		return UnreadConfig
	}

	config, err := ParseConfig(buf)
	if err != nil {
		return InvalidConfig
	}

	if option.WillBeShowTasks() {
		config.ShowAllDefinedTasks()
		return Succeeded
	}

	runner := NewRunner(option, config)
	if err := runner.Run(); err != nil {
		return FailedExecute
	}

	return Succeeded
}
