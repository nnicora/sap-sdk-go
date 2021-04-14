package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpevents"
	"testing"
)

func TestGetEvents(t *testing.T) {
	svc := btpevents.New(sess)
	in := &btpevents.GetEventsInput{
		EntityType: []string{"Subaccount", "Directory", "Tenant"},
		PageNum:    1,
		PageSize:   100,
	}
	if out, err := svc.GetEvents(context.Background(), in); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestGetEventsTypes(t *testing.T) {
	svc := btpevents.New(sess)
	if out, err := svc.GetEventsTypes(context.Background()); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}
