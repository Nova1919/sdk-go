package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	// export SCRAPELESS_RUN_ID=21da69d8-e75a-4360-afe8-84aaecd68ea7
	Trace("hello world")
	Debug("hello world")
	Info("hello world")
	Warn("hello world")
	Error("hello world")

	Tracef("say: %s", "hello world")
	Debugf("say: %s", "hello world")
	Infof("say: %s", "hello world")
	Warnf("say: %s", "hello world")
	Errorf("say: %s", "hello world")

	for i := 0; i < 10000; i++ {
		Tracef("hello world: %d", i)
		Debugf("hello world: %d", i)
		Infof("hello world: %d", i)
		Warnf("hello world: %d", i)
		Errorf("hello world: %d", i)
	}

	// _ = archiveCurrentLog()
}
