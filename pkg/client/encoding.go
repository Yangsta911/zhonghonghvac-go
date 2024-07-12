package client

import (
	"encoding/binary"
)

// CalculateByteSum calculates the sum of a byte slice and returns the least significant byte.
func CalculateByteSum(data []byte) uint8 {
	var sum int64
	for _, b := range data {
		sum = sum + int64(b)
	}
	return uint8(sum % 256)
}

// dataBlockArray returns a byteSlice given a uint16 slice.
func dataBlockArray(arr []uint16) []byte {
	byteSlice := make([]byte, len(arr)*2)
	for i, v := range arr {
		binary.BigEndian.PutUint16(byteSlice[i*2:], v)
	}

	return byteSlice
}

// PrependUint16 prepends a uint16 number to a uint16 slice.
func PrependUint16(slice []uint16, element uint16) []uint16 {
	newSlice := append([]uint16{element}, slice...)
	return newSlice
}
