package ranges

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"time"
)

// Ranges represents an https://ip-ranges.amazonaws.com/ip-ranges.json document
type Ranges struct {
	CreateDate    time.Time    `json:"-"`
	CreateDateRaw string       `json:"createDate"`
	PrefixesIPv4  []PrefixIPv4 `json:"prefixes"`
	PrefixesIPv6  []PrefixIPv6 `json:"ipv6_prefixes"`
	SyncToken     time.Time    `json:"-"`
	SyncTokenRaw  string       `json:"syncToken"`
}

// LookupIPv4 returns the Prefix structs that contain a range that includes the
// passed IPv4 address
func (r *Ranges) LookupIPv4(ip net.IP) ([]Prefix, error) {
	var results []Prefix

	for _, p := range r.PrefixesIPv4 {
		_, pIPNet, err := net.ParseCIDR(p.IPPrefix)
		if err != nil {
			return nil, err
		}

		if pIPNet.Contains(ip) {
			results = append(results, p)
		}
	}

	return results, nil
}

// LookupIPv6 returns the Prefix structs that contain a range that includes the
// passed IPv6 address
func (r *Ranges) LookupIPv6(ip net.IP) ([]Prefix, error) {
	var results []Prefix

	for _, p := range r.PrefixesIPv6 {
		_, pIPNet, err := net.ParseCIDR(p.IPPrefix)
		if err != nil {
			return nil, err
		}

		if pIPNet.Contains(ip) {
			results = append(results, p)
		}
	}

	return results, nil
}

//New is a constructor for Ranges
func New(r io.Reader) (*Ranges, error) {
	var ranges Ranges

	doc, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(doc, &ranges)
	if err != nil {
		return nil, err
	}

	d, err := parseCreateDate(&ranges.CreateDateRaw)
	if err != nil {
		return nil, err
	}
	ranges.CreateDate = d

	s, err := strconv.ParseInt(ranges.SyncTokenRaw, 10, 64)
	if err != nil {
		return nil, err
	}
	ranges.SyncToken = time.Unix(s, 0).UTC()

	if ranges.CreateDate != ranges.SyncToken {
		return nil, fmt.Errorf(
			"syncToken and createDate do not match: %s, %s",
			ranges.SyncToken,
			ranges.CreateDate,
		)
	}

	return &ranges, nil
}

func parseCreateDate(s *string) (time.Time, error) {
	const createDateFormat = "2006-01-02-15-04-05"
	t, err := time.Parse(createDateFormat, *s)
	return t, err
}
