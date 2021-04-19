package btpsaasprovisioning

import (
	"context"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/endpoints"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/sap/http/request/processors/jsonbuiltin"
	"github.com/nnicora/sap-sdk-go/sap/metainfo"
	"github.com/nnicora/sap-sdk-go/sap/service"
)

type SaaSProvisioningV1 struct {
	*service.Requester
}

const (
	ServiceName = "SaaS Provisioning V1" // Label of service.
	EndpointsID = "saas-manager"         // ID to lookup a service endpoint with.
	ServiceID   = "saas-manager"         // ServiceID is a unique identifier of a specific service.
)

func New(p service.RequesterConfig) *SaaSProvisioningV1 {
	c, err := p.ServiceConfig(EndpointsID)
	if err != nil {
		c.Processors.Using(request.Validate).PushFrontHandler(func(t interface{}) {
			r := t.(*request.Request)
			r.Error = err
		})
	}
	return newRequester(c.RuntimeConfig, c.Processors, c.Endpoint)
}

func newRequester(cfg *sap.RuntimeConfig, p *processors.Processors, endpoint *endpoints.Endpoint) *SaaSProvisioningV1 {
	svc := &SaaSProvisioningV1{
		Requester: service.NewRequester(
			cfg,
			metainfo.ServiceInfo{
				ServiceName: ServiceName,
				ServiceID:   ServiceID,
				Endpoint:    endpoint,
				APIVersion:  "v1",
			},
			p,
		),
	}

	// Processors
	p.Using(request.Build).
		PushBack(&jsonbuiltin.BuildProcessor).
		PushBack(&jsonbuiltin.MarshalToRequestJSONBodyProcessor)

	p.Using(request.Unmarshal).
		PushBack(&jsonbuiltin.UnmarshalResponseJSONBodyProcessor)

	p.Using(request.UnmarshalError).
		PushBack(&jsonbuiltin.UnmarshalErrorResponseJSONBodyProcessor)

	p.Using(request.UnmarshalMeta).
		PushBack(&jsonbuiltin.UnmarshalMetaProcessor)

	return svc
}

func (svc *SaaSProvisioningV1) newRequest(ctx context.Context, op *request.Operation, in, out interface{}) *request.Request {
	return svc.NewRequest(ctx, op, in, out)
}
