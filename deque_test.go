package deque

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	q := NewDeque[int](10)
	q.PutFront(1)
	q.PutFront(2)
	q.PutBack(3)
	q.PutBack(4)

	for i := 0; i < q.Size(); i++ {
		fmt.Print(q.At(i), " ")
	}
	fmt.Println()
	q.TestPrint()
	fmt.Println(q.At(1), " ")
	q.InsertAt(3, 5)
	for i := 0; i < q.Size(); i++ {
		fmt.Print(q.At(i), " ")
	}
	fmt.Println()
	q.TestPrint()
	q.InsertAt(1, 6)
	for i := 0; i < q.Size(); i++ {
		fmt.Print(q.At(i), " ")
	}
	fmt.Println()
	q.TestPrint()
}
