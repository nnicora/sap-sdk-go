package utils

import (
	"net"
	"net/url"
)

func IsDnsError(err error) bool {
	if urlError, ok := err.(*url.Error); ok {
		if opError, ok := urlError.Err.(*net.OpError); ok {
			if dnsError, ok := opError.Err.(*net.DNSError); ok {
				return dnsError.IsNotFound
			}
		}
	}
	return false
}

func GetDnsError(err error) error {
	if urlError, ok := err.(*url.Error); ok {
		if opError, ok := urlError.Err.(*net.OpError); ok {
			if dnsError, ok := opError.Err.(*net.DNSError); ok {
				return dnsError
			}
		}
	}
	return err
}
