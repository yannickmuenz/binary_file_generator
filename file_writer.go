package binaryfilegenerator

import (
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
)

type FileWriter struct {
	Filename *string
}

/*
Generates a binary file of size specified in the FileWriter object
*/
func (filewriter FileWriter) GenerateBinaryFile(file_size_bytes int) error {
	file_size := Converter.GetNumberOfElements(file_size_bytes)

	//Set buffer size at most to 4KiB
	buffer_size := min(2048, file_size)

	absPath, err := filepath.Abs(*filewriter.Filename)
	if err != nil {
		return fmt.Errorf("could not find file %w", err)
	}
	file, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	for file_size > 0 {
		buffer := make([]int32, min(buffer_size, file_size))
		fillBuffer(buffer, file_size)
		err = binary.Write(file, binary.BigEndian, buffer)
		if err != nil {
			return err
		}
		file_size = max(file_size-len(buffer), 0)

		//Clears slice array
		buffer = nil
	}
	return nil
}

func fillBuffer(buffer []int32, number_elements_remaining int) {
	for i := 0; i < min(cap(buffer), number_elements_remaining); i++ {
		buffer[i] = int32(rand.Uint32())
	}
}
