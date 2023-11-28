package errorhandling

import (
	"errors"
	"fmt"
	"testing"
)

// ========================================
//
// Notes
//
// * According to uber style guide [ https://github.com/uber-go/guide/blob/master/style.md#errors ]
//
// | static error  | no matching | var errInvalid = errors.New("invalid") |
// | static error  | matching    | var ErrInvalid = errors.New("invalid") |
// | dynamic error | no matching | fmt.Errorf("invalid file: %f", file)   |
// | dynamic error | matching    | custom error type                      |
//
// `Err` or `err` prefix for variables
// `Error` suffix for custom types
//
//
// ========================================

//
// Assuming following chain of errors:
// UnexpectedResponseError
//     ConnectionError
//         ConnectionResetByPeer
// What would be the proper way to handle this?
//

// ---------------------
// A. using structs
// ---------------------

//
// UnexpectedResponseError
//

type UnexpectedResponseError struct {
	// some additional data
	err error
}

func (e UnexpectedResponseError) Error() string {
	return fmt.Sprintf("Unexpected response: %s", e.err.Error())
}

func (e UnexpectedResponseError) Is(target error) bool {
	_, b := target.(UnexpectedResponseError)

	return b
}

//
// ConnectionError
//

type ConnectionError struct {
	// some additional data
	err error
}

func (e ConnectionError) Error() string {
	return fmt.Sprintf("Connection error: %s", e.err.Error())
}

func (e ConnectionError) Is(target error) bool {
	_, b := target.(ConnectionError)
	return b
}

//
// ConnectionResetByPeerError
//

type ConnectionResetByPeerError struct{}

func (e ConnectionResetByPeerError) Error() string {
	return fmt.Sprintf("Connection reset by peer")
}

func (e ConnectionResetByPeerError) Is(target error) bool {
	_, b := target.(ConnectionResetByPeerError)
	return b
}

//
// Test error structs
//

func TestErrorHandling1(t *testing.T) {
	err0 := ConnectionResetByPeerError{}
	err1 := ConnectionError{err: err0}
	err2 := UnexpectedResponseError{err: err1}

	switch {
	case errors.Is(err2, UnexpectedResponseError{}):
		t.Logf(err2.Error())
	default:
		t.Logf("Defaul action")
	}

	//
	// * Error interface requires `Error() string` to be implemented
	//
	// * errors.Is(err, target error) // [ https://pkg.go.dev/errors#Is ]
	//	   Is reports whether any error in err's tree matches target.
	//     An error is considered to match a target if
	//         it is equal to that target or
	//         if it implements a method Is(error) bool such that Is(target) returns true.
	//

	switch {
	case errors.As(err2, &UnexpectedResponseError{}):
		t.Logf(err2.Error())
	default:
		t.Logf("Defaul action")
	}

	//
	// * errors.As(err, target interface{}) // [ https://pkg.go.dev/errors#As ]
	//    As finds the first error in err's tree that matches target, and
	//    if one is found, sets target to that error value and returns true.
	//    Otherwise, it returns false.
	//

	//
	// Main difference for me:
	//	* `Is` will require implement interface, in order to properlly compare error structs
	//	* `As` will not. We also discard the `target`
	//
}

// ---------------------
// B. avoiding structs
// ---------------------

var (
	ErrConnectionResetByPeer = errors.New("connection reset by peer")
)

//
// Test avoiding structs
//

func TestErrorHandling2(t *testing.T) {
	err0 := fmt.Errorf("Connection error: %w", ErrConnectionResetByPeer)
	err1 := fmt.Errorf("Unexpected response: %w", err0)

	switch {
	case errors.Is(err1, ErrConnectionResetByPeer):
		t.Logf(err1.Error())
	default:
		t.Logf("Defaul action")
	}
}

//
// better example for static errors: https://pkg.go.dev/syscall#Errno.Is
//
// -----
//
// package oserror
//
// var (
// 	ErrInvalid    = errors.New("invalid argument")
// 	ErrPermission = errors.New("permission denied")
// 	ErrExist      = errors.New("file already exists")
// 	ErrNotExist   = errors.New("file does not exist")
// 	ErrClosed     = errors.New("file already closed")
// )
//
// -----
//
// package syscall
//
// func (e Errno) Is(target error) bool {
// 	switch target {
// 	case oserror.ErrPermission:
// 		return e == EACCES || e == EPERM
// 	case oserror.ErrExist:
// 		return e == EEXIST || e == ENOTEMPTY
// 	case oserror.ErrNotExist:
// 		return e == ENOENT
// 	case errorspkg.ErrUnsupported:
// 		return e == ENOSYS || e == ENOTSUP || e == EOPNOTSUPP
// 	}
// 	return false
// }
//
// -----
//
