package main

import (
	"log"

	"github.com/nobbs/inwx-dyndns/config"
	"github.com/nobbs/inwx-dyndns/internal/dns"
	"github.com/nobbs/inwx-dyndns/internal/inwx"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func main() {
	// setup new viper config object
	v := viper.New()

	// lets get the configuration
	c := config.Configure(v)

	// lets create the inwx object with the given credentials
	inwx := inwx.New(c.User, c.Password)
	// todo: validate that the credentials are correct

	// set up the cron object
	cron := cron.New()
	cron.AddFunc("@every 1m", func() {
		ipDNS, err := dns.GetIPByDNS("k3s.nobbs.eu")
		if err != nil {
			log.Printf("An error occured: %v", err)
		}

		ipService, err := dns.GetIPByService()
		if err != nil {
			log.Printf("An error occured: %v", err)
		}

		if !ipDNS.Equal(ipService) {
			log.Printf("IPs don't match, updating DNS.")
			inwx.UpdateIP(ipService)
		} else {
			log.Printf("IPs match, sleeping.")
		}
	})

	log.Print("Starting cron service")
	cron.Run()
}
