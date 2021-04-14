package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpaccounts"
	"testing"
)

func TestGetGlobalAccount(t *testing.T) {
	svc := btpaccounts.New(sess)
	if out, err := svc.GetGlobalAccount(context.Background(), &btpaccounts.GetGlobalAccountInput{Expand: false}); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestUpdateGlobalAccount(t *testing.T) {
	svc := btpaccounts.New(sess)
	if out, err := svc.UpdateGlobalAccount(context.Background(), &btpaccounts.UpdateGlobalAccountInput{
		Description: "Technical Field Enablement",
		DisplayName: "Technical Field Enablement",
	}); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}
