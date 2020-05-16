package board

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	_, err := NewBoard(-1, 5)
	require.Equal(t, ErrIllegalSizeBoard, err)

	_, err = NewBoard(3, -4)
	require.Equal(t, ErrIllegalSizeBoard, err)

	expBoard := Board{
		items:      map[twoDCoordinates]byte{},
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

func TestGetItem(t *testing.T) {
	testItem := byte('*')

	tests := []struct {
		name  string
		coord []int
		err   error
		item  byte
	}{
		{
			name:  "Ok",
			coord: []int{2, 1},
			item:  testItem,
		},
		{
			name:  "Not inbound",
			coord: []int{2, 7},
			err:   ErrInvalidDimensionsError,
		},
		{
			name:  "No item in coordinates",
			coord: []int{2, 2},
			item:  NoItem,
		},
	}

	b, err := NewBoard(3, 5)
	require.NoError(t, err)
	err = b.SetItem([]int{2, 1}, testItem)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item, err := b.getItem(test.coord)
			require.True(t, errors.Is(err, test.err))
			require.Equal(t, test.item, item)
		})
	}
}

func TestSetItem(t *testing.T) {
	testItem := byte('*')

	tests := []struct {
		name  string
		coord []int
		item  byte
		err   error
	}{
		{
			name:  "Ok",
			coord: []int{2, 1},
		},
		{
			name:  "Not inbound",
			coord: []int{2, 7},
			err:   ErrInvalidDimensionsError,
		},
	}

	b, err := NewBoard(3, 5)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := b.SetItem(test.coord, testItem)
			require.True(t, errors.Is(err, test.err))
		})
	}
}

func TestIsSet(t *testing.T) {
	testItem := byte('*')
	setCoordinate := []int{3, 2}

	tests := []struct {
		name  string
		coord []int
		isSet bool
	}{
		{
			name:  "Set",
			coord: setCoordinate,
			isSet: true,
		},
		{
			name:  "Not set",
			coord: []int{2, 2},
		},
		{
			name:  "Not inbound",
			coord: []int{2, 7},
		},
	}

	b, err := NewBoard(3, 5)
	require.NoError(t, err)
	err = b.SetItem(setCoordinate, testItem)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isSet := b.IsSet(test.coord)
			require.Equal(t, test.isSet, isSet)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	testItem := byte('*')
	setCoordinate := []int{2, 0}

	tests := []struct {
		name    string
		coord   []int
		isEmpty bool
	}{
		{
			name:    "Empty",
			coord:   []int{2, 2},
			isEmpty: true,
		},
		{
			name:  "Not empty",
			coord: setCoordinate,
		},
		{
			name:    "Not inbound",
			coord:   []int{2, 7},
			isEmpty: true,
		},
	}

	b, err := NewBoard(3, 5)
	require.NoError(t, err)
	err = b.SetItem(setCoordinate, testItem)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isEmpty := b.IsEmpty(test.coord)
			require.Equal(t, test.isEmpty, isEmpty)
		})
	}
}

func TestToPass(t *testing.T) {
	require.Equal(t, true, 2 == 2)
}
