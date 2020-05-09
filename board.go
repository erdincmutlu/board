package board

import (
	"fmt"

	"github.com/pkg/errors"
)

type coordinates struct {
	x int
	y int
}

const (
	// NoItem means there is no item in that coordinate
	NoItem = 0
)

// ErrIllegalSizeBoard is to error board should be at least 1x1
var ErrIllegalSizeBoard error = errors.New("Board should be at least 1x1")

// ErrInvalidDimensionsError is given dimensions are not valid
var ErrInvalidDimensionsError error = errors.New("Invalid dimensions")

// The Board structure to be hold
type Board struct {
	items      map[coordinates]byte
	dimensions []int
}

// NewBoard will initialize the board with the given dimensions
func NewBoard(n, m int) (*Board, error) {
	// Board will have 2 dimensions initially. Then can work on more dimensions
	if n < 1 || m < 1 {
		return nil, ErrIllegalSizeBoard
	}

	return &Board{
		items:      map[coordinates]byte{},
		dimensions: []int{n, m},
	}, nil
}

// Checks whether the given coordinate is in bounds of the board
func (b *Board) isInBounds(coord []int) (bool, error) {
	if len(coord) != len(b.dimensions) {
		msg := fmt.Sprintf("Coordinate %v has %d dimensions; the board has %d",
			coord, len(coord), len(b.dimensions))
		return false, errors.Wrap(ErrInvalidDimensionsError, msg)
	}
	for i := 0; i < len(coord); i++ {
		if coord[i] < 0 || coord[i] > b.dimensions[i] {
			return false, nil
		}
	}

	return true, nil
}

// Helper function, just to get a value from
func (b *Board) getItem(coord []int) (uint8, error) {
	inBounds, err := b.isInBounds(coord)
	if err != nil || !inBounds {
		return 0, ErrInvalidDimensionsError
	}

	val, ok := b.items[coordinates{x: coord[0], y: coord[1]}]
	if !ok {
		return NoItem, nil
	}

	return val, nil
}

// SetItem sets the item in the given coordinates of the board
func (b *Board) SetItem(coord []int, item byte) error {
	inBounds, err := b.isInBounds(coord)
	if err != nil || !inBounds {
		return ErrInvalidDimensionsError
	}

	b.items[coordinates{x: coord[0], y: coord[1]}] = item
	return nil
}
