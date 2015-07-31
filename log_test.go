// Package logger provides ...
package logger

import (
	// "fmt"
	"bitbucket.org/yanghu/logger/redis"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLevels(t *testing.T) {
	assert.True(t, LevelTrace&LevelErrorOnly != 0, "should not be 0")
	assert.True(t, LevelWarning&LevelTraceOnly == 0, "should not be 0")
}

func TestTurnOn(t *testing.T) {
	// use pipes to capture stdout and stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	re, we, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = we

	turnOnLogging(LevelWarning, nil)
	Info("info %d", 123)
	Warning("warning %d", 123)
	Error("error %d", 123)
	we.Close()
	w.Close()

	out, _ := ioutil.ReadAll(r)
	errOut, _ := ioutil.ReadAll(re)
	os.Stdout = oldStdout
	os.Stderr = oldStderr
	outStr := string(out[:])
	errStr := string(errOut[:])

	assert.True(t, strings.Contains(outStr, "warning 123"), "Should have warning printed")
	assert.True(t, strings.Contains(errStr, "error 123"), "Should have error")
	assert.False(t, strings.Contains(outStr, "info"), "Should not have info printed")
}

func TestRedis(t *testing.T) {
	rw, err := redis.NewWriter(redis.ADDRESS, "test_log", 100)
	if err != nil {
		log.Fatal(err)
	}
	turnOnLogging(LevelWarning, rw)
	Warning("Warning in redis")
	Error("Error in redis?")
	Trace("How about trace")
	logs := redis.ReadLog(rw.Conn, rw.Logname, rw.EntryLimit)
	assert.True(t, len(logs) == 2)
	assert.True(t, strings.Contains(logs[0], "Error"), "last in first out")

	//cleanup
	rw.FlushLog()
}
