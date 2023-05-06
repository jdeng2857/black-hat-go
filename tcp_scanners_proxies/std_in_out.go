package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{}

// Read reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

// FooWrite defines an io.Write to write to stdout
type FooWriter struct{}

// Write writes data to stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)

	// Use built-in Copy function
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
