package btpresources

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

type ResourceV1 struct {
	*service.Requester
}

const (
	ServiceName = "Resource Consumption V1" // Label of service.
	EndpointsID = "resources"               // ID to lookup a service endpoint with.
	ServiceID   = "reports"                 // ServiceID is a unique identifier of a specific service.
)

func New(p service.RequesterConfig) *ResourceV1 {
	c, err := p.ServiceConfig(EndpointsID)
	if err != nil {
		c.Processors.Using(request.Validate).PushBack(func(t interface{}) {
			r := t.(*request.Request)
			r.Error = err
		})
	}
	return newRequester(c.RuntimeConfig, c.Processors, c.Endpoint)
}

func newRequester(cfg *sap.RuntimeConfig, p *processors.Processors, endpoint *endpoints.Endpoint) *ResourceV1 {
	svc := &ResourceV1{
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
		PushBackNamed(&jsonbuiltin.BuildProcessor).
		PushBackNamed(&jsonbuiltin.MarshalToRequestJSONBodyProcessor)

	p.Using(request.Unmarshal).
		PushBackNamed(&jsonbuiltin.UnmarshalResponseJSONBodyProcessor)

	p.Using(request.UnmarshalMeta).
		PushBackNamed(&jsonbuiltin.UnmarshalMetaProcessor)

	return svc
}

func (svc *ResourceV1) newRequest(ctx context.Context, op *request.Operation, in, out interface{}) *request.Request {
	return svc.NewRequest(ctx, op, in, out)
}
