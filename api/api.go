package api

import (
	std "log"
	"sync"

	"github.com/RRRRomeo/QLog/internal"
)

type LogApi struct {
	sync.Mutex
	Logger *internal.Logger
}

var logApi LogApi = LogApi{
	Logger: internal.NewLogger(internal.LOG_LEVEL_DBG, 4),
}

func NewLogApi(log_level uint8, isout uint8, out_ffp string) *LogApi {
	return &LogApi{
		Logger: internal.NewLoggerWithOutter(internal.LOG_LEVEL(log_level), isout, out_ffp),
	}
}

func LoggerInit(l uint8, isout uint8, out_ffp string) *internal.Logger {
	return internal.NewLoggerWithOutter(internal.LOG_LEVEL(l), 1, out_ffp)
}

func LoggerDeinit(l *internal.Logger) {
	internal.DelLogger(l)
}

func (log *LogApi) Debugf(fmt string, v ...any) {
	log.Lock()
	defer log.Unlock()

	log.Logger.Debugf(fmt, v...)
}

func Debugf(fmt string, v ...any) {
	logApi.Lock()
	defer logApi.Unlock()

	logApi.Logger.Debugf(fmt, v...)
}

func (log *LogApi) Infof(fmt string, v ...any) {
	log.Lock()
	defer log.Unlock()

	log.Logger.Infof(fmt, v...)
}

func Infof(fmt string, v ...any) {
	logApi.Lock()
	defer logApi.Unlock()

	logApi.Logger.Infof(fmt, v...)
}

func (log *LogApi) Warnf(fmt string, v ...any) {
	log.Lock()
	defer log.Unlock()
	log.Logger.Warnf(fmt, v...)
}

func Warnf(fmt string, v ...any) {
	logApi.Lock()
	defer logApi.Unlock()
	logApi.Logger.Warnf(fmt, v...)
}

func (log *LogApi) Errf(fmt string, v ...any) {
	log.Lock()
	defer log.Unlock()
	log.Logger.Errf(fmt, v...)
}

func Errf(fmt string, v ...any) {
	logApi.Lock()
	defer logApi.Unlock()
	logApi.Logger.Errf(fmt, v...)
}

func (log *LogApi) Fatalf(fmt string, v ...any) {
	log.Lock()
	defer log.Unlock()
	log.Logger.Fatalf(fmt, v...)
}

func Fatalf(fmt string, v ...any) {
	logApi.Lock()
	defer logApi.Unlock()
	logApi.Logger.Fatalf(fmt, v...)
}

func Printf(fmt string, v ...any) {
	std.Printf(fmt, v...)
}

func Println(fmt string, v ...any) {
	std.Println(v...)
}
