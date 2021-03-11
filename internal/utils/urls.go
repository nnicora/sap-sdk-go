package utils

import (
	"net"
	"net/url"
)

func HostAlive(host string) (bool, error) {
	u, err := url.Parse(host)
	if err != nil {
		return false, err
	}
	if u.Host == "" {
		return ipLookup(u.Path)
	}
	return ipLookup(u.Host)
}

func ipLookup(host string) (bool, error) {
	_, err := net.LookupIP(host)
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsUrl(str string) (bool, error) {
	if u, err := url.Parse(str); err != nil {
		return false, err
	} else {
		return u.Scheme != "" && u.Host != "", nil
	}
}
