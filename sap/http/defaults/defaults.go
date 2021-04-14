package defaults

import (
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/sap/http/request/coreprocessors"
)

func Processors() processors.Processors {
	ps := processors.New()
	ps.Using(request.Validate).
		PushBack(&coreprocessors.ValidateEndpointProcessor).
		StopOnError()
	ps.Using(request.Build).
		StopOnError()
	ps.Using(request.Sign).
		PushBack(&coreprocessors.BuildContentLengthProcessor)
	ps.Using(request.Send).
		PushBack(&coreprocessors.ValidateReqSigProcessor).
		PushBack(&coreprocessors.SendProcessor)
	ps.Using(request.ValidateResponse).
		PushBack(&coreprocessors.ValidateResponseProcessor)
	return ps
}
