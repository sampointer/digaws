package ranges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"

// Ranges represents an https://ip-ranges.amazonaws.com/ip-ranges.json document
type Ranges struct {
	CreateDate    time.Time    `json:"-"`
	CreateDateRaw string       `json:"createDate"`
	PrefixesIPv4  []PrefixIPv4 `json:"prefixes"`
	PrefixesIPv6  []PrefixIPv6 `json:"ipv6_prefixes"`
	SyncToken     time.Time    `json:"-"`
	SyncTokenRaw  string       `json:"syncToken"`
}

// PrefixIPv4 holds the detail of a given AWS IPv4 prefix
type PrefixIPv4 struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

// PrefixIPv6 holds the detail of a given AWS IPv6 prefix
type PrefixIPv6 struct {
	IPPrefix           string `json:"ipv6_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

// FindForIPv4 returns the Prefix structs that contain a range that includes the
// passed IPv4 address
func (r *Ranges) FindForIPv4(ip net.IP) ([]PrefixIPv4, error) {
	var results []PrefixIPv4

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

//New is a constructor for Ranges
func New(client *http.Client) (*Ranges, error) {
	var ranges Ranges

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ranges)
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
