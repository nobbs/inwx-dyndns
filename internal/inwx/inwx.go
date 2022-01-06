package inwx

import (
	"context"
	"log"
	"net"

	"github.com/carlmjohnson/requests"
)

type INWX struct {
	endpoint string
	user     string
	password string
}

func New(user, password string) *INWX {
	inwx := &INWX{
		endpoint: "https://dyndns.inwx.com/nic/update?myip=<ipaddr>",
		user:     user,
		password: password,
	}

	return inwx
}

func (inwx *INWX) UpdateIP(ip net.IP) {
	err := requests.URL(inwx.endpoint).
		Param("myip", ip.To4().String()).
		BasicAuth(inwx.user, inwx.password).
		Fetch(context.Background())

	if err != nil {
		log.Fatalf("Update of DNS entry failed, %v", err)
	}
}
