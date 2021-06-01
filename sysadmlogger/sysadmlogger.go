/**
	* SYSADM Server1
	* @Author  Wayne Wang <net_use@bzhy.com>
	* @Copyright Bzhy Network
	* @HomePage http://www.sysadm.cn
	* @Version 0.21.03
	* Licensed under the Apache License, Version 2.0 (the "License");
	* you may not use this file except in compliance with the License.
	* You may obtain a copy of the License at
	* http://www.apache.org/licenses/LICENSE-2.0
	* Unless required by applicable law or agreed to in writing, software
	* distributed under the License is distributed on an "AS IS" BASIS,
	* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	* See the License for the specific language governing permissions and
	* limitations under the License.
	*       @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
	* @Modified May 03 2021
**/

package sysadmLogger

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// LogLevel: log level string which will output prefix to log message
var LogLevel = [7]string{"panic", "fatal", "error", "warn", "info", "debug", "trace"} 

// a interface for log writer 
// the field in sysadmServer.Engine.logWriter point to implementation of interface sysadmLogWriter
// other packages use this implementation to log log message

type SysadmLogWriter interface {
    stdoutWriter(string, string)
    accessWriter(string, string)
    errorWriter(string, string)
    setLogFormat(string)
    setLogLevel(string)
}

// SysadmLogger struct used to save parameters for logger, such as access log file error log file
type SysadmLogger struct {
	// accessLoggerFile: is the path of access log file when log access log to a file.
	accessLoggerFile string 
	// errorLoggerFile: is the path of error log file ,if logger error log to a file.
	errorLoggerFile  string 

	// accessFp: file descriptor of the access log file
	accessFp *os.File 
	// errorFp: file descriptor of the error log file
	errorFp  *os.File 

	// Logger for access log
	accessLogger *log.Logger 
	errorLogger  *log.Logger //Logger for error log
	stdoutLogger *log.Logger //Logger for stdout

	//log formate is text or json
	loggerFormat string
	//TODO
	DateFormat   string //set date formate

	// Allstdout: log message into access(or error) log file and stdout if this value is ture; 
	// or log message into access(or error) log file without stdout.
	// when access(or error) logger no initated,messages will be log to stdout even if this value is false 
	Allstdout bool   
	//Only log the `logLevel` severity or above.
	logLevel  log.Level
}

// Set global variables and its default value
var sysadmLogger = SysadmLogger{
	accessLoggerFile: "",
	errorLoggerFile:  "",

	accessFp: nil,
	errorFp:  nil,

	accessLogger: nil,
	errorLogger:  nil,
	stdoutLogger: nil,

	loggerFormat: "Text",
	DateFormat:   time.RFC3339, //Ref: https://studygolang.com/static/pkgdoc/pkg/time.htm#Time.Format

	Allstdout: true,
	logLevel:  5,
}

//Return a instance of SysadmLogger
func New() *SysadmLogger {
	return &sysadmLogger
}

// InitStdoutLogger initated a logger to logging log message to stdout
// accroding SysadmLogger.LoggerFormat set log formate and set SysadmLogger.stdoutLogger with the initated instance
func (sysadmLogger *SysadmLogger) InitStdoutLogger() (stdoutLogger *log.Logger, err error) {
	stdoutLogger = log.New()
	stdoutLogger.Out = os.Stdout
	stdoutLogger = sysadmLogger.SetFormat(stdoutLogger, "stdout")

	sysadmLogger.stdoutLogger = stdoutLogger
	if sysadmLogger.accessLogger == nil && sysadmLogger.errorLogger == nil {
		sysadmLogger.Allstdout = true
	}

	return stdoutLogger, nil

}

//Init logger instance  for access or error.
//before call this func, openLogFile(logType, logFile) should be called
func (sysadmLogger *SysadmLogger) InitLogger(logType string, toStdout bool) (logger *log.Logger, err error) {
	err = nil
	if strings.ToLower(logType) != "access" && strings.ToLower(logType) != "error" {
		err = fmt.Errorf("LogType must be access or error.You input is: %s", logType)
		return nil, err
	}

	if strings.ToLower(logType) == "access" {
		if sysadmLogger.accessFp == nil {
			err = fmt.Errorf("May be not set access log, you should call openLogFile(%s, logFile) before call InitLogger", logType)
			return nil, err
		}
		logger = log.New()
		logger.Out = sysadmLogger.accessFp
		logger = sysadmLogger.SetFormat(logger, logType)
		sysadmLogger.accessLogger = logger
		if toStdout {
			sysadmLogger.Allstdout = true
		} else {
			sysadmLogger.Allstdout = false
		}
		return logger, nil
	}

	if sysadmLogger.errorFp == nil {
		err = fmt.Errorf("May be not set error log, you should call openLogFile (%s, logFile) before call InitLogger", logType)
		return nil, err
	}

	logger = log.New()
	logger.Out = sysadmLogger.errorFp
	logger = sysadmLogger.SetFormat(logger, logType)
	sysadmLogger.errorLogger = logger

	return logger, nil
}

/*
* EndLogger function will be call by defer
* EndLogger will close file descriptor of access or error
* and reset logger of access or error to nil
 */
func (sysadmLogger *SysadmLogger) EndLogger(logType string) (err error) {
	err = nil
	var fp *os.File

	switch strings.ToLower(logType) {
	case "access":
		fp = sysadmLogger.accessFp
		if fp != nil {
			err = fp.Close()
		} else {
			err = fmt.Errorf("Access logger have closed")
		}

		if err == nil {
			sysadmLogger.accessFp = nil
			sysadmLogger.accessLoggerFile = ""
			sysadmLogger.accessLogger = nil
			if sysadmLogger.accessLogger == nil && sysadmLogger.errorLogger == nil {
				sysadmLogger.Allstdout = true
			}
		}
		break
	case "error":
		fp = sysadmLogger.errorFp
		if fp != nil {
			err = fp.Close()
		} else {
			err = fmt.Errorf("Error logger have closed")
		}

		if err == nil {
			sysadmLogger.errorFp = nil
			sysadmLogger.errorLogger = nil
			sysadmLogger.errorLoggerFile = ""
			if sysadmLogger.accessLogger == nil && sysadmLogger.errorLogger == nil {
				sysadmLogger.Allstdout = true
			}
		}
		break
	case "stdout":
		if sysadmLogger.accessLogger != nil || sysadmLogger.errorLogger != nil {
			err = fmt.Errorf("Access logger and Error logger should be end first")
		}

		if err == nil {
			sysadmLogger.stdoutLogger = nil
		}
		break
	default:
		err = fmt.Errorf("logType: %s is invalid", logType)
		break
	}

	return err
}

/*
* set log format to text or json and set loggerFormat
* the fields of the struct of loggerFormat refer to :https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter
 */
func (sysadmLogger *SysadmLogger) SetFormat(Logger *log.Logger, logType string) (logger *log.Logger) {
	if strings.ToLower(logType) == "access" || strings.ToLower(logType) == "error" {
		if strings.ToLower(sysadmLogger.loggerFormat) == "text" {
			Logger.SetFormatter(&log.TextFormatter{
				ForceColors:               false, //Ref: https://pkg.go.dev/github.com/sirupsen/logrus#pkg-functions
				DisableColors:             true,
				ForceQuote:                false,
				DisableQuote:              true,
				EnvironmentOverrideColors: true,
				DisableTimestamp:          false,
				FullTimestamp:             true,
				TimestampFormat:           sysadmLogger.DateFormat,
				DisableSorting:            true,
				DisableLevelTruncation:    true,
				PadLevelText:              true,
			})
		} else {
			Logger.SetFormatter(&log.JSONFormatter{
				TimestampFormat:  sysadmLogger.DateFormat,
				DisableTimestamp: false,
			})
		}
	} else {
		if strings.ToLower(sysadmLogger.loggerFormat) == "text" {
			Logger.SetFormatter(&log.TextFormatter{
				ForceColors:               true, //Ref: https://pkg.go.dev/github.com/sirupsen/logrus#pkg-functions
				DisableColors:             false,
				ForceQuote:                true,
				DisableQuote:              false,
				EnvironmentOverrideColors: true,
				DisableTimestamp:          false,
				FullTimestamp:             true,
				TimestampFormat:           sysadmLogger.DateFormat,
				DisableSorting:            true,
				DisableLevelTruncation:    true,
				PadLevelText:              true,
			})
		} else {
			Logger.SetFormatter(&log.JSONFormatter{
				TimestampFormat:  sysadmLogger.DateFormat,
				DisableTimestamp: false,
			})
		}
	}

	return Logger

}

/*
* set logger level to sysadmLogger.loggerLevel
 */
func (sysadmLogger *SysadmLogger) Setlevel(loggerLevel string, Logger *log.Logger) (logger *log.Logger) {

	switch strings.ToLower(loggerLevel) {
	case "panic":
		Logger.SetLevel(log.PanicLevel)
	case "fatal":
		Logger.SetLevel(log.FatalLevel)
	case "error":
		Logger.SetLevel(log.ErrorLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	case "info":
		Logger.SetLevel(log.InfoLevel)
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "trace":
		Logger.SetLevel(log.TraceLevel)
	default:
		Logger.SetLevel(log.DebugLevel)
	}

	return Logger
}


// according to logType, openLogFile set logFile to sysadmLogger.accessLoggerFile or sysadmLogger.errorLoggerFile
// and set file descriptor to accessFp or errorFp if logFile can be opened.
// to close the openned file on time, a defer function should be called following call this function if this return successful.
func (sysadmLogger *SysadmLogger) openLogFile(logType string, logFile string) (fp *os.File, err error) {

	err = nil
	if strings.ToLower(logType) != "access" && strings.ToLower(logType) != "error" {
		err = fmt.Errorf("LogType must be access or error.You input is: %s", logType)
		return nil, err
	}

	fp, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		err = fmt.Errorf("Open %s log file %s error: %s", logType, logFile, fmt.Sprintf("%s", err))
		return nil, err
	}

	if strings.ToLower(logType) == "access" {
		sysadmLogger.accessFp = fp
		sysadmLogger.accessLoggerFile = logFile
		_, err = sysadmLogger.InitLogger("access", sysadmLogger.Allstdout)
	} else {
		sysadmLogger.errorFp = fp
		sysadmLogger.errorLoggerFile = logFile
		_, err = sysadmLogger.InitLogger("error", sysadmLogger.Allstdout)
	}

	return fp, err
}

// Logging a message to Logger
// if the sysadmLogger.Allstdout ,then logging the log messages to stdout
func (sysadmLogger *SysadmLogger) LoggingLog(logType string, logLevel string, args ...interface{}) {

	var logger *log.Logger
	var tostdout bool
	var stdLogger *log.Logger

	tostdout = sysadmLogger.Allstdout
	logger = nil
	stdLogger = nil

	switch strings.ToLower(logType) {
	case "access":
		logger = sysadmLogger.accessLogger
		if logger == nil {
			tostdout = true
		}
	case "error":
		logger = sysadmLogger.errorLogger
		if logger == nil {
			tostdout = true
		}
	case "stdout":
		tostdout = false
		logger = sysadmLogger.stdoutLogger
	default:
		tostdout = false
		logger = sysadmLogger.stdoutLogger
	}
	

	if tostdout {
		stdLogger = sysadmLogger.stdoutLogger
	}

	if stdLogger != nil {
		stdLogger = sysadmLogger.Setlevel(logLevel, stdLogger)
	}

	if stdLogger != nil {
		level := stdLogger.GetLevel()
		if level != sysadmLogger.logLevel {
			sysadmLogger.Setlevel(LogLevel[sysadmLogger.logLevel], stdLogger)
		}
		sysadmLogger.SetFormat(stdLogger, "stdout")
	}

	if logger != nil {
		level := logger.GetLevel()
		if level != sysadmLogger.logLevel {
			sysadmLogger.Setlevel(LogLevel[sysadmLogger.logLevel], logger)
		}
		sysadmLogger.SetFormat(logger, "access")
	}
	switch strings.ToLower(logLevel) {
	case "panic":
		if logger != nil {
			logger.Panic(args...)
		}
		if stdLogger != nil {
			stdLogger.Panic(args...)
		}
		break
	case "fatal":
		if logger != nil {
			logger.Fatal(args...)
		}
		if stdLogger != nil {
			stdLogger.Fatal(args...)
		}
		break
	case "error":
		if logger != nil {
			logger.Error(args...)
		}
		if stdLogger != nil {
			stdLogger.Error(args...)
		}
		break
	case "warn":
		if logger != nil {
			logger.Warn(args...)
		}
		if stdLogger != nil {
			stdLogger.Warn(args...)
		}
		break
	case "info":
		if logger != nil {
			logger.Info(args...)
		}
		if stdLogger != nil {
			stdLogger.Info(args...)
		}
		break
	case "debug":
		if logger != nil {
			logger.Debug(args...)
		}
		if stdLogger != nil {
			stdLogger.Debug(args...)
		}
		break
	case "trace":
		if logger != nil {
			logger.Trace(args...)
		}
		if stdLogger != nil {
			stdLogger.Trace(args...)
		}
		break
	}
}

// checking level name  is valid. 
// if the level name  is invalid, then return `debug`
func checkLevelName(level string) string {

	var levelName = ""
	for _, value := range LogLevel {
		if strings.ToLower(level) == value {
			levelName = level
		}
	}

	if levelName == "" {
		levelName = "debug"
	}

	return levelName
}

// stdoutWriter is a implementation of the interface sysadmLogWriter
func (sysadmLogger *SysadmLogger) stdoutWriter(level string, msg string) {
    
	levelName := checkLevelName(level)
	
	sysadmLogger.LoggingLog("stdout", levelName, msg)
}

// accessWriter is a implementation of the interface sysadmLogWriter 
// for log message to access log file
func (sysadmLogger *SysadmLogger) accessWriter(level string, msg string) {

	levelName := checkLevelName(level)
	sysadmLogger.LoggingLog("access", levelName, msg)

}

// errorWriter is a implementation of the interface sysadmLogWriter
// for log message to error log file
func (sysadmLogger *SysadmLogger) errorWriter(level string, msg string) {
	
	levelName := checkLevelName(level)
	sysadmLogger.LoggingLog("error", levelName, msg)
}

// setLogFormat is a implementation of the interface sysadmLogWriter
// for setting log format to SysadmLogger
func (sysadmLogger *SysadmLogger) setLogFormat(format string) {
	if strings.ToLower(format) != "text" && strings.ToLower(format) != "json" {
		format = "json"
	}

	sysadmLogger.loggerFormat = format
}

// setLogLevel is a implementation of the interface sysadmLogWriter
// for setting logLevel to SysadmLogger
func (sysadmLogger *SysadmLogger) setLogLevel(level string) {

	logLevel := 5 
	for index, value := range LogLevel {
		if strings.ToLower(level) == value {
			logLevel = index
		}
	}
	
	sysadmLogger.logLevel = log.Level(logLevel)
}
