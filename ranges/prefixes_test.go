package ranges

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var prefixIPv4 = PrefixIPv4{
	IPPrefix:           "52.94.76.0/22",
	Region:             "us-west-2",
	Service:            "AMAZON",
	NetworkBorderGroup: "us-west-2",
}

var prefixIPv6 = PrefixIPv6{
	IPPrefix:           "2a05:d07a:a000::/40",
	Region:             "eu-south-1",
	Service:            "AMAZON",
	NetworkBorderGroup: "eu-south-1",
}

func TestPrefixString(t *testing.T) {
	t.Parallel()
	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t,
			"prefix: 52.94.76.0/22 region: us-west-2 service: AMAZON network_border_group: us-west-2",
			prefixIPv4.String(),
		)
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t,
			"prefix: 2a05:d07a:a000::/40 region: eu-south-1 service: AMAZON network_border_group: eu-south-1",
			prefixIPv6.String(),
		)
	})
}

func TestPrefixJSON(t *testing.T) {
	t.Parallel()
	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		out, err := prefixIPv4.JSON()
		require.NoError(t, err)
		require.Equal(
			t,
			`{"ip_prefix":"52.94.76.0/22","region":"us-west-2","service":"AMAZON","network_border_group":"us-west-2"}`,
			out,
		)
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		out, err := prefixIPv6.JSON()
		require.NoError(t, err)
		require.Equal(
			t,
			`{"ipv6_prefix":"2a05:d07a:a000::/40","region":"eu-south-1","service":"AMAZON","network_border_group":"eu-south-1"}`,
			out,
		)
	})
}

func TestPrefixGetRegion(t *testing.T) {
	t.Parallel()
	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, prefixIPv4.GetRegion(), prefixIPv4.Region)
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, prefixIPv6.GetRegion(), prefixIPv6.Region)
	})
}
