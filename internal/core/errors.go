package core

import "errors"

var (
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrInvalidFileExtension = errors.New("invalid file extension")
	ErrOpenFile             = errors.New("error while opening file")
	ErrWriteFile            = errors.New("error while writing file")
)
