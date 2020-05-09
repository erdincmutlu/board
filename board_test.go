package board

import (
	"testing"

	//"github.com/pkg/errors"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	_, err := NewBoard(-1, 5)
	require.Equal(t, ErrIllegalSizeBoard, err)

	_, err = NewBoard(3, -4)
	require.Equal(t, ErrIllegalSizeBoard, err)

	expBoard := Board{
		items:      []uint8{},
		dimensions: []int{3, 4},
	}
	b, err := NewBoard(3, 4)
	require.NoError(t, err)
	require.Equal(t, &expBoard, b)
}

func TestIsInBounds(t *testing.T) {
	tests := []struct {
		name     string
		coord    []int
		err      error
		errText  string
		InBounds bool
	}{
		{
			name:    "Invalid dimension",
			coord:   []int{1, 2, 3},
			err:     ErrInvalidDimensionsError,
			errText: "Coordinate [1 2 3] has 3 dimensions; the board has 2: Invalid dimensions",
		},
		{
			name:     "Ok",
			coord:    []int{2, 1},
			InBounds: true,
		},
		{
			name:  "Bigger first",
			coord: []int{7, 2},
		},
		{
			name:  "Bigger second",
			coord: []int{2, 8},
		},
		{
			name:  "Negative first",
			coord: []int{-5, 2},
		},
		{
			name:  "Negative second",
			coord: []int{1, -3},
		},
	}

	b, err := NewBoard(3, 5)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inBounds, err := b.isInBounds(test.coord)
			require.True(t, errors.Is(err, test.err))
			if err != nil {
				require.Equal(t, test.errText, err.Error())
			}

			require.Equal(t, test.InBounds, inBounds)
		})
	}
}
