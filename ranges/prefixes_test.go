package ranges

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixString(t *testing.T) {
	t.Run("IPv4", func(t *testing.T) {
		prefix := PrefixIPv4{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		require.Equal(
			t,
			"prefix: 52.94.76.0/22 region: us-west-2 service: AMAZON, network_border_group: us-west-2",
			prefix.String(),
		)
	})

	t.Run("IPv6", func(t *testing.T) {
		prefix := PrefixIPv6{
			IPPrefix:           "2a05:d07a:a000::/40",
			Region:             "eu-south-1",
			Service:            "AMAZON",
			NetworkBorderGroup: "eu-south-1",
		}

		require.Equal(
			t,
			"prefix: 2a05:d07a:a000::/40 region: eu-south-1 service: AMAZON, network_border_group: eu-south-1",
			prefix.String(),
		)
	})

}
