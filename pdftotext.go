//go:build pdftotext

package main

import (
	"errors"
	"os"

	"github.com/heussd/pdftotext-go"
)

func usePDFtoText(filename string) (string, error) {
	var content []byte
	onErr := func(e1, e2 error) error {
		return errors.Join(e1, e2,
			ErrorArg("filename", filename),
			ErrorArg("error", err),
		)
	}
	// Open the PDF file
	f, err := os.Open(filename)
	if err != nil {
		err = onErr(err, ErrFailedToOpenPDF)
		goto end
	}
	defer mustClose(f)
	// Read the file content
	content, err = os.ReadFile(filename)
	if err != nil {
		err = onErr(err, ErrFailedToReadPDF)
		goto end
	}
	// Extract text from the PDF file
	text, err = pdftotext.Extract(content)
	if err != nil {
		err = onErr(err, ErrFailedToExtractTextFromPDF)
		goto end
	}
	// Print the extracted text
	return text, err
}
