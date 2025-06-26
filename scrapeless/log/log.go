package log

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/helper"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// FileName is the log file name
	fileName = "scrapeless.log"
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	maxSizeOfLog = 100
	// MaxBackups is the maximum number of old log files to retain.  The default
	maxBackupsOfLog = 5
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	maxAgeOfLog = 7
	// LogRootDir is the log file path
	logRootDir = "/var/log/scrapeless"
)

const (
	traceKey   = "trace-id"
	timeFormat = "2006-01-02T15:04:05Z"
)

var (
	logger zerolog.Logger
	lj     *lumberjack.Logger
)

func init() {
	logDir := helper.Coalesce(env.GetLogEnv().LogRootDir, logRootDir)

	filename := fmt.Sprintf("%s/%s", logDir, fileName)

	_, err := os.Stat(logDir)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(logDir, os.ModePerm)
	}

	// Set defaults
	setDefaults()

	// Get MultiWriter
	multi := getWriter(filename)

	// New logger
	logger = zerolog.New(multi).
		With().
		Timestamp().
		Caller().
		Logger()

	logger = logger.Hook(tracingHook{})
}

type tracingHook struct{}

func (t tracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceVal := ctx.Value(traceKey)
	if traceVal != nil {
		if val, ok := traceVal.(string); ok {
			e.Str(traceKey, val)
		}
	}
}

func archiveCurrentLog() error {
	logPath := lj.Filename
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		fmt.Printf("Log file does not exist: %s\n", logPath)
		return nil
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.999")

	dir := filepath.Dir(logPath)
	base := filepath.Base(logPath)
	ext := filepath.Ext(base)
	nameOnly := base[:len(base)-len(ext)]

	// New compressed file name
	gzFilename := fmt.Sprintf("%s-%s%s.gz", nameOnly, timestamp, ext)
	gzPath := filepath.Join(dir, gzFilename)
	gzFile, err := os.Create(gzPath)
	if err != nil {
		return fmt.Errorf("failed to create gzip file: %w", err)
	}
	defer gzFile.Close()

	// open the source file
	srcFile, err := os.Open(logPath)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer srcFile.Close()

	// Write gzip
	gzWriter := gzip.NewWriter(gzFile)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, srcFile)
	if err != nil {
		return fmt.Errorf("failed to compress log file: %w", err)
	}

	// Delete the original log file
	err = os.Remove(logPath)
	if err != nil {
		return fmt.Errorf("failed to remove original log file: %w", err)
	}

	fmt.Printf("Archived and removed: %s → %s\n", logPath, gzPath)
	return nil
}

func customFormatLevel(i interface{}) string {
	l, ok := i.(string)
	if ok {
		return l
	}
	return "unknown"
}

func getWriter(filename string) zerolog.LevelWriter {
	maxSize := helper.Coalesce(env.GetLogEnv().MaxSize, maxSizeOfLog)
	maxBackups := helper.Coalesce(env.GetLogEnv().MaxBackups, maxBackupsOfLog)
	maxAge := helper.Coalesce(env.GetLogEnv().MaxAge, maxAgeOfLog)

	// MultiWriter to write logs to both console and file simultaneously
	lj = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}

	// Set Lumberjack's Writer to zero log output
	consoleWriter := &zerolog.ConsoleWriter{
		Out:         os.Stdout,
		TimeFormat:  timeFormat,
		NoColor:     true,
		FormatLevel: customFormatLevel,
	}
	fileWriter := &zerolog.ConsoleWriter{
		Out:         lj,
		TimeFormat:  timeFormat,
		NoColor:     true,
		FormatLevel: customFormatLevel,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)

	return multi
}

func setDefaults() {
	zerolog.TimestampFieldName = "ts"
	zerolog.MessageFieldName = "msg"
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
	zerolog.TimeFieldFormat = timeFormat

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
}
