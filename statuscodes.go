package rtsp

// Codes as defined in https://tools.ietf.org/html/rfc2326

//go:generate stringer -output statuslines.go -type Status -linecomment
const (
	// 1xx
	StatusContinue Status = 100 // 100 Continue
)

const (
	// 2xx
	StatusOK      Status = 200 + iota // 200 OK
	StatusCreated                     // 201 Created
)

const (
	// 2xx RTSP specific
	StatusLowOnStorageSpace Status = 250 + iota // 250 Low on Storage Space
)

const (
	// 3xx
	StatusMultipleChoices Status = 300 + iota // 300 Multiple Choices

	StatusMovedPermanently // 301 Moved Permanently
	StatusMovedTemporarily // 302 Moved Temporarily
	StatusSeeOther         // 303 See Other
	StatusNotModified      // 304 Not Modified
	StatusUseProxy         // 305 Use Proxy
)

const (
	// 4xx
	StatusBadRequest Status = 400 + iota // 400 Bad Request

	StatusUnauthorized                // 401 Unauthorized
	StatusPaymentRequired             // 402 Payment Required
	StatusForbidden                   // 403 Forbidden
	StatusNotFound                    // 404 Not Found
	StatusMethodNotAllowed            // 405 Method Not Allowed
	StatusNotAcceptable               // 406 Not Acceptable
	StatusProxyAuthenticationRequired // 407 Proxy Authentication Required
	StatusRequestTimeout              // 408 Request Time-out
	Status409                         // 409 UNDEFINED
	StatusGone                        // 410 Gone
	StatusLengthRequired              // 411 Length Required
	StatusPreconditionFailed          // 412 Precondition Failed
	StatusRequestEntityTooLarge       // 413 Request Entity Too Large
	StatusRequestURITooLarge          // 414 Request-URI Too Large
	StatusUnsupportedMediaType        // 415 Unsupported Media Type
)

const (
	// 4xx RTSP specific
	StatusParameterNotUnderstood Status = 451 + iota // 451 Parameter Not Understood

	StatusConferenceNotFound             // 452 Conference Not Found
	StatusNotEnoughBandwidth             // 453 Not Enough Bandwidth
	StatusSessionNotFound                // 454 Session Not Found
	StatusMethodNotValidInThisState      // 455 Method Not Valid in This State
	StatusHeaderFieldNotValidForResource // 456 Header Field Not Valid for Resource
	StatusInvalidRange                   // 457 Invalid Range
	StatusParameterIsReadOnly            // 458 Parameter Is Read-Only
	StatusAggregateoperationNotAllowed   // 459 Aggregate operation not allowed
	StatusOnlyAggregateOperationAllowed  // 460 Only aggregate operation allowed
	StatusUnsupportedTransport           // 461 Unsupported transport
	StatusDestinationUnreachable         // 462 Destination unreachable
)

const (
	// 5xx
	StatusInternalServerError Status = 500 + iota // 500 Internal Server Error

	StatusNotImplemented          // 501 Not Implemented
	StatusBadGateway              // 502 Bad Gateway
	StatusServiceUnavailable      // 503 Service Unavailable
	StatusGatewayTimeout          // 504 Gateway Time-out
	StatusRTSPVersionNotSupported // 505 RTSP Version not supported
)

const (
	// 5xx RTSP specific
	StatusOptionNotSupported Status = 551 + iota // 551 Option not supported
)

type Status int
