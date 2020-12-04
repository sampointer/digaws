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
	CreateDateRaw string    `json:"createDate"`
	CreateDate    time.Time `json:"-"`
	Prefixes      []Prefix  `json:"prefixes"`
	IPv6Prefixes  []Prefix  `json:"ipv6_prefixes"`
	SyncTokenRaw  string    `json:"syncToken"`
	SyncToken     time.Time `json:"-"`
}

// Prefix holds the detail of a given AWS prefix
type Prefix struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

// FindForIP returns the Prefix structs that contain a range that includes the
// passed IP address
func (r *Ranges) FindForIP(ip net.IP) ([]Prefix, error) {
	var results []Prefix

	for _, p := range r.Prefixes {
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
