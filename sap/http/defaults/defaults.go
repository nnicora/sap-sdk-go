package defaults

import (
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/sap/http/request/coreprocessors"
)

func Processors() processors.Processors {
	ps := processors.New()
	ps.Using(request.Validate).
		PushBackNamed(&coreprocessors.ValidateEndpointProcessor).
		StopOnError()
	ps.Using(request.Build).
		StopOnError()
	ps.Using(request.Sign).
		PushBackNamed(&coreprocessors.BuildContentLengthProcessor)
	ps.Using(request.Send).
		PushBackNamed(&coreprocessors.ValidateReqSigProcessor).
		PushBackNamed(&coreprocessors.SendProcessor)
	ps.Using(request.ValidateResponse).
		PushBackNamed(&coreprocessors.ValidateResponseProcessor)
	return ps
}
