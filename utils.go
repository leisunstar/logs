package logs

import (
	"bytes"
	"os"
	"runtime"
	"strconv"
	"time"
)

func appendOut(level int, enable bool, logType, logFile string) error {
	if !enable {
		return nil
	}
	var o *out
	if logType == "console" {
		o = &out{level, levelMap[level], os.Stdout}
	} else {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		o = &out{level, levelMap[level], file}
	}
	log.outs = append(log.outs, o)
	return nil
}

func fmtMsg(prefix, content string) []byte {
	pc, _, line, _ := runtime.Caller(3)
	buffer := bytes.NewBufferString(time.Now().Format(dateTimeFmt))
	buffer.WriteString(" [" + prefix + "] ")
	buffer.WriteString(runtime.FuncForPC(pc).Name())
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(line))
	buffer.WriteString(" ")
	buffer.WriteString(content)
	buffer.WriteString("\n")
	return buffer.Bytes()
}
