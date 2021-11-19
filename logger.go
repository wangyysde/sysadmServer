// sysadmServer
// @Author  Wayne Wang <net_use@bzhy.com>
// @Copyright Bzhy Network
// @HomePage http://www.sysadm.cn
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
//  @Modified on Sep 19 2021

package sysadmServer

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/wangyysde/sysadmLog"
)

type LogFields map[string]interface{}

type LoggerConfig struct {
	// DefaultLogger is a instance of logger. We use defaultLogger to log log message
	// when AccessLogger  or ErrorLogger is nil
	DefaultLogger *log.Logger

	// AccessLogger is a instance of logrus
	// and this instance if for logging access log
	AccessLogger *log.Logger

	// ErrorLogger is a instance of logrus
	// and this instance if for logging error log
	ErrorLogger *log.Logger

	// Kind specifies the format of the log where be log to
	// kind is one of text or json
	Kind string

	// AccessLogFile records the path of log file for access
	// if the access log and error log log into difference files
	AccessLogFile string

	// ErrorLogFile records the path of log file for error
	// if the access log and error log log into difference files
	ErrorLogFile string

	// Level specifies which level log will be logged
	Level string

	// SplitAccessAndError identify if log access log and error log
	// into difference io.Writer. 
	// Logs will be log into  AccessLogger if SplitAccessAndError is false
	// otherwise access logs  will be log into AccessLogger and error logs will be log into ErrorLogger.
	SplitAccessAndError bool

	// Specifies the format of the log timestamp,like: "2021/09/02 - 15:04:05"
	TimeStampFormat string

	// Flag for whether to log caller info (off by default)
	ReportCaller bool

	// SkipPaths is a url path array which logs are not written.
	// Optional.
	SkipPaths []string
}

var	Levels = []string{
		"trace",
		"debug",
		"info",
		"warning",
		"error",
		"fatal",
		"panic",
	}

var TimestampFormat = map[string]string{
	"ANSIC":       	"Mon Jan _2 15:04:05 2006",
    "UnixDate":     "Mon Jan _2 15:04:05 MST 2006",
	"RubyDate":     "Mon Jan 02 15:04:05 -0700 2006",
   	"RFC822":      	"02 Jan 06 15:04 MST",
    "RFC822Z":      "02 Jan 06 15:04 -0700",  // 使用数字表示时区的RFC822
    "RFC850":       "Monday, 02-Jan-06 15:04:05 MST",
    "RFC1123":      "Mon, 02 Jan 2006 15:04:05 MST",
    "RFC1123Z":     "Mon, 02 Jan 2006 15:04:05 -0700", // 使用数字表示时区的RFC1123
    "RFC3339":      "2006-01-02T15:04:05Z07:00",
    "RFC3339Nano":  "2006-01-02T15:04:05.999999999Z07:00",
    "Kitchen":      "3:04PM",
    "Stamp":        "Jan _2 15:04:05",
    "StampMilli":   "Jan _2 15:04:05.000",
    "StampMicro":   "Jan _2 15:04:05.000000",
    "StampNano":    "Jan _2 15:04:05.000000000",
	"DateTime":     "2006-01-02 15:04:05",
}


var LoggerConfigVar = LoggerConfig{
	DefaultLogger:       nil,
	AccessLogger:        nil,
	ErrorLogger:         nil,
	Kind:                "text",
	AccessLogFile:       "",
	ErrorLogFile:        "",
	Level: 				 "debug",
	SplitAccessAndError: false,
	TimeStampFormat:     TimestampFormat["RFC3339"],
	ReportCaller: 		 true,			
	SkipPaths:           nil,
}


var textFormatter = &log.TextFormatter{
	ForceColors:               false,
	DisableColors:             false,
	ForceQuote:                false,
	DisableQuote:              true,
	EnvironmentOverrideColors: true,
	DisableTimestamp:          false,                                                                                                                                                                                                      
	FullTimestamp:             true,
	TimestampFormat:           TimestampFormat["RFC3339"],
	DisableSorting:            true,
	DisableLevelTruncation:    true,
	PadLevelText:              false,
	QuoteEmptyFields: 		   true,
}

var jsonFormatter = &log.JSONFormatter{
	TimestampFormat:		   TimestampFormat["RFC3339"],
	DisableTimestamp: 		   false,
	DisableHTMLEscape: 		   true,
}


func init() {
	logger := LoggerConfigVar.DefaultLogger
	if logger == nil {
		logger = log.New()
	}

	LoggerConfigVar.DefaultLogger = logger
	logger.Out = DefaultWriter
	setLoggerLevel()
	setLoggerKind()

}

// SetLogLevel  set the value  of LoggerConfigVar.Level to "debug" if the value of it is ""
// Then SetLogLevel  sets the the levels for DefaultLoger, AccessLogger and ErrorLogger
//  This function should be called during LoggerConfigVar.Level is setting and a new logger is initating.
func setLoggerLevel() {

	if _, err := log.ParseLevel(LoggerConfigVar.Level); err != nil {
		LoggerConfigVar.Level = "debug"
	}

	loggerLevel, _ := log.ParseLevel(LoggerConfigVar.Level)

	if LoggerConfigVar.DefaultLogger != nil {
		LoggerConfigVar.DefaultLogger.SetLevel(loggerLevel)
	}

	if LoggerConfigVar.AccessLogger != nil {
		LoggerConfigVar.AccessLogger.SetLevel(loggerLevel)
	}

	if LoggerConfigVar.ErrorLogger != nil {
		LoggerConfigVar.ErrorLogger.SetLevel(loggerLevel)
	}

}

// SetLogLevel set the vavlue of level to LoggerConfigVar and set the value of level of Loggers
func SetLogLevel(level string) error {
	if strings.TrimSpace(level) == "" {
		return fmt.Errorf("The value of log level is null")
	}

	found := false
	
	for _, value := range Levels {
		if strings.ToLower(level) == value {
			LoggerConfigVar.Level = value
			found = true
			setLoggerLevel()
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("The value of log level: %s is invalid",level)
}

func (*Engine) SetLogLevel(level string) error {
	return SetLogLevel(level)
}


// setLoggerKind  set the value  of LoggerConfigVar.Kind to "text" if the value of it is ""
// Then setLoggerKind  sets the the formatter for DefaultLoger, AccessLogger and ErrorLogger
//  This function should be called during LoggerConfigVar.kind is setting and a new logger is initating.
func setLoggerKind()  {
	if strings.ToLower(LoggerConfigVar.Kind) != "text" && strings.ToLower(LoggerConfigVar.Kind) != "json" {
		LoggerConfigVar.Kind = "text"
	}

	if logger := LoggerConfigVar.DefaultLogger; logger != nil {
		if strings.ToLower(LoggerConfigVar.Kind) == "text" {
			formatter := *textFormatter
			formatter.DisableColors = false 
			logger.SetFormatter(&formatter)
		} else {
			formatter := *jsonFormatter
			logger.SetFormatter(&formatter)
		}
	} 
	
	if logger := LoggerConfigVar.AccessLogger; logger != nil {
		if strings.ToLower(LoggerConfigVar.Kind) == "text" {
			logger.SetFormatter(textFormatter)
		} else {
			logger.SetFormatter(jsonFormatter)
		}
	}

	if logger := LoggerConfigVar.ErrorLogger; logger != nil {
		if strings.ToLower(LoggerConfigVar.Kind) == "text" {
			logger.SetFormatter(textFormatter)
		} else {
			logger.SetFormatter(jsonFormatter)
		}
	}
}

// SetLogLevel set the vavlue of Kind to LoggerConfigVar and set the value of Kind of Loggers
func SetLoggerKind(kind string) error {
	if strings.ToLower(kind) != "text" && strings.ToLower(kind) != "json" {
		return fmt.Errorf("The kind(%s) of log is invalid", kind)
	}

	LoggerConfigVar.Kind = kind

	setLoggerKind()

	return nil
}

func (*Engine) SetLoggerKind(kind string) error {
	return SetLoggerKind(kind)
}

// SetAccessLogFile set file to access log file and then open the file for access logger output.
// This function returns func (fd *os.File) should be called by defer following this function  calling to 
// close the file descriptor which opened by this function. *os.File: the file descriptor by this function opened.
func SetAccessLogFile(file string) (func (fd *os.File) error, *os.File, error) {
	if strings.TrimSpace(file) == "" {
		return nil, nil, fmt.Errorf("The length of access log file should be larger 1")
	}

	file = strings.TrimSpace(file)
	fp, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return nil, nil, fmt.Errorf("Open log file %s error: %s", file, fmt.Sprintf("%s", err))
	}

	logger := log.New()
	logger.Out = fp
	if LoggerConfigVar.AccessLogger != nil {
		oldfp := LoggerConfigVar.AccessLogger.Out
		switch v := oldfp.(type) {
		case *os.File:
			_ = v.Close()
		default:
			LoggerConfigVar.AccessLogger.Out = nil
		}
	}

	LoggerConfigVar.AccessLogger = logger

	setLoggerLevel()
	setLoggerKind()

	retFun :=  func(fd *os.File) error {
		err := fd.Close()
		LoggerConfigVar.AccessLogger = nil
		LoggerConfigVar.AccessLogFile = ""
		return err
	}

	return retFun, fp, nil
}

// (*Engine)SetAccessLogFile is a method to SetAccessLogFile function
func (*Engine)SetAccessLogFile(file string) (func (fd *os.File) error, *os.File, error) {
	return SetAccessLogFile(file)
}


// SetErrorLogFile set file to error log file and then open the file for error logger output.
// This function returns func (fd *os.File) should be called by defer following this function  calling to 
// close the file descriptor which opened by this function. *os.File: the file descriptor by this function opened.
func SetErrorLogFile(file string) (func (fd *os.File) error, *os.File, error) {
	if strings.TrimSpace(file) == "" {
		return nil, nil, fmt.Errorf("The length of error log file should be larger 1")
	}

	file = strings.TrimSpace(file)
	fp, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return nil, nil, fmt.Errorf("Open log file %s error: %s", file, fmt.Sprintf("%s", err))
	}

	logger := log.New()
	logger.Out = fp
	if LoggerConfigVar.ErrorLogger != nil {
		oldfp := LoggerConfigVar.ErrorLogger.Out
		switch v := oldfp.(type) {
		case *os.File:
			_ = v.Close()
		default:
			LoggerConfigVar.ErrorLogger.Out = nil
		}
	}

	LoggerConfigVar.ErrorLogger = logger

	setLoggerLevel()
	setLoggerKind()

	retFun :=  func(fd *os.File) error {
		err := fd.Close()
		LoggerConfigVar.ErrorLogger = nil
		LoggerConfigVar.ErrorLogFile = ""
		return err
	}

	return retFun, fp, nil
}

// (*Engine)SetErrorLogFile is a method to SetErrorLogFile function
func (*Engine)SetErrorLogFile(file string) (func (fd *os.File) error, *os.File, error) {
	return SetErrorLogFile(file)
}

// SetIsSplitLog set IsSplitLog  to Logger configuration
func SetIsSplitLog(isSplit bool){
	if isSplit {
		if LoggerConfigVar.AccessLogger == nil || LoggerConfigVar.ErrorLogger == nil {
			Log("You try to set SplitAccessAndError to true, but access log or error logger have not opened.", "warning")
		}
	} else {
		if LoggerConfigVar.AccessLogger == nil {
			Log("You try to set SplitAccessAndError to false, but access logger have not opened. All log message will be log to defaultOutput and error logger", "warning")
		}
	}
	LoggerConfigVar.SplitAccessAndError = isSplit
}

func (*Engine)SetIsSplitLog(isSplit bool){
	LoggerConfigVar.SplitAccessAndError = isSplit
}

// SetTimestampFormat set timeStampFormat to the LoggerConfig and then set it to all the Loggers
// The value of format is one of constants of time package and "DateTime"
func SetTimestampFormat(format string) error {
	if strings.TrimSpace(format) == "" {
		return fmt.Errorf("The length of format should be larger 1")
	}

	for k, v := range TimestampFormat {
		if strings.ToLower(k) == strings.ToLower(format) {
			LoggerConfigVar.TimeStampFormat = v
			textFormatter.TimestampFormat = v
			jsonFormatter.TimestampFormat = v
			fmt.Printf("v:%s\n",v)
			setLoggerKind()
			return nil
		}
	}

	return fmt.Errorf("The TimeStampFormat(%s) is invalid.",format)
}

// (*Engine)SetTimestampFormat is the method to SetTimestampFormat function
func  (*Engine)SetTimestampFormat(format string) error {
	return SetTimestampFormat(format)
}

// if isDisable is false, then timestamp message will be added to log messages automatically.
// Otherwise no timestamp will be added.
func DisableTimestamp(isDisable bool) {
	textFormatter.DisableTimestamp = isDisable
	jsonFormatter.DisableTimestamp = isDisable
	setLoggerKind()	
}

// (*Engine)DisableTimestamp is the method to DisableTimestamp function
func (*Engine)DisableTimestamp(isDisable bool){
	DisableTimestamp(isDisable)
}

// SetReportCaller sets ReportCaller of LoggerConfig to true or false.
// if the value of ReportCaller is true, then callers name and the file path information which the caller in will be
// added to log messages. 
func SetReportCaller(isReportCaller bool){
	LoggerConfigVar.ReportCaller = isReportCaller

	if LoggerConfigVar.DefaultLogger != nil {
		LoggerConfigVar.DefaultLogger.ReportCaller = isReportCaller
	}

	if LoggerConfigVar.AccessLogger != nil {
		LoggerConfigVar.AccessLogger.ReportCaller = isReportCaller
	}

	if LoggerConfigVar.ErrorLogger != nil {
		LoggerConfigVar.ErrorLogger.ReportCaller = isReportCaller
	}
	
}

// (*Engine)SetReportCaller is the method to SetReportCaller function
func (*Engine) SetReportCaller(isReportCaller bool){
	SetReportCaller(isReportCaller)
}

// WriteLogWithFields write fields to the logger. logLevel will be set to "info" when its value is ""
// the value of logLevel will be set to "error" if fields["ErrorMessage"] is not ""
// fields["ErrorMessage"] will write to logger when  fields["ErrorMessage"] and  fields["Message"] are not ""
// the value of fields["Message"] will be igored.
func WriteLogWithFields(logger *log.Logger, fields LogFields, logLevel string){
	if logger == nil {
		return 
	}

	if strings.TrimSpace(logLevel) == "" {
		logLevel = "info"
	}

	found := false

	for _,value := range Levels {
		if strings.ToLower(strings.TrimSpace(logLevel)) == value {
			found = true
		}
	}

	errormsg := ""
	if v,ok := fields["ErrorMessage"]; ok {
		errormsg = v.(string)
	}

	if strings.TrimSpace(errormsg) != "" && !found { 
		logLevel = "error"
	} else if !found { 
		logLevel = "info"
	}
 
	switch strings.ToLower(logLevel){
	case "trace":
		logger.WithFields(log.Fields(fields)).Trace("")
	case "debug":
		logger.WithFields(log.Fields(fields)).Debug("")
	case "info":
		logger.WithFields(log.Fields(fields)).Info("")
	case "warning":
		logger.WithFields(log.Fields(fields)).Warn("")
	case "error":
		logger.WithFields(log.Fields(fields)).Error("")
	case "fatal":
		fd := logger.Out
		switch v := fd.(type) {
		case *os.File:
			_ = v.Close()
		}
		logger.WithFields(log.Fields(fields)).Fatal("") 
	case "panic":
		logger.WithFields(log.Fields(fields)).Panic("")  
		panic("")
	}
}

// (*Engine)WriteLogWithFields is the method to WriteLogWithFields function
func (*Engine)WriteLogWithFields(logger *log.Logger, fields LogFields, logLevel string){
	WriteLogWithFields(logger, fields,logLevel)
}

// LoggerWithConfig instance a Logger middleware with config.
func Logger() HandlerFunc {
	
	notlogged := LoggerConfigVar.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {

		fields := make(LogFields)

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		
		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {

			fields["Request"] =  c.Request
			fields["Keys"] = c.Keys
			fields["ClientIP"] = c.ClientIP()
			fields["Method"] = c.Request.Method
			fields["StatusCode"] = c.Writer.Status()
			fields["ErrorMessage"] =  c.Errors.ByType(ErrorTypePrivate).String()
			fields["BodySize"] = c.Writer.Size()
			if raw != "" {
				path = path + "?" + raw
			}
			fields["Path"] = path

			LogWithFields(fields,"info")
		}
	}
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {

		fields := make(LogFields)

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		
		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {

			fields["Request"] =  c.Request
			fields["Keys"] = c.Keys
			fields["ClientIP"] = c.ClientIP()
			fields["Method"] = c.Request.Method
			fields["StatusCode"] = c.Writer.Status()
			fields["ErrorMessage"] =  c.Errors.ByType(ErrorTypePrivate).String()
			fields["BodySize"] = c.Writer.Size()
			if raw != "" {
				path = path + "?" + raw
			}
			fields["Path"] = path

			logger := log.New()
			logger.Out = out
			loggerLevel, _ := log.ParseLevel("debug")
			logger.SetLevel(loggerLevel)
			formatter := *textFormatter
			formatter.DisableColors = false 
			logger.SetFormatter(&formatter)
			WriteLogWithFields(logger, fields, "info")
		}
	}
}

// WriteLog write message to the logger. logLevel will be set to "error" when its value is ""
func WriteLog(logger *log.Logger, message string, logLevel string){
	if logger == nil {
		return 
	}

	if strings.TrimSpace(logLevel) == "" {
		logLevel = "info"
	}

	found := false

	for _,value := range Levels {
		if strings.ToLower(strings.TrimSpace(logLevel)) == value {
			found = true
		}
	}

	if !found {
		logLevel = "error"
	}

	switch strings.ToLower(logLevel){
	case "trace":
		logger.Trace(message)
	case "debug":
		logger.Debug(message)
	case "info":
		logger.Info(message)
	case "warning":
		logger.Warn(message)
	case "error":
		logger.Error(message)
	case "fatal":
		fd := logger.Out
		switch v := fd.(type) {
		case *os.File:
			_ = v.Close()
		}
		logger.Fatal(message) 
	case "panic":
		logger.Panic(message)  
		panic("")
	}
}

// (*Engine)WriteLog is the method to WriteLog function
func (*Engine)WriteLog(logger *log.Logger, message string, logLevel string){
	WriteLog(logger,message,logLevel)
}

// LogWithFields log LogFields to DefaultLogger, AccessLogger and ErrorLogger. 
// if SplitAccessAndError is true, then log info log message to AccessLogger and other log message to ErrorLogger.
// otherwise  all log message will be log to AccessLogger.  If  SplitAccessAndError is true but ErrorLogger is nil 
// then all log message will be log to AccessLogger and a additional warning log message will be log to AccessLogger.
// If SplitAccessAndError is false but AccessLogger is nil then all log message will be log to ErrorLogger and a additional 
// warning log message will be log to ErrorLogger.
func LogWithFields(fields LogFields, logLevel string){
	
	if strings.TrimSpace(logLevel) == "" {
		if _,ok := fields["ErrorMessage"]; ok {
			logLevel = "error"
		} else {
			logLevel = "info"
		}
	}

	found := false
	for _,value := range Levels {
		if strings.ToLower(logLevel) == strings.ToLower(value) {
			found = true
		}
	}

	if !found {
		if _,ok := fields["ErrorMessage"]; ok {
			logLevel = "error"
		} else {
			logLevel = "info"
		}
	}

	logger := LoggerConfigVar.DefaultLogger
	if logger != nil {
		WriteLogWithFields(logger,fields,logLevel)
	}

	accessLogger := LoggerConfigVar.AccessLogger
	errorLogger := LoggerConfigVar.ErrorLogger

	if  strings.ToLower(logLevel) == "info" {
		if  accessLogger != nil {
			WriteLogWithFields(accessLogger,fields,logLevel)
		}
	} else {
		if LoggerConfigVar.SplitAccessAndError {
			if errorLogger != nil {
				WriteLogWithFields(errorLogger,fields,logLevel)
			} else if accessLogger != nil {
				WriteLogWithFields(accessLogger,fields,logLevel)
				f := make(LogFields)
				f["ErrorMessage"] = "SplitAccessAndError has be set to true, but  error log file has not be set."
				logLevel = "warning"
				WriteLogWithFields(accessLogger,f,logLevel)
			}
		} else {
			if accessLogger != nil {
				WriteLogWithFields(accessLogger,fields,logLevel)
			} else if errorLogger != nil {
				WriteLogWithFields(errorLogger,fields,logLevel)
				f := make(LogFields)
				f["ErrorMessage"] = "SplitAccessAndError has be set to false, but  access log file has not be set."
				logLevel = "warning"
				WriteLogWithFields(errorLogger,f,logLevel)
			}
		}
	}

}

// (*Engine)LogWithFields is the method to LogWithFields function
func (*Engine)LogWithFields(fields LogFields, logLevel string){
	LogWithFields(fields,logLevel)
}

// Log log message to DefaultLogger, AccessLogger and ErrorLogger. 
// if SplitAccessAndError is true, then log info log message to AccessLogger and other log message to ErrorLogger.
// otherwise  all log message will be log to AccessLogger.  If  SplitAccessAndError is true but ErrorLogger is nil 
// then all log message will be log to AccessLogger and a additional warning log message will be log to AccessLogger.
// If SplitAccessAndError is false but AccessLogger is nil then all log message will be log to ErrorLogger and a additional 
// warning log message will be log to ErrorLogger.
func Log(message string, logLevel string){
	if strings.TrimSpace(logLevel) == "" {		
		logLevel = "info"
	}

	found := false
	for _,value := range Levels {
		if strings.ToLower(logLevel) == strings.ToLower(value) {
			found = true
		}
	}
	
	if !found {
		logLevel = "info"
	}

	logger := LoggerConfigVar.DefaultLogger
	if logger != nil {
		WriteLog(logger,message,logLevel)
	}

	accessLogger := LoggerConfigVar.AccessLogger
	errorLogger := LoggerConfigVar.ErrorLogger

	if  strings.ToLower(logLevel) == "info" {
		if  accessLogger != nil {
			WriteLog(accessLogger,message,logLevel)
		}
	} else {
		if LoggerConfigVar.SplitAccessAndError {
			if errorLogger != nil {
				WriteLog(errorLogger,message,logLevel)
			} else if accessLogger != nil {
				WriteLog(accessLogger,message,logLevel)
				logLevel = "warning"
				WriteLog(accessLogger,"SplitAccessAndError has be set to true, but  error log file has not be set.",logLevel)
			}
		} else {
			if accessLogger != nil {
				WriteLog(accessLogger,message,logLevel)
			} else if errorLogger != nil {
				WriteLog(errorLogger,message,logLevel)
				logLevel = "warning"
				WriteLog(errorLogger,"SplitAccessAndError has be set to false, but  access log file has not be set.",logLevel)
			}
		}
	}

}

// (*Engine)Log is the method to Log function
func (*Engine)Log(message string, logLevel string){
	Log(message, logLevel )
}


// Logf log message to DefaultLogger, AccessLogger and ErrorLogger. 
// if SplitAccessAndError is true, then log info log message to AccessLogger and other log message to ErrorLogger.
// otherwise  all log message will be log to AccessLogger.  If  SplitAccessAndError is true but ErrorLogger is nil 
// then all log message will be log to AccessLogger and a additional warning log message will be log to AccessLogger.
// If SplitAccessAndError is false but AccessLogger is nil then all log message will be log to ErrorLogger and a additional 
// warning log message will be log to ErrorLogger.
func Logf(logLevel string , format string , a ...interface{}){
	if strings.TrimSpace(logLevel) == "" {		
		logLevel = "info"
	}

	found := false
	for _,value := range Levels {
		if strings.ToLower(logLevel) == strings.ToLower(value) {
			found = true
		}
	}
	
	if !found {
		logLevel = "info"
	}
	
	logMsg := fmt.Sprintf(format, a...)
	logger := LoggerConfigVar.DefaultLogger
	if logger != nil {
		WriteLog(logger,logMsg,logLevel)
	}

	accessLogger := LoggerConfigVar.AccessLogger
	errorLogger := LoggerConfigVar.ErrorLogger

	if  strings.ToLower(logLevel) == "info" {
		if  accessLogger != nil {
			WriteLog(accessLogger,logMsg,logLevel)
		}
	} else {
		if LoggerConfigVar.SplitAccessAndError {
			if errorLogger != nil {
				WriteLog(errorLogger,logMsg,logLevel)
			} else if accessLogger != nil {
				WriteLog(accessLogger,logMsg,logLevel)
				logLevel = "warning"
				WriteLog(accessLogger,"SplitAccessAndError has be set to true, but  error log file has not be set.",logLevel)
			}
		} else {
			if accessLogger != nil {
				WriteLog(accessLogger,logMsg,logLevel)
			} else if errorLogger != nil {
				WriteLog(errorLogger,logMsg,logLevel)
				logLevel = "warning"
				WriteLog(errorLogger,"SplitAccessAndError has be set to false, but  access log file has not be set.",logLevel)
			}
		}
	}

}

// (*Engine)Logf is the method to Log function
func (*Engine)Logf(logLevel string , format string , a ...interface{}){
	Logf(logLevel,format, a... )
}

