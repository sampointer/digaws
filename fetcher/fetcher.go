package fetcher

import (
	"io"
	"net/http"
)

const url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"

//Fetch downloads the ip ranages
func Fetch() (io.Reader, error) {
	var blank io.Reader
	var client http.Client

	res, err := client.Get(url)
	if err != nil {
		return blank, err
	}

	return res.Body, nil
}
