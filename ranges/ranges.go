package ranges

import (
	"fmt"
	"ioutil"
	"json"
	"net/http"
	"time"
)

const url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"

// Ranges represents an https://ip-ranges.amazonaws.com/ip-ranges.json document
type Ranges struct {
	CreateDate    time.Time
	CreateDateRaw string `json::createDate`
	Prefixes      []Prefix
	SyncToken     time.Time
	SyncTokenRaw  int64 `json:syncToken`
}

// Prefix holds the detail of a given AWS IPv4 prefix
type Prefix struct {
	IPPrefix           string `json:ip_prefix`
	Region             string `json:region`
	Service            string `json:service`
	NetworkBorderGroup string `json:network_border_group`
}

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

	ranges.SyncToken = time.Unix(ranges.SyncTokenRaw, 0)

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
	const createDateFormat = "2006-01-02-05-04-15"
	t, err := time.Parse(createDateFormat, *s)
	return t, err
}
