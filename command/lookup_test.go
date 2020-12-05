package command

import (
	"testing"

	"github.com/sampointer/digaws/ranges"
	"github.com/stretchr/testify/require"
)

const ip = "52.94.76.5"

func TestLookup(t *testing.T) {
	t.Run("looks up IPv4 address", func(t *testing.T) {
		prefix := ranges.PrefixIPv4{
			IPPrefix:           "52.94.76.0/22",
			Region:             "us-west-2",
			Service:            "AMAZON",
			NetworkBorderGroup: "us-west-2",
		}

		p, err := Lookup(ip)
		require.NoError(t, err)
		require.Equal(t, 1, len(p))
		require.Equal(t, prefix, p[0])
	})
}
