package main

import (
	"errors"
	"fmt"
)

var (
	ErrReadingFileContext         = errors.New("error reading file context")
	ErrExtractingContent          = errors.New("error extracting content")
	ErrReadingContent             = errors.New("error reading content")
	ErrFailedToOpenPDF            = errors.New("failed to open PDF file")
	ErrFailedToReadPDF            = errors.New("failed to read PDF file")
	ErrFailedToExtractTextFromPDF = errors.New("failed to extract text from PDF file")
)

type errorArg struct {
	Key   string
	Value any
}

func (e errorArg) Error() string {
	return fmt.Sprintf("%s=%v", e.Key, e.Value)
}

func ErrorArg(key string, value any) error {
	return &errorArg{Key: key, Value: value}
}
