package ports

type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
	Debug(string, ...interface{})
}
