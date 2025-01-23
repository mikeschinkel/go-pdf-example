package main

import (
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func useUnidoc(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer mustClose(f)

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return "", err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var text string
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		pageText, err := ex.ExtractText()
		if err != nil {
			return "", err
		}

		text += pageText
	}
	return text, nil
}
