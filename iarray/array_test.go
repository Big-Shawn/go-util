package iarray

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	newArr := Push(arr, 3)
	fmt.Println(newArr)
}

func TestPushArr(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	newArr := Push(arr, []int{1, 2, 3, 4}...)
	fmt.Println(newArr)
}

func TestChunk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 78, 89, 7, 6, 3, 2, 4}
	chunked := Chunk(arr, 10)
	fmt.Println(chunked)
}
