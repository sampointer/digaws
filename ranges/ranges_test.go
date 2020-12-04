package ranges

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRanges(t *testing.T) {
	client := new(http.Client)
	ranges, err := New(client)
	require.NoError(t, err)

	t.Run("has IPv4 prefixes", func(t *testing.T) {
		require.NotZero(t, ranges.Prefixes, "should have 1 or more prefixes")
	})

	t.Run("has IPv6 prefixes", func(t *testing.T) {
		require.NotZero(t, ranges.IPv6Prefixes, "should have 1 or more prefixes")
	})
}
