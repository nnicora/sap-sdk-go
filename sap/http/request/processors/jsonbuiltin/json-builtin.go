package jsonbuiltin

import (
	"bytes"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/internal/saperr"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
)

var BuildProcessor = processors.DefaultProcessor{
	Name:    "sap.json.builtin.Build",
	Handler: Build,
}

var MarshalToRequestJSONBodyProcessor = processors.DefaultProcessor{
	Name:    "sap.json.builtin.MarshalToJSONRequestBody",
	Handler: MarshalToJSONRequestBody,
}

var UnmarshalResponseJSONBodyProcessor = processors.DefaultProcessor{
	Name:    "sap.json.builtin.UnmarshalJSONResponseBody",
	Handler: UnmarshalJSONResponseBody,
}

var UnmarshalMetaProcessor = processors.DefaultProcessor{
	Name:    "sap.json.builtin.UnmarshalMeta",
	Handler: UnmarshalMeta,
}

func Build(t interface{}) {
	r := t.(*request.Request)
	if v := r.HTTPRequest.Header.Get("Content-Type"); len(v) == 0 {
		r.HTTPRequest.Header.Set("Content-Type", "application/json")
	}
}

func MarshalToJSONRequestBody(t interface{}) {
	r := t.(*request.Request)
	if r.InputData != nil {
		body := &bytes.Buffer{}
		if err := json.NewEncoder(body).Encode(r.InputData); err != nil {
			r.Error = err
			return
		} else {
			r.SetBytesBody(body.Bytes())
		}
	}
}

func UnmarshalJSONResponseBody(t interface{}) {
	r := t.(*request.Request)
	if r.OutputData != nil {
		if err := json.Unmarshal(r.ResponseBody, r.OutputData); err != nil {
			r.Error = saperr.New(saperr.Serialization, "failed reading request body", err)
			return
		}
	}
}

func UnmarshalMeta(t interface{}) {
	//r := t.(*request.Request)
}
