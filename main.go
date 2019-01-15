package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	geoip2 "github.com/oschwald/geoip2-golang"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ipSlice := r.URL.Query()["ip"]
	if len(ipSlice) != 1 {
		return
	}

	ip := ipSlice[0]
	whitelist := r.URL.Query()["whitelist"][0]
	c := getCountry(ip)

	valid := inWhiteList(c, whitelist)
	if valid == true {
		fmt.Fprintf(w, " STATUS = APPROVED, ")
	} else {
		fmt.Fprintf(w, " STATUS = DENIED, ")
	}

	fmt.Fprintf(w, "Incoming IP "+ip)

	fmt.Fprintf(w, " maps to ISO country code "+c)

	fmt.Fprintf(w, " WHITELIST REFERENCE = "+whitelist)

	return

}

// Compares incoming IP address country_iso_code to whitelisted country_iso_codes
func inWhiteList(c string, whitelist string) bool {
	inList := (strings.Contains(whitelist, c))
	fmt.Println(inList)
	return inList
}

// Checks the geoip2 database for the incoming IP address and maps it to the country_iso_code
func getCountry(ip string) string {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	sentIP := net.ParseIP(ip)
	record, err := db.Country(sentIP)
	if err != nil {
		log.Fatal(err)
	}

	return record.Country.IsoCode
}
