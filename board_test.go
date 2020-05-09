package board

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoard(t *testing.T) {
	err := Board()
	require.NoError(t, err)
}
