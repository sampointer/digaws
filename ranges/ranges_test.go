package ranges

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRanges(t *testing.T) {
	t.Parallel()
	doc, err := os.Open("../test/ip-ranges.json")
	require.NoError(t, err)

	ranges, err := New(doc)
	require.NoError(t, err)

	t.Run("has IPv4 prefixes", func(t *testing.T) {
		t.Parallel()
		require.NotZero(t, ranges.PrefixesIPv4, "should have 1 or more prefixes")
	})

	t.Run("has IPv6 prefixes", func(t *testing.T) {
		t.Parallel()
		require.NotZero(t, ranges.PrefixesIPv6, "should have 1 or more prefixes")
	})

	t.Run("returns a Prefix struct for an IPv4 address", func(t *testing.T) {
		t.Parallel()
		prefix := PrefixIPv4{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		ip := net.ParseIP("52.94.76.5")
		results, err := ranges.LookupIPv4(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for non-AWS IPv4 address", func(t *testing.T) {
		t.Parallel()
		ip := net.ParseIP("1.2.3.4")
		results, err := ranges.LookupIPv4(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})

	t.Run("returns a Prefix struct for an IPv6 address", func(t *testing.T) {
		t.Parallel()
		prefix := PrefixIPv6{
			IPPrefix:           "2a05:d07a:a000::/40",
			Region:             "eu-south-1",
			Service:            "AMAZON",
			NetworkBorderGroup: "eu-south-1",
		}

		ip := net.ParseIP("2a05:d07a:a0ff:ffff:ffff:ffff:ffff:aaaa")
		results, err := ranges.LookupIPv6(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for non-AWS IPv6 address", func(t *testing.T) {
		t.Parallel()
		ip := net.ParseIP("1:2:3:4:5")
		results, err := ranges.LookupIPv6(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})
}
