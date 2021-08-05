package fetcher

import (
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestFetcher(t *testing.T) {
	t.Parallel()

	// Setup mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://ip-ranges.amazonaws.com/ip-ranges.json",
		httpmock.NewBytesResponder(
			200, httpmock.File("../testdata/ip-ranges.json").Bytes(),
		),
	)

	// Perform the request
	f, err := Fetch()
	require.NoError(t, err)

	d, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	re := regexp.MustCompile("syncToken")
	res := re.FindString(string(d))
	require.NotZero(t, res)

	require.Equal(t, 1, httpmock.GetTotalCallCount())
}
