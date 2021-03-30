package service

import (
	"context"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/endpoints"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/sap/metainfo"
)

type Config struct {
	RuntimeConfig *sap.RuntimeConfig

	Processors *processors.Processors
	Endpoint   *endpoints.Endpoint
}

type Requester struct {
	ServiceInfo metainfo.ServiceInfo

	RuntimeConfig *sap.RuntimeConfig
	Processors    *processors.Processors
}

type RequesterConfig interface {
	ServiceConfig(serviceId string) (*Config, error)
}

func NewRequester(cfg *sap.RuntimeConfig, info metainfo.ServiceInfo, processors *processors.Processors, options ...func(*Requester)) *Requester {
	svc := &Requester{
		RuntimeConfig: cfg,
		ServiceInfo:   info,
		Processors:    processors,
	}

	for _, option := range options {
		option(svc)
	}
	return svc
}

func (r *Requester) NewRequest(ctx context.Context, op *request.Operation, in interface{}, out interface{}) *request.Request {
	return request.New(ctx, r.RuntimeConfig, r.ServiceInfo, r.Processors, op, in, out)
}
