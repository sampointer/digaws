package fetcher

import (
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetcher(t *testing.T) {
	f, err := Fetch()
	require.NoError(t, err)

	d, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	re := regexp.MustCompile("syncToken")
	res := re.FindString(string(d))
	require.NotZero(t, res)
}
