package main

import (
	"log"
	"os"
	"redis_go/store"
)

type LoggingMiddleware struct {
	inner  store.Storer //any storer,the thing we are decorating
	logger *log.Logger  //log.Logger is
}

// log shortfile is used to log the file name and line number of the log message.
func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *LoggingMiddleware) Set(key, value string) error {
	l.logger.Printf("Set called with key: %s, value: %s", key, value)
	return l.inner.Set(key, value)
}

func (l *LoggingMiddleware) Get(key string) (string, bool, error) {
	l.logger.Printf("Get called with key: %s", key)
	return l.inner.Get(key)
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("Delete called with key: %s", key)
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Len() int {
	return l.inner.Len()
}

func (l *LoggingMiddleware) Keys() []string {
	l.logger.Printf("Keys called")
	return l.inner.Keys()

}

func (m *LoggingMiddleware) Incr(key string) (int, error) {

	return 0, nil
}
