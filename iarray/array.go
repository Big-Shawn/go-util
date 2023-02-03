package iarray

import "github.com/Big-Shawn/go-util/itype"

func Sum[T itype.Numerical](arr []T) T {
	var sum T
	for _, val := range arr {
		sum += val
	}
	return sum
}

func Push[T itype.Numerical](arr []T, elem ...T) []T {
	return append(arr, elem...)
}

func Chunk[T itype.Numerical](arr []T, length int) [][]T {
	var chunk [][]T
	var start int = 0
	var end int = length

	for {
		if len(arr) > end {
			chunk = append(chunk, arr[start:end])
			start = end
			end += length
		} else {
			chunk = append(chunk, arr[start:])
			break
		}
	}

	return chunk
}
