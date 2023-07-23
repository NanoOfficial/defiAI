package logger

import "time"

type Record struct {
	Time     time.Time
	Lvl      Lvl
	Msg      string
	Ctx      []interface{}
	Call     stack.Call
	KeyNames RecordKeyNames
}

type RecordKeyNames struct {
	Time string
	Msg  string
	Lvl  string
	Ctx  string
}
