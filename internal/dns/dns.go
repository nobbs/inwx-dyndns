package dns

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"
)

var (
	nameservers []string = []string{
		"9.9.9.9",
		"8.8.8.8",
		"1.1.1.1",
		"8.8.4.4",
		"1.0.0.1",
	}
)

func GetIPByDNS(host string) (net.IP, error) {
	nsIndex := rand.Intn(len(nameservers))
	ns := fmt.Sprintf("%s:%d", nameservers[nsIndex], 53)

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, ns)
		},
	}

	ipStrings, err := r.LookupHost(context.Background(), host)
	if err != nil {
		e := fmt.Errorf("unable to obtain IP address by dns resolver, %w", err)
		return nil, e
	}
	if len(ipStrings) == 0 {
		return nil, errors.New("the DNS resolver returned no IP address")
	}

	ip := net.ParseIP(ipStrings[0])
	return ip, nil
}
