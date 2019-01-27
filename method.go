package rtsp

//go:generate stringer -type Method -output methodlines.go

const (
	DESCRIBE Method = iota
	ANNOUNCE
	GET_PARAMETER
	OPTIONS
	PAUSE
	PLAY
	RECORD
	REDIRECT
	SETUP
	SET_PARAMETER
	TEARDOWN
)

type Method int
