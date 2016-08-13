package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

var _ = fmt.Printf

type Causer interface {
	Cause() error
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}

type BaseError struct {
	wrapped    error
	cause      error
	stackTrace errors.StackTrace
	outpuStackTrace bool
}

func (be *BaseError) Cause() error {
	return be.cause
}

func (be *BaseError) StackTrace() errors.StackTrace {
	return be.stackTrace
}

func (be *BaseError) SetOutputStackTrace(output bool) {
	be.outpuStackTrace = output
}

func (be *BaseError) GetOutputStackTrace() bool {
	return be.outpuStackTrace
}

func NewBaseError(wrapped error) *BaseError {
	be := &BaseError{wrapped: wrapped}
	if c, ok := be.wrapped.(Causer); ok {
		be.cause = c.Cause()
	}
	if st, ok := be.wrapped.(StackTracer); ok {
		be.stackTrace = st.StackTrace()
	}
	return be
}

type Internal struct {
	*BaseError
}

func (e *Internal) Error() string {
	return fmt.Sprintf("errors.Internal: %s", e.wrapped.Error())
}

type NotFound struct {
	*BaseError
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("errors.NotFound: %s", e.wrapped.Error())
}

func Internalf(format string, args ...interface{}) *Internal {
	return &Internal{NewBaseError(fmt.Errorf(format, args...))}
}

func InternalWrapf(err error, format string, args ...interface{}) *Internal {
	return &Internal{NewBaseError(errors.Wrapf(err, format, args...))}
}

func NotFoundf(format string, args ...interface{}) *NotFound {
	return &NotFound{NewBaseError(fmt.Errorf(format, args...))}
}

func NotFoundWrapf(err error, format string, args ...interface{}) *NotFound {
	return &NotFound{NewBaseError(errors.Wrapf(err, format, args...))}
}