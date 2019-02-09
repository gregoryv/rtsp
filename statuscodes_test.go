package rtsp

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_status_descriptions(t *testing.T) {
	assert := asserter.New(t)
	cases := []struct {
		code int
		exp  string
	}{
		{0, "Status(0)"},
		{100, "100 Continue"},
		{200, "200 OK"},
		{201, "201 Created"},
		{250, "250 Low on Storage Space"},
		{300, "300 Multiple Choices"},
		{301, "301 Moved Permanently"},
		{302, "302 Moved Temporarily"},
		{303, "303 See Other"},
		{304, "304 Not Modified"},
		{305, "305 Use Proxy"},
		{400, "400 Bad Request"},
		{401, "401 Unauthorized"},
		{402, "402 Payment Required"},
		{403, "403 Forbidden"},
		{404, "404 Not Found"},
		{405, "405 Method Not Allowed"},
		{406, "406 Not Acceptable"},
		{407, "407 Proxy Authentication Required"},
		{408, "408 Request Time-out"},
		{410, "410 Gone"},
		{411, "411 Length Required"},
		{412, "412 Precondition Failed"},
		{413, "413 Request Entity Too Large"},
		{414, "414 Request-URI Too Large"},
		{415, "415 Unsupported Media Type"},
		{451, "451 Parameter Not Understood"},
		{452, "452 Conference Not Found"},
		{453, "453 Not Enough Bandwidth"},
		{454, "454 Session Not Found"},
		{455, "455 Method Not Valid in This State"},
		{456, "456 Header Field Not Valid for Resource"},
		{457, "457 Invalid Range"},
		{458, "458 Parameter Is Read-Only"},
		{459, "459 Aggregate operation not allowed"},
		{460, "460 Only aggregate operation allowed"},
		{461, "461 Unsupported transport"},
		{462, "462 Destination unreachable"},
		{500, "500 Internal Server Error"},
		{501, "501 Not Implemented"},
		{502, "502 Bad Gateway"},
		{503, "503 Service Unavailable"},
		{504, "504 Gateway Time-out"},
		{505, "505 RTSP Version not supported"},
		{551, "551 Option not supported"},
		{999, "Status(999)"},
	}
	for _, c := range cases {
		assert().Equals(Status(c.code), c.exp)
	}
}
