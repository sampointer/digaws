package ranges

import (
	"net"
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

	t.Run("returns a Prefix struct for an IPv4 address", func(t *testing.T) {
		prefix := Prefix{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		ip := net.ParseIP("52.94.76.5")
		results, err := ranges.FindForIP(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for invalid IPv4 address", func(t *testing.T) {
		ip := net.ParseIP("1.2.3.4")
		results, err := ranges.FindForIP(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})

}
