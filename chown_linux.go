package lumberjack

import (
	"os"
)

var osChown = os.Chown

func chown(name string, info os.FileInfo) error { _ = "STUB: not implemented"; return nil }
