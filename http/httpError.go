package http

type HttpError struct {
	statusCode  int
	level       int8
	error       string
	callerError string
}

func (e HttpError) Error() string {
	return e.callerError + ": " + e.error
}

func (e *HttpError) OrigError() string {
	return e.error
}

func (e *HttpError) CallerError() string {
	return e.callerError
}

func (e *HttpError) Level() int8 {
	return e.level
}

func (e *HttpError) StatusCode() int {
	return e.statusCode
}
