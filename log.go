package main

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
)

var (
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
)

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(Stdout, format, a...)
	fmt.Fprintln(Stdout)
}

func Debug(format string, a ...interface{}) {
	if DEBUG {
		fmt.Fprintf(Stdout, color.HiBlackString("[DEBUG]"))
		fmt.Fprintf(Stdout, color.HiWhiteString("[%s] ", TimeStamp()))
		fmt.Fprintf(Stdout, format, a...)
		fmt.Fprintln(Stdout)
	}
}

func Info(format string, a ...interface{}) {
	fmt.Fprintf(Stdout, color.HiCyanString("[INFO]"))
	fmt.Fprintf(Stdout, color.HiWhiteString("[%s] ", TimeStamp()))
	fmt.Fprintf(Stdout, format, a...)
	fmt.Fprintln(Stdout)
}

func Warn(format string, a ...interface{}) {
	fmt.Fprintf(Stderr, color.HiYellowString("[WARN]"))
	fmt.Fprintf(Stderr, color.HiWhiteString("[%s] ", TimeStamp()))
	fmt.Fprintf(Stderr, format, a...)
	fmt.Fprintln(Stderr)
}

func Error(format string, a ...interface{}) {
	fmt.Fprintf(Stderr, color.HiRedString("[ERROR]"))
	fmt.Fprintf(Stderr, color.HiWhiteString("[%s] ", TimeStamp()))
	fmt.Fprintf(Stderr, format, a...)
	fmt.Fprintln(Stderr)
}

func TimeStamp() string {
	return Now().Format("15:04:05")
}
