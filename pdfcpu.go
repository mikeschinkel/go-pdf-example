package main

import (
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func usePdfCPU(filename string) (s string, err error) {
	var reader io.Reader
	var ctx *model.Context
	var pageNum int
	var content []byte

	ctx, err = api.ReadContextFile(filename)
	if err != nil {
		err = errors.Join(ErrReadingFileContext, err)
		goto end
	}
	pageNum = 1 // For first page
	reader, err = pdfcpu.ExtractPageContent(ctx, pageNum)
	if err != nil {
		err = errors.Join(ErrExtractingContent, err)
		goto end
	}
	content, err = io.ReadAll(reader)
	if err != nil {
		err = errors.Join(ErrReadingContent, err)
		goto end
	}
	s = string(content)
end:
	return s, err
}
