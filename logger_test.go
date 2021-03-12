package go_logger

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	NewLogger()
}

func TestLogger_Attach(t *testing.T) {

	logger := NewLogger()
	fileConfig := &FileConfig{
		Filename: "./test.log",
	}
	logger.Attach("file", LoggerLevelDebug, fileConfig)
	outputs := logger.outputs
	for _, outputLogger := range outputs {
		if outputLogger.Name != "file" {
			t.Error("file attach failed")
		}
	}
}

func TestLogger_Detach(t *testing.T) {

	logger := NewLogger()
	logger.Detach("console")

	outputs := logger.outputs

	if len(outputs) > 0 {
		t.Error("logger detach error")
	}
}

func TestLogger_LoggerLevel(t *testing.T) {

	logger := NewLogger()

	level := logger.LoggerLevel("emerGENCY")
	if level != LoggerLevelEmergency {
		t.Error("logger loggerLevel error")
	}
	level = logger.LoggerLevel("ALERT")
	if level != LoggerLevelAlert {
		t.Error("logger loggerLevel error")
	}
	level = logger.LoggerLevel("cRITICAL")
	if level != LoggerLevelCritical {
		t.Error("logger loggerLevel error")
	}
	level = logger.LoggerLevel("DEBUG")
	if level != LoggerLevelDebug {
		t.Error("logger loggerLevel error")
	}
	level = logger.LoggerLevel("ooox")
	if level != LoggerLevelDebug {
		t.Error("logger loggerLevel error")
	}
}

func TestLogger_loggerMessageFormat(t *testing.T) {

	loggerMsg := &loggerMessage{
		Timestamp:         time.Now().Unix(),
		TimestampFormat:   time.Now().Format("2006-01-02 15:04:05"),
		Millisecond:       time.Now().UnixNano() / 1e6,
		MillisecondFormat: time.Now().Format("2006-01-02 15:04:05.999"),
		Level:             LoggerLevelDebug,
		LevelString:       "debug",
		Body:              "logger console adapter test",
		File:              "console_test.go",
		Line:              77,
		Function:          "TestAdapterConsole_WriteJsonFormat",
	}

	format := "%millisecond_format% [%level_string%] [%file%:%line%] %body%"
	str := loggerMessageFormat(format, loggerMsg)

	fmt.Println(str)
}
