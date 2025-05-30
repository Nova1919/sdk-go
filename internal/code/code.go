package code

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrCodeSystem          codes.Code = 500000
	ErrCodeUnavailable                = 500014
	ErrCodeUnauthorized               = 500401
	ErrCodeInvalidArgument            = 400403
	ErrCodeNotFound                   = 200005
	ErrCodeAlreadyExists              = 200006
)

const (
	success        codes.Code = iota
	ErrCodeDefault            = iota + 199999
	ErrCodeDBIOFailed
)

var (
	ErrInvalidArgument = Error(ErrCodeInvalidArgument)
	ErrNotFound        = Error(ErrCodeNotFound)
	ErrAlreadyExists   = Error(ErrCodeAlreadyExists)
	ErrDBIOFailure     = Error(ErrCodeDBIOFailed)
	ErrSystem          = Error(ErrCodeSystem)
)

var (
	code2msg = map[codes.Code]string{
		ErrCodeUnavailable:     "service unavailable",
		ErrCodeUnauthorized:    "unauthorized",
		ErrCodeSystem:          "system error",
		ErrCodeInvalidArgument: "param invalid",
		ErrCodeNotFound:        "not found",
		ErrCodeAlreadyExists:   "can not be duplicated",
		ErrCodeDBIOFailed:      "operation failed",
	}
)

func ErrorDefault(err error) error {
	msg := "unknown error"
	if err != nil {
		msg = err.Error()
	}
	return status.Error(ErrCodeDefault, msg)
}

func Error(c codes.Code) error {
	msg, ok := code2msg[c]
	if ok {
		return status.Error(c, msg)
	}
	return status.Error(c, "some thing wrong")
}

func ErrorMsg(c codes.Code, msg string) error {
	return status.Error(c, msg)
}

func ErrParamInvalidMsg(msg string) error {
	if msg == "" {
		return ErrInvalidArgument
	}
	return status.Error(ErrCodeInvalidArgument, msg)
}

func Format(err error) error {
	if err == nil {
		return nil
	}

	s, ok := status.FromError(err)
	if ok && s != nil {
		return errors.New(fmt.Sprintf("%d | %s", s.Proto().GetCode(), s.Proto().GetMessage()))
	}

	return errors.New(fmt.Sprintf("%d | %s", ErrCodeDefault, err.Error()))
}
