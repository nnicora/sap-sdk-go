package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btp"
	"testing"
)

func TestGetRegions(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		if acc, errAcc := client.Entitlements.GetDataCentersRestApi(ctx); errAcc != nil {
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

func TestGetProviderRegions(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		if acc, errAcc := client.Entitlements.GetProvidersRegionsRestApi(ctx); errAcc != nil {
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

func TestGetProviderRegionsFilter(t *testing.T) {
	ctx := context.Background()

	if client, err := btp.NewBtpClient(ctx, cfg); err == nil {
		if acc, errAcc := client.Entitlements.GetProviderRegionsRestApi(ctx, "AWS"); errAcc != nil {
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
