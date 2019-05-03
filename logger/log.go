package logger

import (
	log "github.com/Sirupsen/logrus"
)

// LogMsgInfo LogInfo
func LogMsgInfo(msg string) {
	log.Info(msg)
}

// LogMsgWarning LogWarning
func LogMsgWarning(msg string) {
	log.Warning(msg)
}

// LogMsgError LogError
func LogMsgError(msg string) {
	log.Error(msg)
}

// LogMsgFatal LogFatal
func LogMsgFatal(msg string) {
	log.Fatal(msg)
}

// LogMsgPanic LogPanic
func LogMsgPanic(msg string) {
	log.Panic(msg)
}

// LogFormat Formatacao padrao
func LogFormat() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.Info("Iniciando o log")
}
