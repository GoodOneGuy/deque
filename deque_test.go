package deque

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	q := NewDeque[int](10)

	for i := 0; i < q.Size(); i++ {
		fmt.Print(q.At(i), " ")
	}

	for i := 1; i < 20; i++ {
		q.PutFront(50 - i)
		q.PutBack(50 + i)
		q.TestPrint()
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

	q.DeleteAt(2)
	for i := 0; i < q.Size(); i++ {
		fmt.Print(q.At(i), " ")
	}
	fmt.Println()
	q.TestPrint()

}
