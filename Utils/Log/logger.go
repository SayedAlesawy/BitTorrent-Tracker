package logger

import "log"

// LogInfo Used to customize logging info insider generic functions
type LogInfo struct {
	Success string //The message in case of success
	Error   string //The message in case of failure
}

// LogMsg A function to log messages
func LogMsg(sign string, msg string) {
	log.Println(sign, "#", msg)
}

// LogSuccess A function to log success messages
func LogSuccess(err error, sign string, msg string) {
	if err == nil {
		LogMsg(sign, msg)
	}
}

// LogErr A function to log error messages
func LogErr(err error, sign string, msg string, abort bool) {
	if err != nil {
		LogMsg(sign, msg)
		if abort == true {
			panic(err)
		}
	}
}
