package lib

import (
	logpkg "github.com/ckeyer/go-log"
	"os"
)

var (
	log    = logpkg.MustGetLogger("example")
	format = logpkg.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
)

func init() {
	backend1 := logpkg.NewLogBackend(os.Stderr, "", 0)
	backend2 := logpkg.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logpkg.NewBackendFormatter(backend2, format)

	backend1Leveled := logpkg.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logpkg.ERROR, "")

	logpkg.SetBackend(backend1Leveled, backend2Formatter)
}

func GetLogger() *logpkg.Logger {
	return log
}
