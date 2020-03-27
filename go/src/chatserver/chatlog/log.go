/*
 * @Author: dzw
 * @Date: 2020-03-02 22:05:40
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-06 18:00:58
 */

package chatlog

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// log level
// Trace Debug Info Warning Error Fault
type logLevel uint8

const (
	// INVALID ...
	INVALID logLevel = iota
	// TRACE ...
	TRACE
	// DEBUG ...
	DEBUG
	// INFO ...
	INFO
	// WARNING ...
	WARNING
	// ERROR ...
	ERROR
	// FATAL ...
	FATAL
)

var (
	log = New("./chatlog/log.ini")
)

type chLogMsg struct {
	msg string
}

// Logger ...
type Logger struct {
	// file
	fileflag   bool
	filelevel  logLevel
	filepath   string
	maxsize    int64
	filelog    *os.File
	fileoutput io.Writer

	// console
	consoleflag   bool
	consolelevel  logLevel
	consoleoutput io.Writer

	// channel
	chLog  chan *chLogMsg // 异步记日志channel
	chSize int            // channel size
}

// default logger
func defaultLogger() *Logger {
	l := &Logger{
		fileflag:   false,
		filelevel:  TRACE,
		filepath:   "",
		maxsize:    0,
		filelog:    nil,
		fileoutput: nil,

		// console
		consoleflag:   true,
		consolelevel:  TRACE,
		consoleoutput: os.Stdout,

		// channel
		chLog:  nil,
		chSize: 0,
	}
	return l
}

// New Logger
func New(confpath string) *Logger {
	if confpath == "" {
		return defaultLogger()
	}
	l := &Logger{}
	l.setConfig(confpath)
	// channel init
	l.chLog = make(chan *chLogMsg, l.chSize)
	go l.processLogData()
	return l
}

// setConfig ...
func (l *Logger) setConfig(path string) {
	// 打开配置文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// 关闭配置文件
	defer file.Close()

	// 读取配置文件信息
	conf := &Config{}
	err = LoadConfig(path, conf)
	if err != nil {
		panic(err)
	}

	// console
	l.consoleflag = conf.ConsoleConfig.Flag
	if l.consoleflag {
		l.consolelevel = getLevel(conf.ConsoleConfig.Level)
		l.consoleoutput = os.Stdout
	}

	// file
	l.fileflag = conf.FileConfig.Flag
	if l.fileflag {
		l.filelevel = getLevel(conf.FileConfig.Level)
		l.filepath = conf.FileConfig.Filepath
		max, err := GetMaxBytesSize(conf.FileConfig.Maxsize)
		if err != nil {
			panic(err)
		}
		l.maxsize = max
		err = l.openLogFile()
		if err != nil {
			panic(err)
		}
		l.chSize = conf.Buffer.BufSize
	}
}

func (l *Logger) enableFile(level logLevel) bool {
	return l.fileflag && level >= l.filelevel
}

func (l *Logger) enableConsole(level logLevel) bool {
	return l.consoleflag && level >= l.consolelevel
}

// openLogFile ...
func (l *Logger) openLogFile() error {
	file, err := os.OpenFile(l.filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	l.filelog = file
	l.fileoutput = file
	return nil
}

// closeLogFile ...
func (l *Logger) closeLogFile() {
	if l.filelog != nil {
		l.filelog.Close()
	}
}

// checkLogFileSize ...
// true ----- write log file
// false ---- backup, create new log file
func (l *Logger) checkLogFileSize() bool {
	size, err := l.getLogFileSize()
	if err != nil {
		return false
	}

	return size < l.maxsize
}

func (l *Logger) getLogFileSize() (size int64, err error) {
	if l.filelog != nil {
		fileInfo, err := l.filelog.Stat()
		if err != nil {
			return 0, err
		}
		size = fileInfo.Size()
	}

	return size, err
}

func (l *Logger) backupLogFile() error {
	if l.filelog != nil {
		oldName := l.filelog.Name() // 带路径的name
		// fmt.Println("old name:", oldName)
		oldFileAbsPath := oldName

		if filepath.IsAbs(oldName) == false {
			oldNameAbs, err := filepath.Abs(oldName)
			if err != nil {
				return err
			}
			// fmt.Println("old name abs =", oldNameAbs, ",dir =", filepath.Dir(oldNameAbs))
			oldFileAbsPath = oldNameAbs
		}

		newFileAbsPath := oldFileAbsPath + time.Now().Format("20060102150405000")
		// fileInfo, err := l.filelog.Stat()
		// if err != nil {
		// 	return
		// }
		// fmt.Println("fileInfo.Name() = ", fileInfo.Name()) // 只有文件名
		l.closeLogFile()                                 // 关闭文件
		err := os.Rename(oldFileAbsPath, newFileAbsPath) //重命名文件
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

// processLogFile ...
// 检查是否需要备份日志文件
// 需要备份的话，备份日志文件，然后重新生成日志文件，设置新的*os.File和io.Writer
func (l *Logger) processLogFile() {
	// 需要备份日志文件
	if l.checkLogFileSize() == false {
		if err := l.backupLogFile(); err == nil {
			l.openLogFile() // 备份成功，重新打开文件
		}
	}
}

// Output ...
func (l *Logger) Output(level logLevel, format string, a ...interface{}) {
	levelStr := getLevelString(level)
	filename, funcname, linenum := getRuntimeInfo(3)
	msg := fmt.Sprintf(format, a...)                    // 格式化输入的日志
	now := time.Now().Format("2006-01-02 15:04:05.000") // 格式化的时间
	// 格式化日志信息
	data := fmt.Sprintf("[%s][%s](%s:%s:%d) %s\n", now, levelStr, filename, funcname, linenum, msg)

	if l.enableFile(level) {
		l.processLogFile()
		l.writeDataToChan(&chLogMsg{data})
		// fmt.Fprintf(l.fileoutput, "[%s][%s](%s:%s:%d) %s\n", now, levelStr, filename, funcname, linenum, msg) // 格式化日志输出
	}

	if l.enableConsole(level) {
		// fmt.Fprintf(l.consoleoutput, "[%s][%s](%s:%s:%d) %s\n", now, levelStr, filename, funcname, linenum, msg) // 格式化日志输出
		fmt.Fprint(l.consoleoutput, data)
	}
}

// 异步写日志
func (l *Logger) writeDataToChan(msg *chLogMsg) {
	select {
	case l.chLog <- msg: // 缓冲区满，会阻塞
	default: //缓冲区满,丢掉日志
	}
}

func (l *Logger) readDataFromChan() *chLogMsg {
	msg := <-l.chLog

	return msg
}

func (l *Logger) writeDataToFile(msg *string) {
	// msg := <-l.chLog
	fmt.Fprint(l.fileoutput, *msg)
}

func (l *Logger) processLogData() {
	for {
		data := l.readDataFromChan()
		l.writeDataToFile(&data.msg)
	}
}

// Trace ...
func (l *Logger) Trace(format string, a ...interface{}) {
	l.Output(TRACE, format, a...)
}

// Debug ...
func (l *Logger) Debug(format string, a ...interface{}) {
	l.Output(DEBUG, format, a...)
}

// Info ...
func (l *Logger) Info(format string, a ...interface{}) {
	l.Output(INFO, format, a...)
}

// Warning ...
func (l *Logger) Warning(format string, a ...interface{}) {
	l.Output(WARNING, format, a...)
}

// Error ...
func (l *Logger) Error(format string, a ...interface{}) {
	l.Output(ERROR, format, a...)
}

// Fatal ...
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.Output(FATAL, format, a...)
}

// Trace ...
func Trace(format string, a ...interface{}) {
	log.Output(TRACE, format, a...)
}

// Debug ...
func Debug(format string, a ...interface{}) {
	log.Output(DEBUG, format, a...)
}

// Info ...
func Info(format string, a ...interface{}) {
	log.Output(INFO, format, a...)
}

// Warning ...
func Warning(format string, a ...interface{}) {
	log.Output(WARNING, format, a...)
}

// Error ...
func Error(format string, a ...interface{}) {
	log.Output(ERROR, format, a...)
}

// Fatal ...
func Fatal(format string, a ...interface{}) {
	log.Output(FATAL, format, a...)
}

// getRuntimeInfo ...
func getRuntimeInfo(skip int) (filename, funcname string, linenum int) {
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {
		fmt.Println("getRuntimeInfo failed.")
		return
	}

	filename = path.Base(file) // 文件名
	linenum = line             //行号

	funcName := runtime.FuncForPC(pc).Name() // 函数名
	s := strings.Split(funcName, ".")
	if len(s) == 2 {
		funcname = s[1]
	} else {
		funcname = funcName
	}

	return
}

func getLevelString(level logLevel) string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// getLevel ...
func getLevel(level string) logLevel {
	level = strings.ToLower(level)
	switch level {
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INVALID
	}
}
