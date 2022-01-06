package dns

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/carlmjohnson/requests"
)

var (
	IPServices []string = []string{
		"ifconfig.co",
		"ipify.org",
	}
)

func callService(service string) (net.IP, error) {
	var response string

	switch service {
	case "ifconfig.co":
		err := requests.
			URL("https://ifconfig.co/ip").
			ToString(&response).
			Fetch(context.Background())

		if err != nil {
			return nil, err
		}

		response = strings.TrimSpace(response)
	case "ipify.org":
		err := requests.
			URL("https://api.ipify.org").
			ToString(&response).
			Fetch(context.Background())

		if err != nil {
			return nil, err
		}

		response = strings.TrimSpace(response)
	}

	ip := net.ParseIP(response)
	if ip == nil {
		return nil, errors.New("could not parse ip from service response")
	}

	return ip, nil
}

func GetIPByService() (net.IP, error) {
	for _, service := range IPServices {
		ip, _ := callService(service)
		if ip != nil {
			return ip, nil
		}
	}

	return nil, errors.New("could not get an IP address via external services")
}
