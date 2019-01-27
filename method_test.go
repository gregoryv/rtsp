package rtsp

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_method_names(t *testing.T) {
	assert := asserter.New(t)
	cases := []struct {
		method Method
		exp    string
	}{
		{DESCRIBE, "DESCRIBE"},
		{ANNOUNCE, "ANNOUNCE"},
		{GET_PARAMETER, "GET_PARAMETER"},
		{OPTIONS, "OPTIONS"},
		{PAUSE, "PAUSE"},
		{PLAY, "PLAY"},
		{RECORD, "RECORD"},
		{REDIRECT, "REDIRECT"},
		{SETUP, "SETUP"},
		{SET_PARAMETER, "SET_PARAMETER"},
		{TEARDOWN, "TEARDOWN"},
		{Method(9999), "Method(9999)"},
	}
	for _, c := range cases {
		assert().Equals(c.method.String(), c.exp)
	}
}
