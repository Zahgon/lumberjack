//go:build !linux
// +build !linux

package lumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error { _ = "STUB: not implemented"; return nil }
