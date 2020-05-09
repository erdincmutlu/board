package board

import (
	"errors"
	"fmt"
)

var IllegalSizeBoard error = errors.New("Board should be at least 1x1")

// The Board structure to be hold
type Board struct {
	items      []uint8
	dimensions []int
}

// Board will initialize the board with the given dimensions
func NewBoard(n, m int) (*Board, error) {
	// Board will have 2 dimensions initially. Then can work on more dimensions
	if n < 1 || m < 1 {
		return nil, IllegalSizeBoard
	}

	fmt.Println("Board initialized")
	return &Board{items: []uint8{}, dimensions: []int{n, m}}, nil
}
