package manifest

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetManifest(t *testing.T) {
	t.Parallel()

	t.Run("manifest is valid JSON", func(t *testing.T) {
		res, err := ioutil.ReadAll(GetManifest())
		require.NoError(t, err)
		require.True(t, json.Valid(res))
	})
}
