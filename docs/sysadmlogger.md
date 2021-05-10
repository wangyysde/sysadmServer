package sysadmlogger // import "github.com/wangyysde/sysadmServer/logger"

var sysadmLogger = SysadmLogger{
	accessLoggerFile: "",
	errorLoggerFile:  "",

	accessFp: nil,
	errorFp:  nil,

	accessLogger: nil,
	errorLogger:  nil,
	stdoutLogger: nil,

	LoggerFormat: "Text",
	DateFormat:   time.RFC3339,

	Allstdout: true,
	logLevel:  5,
}
    Set global variable config and its default value

type SysadmLogger struct {
	accessLoggerFile string //accessLoggerFile is the path of access log file ,if logger access log to a file
	errorLoggerFile  string //errorLoggerFile is the path of error log file ,if logger error log to a file

	accessFp *os.File //The file descriptor of the access log file
	errorFp  *os.File //The file descriptor of the error log file

	accessLogger *log.Logger //Logger for access log
	errorLogger  *log.Logger //Logger for error log
	stdoutLogger *log.Logger //Logger for stdout

	LoggerFormat string //set log format for output
	DateFormat   string //set date formate

	Allstdout bool   //If all log message log to stdout ,then Allstdout should be set to True
	logLevel  uint32 //Only log the `logLevel` severity or above.
}
    SysadmLogger struct used to save parameters for logger, such as access log
    file error log file

func New() *SysadmLogger
    Return a instance of SysadmLogger

func (sysadmLogger *SysadmLogger) EndLogger(logType string) (err error)
    * EndLogger function will be call by defer * EndLogger will close file
    descriptor of access or error * and reset logger of access or error to nil

func (sysadmLogger *SysadmLogger) InitLogger(logType string, toStdout bool) (logger *log.Logger, err error)
    Init logger instance for access or error. before call this func,
    sysadmLoggerLogfile(logType, logFile) should be called

func (sysadmLogger *SysadmLogger) InitStdoutLogger() (stdoutLogger *log.Logger, err error)
    InitStdoutLogger initated a logger to logging log message to stdout
    accroding SysadmLogger.LoggerFormat set log formate and set
    SysadmLogger.stdoutLogger with the initated instance

func (sysadmLogger *SysadmLogger) LoggingLog(logType string, logLevel string, args ...interface{})
    * * Logging a message to Logger * if the sysadmLogger.Allstdout ,then
    logging the log messages to stdout

func (sysadmLogger *SysadmLogger) OpenLogfile(logType string, logFile string) (fp *os.File, err error)
    * according to logType, sysadmLoggerLogfile set logFile to
    sysadmLogger.accessLoggerFile or sysadmLogger.errorLoggerFile * and set file
    descriptor to accessFp or errorFp if logFile can be opened. * to close the
    openned file on time, a defer function should be called following call this
    function if this return successful.

func (sysadmLogger *SysadmLogger) SetFormat(Logger *log.Logger, logType string) (logger *log.Logger)
    * set log format to text or json and set loggerFormat * the fields of the
    struct of loggerFormat refer to
    :https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter

func (sysadmLogger *SysadmLogger) Setlevel(loggerLevel string, Logger *log.Logger) (logger *log.Logger)
    * set logger level to sysadmLogger.loggerLevel

func (sysadmLogger *SysadmLogger) accessWriter(level int, msg string)

func (sysadmLogger *SysadmLogger) errorWriter(level int, msg string)

func (sysadmLogger *SysadmLogger) setLogFormat(format string)

func (sysadmLogger *SysadmLogger) setLogLevel(level int)

func (sysadmLogger *SysadmLogger) stdoutWriter(level int, msg string)

