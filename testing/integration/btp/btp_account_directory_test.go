package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpaccounts"
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
