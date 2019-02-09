package rtsp

import "fmt"

// Codes as defined in https://tools.ietf.org/html/rfc2326

const (
	StatusContinue                       = 100        // 100 Continue
	StatusOK                             = 200 + iota // 200 OK
	StatusCreated                                     // 201 Created
	StatusLowOnStorageSpace              = 250 + iota // 250 Low on Storage Space
	StatusMultipleChoices                = 300 + iota // 300 Multiple Choices
	StatusMovedPermanently                            // 301 Moved Permanently
	StatusMovedTemporarily                            // 302 Moved Temporarily
	StatusSeeOther                                    // 303 See Other
	StatusNotModified                                 // 304 Not Modified
	StatusUseProxy                                    // 305 Use Proxy
	StatusBadRequest                     = 400 + iota // 400 Bad Request
	StatusUnauthorized                                // 401 Unauthorized
	StatusPaymentRequired                             // 402 Payment Required
	StatusForbidden                                   // 403 Forbidden
	StatusNotFound                                    // 404 Not Found
	StatusMethodNotAllowed                            // 405 Method Not Allowed
	StatusNotAcceptable                               // 406 Not Acceptable
	StatusProxyAuthenticationRequired                 // 407 Proxy Authentication Required
	StatusRequestTimeout                              // 408 Request Time-out
	Status409                                         // 409 UNDEFINED
	StatusGone                                        // 410 Gone
	StatusLengthRequired                              // 411 Length Required
	StatusPreconditionFailed                          // 412 Precondition Failed
	StatusRequestEntityTooLarge                       // 413 Request Entity Too Large
	StatusRequestURITooLarge                          // 414 Request-URI Too Large
	StatusUnsupportedMediaType                        // 415 Unsupported Media Type
	StatusParameterNotUnderstood         = 451 + iota // 451 Parameter Not Understood
	StatusConferenceNotFound                          // 452 Conference Not Found
	StatusNotEnoughBandwidth                          // 453 Not Enough Bandwidth
	StatusSessionNotFound                             // 454 Session Not Found
	StatusMethodNotValidInThisState                   // 455 Method Not Valid in This State
	StatusHeaderFieldNotValidForResource              // 456 Header Field Not Valid for Resource
	StatusInvalidRange                                // 457 Invalid Range
	StatusParameterIsReadOnly                         // 458 Parameter Is Read-Only
	StatusAggregateoperationNotAllowed                // 459 Aggregate operation not allowed
	StatusOnlyAggregateOperationAllowed               // 460 Only aggregate operation allowed
	StatusUnsupportedTransport                        // 461 Unsupported transport
	StatusDestinationUnreachable                      // 462 Destination unreachable
	StatusInternalServerError            = 500 + iota // 500 Internal Server Error
	StatusNotImplemented                              // 501 Not Implemented
	StatusBadGateway                                  // 502 Bad Gateway
	StatusServiceUnavailable                          // 503 Service Unavailable
	StatusGatewayTimeout                              // 504 Gateway Time-out
	StatusRTSPVersionNotSupported                     // 505 RTSP Version not supported
	StatusOptionNotSupported             = 551 + iota // 551 Option not supported
)

var lines map[int]string

func init() {
	lines = make(map[int]string)
	lines[100] = "100 Continue"
	lines[200] = "200 OK"
	lines[201] = "201 Created"
	lines[250] = "250 Low on Storage Space"
	lines[300] = "300 Multiple Choices"
	lines[301] = "301 Moved Permanently"
	lines[302] = "302 Moved Temporarily"
	lines[303] = "303 See Other"
	lines[304] = "304 Not Modified"
	lines[305] = "305 Use Proxy"
	lines[400] = "400 Bad Request"
	lines[401] = "401 Unauthorized"
	lines[402] = "402 Payment Required"
	lines[403] = "403 Forbidden"
	lines[404] = "404 Not Found"
	lines[405] = "405 Method Not Allowed"
	lines[406] = "406 Not Acceptable"
	lines[407] = "407 Proxy Authentication Required"
	lines[408] = "408 Request Time-out"
	lines[410] = "410 Gone"
	lines[411] = "411 Length Required"
	lines[412] = "412 Precondition Failed"
	lines[413] = "413 Request Entity Too Large"
	lines[414] = "414 Request-URI Too Large"
	lines[415] = "415 Unsupported Media Type"
	lines[451] = "451 Parameter Not Understood"
	lines[452] = "452 Conference Not Found"
	lines[453] = "453 Not Enough Bandwidth"
	lines[454] = "454 Session Not Found"
	lines[455] = "455 Method Not Valid in This State"
	lines[456] = "456 Header Field Not Valid for Resource"
	lines[457] = "457 Invalid Range"
	lines[458] = "458 Parameter Is Read-Only"
	lines[459] = "459 Aggregate operation not allowed"
	lines[460] = "460 Only aggregate operation allowed"
	lines[461] = "461 Unsupported transport"
	lines[462] = "462 Destination unreachable"
	lines[500] = "500 Internal Server Error"
	lines[501] = "501 Not Implemented"
	lines[502] = "502 Bad Gateway"
	lines[503] = "503 Service Unavailable"
	lines[504] = "504 Gateway Time-out"
	lines[505] = "505 RTSP Version not supported"
	lines[551] = "551 Option not supported"
}

func Status(code int) (line string) {
	line, found := lines[code]
	if !found {
		return fmt.Sprintf("Status(%v)", code)
	}
	return
}
