package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpaccounts"
	"github.com/nnicora/sap-sdk-go/service/btpentitlements"
	"testing"
)

func TestCreateAccountDirectory(t *testing.T) {
	svc := btpaccounts.New(sess)

	features := make([]string, 2)
	features[0] = "DEFAULT"
	features[1] = "ENTITLEMENTS"
	dirInput := &btpaccounts.CreateDirectoryInput{
		DisplayName:       "data_terraform",
		Subdomain:         "terraform",
		DirectoryFeatures: features,
	}
	if out, err := svc.CreateDirectory(context.Background(), dirInput); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestEntitlementsGlobalAccount(t *testing.T) {
	svc := btpentitlements.New(sess)
	if out, err := svc.GetGlobalAccountAssignments(context.Background(), &btpentitlements.GlobalAccountAssignmentsInput{
		AcceptLanguage:          "en",
		SubAccountGuid:          "24d360fd-8e28-48a3-ab69-f574f388761b",
		IncludeAutoManagedPlans: true,
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

func TestEntitlements(t *testing.T) {
	svc := btpentitlements.New(sess)
	if out, err := svc.GetAssignments(context.Background(), &btpentitlements.GetAssignmentsInput{
		SubAccountGuid:          "6cd6350e-d589-48ca-94f0-8d68c8559c01",
		IncludeAutoManagedPlans: true,
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
