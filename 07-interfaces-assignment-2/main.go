package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct {}

func main() {

	fileName := os.Args[1]
	data, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error while opening file: ", fileName)
	}

	fs := fileWriter{}
	io.Copy(fs, data)
}

func (fileWriter) Write(b []byte) (int, error){
	fmt.Println("Printing file data: ")
	fmt.Println(string(b))
	return len(b), nil
}