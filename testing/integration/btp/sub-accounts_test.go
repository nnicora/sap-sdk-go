package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btp"
	"github.com/nnicora/sap-sdk-go/service/btp/btpsubaccounts"
	"testing"
)

func TestGetSubAccounts(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		if acc, errAcc := client.SubAccounts.GetAllRestApi(ctx); errAcc != nil {
			t.Error(errAcc)
		} else {
			if data, err := json.Marshal(acc); err != nil {
				t.Error(err)
			} else {
				t.Logf("\nSuccess Response: %s\n", string(data))
			}
		}
	} else {
		t.Error(err)
	}
}

func TestCreateSubAccounts(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		req := &btpsubaccounts.SubAccountRequest{
			BetaEnabled:       true,
			Description:       "Terraform Managed",
			DisplayName:       "Terraform",
			ParentGUID:        globalAccountGuuid,
			Region:            "eu10",
			Subdomain:         "terraform-create",
			UsedForProduction: "UNSET",
			Origin:            "REGION_SETUP",
		}
		if acc, errAcc := client.SubAccounts.CreateRestApi(ctx, req); errAcc != nil {
			t.Error(errAcc)
		} else {
			client.SubAccounts.DeleteRestApi(ctx, acc.Guid)
			if data, err := json.Marshal(acc); err != nil {
				t.Error(err)
			} else {
				t.Logf("\nSuccess Response: %s\n", string(data))
			}
		}
	} else {
		t.Error(err)
	}
}

func TestDeleteSubAccounts(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		req := &btpsubaccounts.SubAccountRequest{
			BetaEnabled:       true,
			Description:       "Terraform Managed",
			DisplayName:       "Terraform",
			ParentGUID:        globalAccountGuuid,
			Region:            "eu10",
			Subdomain:         "terraform-delete",
			UsedForProduction: "UNSET",
		}
		if acc, err := client.SubAccounts.CreateRestApi(ctx, req); err != nil {
			t.Error(err)
		} else {
			if delAcc, err := client.SubAccounts.DeleteRestApi(ctx, acc.Guid); err != nil {
				t.Error(err)
			} else {
				if data, err := json.Marshal(delAcc); err != nil {
					t.Error(err)
				} else {
					t.Logf("\nSuccess Response: %s\n", string(data))
				}
			}
		}
	} else {
		t.Error(err)
	}
}
