package binaryfilegenerator

type converter struct {
	data_type_number_of_bytes int
}

var Converter = converter{
	data_type_number_of_bytes: 4,
}

// Checks if input is a multiple of the number of bytes of the underlying data type
// Input: number of bytes
func (conv converter) IsMultiple(input int) bool {
	return input%conv.data_type_number_of_bytes == 0
}

// Input: Number of elements
// Output: Number of bytes needed to represent elements
func (conv converter) GetNumberOfBytes(input int) int {
	return input * conv.data_type_number_of_bytes
}

// Input: number of bytes
// Output: Number of elements the bytes represent
func (conv converter) GetNumberOfElements(input int) int {
	return (input - input%conv.data_type_number_of_bytes) / conv.data_type_number_of_bytes
}
