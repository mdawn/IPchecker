# IP Capture & Check 

An API which uses [GeoLite2 Databases](https://dev.maxmind.com/geoip/geoip2/geolite2/).

## Description

Send this API an IP address and a list of [GeoName country codes](http://www.geonames.org/countries/), and it will answer two questions:
1. What country code does this IP address come from?
2. Does this country's IP address match the countries in my list?

This can be useful if you want to restrict users from logging into an account. Practical applications include general compliance as well as preventing labor outsourcing.

## What it Does

At http://localhost:8080, this API will:

1. Map the IP address to the country
2. Compare that country to a given list, 
and 
3. Return a status of `APPROVED` or `DENIED` 

## Example Scenario

I want to approve only IP addresses from Canada, the US, the UK, Sweden, and Australia.
A user with the IP address of `1.1.4.22` wants to log into my UI. 

**STEP 1**: Clone this repository. Cd into the directory. Run `go get github.com/oschwald/geoip2-golang` & `go run main.go`.

**STEP 2**: Look up the [2-letter GeoName country codes](http://www.geonames.org/countries/) for the countries on the whitelist. 

In our case: 
* Canada = CA
* the US = US
* the UK = GB
* Sweden = SE
* Australia = AU

**STEP 3**: Enter the country codes (separated by commas, no spaces) and incoming IP address as a URL like so:

 `http://localhost:8080/?whitelist=CA,US,GB,SE,AU&ip=1.1.4.22` 

The browser returns:

`STATUS = DENIED, Incoming IP 1.1.4.22 maps to ISO country code CN WHITELIST REFERENCE = CA,CN,US,GB,AU`

Uh oh! Looks like our incoming IP address is from ountry code `CN`. That's not in our list. DENIED!


## What Would Make It Better <br>

* Robust error handling
* A [GraphQL](https://graphql.org/) endpoint