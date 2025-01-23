//go:build !pdftotext

package main

func usePDFtoText(filename string) (s string, err error) {
	s = filename
	return s, err
}
