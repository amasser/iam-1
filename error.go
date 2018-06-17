package iam

import (
	"bytes"
	"fmt"
)

// ECONFLICT is the error code for conflicts.
// EINTERNAL is the error code for internal errors.
// EINVALID is the error code for invalid data.
// ENOTFOUND is the error code for a not found object.
const (
	ECONFLICT = "conflict"
	EINTERNAL = "internal"
	EINVALID  = "invalid"
	ENOTFOUND = "notfound"
)

// Error represents all the IAM related error codes.
type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// ErrorCode extract the error code of supplied error.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

// ErrorMessage extract the error message of supplied error.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}
