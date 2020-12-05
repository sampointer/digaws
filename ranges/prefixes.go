package ranges

import (
	"fmt"
)

// Prefix is an interface that both IPv4 and IPv6 structs implement
type Prefix interface {
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
		"prefix: %s region: %s service: %s, network_border_group: %s",
		p.IPPrefix,
		p.Region,
		p.Service,
		p.NetworkBorderGroup,
	)
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
		"prefix: %s region: %s service: %s, network_border_group: %s",
		p.IPPrefix,
		p.Region,
		p.Service,
		p.NetworkBorderGroup,
	)
}
