package ranges

import (
	"encoding/json"
	"fmt"
)

// Prefix is an interface that both IPv4 and IPv6 structs implement
type Prefix interface {
	GetRegion() string
	JSON() (string, error)
	String() string
}

// PrefixIPv4 holds the detail of a given AWS IPv4 prefix
type PrefixIPv4 struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

//String returns a column-format representation of the Prefix
func (p PrefixIPv4) String() string {
	return fmt.Sprintf(
		"prefix: %s region: %s service: %s network_border_group: %s",
		p.IPPrefix,
		p.Region,
		p.Service,
		p.NetworkBorderGroup,
	)
}

//JSON returns a string of JSON representing the Prefix
func (p PrefixIPv4) JSON() (string, error) {
	out, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

//GetRegion is an accessor method for the Region field, present in order that
//the Prefix interface can expose it.
func (p PrefixIPv4) GetRegion() string {
	return p.Region
}

// PrefixIPv6 holds the detail of a given AWS IPv6 prefix
type PrefixIPv6 struct {
	IPPrefix           string `json:"ipv6_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

//String returns a column-format representation of the Prefix
func (p PrefixIPv6) String() string {
	return fmt.Sprintf(
		"prefix: %s region: %s service: %s network_border_group: %s",
		p.IPPrefix,
		p.Region,
		p.Service,
		p.NetworkBorderGroup,
	)
}

//JSON returns a string of JSON representing the Prefix
func (p PrefixIPv6) JSON() (string, error) {
	out, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

//GetRegion is an accessor method for the Region field, present in order that
//the Prefix interface can expose it.
func (p PrefixIPv6) GetRegion() string {
	return p.Region
}
