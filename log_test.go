// Package levelLog provides ...
package levelLog

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLevels(t *testing.T) {
	assert.True(t, LevelTrace&LevelErrorOnly != 0, "should not be 0")
	assert.True(t, LevelWarning&LevelTraceOnly == 0, "should not be 0")
}

func TestTurnOn(t *testing.T) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	re, we, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = we

	turnOnLogging(LevelWarning, nil)
	logger.Info.Output(2, "infoout")
	logger.Warning.Output(2, "warningout")
	logger.Error.Output(2, "errorout")
	we.Close()
	w.Close()

	out, _ := ioutil.ReadAll(r)
	errOut, _ := ioutil.ReadAll(re)
	os.Stdout = oldStdout
	os.Stderr = oldStderr
	outStr := string(out[:])
	errStr := string(errOut[:])
	assert.True(t, strings.Contains(outStr, "warningout"), "Should have warning printed")
	assert.True(t, strings.Contains(errStr, "errorout"), "Should have error printed")
	assert.False(t, strings.Contains(outStr, "infoout"), "Should not have info printed")
}
