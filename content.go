package main

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/go-pdf-content-reader"
)

func useContentPDF(filename string) (s string, err error) {
	var pdfRdr *pdf.Reader
	//var ioRdr io.Reader
	var lines []string
	var sb strings.Builder

	pdfRdr, err = pdf.Open(filename)
	// remember close file
	defer mustClose(pdfRdr)
	if err != nil {
		goto end
	}

	lines, err = pdfRdr.GetPlainTextLines()
	if err != nil {
		goto end
	}
	for i, line := range lines {
		sb.WriteString(fmt.Sprintf("%d. — %s\n", i, strings.TrimRight(line, "\n")))
	}
	//_, err = buf.ReadFrom(ioRdr)
	//if err != nil {
	//	goto end
	//}
	//s = buf.String()
	s = sb.String()
end:
	return s, err
}
