package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	var s string
	var err error
	var cmd string

	if len(os.Args) < 2 {
		cmd = "content"
		//fmt.Println("Usage: pdf-archiver <method>")
	} else {
		cmd = strings.ToLower(os.Args[1])
	}
	filename := "test.pdf"
	switch cmd {
	case "content":
		s, err = useContentPDF(filename)
	case "cpu":
		s, err = usePdfCPU(filename)
	case "unidoc":
		s, err = useUnidoc(filename)
	case "totext":
		s, err = usePDFtoText(filename)
	}
	if err != nil {
		slog.Error("Failed to read PDF file", "error", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
