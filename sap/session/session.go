package session

import (
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/endpoints"
	"github.com/nnicora/sap-sdk-go/sap/http/defaults"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
	"github.com/nnicora/sap-sdk-go/sap/service"
	"net/http"
)

type RuntimeSession struct {
	RuntimeConfig *sap.RuntimeConfig
	Processors    processors.Processors
}

// Preparing the service configuration, coming out of runtime session
func (s *RuntimeSession) ServiceConfig(serviceId string) (*service.Config, error) {
	baseProcessors := s.Processors.Copy()
	cfg := &service.Config{
		RuntimeConfig: s.RuntimeConfig,
		Processors:    &baseProcessors,

		Endpoint: &endpoints.Endpoint{},
	}

	v, ok := s.RuntimeConfig.Endpoints[serviceId]
	if !ok {
		return cfg, fmt.Errorf("endpoint not identified for service '%s'", serviceId)
	}
	if err := utils.IsValidUrl(v.Host); err != nil {
		return cfg, err
	}
	cfg.Endpoint.Host = v.Host
	cfg.Endpoint.Client = v.Client
	return cfg, nil
}

// Light Update of RuntimeSession Configuration, by skiping existent endpoint, added only new one from config.
func (s *RuntimeSession) LightUpdate(c *sap.Config) error {
	return s.update(c, true)
}

// Hard Update of RuntimeSession Configuration, by updating all configuration, replacing and old one.
func (s *RuntimeSession) HardUpdate(c *sap.Config) error {
	return s.update(c, false)
}

func (s *RuntimeSession) AddEndpoint(serviceId string, endpointConfig *sap.EndpointConfig) error {
	if endpoint, err := s.createEndpoint(endpointConfig); err != nil {
		return err
	} else {
		if _, ok := s.RuntimeConfig.Endpoints[serviceId]; ok {
			return fmt.Errorf("endpoint mapped to '%s' id already exist", serviceId)
		}
		s.RuntimeConfig.Endpoints[serviceId] = endpoint
		return nil
	}
}

// Update the existent RuntimeSession Configuration; Used for cases when new endpoints was added into configuration and
// session should be updated to have it too.
func (s *RuntimeSession) update(c *sap.Config, light bool) error {
	defaultHttpClient, defErr := oauth2.NewOAuth2Client(c.DefaultOAuth2)
	for k, v := range c.Endpoints {
		if _, ok := s.RuntimeConfig.Endpoints[k]; light && ok {
			continue
		}
		if endpoint, err := s.createEndpointWithDefault(v, defaultHttpClient); err != nil {
			return err
		} else {
			if endpoint.Client == defaultHttpClient && defErr != nil {
				return defErr
			}
			s.RuntimeConfig.Endpoints[k] = endpoint
		}
	}

	keys := make([]string, 0)
	for k, _ := range s.RuntimeConfig.Endpoints {
		if _, ok := c.Endpoints[k]; ok {
			continue
		} else {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(s.RuntimeConfig.Endpoints, k)
	}

	return nil
}
func (s *RuntimeSession) createEndpointWithDefault(ec *sap.EndpointConfig, defaultHttpClient *http.Client) (*endpoints.Endpoint, error) {
	if httpClient, err := createOAuth2ClientWithDefault(ec.OAuth2, defaultHttpClient); err != nil {
		return nil, err
	} else {
		return &endpoints.Endpoint{
			Host:   ec.Host,
			Client: httpClient,
		}, nil
	}
}
func (s *RuntimeSession) createEndpoint(ec *sap.EndpointConfig) (*endpoints.Endpoint, error) {
	if httpClient, err := createOAuth2Client(ec.OAuth2); err != nil {
		return nil, err
	} else {
		return &endpoints.Endpoint{
			Host:   ec.Host,
			Client: httpClient,
		}, nil
	}
}

// Build new RuntimeSession from sap.Config data
func BuildFromConfig(c *sap.Config) (*RuntimeSession, error) {
	rs := &RuntimeSession{
		RuntimeConfig: &sap.RuntimeConfig{
			Endpoints:  make(map[string]*endpoints.Endpoint),
			MaxRetries: c.MaxRetries,
		},
		Processors: defaults.Processors(),
	}
	if err := rs.HardUpdate(c); err != nil {
		return nil, err
	}
	return rs, nil
}

// Create http.Client from oauth2.Config
func createOAuth2ClientWithDefault(cfg *oauth2.Config, defaultHttpClient *http.Client) (*http.Client, error) {
	var client *http.Client
	if cfg != nil {
		if httpClient, err := createOAuth2Client(cfg); err != nil {
			return nil, err
		} else {
			client = httpClient
		}
	} else {
		client = defaultHttpClient
	}
	return client, nil
}

func createOAuth2Client(cfg *oauth2.Config) (*http.Client, error) {
	if httpClient, err := oauth2.NewOAuth2Client(cfg); err != nil {
		return nil, err
	} else {
		return httpClient, nil
	}
}
