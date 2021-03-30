package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpaccounts"
	"testing"
)

func TestListSubAccounts(t *testing.T) {
	svc := btpaccounts.New(sess)
	if out, err := svc.GetSubAccounts(context.Background(), &btpaccounts.GetSubAccountsInput{}); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess Response: %s\n", string(data))
		}
	}
}

func TestCreateSubAccounts(t *testing.T) {
	svc := btpaccounts.New(sess)

	input := &btpaccounts.CreateSubAccountInput{
		BetaEnabled:       true,
		Description:       "Terraform Managed",
		DisplayName:       "Terraform",
		ParentGuid:        globalAccountGuid,
		Region:            "eu10",
		Subdomain:         "terraform-create",
		UsedForProduction: "UNSET",
		Origin:            "REGION_SETUP",
	}
	if out, err := svc.CreateSubAccount(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess Response: %s\n", string(data))
		}
	}
}

func TestDeleteSubAccounts(t *testing.T) {
	svc := btpaccounts.New(sess)

	input := &btpaccounts.CreateSubAccountInput{
		BetaEnabled:       true,
		Description:       "Terraform Managed",
		DisplayName:       "Terraform",
		ParentGuid:        globalAccountGuid,
		Region:            "eu10",
		Subdomain:         "terraform-create",
		UsedForProduction: "UNSET",
		Origin:            "REGION_SETUP",
	}
	if acc, err := svc.CreateSubAccount(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		input := &btpaccounts.DeleteSubAccountInput{
			SubAccountGuid: acc.Guid,
		}
		if outDeleted, err := svc.DeleteSubAccount(context.Background(), input); err != nil {
			t.Error(err)
		} else {
			if data, err := json.Marshal(outDeleted); err != nil {
				t.Error(err)
			} else {
				t.Logf("\nSuccess Response: %s\n", string(data))
			}
		}

	}
}

func TestGetSubAccountServiceManagementBinding(t *testing.T) {
	svc := btpaccounts.New(sess)

	input := &btpaccounts.GetServiceManagementBindingInput{
		SubAccountGuid: "bac3ed4a-d6f4-46f2-8fc9-2098ca4743a8",
	}
	if acc, err := svc.GetSubAccountServiceManagementBinding(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(acc); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess Response: %s\n", string(data))
		}
	}
}

func TestCreateSubAccountServiceManagementBinding(t *testing.T) {
	svc := btpaccounts.New(sess)

	input := &btpaccounts.CreateServiceManagementBindingInput{
		SubAccountGuid: "24d360fd-8e28-48a3-ab69-f574f388761b",
	}
	if acc, err := svc.CreateSubAccountServiceManagementBinding(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(acc); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess Response: %s\n", string(data))
		}
	}
}

func TestDeleteSubAccountServiceManagementBinding(t *testing.T) {
	svc := btpaccounts.New(sess)

	input := &btpaccounts.DeleteServiceManagementBindingInput{
		SubAccountGuid: "24d360fd-8e28-48a3-ab69-f574f388761b",
	}
	if acc, err := svc.DeleteSubAccountServiceManagementBinding(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(acc); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess Response: %s\n", string(data))
		}
	}
}
