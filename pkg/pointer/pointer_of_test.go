package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGetPointerOfString(t *testing.T) {
	foo := "foo"
	require.Equal(t, &foo, Of(foo))
}

func TestCanGetPointerOfBool(t *testing.T) {
	foo := true
	require.Equal(t, &foo, Of(foo))
}

func TestCanGetPointerOfInt(t *testing.T) {
	foo := 42
	require.Equal(t, &foo, Of(foo))
}
