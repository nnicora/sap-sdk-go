package httplight

import "net/http"

func DefaultTransport() *http.Transport {
	transport := getCustomTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1
	return transport
}

func DefaultClient() *http.Client {
	return &http.Client{
		Transport: DefaultTransport(),
	}
}

func DefaultPooledClient() *http.Client {
	return &http.Client{
		Transport: getCustomTransport(),
	}
}
