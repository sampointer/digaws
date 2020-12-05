package command

import (
	"testing"

	"github.com/sampointer/digaws/ranges"
	"github.com/stretchr/testify/require"
)

const ipv4 = "52.94.76.5"
const ipv6 = "2a05:d07a:a0ff:ffff:ffff:ffff:ffff:aaaa"

func TestLookup(t *testing.T) {
	t.Parallel()
	t.Run("looks up IPv4 address", func(t *testing.T) {
		t.Parallel()
		prefix := ranges.PrefixIPv4{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		p, err := Lookup(ipv4)
		require.NoError(t, err)
		require.Equal(t, 1, len(p))
		require.Equal(t, prefix, p[0])
	})

	t.Run("looks up IPv6 address", func(t *testing.T) {
		t.Parallel()
		prefix1 := ranges.PrefixIPv6{
			IPPrefix:           "2a05:d07a:a000::/40",
			Region:             "eu-south-1",
			Service:            "AMAZON",
			NetworkBorderGroup: "eu-south-1",
		}
		prefix2 := ranges.PrefixIPv6{
			IPPrefix:           "2a05:d07a:a000::/40",
			Region:             "eu-south-1",
			Service:            "S3",
			NetworkBorderGroup: "eu-south-1",
		}

		p, err := Lookup(ipv6)
		require.NoError(t, err)
		require.Equal(t, 2, len(p))
		require.Equal(t, prefix1, p[0])
		require.Equal(t, prefix2, p[1])
	})
}
