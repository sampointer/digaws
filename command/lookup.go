package command

import (
	"github.com/sampointer/digaws/ranges"

	"io"
	"net"
	"strings"
)

const url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"

//Lookup returns Prefixes the ranges of which include the passed IP address
func Lookup(q string, doc io.Reader) ([]ranges.Prefix, error) {
	var prefixes []ranges.Prefix

	r, err := ranges.New(doc)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(q)

	if isIPv4(ip) {
		p, err := r.LookupIPv4(ip)
		if err != nil {
			return nil, err
		}
		prefixes = append(prefixes, p...)
	} else {
		p, err := r.LookupIPv6(ip)
		if err != nil {
			return nil, err
		}
		prefixes = append(prefixes, p...,
		)
	}

	return prefixes, nil
}

func isIPv4(ip net.IP) bool {
	return strings.Contains(ip.String(), ".")
}
