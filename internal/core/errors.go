package core

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrOpenFile        = errors.New("error while opening file")
	ErrWriteFile       = errors.New("error while writing file")
)
