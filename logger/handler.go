package logger

type Handler interface {
	Log(r *Record) error
}
