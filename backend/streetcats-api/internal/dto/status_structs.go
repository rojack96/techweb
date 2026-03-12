package dto

// Continue represents HTTP 100 responses.
type Continue struct {
	Status   uint16 `json:"status" example:"100"`
	Message  string `json:"message" example:"Continue"`
	Response any    `json:"response,omitempty"`
} //	@name	ContinueResponse

// SwitchingProtocols represents HTTP 101 responses.
type SwitchingProtocols struct {
	Status   uint16 `json:"status" example:"101"`
	Message  string `json:"message" example:"Switching Protocols"`
	Response any    `json:"response,omitempty"`
} //	@name	SwitchingProtocolsResponse

// Processing represents HTTP 102 responses.
type Processing struct {
	Status   uint16 `json:"status" example:"102"`
	Message  string `json:"message" example:"Processing"`
	Response any    `json:"response,omitempty"`
} //	@name	ProcessingResponse

// EarlyHints represents HTTP 103 responses.
type EarlyHints struct {
	Status   uint16 `json:"status" example:"103"`
	Message  string `json:"message" example:"Early Hints"`
	Response any    `json:"response,omitempty"`
} //	@name	EarlyHintsResponse

// OK represents HTTP 200 responses.
type OK struct {
	Status   uint16 `json:"status" example:"200"`
	Message  string `json:"message" example:"OK"`
	Response any    `json:"response,omitempty"`
} //	@name	OKResponse

// Created represents HTTP 201 responses.
type Created struct {
	Status   uint16 `json:"status" example:"201"`
	Message  string `json:"message" example:"Created"`
	Response any    `json:"response,omitempty"`
} //	@name	CreatedResponse

// Accepted represents HTTP 202 responses.
type Accepted struct {
	Status   uint16 `json:"status" example:"202"`
	Message  string `json:"message" example:"Accepted"`
	Response any    `json:"response,omitempty"`
} //	@name	AcceptedResponse

// NonAuthoritativeInformation represents HTTP 203 responses.
type NonAuthoritativeInformation struct {
	Status   uint16 `json:"status" example:"203"`
	Message  string `json:"message" example:"Non-Authoritative Information"`
	Response any    `json:"response,omitempty"`
} //	@name	NonAuthoritativeInformationResponse

// NoContent represents HTTP 204 responses.
type NoContent struct {
	Status   uint16 `json:"status" example:"204"`
	Message  string `json:"message" example:"No Content"`
	Response any    `json:"response,omitempty"`
} //	@name	NoContentResponse

// ResetContent represents HTTP 205 responses.
type ResetContent struct {
	Status   uint16 `json:"status" example:"205"`
	Message  string `json:"message" example:"Reset Content"`
	Response any    `json:"response,omitempty"`
} //	@name	ResetContentResponse

// PartialContent represents HTTP 206 responses.
type PartialContent struct {
	Status   uint16 `json:"status" example:"206"`
	Message  string `json:"message" example:"Partial Content"`
	Response any    `json:"response,omitempty"`
} //	@name	PartialContentResponse

// MultiStatus represents HTTP 207 responses.
type MultiStatus struct {
	Status   uint16 `json:"status" example:"207"`
	Message  string `json:"message" example:"Multi-Status"`
	Response any    `json:"response,omitempty"`
} //	@name	MultiStatusResponse

// AlreadyReported represents HTTP 208 responses.
type AlreadyReported struct {
	Status   uint16 `json:"status" example:"208"`
	Message  string `json:"message" example:"Already Reported"`
	Response any    `json:"response,omitempty"`
} //	@name	AlreadyReportedResponse

// IMUsed represents HTTP 226 responses.
type IMUsed struct {
	Status   uint16 `json:"status" example:"226"`
	Message  string `json:"message" example:"IM Used"`
	Response any    `json:"response,omitempty"`
} //	@name	IMUsedResponse

// MultipleChoices represents HTTP 300 responses.
type MultipleChoices struct {
	Status   uint16 `json:"status" example:"300"`
	Message  string `json:"message" example:"Multiple Choices"`
	Response any    `json:"response,omitempty"`
} //	@name	MultipleChoicesResponse

// MovedPermanently represents HTTP 301 responses.
type MovedPermanently struct {
	Status   uint16 `json:"status" example:"301"`
	Message  string `json:"message" example:"Moved Permanently"`
	Response any    `json:"response,omitempty"`
} //	@name	MovedPermanentlyResponse

// Found represents HTTP 302 responses.
type Found struct {
	Status   uint16 `json:"status" example:"302"`
	Message  string `json:"message" example:"Found"`
	Response any    `json:"response,omitempty"`
} //	@name	FoundResponse

// SeeOther represents HTTP 303 responses.
type SeeOther struct {
	Status   uint16 `json:"status" example:"303"`
	Message  string `json:"message" example:"See Other"`
	Response any    `json:"response,omitempty"`
} //	@name	SeeOtherResponse

// NotModified represents HTTP 304 responses.
type NotModified struct {
	Status   uint16 `json:"status" example:"304"`
	Message  string `json:"message" example:"Not Modified"`
	Response any    `json:"response,omitempty"`
} //	@name	NotModifiedResponse

// UseProxy represents HTTP 305 responses.
type UseProxy struct {
	Status   uint16 `json:"status" example:"305"`
	Message  string `json:"message" example:"Use Proxy"`
	Response any    `json:"response,omitempty"`
} //	@name	UseProxyResponse

// TemporaryRedirect represents HTTP 307 responses.
type TemporaryRedirect struct {
	Status   uint16 `json:"status" example:"307"`
	Message  string `json:"message" example:"Temporary Redirect"`
	Response any    `json:"response,omitempty"`
} //	@name	TemporaryRedirectResponse

// PermanentRedirect represents HTTP 308 responses.
type PermanentRedirect struct {
	Status   uint16 `json:"status" example:"308"`
	Message  string `json:"message" example:"Permanent Redirect"`
	Response any    `json:"response,omitempty"`
} //	@name	PermanentRedirectResponse

// BadRequest represents HTTP 400 responses.
type BadRequest struct {
	Status   uint16 `json:"status" example:"400"`
	Message  string `json:"message" example:"Bad Request"`
	Response any    `json:"response,omitempty"`
} //	@name	BadRequestResponse

// Unauthorized represents HTTP 401 responses.
type Unauthorized struct {
	Status   uint16 `json:"status" example:"401"`
	Message  string `json:"message" example:"Unauthorized"`
	Response any    `json:"response,omitempty"`
} //	@name	UnauthorizedResponse

// PaymentRequired represents HTTP 402 responses.
type PaymentRequired struct {
	Status   uint16 `json:"status" example:"402"`
	Message  string `json:"message" example:"Payment Required"`
	Response any    `json:"response,omitempty"`
} //	@name	PaymentRequiredResponse

// Forbidden represents HTTP 403 responses.
type Forbidden struct {
	Status   uint16 `json:"status" example:"403"`
	Message  string `json:"message" example:"Forbidden"`
	Response any    `json:"response,omitempty"`
} //	@name	ForbiddenResponse

// NotFound represents HTTP 404 responses.
type NotFound struct {
	Status   uint16 `json:"status" example:"404"`
	Message  string `json:"message" example:"Not Found"`
	Response any    `json:"response,omitempty"`
} //	@name	NotFoundResponse

// MethodNotAllowed represents HTTP 405 responses.
type MethodNotAllowed struct {
	Status   uint16 `json:"status" example:"405"`
	Message  string `json:"message" example:"Method Not Allowed"`
	Response any    `json:"response,omitempty"`
} //	@name	MethodNotAllowedResponse

// NotAcceptable represents HTTP 406 responses.
type NotAcceptable struct {
	Status   uint16 `json:"status" example:"406"`
	Message  string `json:"message" example:"Not Acceptable"`
	Response any    `json:"response,omitempty"`
} //	@name	NotAcceptableResponse

// ProxyAuthenticationRequired represents HTTP 407 responses.
type ProxyAuthenticationRequired struct {
	Status   uint16 `json:"status" example:"407"`
	Message  string `json:"message" example:"Proxy Authentication Required"`
	Response any    `json:"response,omitempty"`
} //	@name	ProxyAuthenticationRequiredResponse

// RequestTimeout represents HTTP 408 responses.
type RequestTimeout struct {
	Status   uint16 `json:"status" example:"408"`
	Message  string `json:"message" example:"Request Timeout"`
	Response any    `json:"response,omitempty"`
} //	@name	RequestTimeoutResponse

// Conflict represents HTTP 409 responses.
type Conflict struct {
	Status   uint16 `json:"status" example:"409"`
	Message  string `json:"message" example:"Conflict"`
	Response any    `json:"response,omitempty"`
} //	@name	ConflictResponse

// Gone represents HTTP 410 responses.
type Gone struct {
	Status   uint16 `json:"status" example:"410"`
	Message  string `json:"message" example:"Gone"`
	Response any    `json:"response,omitempty"`
} //	@name	GoneResponse

// LengthRequired represents HTTP 411 responses.
type LengthRequired struct {
	Status   uint16 `json:"status" example:"411"`
	Message  string `json:"message" example:"Length Required"`
	Response any    `json:"response,omitempty"`
} //	@name	LengthRequiredResponse

// PreconditionFailed represents HTTP 412 responses.
type PreconditionFailed struct {
	Status   uint16 `json:"status" example:"412"`
	Message  string `json:"message" example:"Precondition Failed"`
	Response any    `json:"response,omitempty"`
} //	@name	PreconditionFailedResponse

// PayloadTooLarge represents HTTP 413 responses.
type PayloadTooLarge struct {
	Status   uint16 `json:"status" example:"413"`
	Message  string `json:"message" example:"Payload Too Large"`
	Response any    `json:"response,omitempty"`
} //	@name	PayloadTooLargeResponse

// URITooLong represents HTTP 414 responses.
type URITooLong struct {
	Status   uint16 `json:"status" example:"414"`
	Message  string `json:"message" example:"URI Too Long"`
	Response any    `json:"response,omitempty"`
} //	@name	URITooLongResponse

// UnsupportedMediaType represents HTTP 415 responses.
type UnsupportedMediaType struct {
	Status   uint16 `json:"status" example:"415"`
	Message  string `json:"message" example:"Unsupported Media Type"`
	Response any    `json:"response,omitempty"`
} //	@name	UnsupportedMediaTypeResponse

// RangeNotSatisfiable represents HTTP 416 responses.
type RangeNotSatisfiable struct {
	Status   uint16 `json:"status" example:"416"`
	Message  string `json:"message" example:"Range Not Satisfiable"`
	Response any    `json:"response,omitempty"`
} //	@name	RangeNotSatisfiableResponse

// ExpectationFailed represents HTTP 417 responses.
type ExpectationFailed struct {
	Status   uint16 `json:"status" example:"417"`
	Message  string `json:"message" example:"Expectation Failed"`
	Response any    `json:"response,omitempty"`
} //	@name	ExpectationFailedResponse

// ImATeapot represents HTTP 418 responses.
type ImATeapot struct {
	Status   uint16 `json:"status" example:"418"`
	Message  string `json:"message" example:"I'm a teapot"`
	Response any    `json:"response,omitempty"`
} //	@name	ImATeapotResponse

// MisdirectedRequest represents HTTP 421 responses.
type MisdirectedRequest struct {
	Status   uint16 `json:"status" example:"421"`
	Message  string `json:"message" example:"Misdirected Request"`
	Response any    `json:"response,omitempty"`
} //	@name	MisdirectedRequestResponse

// UnprocessableEntity represents HTTP 422 responses.
type UnprocessableEntity struct {
	Status   uint16 `json:"status" example:"422"`
	Message  string `json:"message" example:"Unprocessable Entity"`
	Response any    `json:"response,omitempty"`
} //	@name	UnprocessableEntityResponse

// Locked represents HTTP 423 responses.
type Locked struct {
	Status   uint16 `json:"status" example:"423"`
	Message  string `json:"message" example:"Locked"`
	Response any    `json:"response,omitempty"`
} //	@name	LockedResponse

// FailedDependency represents HTTP 424 responses.
type FailedDependency struct {
	Status   uint16 `json:"status" example:"424"`
	Message  string `json:"message" example:"Failed Dependency"`
	Response any    `json:"response,omitempty"`
} //	@name	FailedDependencyResponse

// TooEarly represents HTTP 425 responses.
type TooEarly struct {
	Status   uint16 `json:"status" example:"425"`
	Message  string `json:"message" example:"Too Early"`
	Response any    `json:"response,omitempty"`
} //	@name	TooEarlyResponse

// UpgradeRequired represents HTTP 426 responses.
type UpgradeRequired struct {
	Status   uint16 `json:"status" example:"426"`
	Message  string `json:"message" example:"Upgrade Required"`
	Response any    `json:"response,omitempty"`
} //	@name	UpgradeRequiredResponse

// PreconditionRequired represents HTTP 428 responses.
type PreconditionRequired struct {
	Status   uint16 `json:"status" example:"428"`
	Message  string `json:"message" example:"Precondition Required"`
	Response any    `json:"response,omitempty"`
} //	@name	PreconditionRequiredResponse

// TooManyRequests represents HTTP 429 responses.
type TooManyRequests struct {
	Status   uint16 `json:"status" example:"429"`
	Message  string `json:"message" example:"Too Many Requests"`
	Response any    `json:"response,omitempty"`
} //	@name	TooManyRequestsResponse

// RequestHeaderFieldsTooLarge represents HTTP 431 responses.
type RequestHeaderFieldsTooLarge struct {
	Status   uint16 `json:"status" example:"431"`
	Message  string `json:"message" example:"Request Header Fields Too Large"`
	Response any    `json:"response,omitempty"`
} //	@name	RequestHeaderFieldsTooLargeResponse

// UnavailableForLegalReasons represents HTTP 451 responses.
type UnavailableForLegalReasons struct {
	Status   uint16 `json:"status" example:"451"`
	Message  string `json:"message" example:"Unavailable For Legal Reasons"`
	Response any    `json:"response,omitempty"`
} //	@name	UnavailableForLegalReasonsResponse

// InternalServerError represents HTTP 500 responses.
type InternalServerError struct {
	Status   uint16 `json:"status" example:"500"`
	Message  string `json:"message" example:"Internal Server Error"`
	Response any    `json:"response,omitempty"`
} //	@name	InternalServerErrorResponse

// NotImplemented represents HTTP 501 responses.
type NotImplemented struct {
	Status   uint16 `json:"status" example:"501"`
	Message  string `json:"message" example:"Not Implemented"`
	Response any    `json:"response,omitempty"`
} //	@name	NotImplementedResponse

// BadGateway represents HTTP 502 responses.
type BadGateway struct {
	Status   uint16 `json:"status" example:"502"`
	Message  string `json:"message" example:"Bad Gateway"`
	Response any    `json:"response,omitempty"`
} //	@name	BadGatewayResponse

// ServiceUnavailable represents HTTP 503 responses.
type ServiceUnavailable struct {
	Status   uint16 `json:"status" example:"503"`
	Message  string `json:"message" example:"Service Unavailable"`
	Response any    `json:"response,omitempty"`
} //	@name	ServiceUnavailableResponse

// GatewayTimeout represents HTTP 504 responses.
type GatewayTimeout struct {
	Status   uint16 `json:"status" example:"504"`
	Message  string `json:"message" example:"Gateway Timeout"`
	Response any    `json:"response,omitempty"`
} //	@name	GatewayTimeoutResponse

// HTTPVersionNotSupported represents HTTP 505 responses.
type HTTPVersionNotSupported struct {
	Status   uint16 `json:"status" example:"505"`
	Message  string `json:"message" example:"HTTP Version Not Supported"`
	Response any    `json:"response,omitempty"`
} //	@name	HTTPVersionNotSupportedResponse

// VariantAlsoNegotiates represents HTTP 506 responses.
type VariantAlsoNegotiates struct {
	Status   uint16 `json:"status" example:"506"`
	Message  string `json:"message" example:"Variant Also Negotiates"`
	Response any    `json:"response,omitempty"`
} //	@name	VariantAlsoNegotiatesResponse

// InsufficientStorage represents HTTP 507 responses.
type InsufficientStorage struct {
	Status   uint16 `json:"status" example:"507"`
	Message  string `json:"message" example:"Insufficient Storage"`
	Response any    `json:"response,omitempty"`
} //	@name	InsufficientStorageResponse

// LoopDetected represents HTTP 508 responses.
type LoopDetected struct {
	Status   uint16 `json:"status" example:"508"`
	Message  string `json:"message" example:"Loop Detected"`
	Response any    `json:"response,omitempty"`
} //	@name	LoopDetectedResponse

// NotExtended represents HTTP 510 responses.
type NotExtended struct {
	Status   uint16 `json:"status" example:"510"`
	Message  string `json:"message" example:"Not Extended"`
	Response any    `json:"response,omitempty"`
} //	@name	NotExtendedResponse

// NetworkAuthenticationRequired represents HTTP 511 responses.
type NetworkAuthenticationRequired struct {
	Status   uint16 `json:"status" example:"511"`
	Message  string `json:"message" example:"Network Authentication Required"`
	Response any    `json:"response,omitempty"`
} //	@name	NetworkAuthenticationRequiredResponse
