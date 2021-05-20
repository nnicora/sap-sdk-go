package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/service/btpprovisioning"
	"github.com/nnicora/sap-sdk-go/service/btpsaasprovisioning"
	"testing"
)

func TestProv(t *testing.T) {
	svc := btpsaasprovisioning.New(sess)

	input := &btpsaasprovisioning.GetApplicationSubscriptionsInput{}
	if out, err := svc.GetApplicationSubscriptions(context.Background(), input); err != nil {
		t.Error(err)
		t.Logf(out.ErrorDescription)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestAppConsumers(t *testing.T) {
	svc := btpsaasprovisioning.New(sess)

	input := &btpsaasprovisioning.GetEntitledApplicationsInput{}
	if out, err := svc.GetEntitledApplications(context.Background(), input); err != nil {
		t.Error(err)
		t.Logf(out.ErrorDescription)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestRegisterApp(t *testing.T) {
	svc := btpsaasprovisioning.New(sess)

	input := &btpsaasprovisioning.SubscribeToApplicationInput{
		AppName:  "test",
		PlanName: "default",
	}
	if out, err := svc.SubscribeToApplication(context.Background(), input); err != nil {
		t.Error(err)
		if out.Error != nil {
			t.Logf(*out.Error.Message)
		}
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}

func TestGetProvisioningEnv(t *testing.T) {
	svc := btpprovisioning.New(sess)

	if out, err := svc.GetEnvironmentInstances(context.Background()); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}
	}
}
