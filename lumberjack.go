package lumberjack

import (
	"io"
	"os"
	"sync"
	"time"
)

const (
	backupTimeFormat = "2006-01-02T15-04-05.000"
	compressSuffix   = ".gz"
	defaultMaxSize   = 100
)

var _ io.WriteCloser = (*Logger)(nil)

type Logger struct {
	Filename string `json:"filename" yaml:"filename"`

	MaxSize int `json:"maxsize" yaml:"maxsize"`

	MaxAge int `json:"maxage" yaml:"maxage"`

	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	LocalTime bool `json:"localtime" yaml:"localtime"`

	Compress bool `json:"compress" yaml:"compress"`

	size int64
	file *os.File
	mu   sync.Mutex

	millCh    chan bool
	startMill sync.Once
}

var (
	currentTime = time.Now

	osStat = os.Stat

	megabyte = 1024 * 1024
)

func (l *Logger) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (l *Logger) Close() error { _ = "STUB: not implemented"; return nil }

func (l *Logger) close() error { _ = "STUB: not implemented"; return nil }

func (l *Logger) Rotate() error { _ = "STUB: not implemented"; return nil }

func (l *Logger) rotate() error { _ = "STUB: not implemented"; return nil }

func (l *Logger) openNew() error { _ = "STUB: not implemented"; return nil }

func backupName(name string, local bool) string { _ = "STUB: not implemented"; return "" }

func (l *Logger) openExistingOrNew(writeLen int) error { _ = "STUB: not implemented"; return nil }

func (l *Logger) filename() string { _ = "STUB: not implemented"; return "" }

func (l *Logger) millRunOnce() error { _ = "STUB: not implemented"; return nil }

func (l *Logger) millRun() { _ = "STUB: not implemented"; return }

func (l *Logger) mill() { _ = "STUB: not implemented"; return }

func (l *Logger) oldLogFiles() ([]logInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Logger) timeFromName(filename, prefix, ext string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (l *Logger) max() int64 { _ = "STUB: not implemented"; return 0 }

func (l *Logger) dir() string { _ = "STUB: not implemented"; return "" }

func (l *Logger) prefixAndExt() (prefix, ext string) { _ = "STUB: not implemented"; return "", "" }

func compressLogFile(src, dst string) (err error) { _ = "STUB: not implemented"; return nil }

type logInfo struct {
	timestamp time.Time
	os.FileInfo
}

type byFormatTime []logInfo

func (b byFormatTime) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (b byFormatTime) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (b byFormatTime) Len() int { _ = "STUB: not implemented"; return 0 }
