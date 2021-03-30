package metainfo

import (
	"errors"
	"github.com/nnicora/sap-sdk-go/sap/endpoints"
	"net/http"
)

type ServiceInfo struct {
	ServiceName string
	ServiceID   string
	APIVersion  string
	Endpoint    *endpoints.Endpoint
}

func (e *ServiceInfo) EndpointHttpClient() (*http.Client, error) {
	if v, ok := e.Endpoint.Client.(*http.Client); !ok {
		return nil, errors.New("endpoint client is not of type http")
	} else {
		return v, nil
	}
}

func (e *ServiceInfo) EndpointHost() string {
	return e.Endpoint.Host
}
