package logs

import (
	"fmt"
	"github.com/yzw/conf"
	"io"
	"io/ioutil"
	"strings"
	"sync"
)

const (
	levelDebug = iota
	levelInfo
	levelWarning
	levelError
)

var (
	c           *config
	log         *logger
	dateTimeFmt = "2006-01-02 15:04:05"
	levelMap    = map[int]string{
		levelDebug:   "DEBUG  ",
		levelInfo:    "INFO   ",
		levelWarning: "WARNING",
		levelError:   "ERROR  ",
	}
)

type out struct {
	level    int
	levelStr string
	out      io.Writer
}

type logger struct {
	mu   *sync.Mutex
	outs []*out
}

func (l *logger) write(level int, format string, a ...interface{}) {
	l.mu.Lock()
	for _, o := range log.outs {
		if o.level == level {
			o.out.Write(fmtMsg(o.levelStr, fmt.Sprintf(format, a...)))
		}
	}
	l.mu.Unlock()
}

// Init
func Init(confFileName string) error {
	// init config
	c = &config{Level: "debug"}
	b, err := ioutil.ReadFile(confFileName)
	if err != nil {
		return err
	}
	err = conf.Unmarshal(b, c)
	if err != nil {
		return err
	}
	// init datetime format
	if c.DatetimeFmt != "" {
		dateTimeFmt = c.DatetimeFmt
	}
	// init level and logger
	level := levelDebug
	switch strings.ToLower(c.Level) {
	case "info":
		level = levelInfo
	case "warning":
		level = levelWarning
	case "error":
		level = levelError
	}
	log = &logger{mu: &sync.Mutex{}, outs: make([]*out, 0)}
	enableMap := map[int]bool{
		levelDebug:   c.DebugEnable,
		levelInfo:    c.InfoEnable,
		levelWarning: c.WarningEnable,
		levelError:   c.ErrorEnable,
	}
	typeMap := map[int]string{
		levelDebug:   c.DebugType,
		levelInfo:    c.InfoType,
		levelWarning: c.WarningType,
		levelError:   c.ErrorType,
	}
	outMap := map[int]string{
		levelDebug:   c.DebugOut,
		levelInfo:    c.InfoOut,
		levelWarning: c.WarningOut,
		levelError:   c.ErrorOut,
	}
	for i := level; i < 4; i++ {
		if appendOut(i, enableMap[i], typeMap[i], outMap[i]) != nil {
			return err
		}
	}
	return nil
}

func Debug(format string, a ...interface{}) {
	log.write(levelDebug, format, a...)
}

func Info(format string, a ...interface{}) {
	log.write(levelInfo, format, a...)
}

func Warning(format string, a ...interface{}) {
	log.write(levelWarning, format, a...)
}

func Error(format string, a ...interface{}) {
	log.write(levelError, format, a...)
}
