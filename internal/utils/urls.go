package utils

import (
	"errors"
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

func IsValidUrl(str string) error {
	if u, err := url.Parse(str); err != nil {
		return err
	} else if u.Scheme == "" || u.Host == "" {
		return errors.New("invalid url, having schema of host empty")
	}
	return nil
}
