package ranges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"

// Ranges represents an https://ip-ranges.amazonaws.com/ip-ranges.json document
type Ranges struct {
	CreateDate       string `json:createDate`
	CreateDateParsed time.Time
	Prefixes         []Prefix `json:prefixes`
	IPv6Prefixes     []Prefix `json:ipv6_prefixes`
	SyncToken        string   `json:syncToken`
	SyncTokenParsed  time.Time
}

// Prefix holds the detail of a given AWS prefix
type Prefix struct {
	IPPrefix           string `json:ip_prefix`
	Region             string `json:region`
	Service            string `json:service`
	NetworkBorderGroup string `json:network_border_group`
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

	d, err := parseCreateDate(&ranges.CreateDate)
	if err != nil {
		return nil, err
	}
	ranges.CreateDateParsed = d

	s, err := strconv.ParseInt(ranges.SyncToken, 10, 64)
	if err != nil {
		return nil, err
	}
	ranges.SyncTokenParsed = time.Unix(s, 0).UTC()

	if ranges.CreateDateParsed != ranges.SyncTokenParsed {
		return nil, fmt.Errorf(
			"syncToken and createDate do not match: %s, %s",
			ranges.SyncTokenParsed,
			ranges.CreateDateParsed,
		)
	}

	return &ranges, nil
}

func parseCreateDate(s *string) (time.Time, error) {
	const createDateFormat = "2006-01-02-15-04-05"
	t, err := time.Parse(createDateFormat, *s)
	return t, err
}
