package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpsaasprovisioning"
	"testing"
)

func TestProv(t *testing.T) {
	svc := btpsaasprovisioning.New(sess)

	input := &btpsaasprovisioning.GetApplicationSubscriptionsInput{}
	if out, err := svc.GetApplicationSubscriptions(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}
