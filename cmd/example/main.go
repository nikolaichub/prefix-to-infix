package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/nikolaichub/prefix-to-infix"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	fileInput       = flag.String("f", "", "File containing expression to compute")
	fileOutput      = flag.String("o", "", "Expression output to file")
)

func main() {
	flag.Parse()

	var Input io.Reader
	var Output io.Writer

	if len(*fileInput) > 0 {
		file, err := os.Open(*fileInput)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: No such file!")
			return
		}
		Input = file
	} else {
		Input = strings.NewReader(*inputExpression)
	}

	if len(*fileOutput) > 0 {
		file, err := os.OpenFile(*fileOutput, os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: No such file!")
			return
		}
		Output = file
	} else {
		Output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  Input,
		Output: Output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

}
