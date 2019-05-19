package logging

import (
	"github.com/op/go-logging"
	"os"
)

var (
	Log    = logging.MustGetLogger("volvox.fortnite")
	format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)
)

// InitLogger initialises a logger instance for us to use throughout Volvox.Fortnite
func InitLogger() {
	f, err := os.OpenFile("bot.log", os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	stdOut := logging.NewLogBackend(os.Stdout, "", 0)
	logFile := logging.NewLogBackend(f, "", 0)

	stdOutFormatter := logging.NewBackendFormatter(stdOut, format)
	logFileFormatter := logging.NewBackendFormatter(logFile, format)

	stdOutLevel := logging.AddModuleLevel(stdOut)
	stdOutLevel.SetLevel(logging.DEBUG, "")

	logFileLevel := logging.AddModuleLevel(logFile)
	logFileLevel.SetLevel(logging.INFO, "")

	logging.SetBackend(stdOutFormatter, logFileFormatter)

}
