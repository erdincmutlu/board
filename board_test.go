package board

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	_, err := NewBoard(-1, 5)
	require.Equal(t, IllegalSizeBoard, err)

	_, err = NewBoard(3, -4)
	require.Equal(t, IllegalSizeBoard, err)

	expBoard := Board{
		items:      []uint8{},
		dimensions: []int{3, 4},
	}
	b, err := NewBoard(3, 4)
	require.NoError(t, err)
	require.Equal(t, &expBoard, b)
}
