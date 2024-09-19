package log

import (
	"fmt"
	"os"
	"sync/atomic"

	"r3_client/tools"
)

var (
	debug           atomic.Bool
	fileMaxSizeKb   = 10240 // log file max size before rotate in kilobytes
	filePath        = ""    // log file path
	filePathRotated string  // log file path, rotated
)

func SetDebug(v bool) {
	debug.Store(v)
}
func SetFilePath(v string) {
	filePath = v
	filePathRotated = fmt.Sprintf("%s.1", v)
}
func RotateIfNecessary() error {

	fi, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if int((fi.Size() / 1024)) < fileMaxSizeKb {
		return nil
	}

	// rotate log
	exists, err := tools.Exists(filePathRotated)
	if err != nil {
		return err
	}
	if exists {
		if err := os.Remove(filePathRotated); err != nil {
			return err
		}
	}
	if err := os.Rename(filePath, filePathRotated); err != nil {
		return err
	}
	if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
		return err
	}
	return nil
}
func Info(context string, message string) {
	write(3, context, message, nil, true)
}
func Warning(context string, message string, err error) {
	write(2, context, message, err, true)
}
func Error(context string, message string, err error) {
	write(1, context, message, err, true)
}

func write(levelRequested int, context string, message string, err error, writeToFile bool) {

	if levelRequested != 1 && !debug.Load() {
		return
	}

	// append error message, if available
	if err != nil {
		if message != "" {
			message = fmt.Sprintf("%s, %s", message, err.Error())
		} else {
			message = err.Error()
		}
	}
	message = fmt.Sprintf("%s %s %s\n", tools.GetTimeSql(), context, message)

	// log to CLI
	fmt.Print(message)

	// log to file if set
	if writeToFile && filePath != "" {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			write(1, "log", "failed to open log file", err, false)
			return
		}
		if _, err := f.Write([]byte(message)); err != nil {
			write(1, "log", "failed to write to log file", err, false)
			return
		}
		if err := f.Close(); err != nil {
			write(1, "log", "failed to close log file", err, false)
			return
		}
	}
}
