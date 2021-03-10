package utils

import (
	"net"
	"net/url"
	"strings"
)

const (
	http  = "http://"
	https = "https://"
)

func CheckUrlSchema(host string) bool {
	if strings.HasPrefix(host, http) {
		return true
	}
	if strings.HasPrefix(host, https) {
		return true
	}
	return false
}

func VerifyHostIsAlive(host string) bool {
	if strings.HasPrefix(host, http) {
		return ipLookup(strings.TrimPrefix(host, http))
	}
	if strings.HasPrefix(host, https) {
		return ipLookup(strings.TrimPrefix(host, https))
	}
	return ipLookup(host)
}

func ipLookup(host string) bool {
	_, err := net.LookupIP(host)
	if err != nil {
		return false
	}
	return true
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
