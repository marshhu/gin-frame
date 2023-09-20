package log

import "testing"

func TestDebug(t *testing.T) {
	Init(&Settings{
		Path:        DefaultPath,
		FileName:    DefaultFileName,
		Level:       "debug",
		LogCategory: DefaultLogCategory,
		Caller:      true,
	})

	Debug("test debug")
	Error("test error")
	Sync()
}
