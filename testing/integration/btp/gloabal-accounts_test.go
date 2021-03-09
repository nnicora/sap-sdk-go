package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btp"
	"github.com/nnicora/sap-sdk-go/service/btp/btpglobalaccounts"
	"testing"
)

func TestGetGlobalAccounts(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		params := &btpglobalaccounts.GlobalAccountParams{Expand: true}
		if acc, errAcc := client.GlobalAccounts.GetRestApiWithParams(ctx, params); errAcc != nil {
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

func TestUpdateGlobalAccount(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		req := &btpglobalaccounts.GlobalAccountRequest{
			DisplayName: "Technical Field Enablement",
			Description: "Technical Field Enablement",
		}
		if acc, errAcc := client.GlobalAccounts.UpdateRestApi(ctx, req); errAcc != nil {
			t.Error(errAcc)
		} else {
			if data, err := json.Marshal(acc); err != nil {
				t.Error(err)
			} else {
				t.Logf("\nSuccess Response: %s\n", string(data))
			}
			if rscc, err := client.GlobalAccounts.GetRestApiWithParams(ctx, nil); err != nil {
				if rscc.Description != req.Description {
					t.Error("Global account description wasn't updated")
				}
				if rscc.DisplayName != req.DisplayName {
					t.Error("Global account name wasn't updated")
				}
			}
		}
	} else {
		t.Error(err)
	}
}
