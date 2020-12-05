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
		require.NotZero(t, ranges.PrefixesIPv4, "should have 1 or more prefixes")
	})

	t.Run("has IPv6 prefixes", func(t *testing.T) {
		require.NotZero(t, ranges.PrefixesIPv6, "should have 1 or more prefixes")
	})

	t.Run("returns a Prefix struct for an IPv4 address", func(t *testing.T) {
		prefix := PrefixIPv4{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		ip := net.ParseIP("52.94.76.5")
		results, err := ranges.FindForIPv4(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for invalid IPv4 address", func(t *testing.T) {
		ip := net.ParseIP("1.2.3.4")
		results, err := ranges.FindForIPv4(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})

	//	t.Run("returns a Prefix struct for an IPv6 address", func(t *testing.T) {
	//		prefix := PrefixIPv6{
	//			IPPrefix:           "2600:1f70:4000::/56",
	//			Region:             "us-west-2",
	//			Service:            "AMAZON",
	//			NetworkBorderGroup: "us-west-2",
	//		}
	//
	//		ip := net.ParseIP("2600:1f70:4000:3:4")
	//		results, err := ranges.FindForIPv6(ip)
	//		require.NoError(t, err)
	//		require.Equal(t, prefix, results[0])
	//	})
	//
	t.Run("returns no Prefix struct for invalid IPv6 address", func(t *testing.T) {
		ip := net.ParseIP("1:2:3:4:5")
		results, err := ranges.FindForIPv6(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})
}
