package logs

import (
	"fmt"
	"github.com/yzw/conf"
	"io"
	"io/ioutil"
	"strings"
	"time"
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
	logsChan = make(chan *logChan, 10000)
)

type logChan struct {
	level int
	msg   []byte
}

type out struct {
	level    int
	levelStr string
	out      io.Writer
}

type logger struct {
	outs map[int]*out
}

func (l *logger) put(level int, format string, a ...interface{}) {
	o, ok := l.outs[level]
	if ok {
		logsChan <- &logChan{level: level, msg: fmtMsg(o.levelStr, fmt.Sprintf(format, a...))}
	}
}

func (l *logger) write() {
	var chLog *logChan
	for {
		chLog = <-logsChan
		l.outs[chLog.level].out.Write(chLog.msg)
	}
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
	log = &logger{outs: make(map[int]*out, 0)}
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
		if appendOut(i, levelMap[i], enableMap[i], typeMap[i], outMap[i]) != nil {
			return err
		}
	}
	go log.write()
	return nil
}

func Close() {
	time.Sleep(1 * time.Second)
}

func Debug(format string, a ...interface{}) {
	log.put(levelDebug, format, a...)
}

func Info(format string, a ...interface{}) {
	log.put(levelInfo, format, a...)
}

func Warning(format string, a ...interface{}) {
	log.put(levelWarning, format, a...)
}

func Error(format string, a ...interface{}) {
	log.put(levelError, format, a...)
}
