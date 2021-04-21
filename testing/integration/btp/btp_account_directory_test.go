package btp

import (
	"context"
	"encoding/json"
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/service/btpaccounts"
	"github.com/nnicora/sap-sdk-go/service/btpentitlements"
	"testing"
)

func TestCreateAccountDirectory(t *testing.T) {
	svc := btpaccounts.New(sess)

	//{"directoryFeatures":["CRM","AUTHORIZATIONS","ENTITLEMENTS","DEFAULT"],"displayName":"resource from terraform"}

	features := make([]string, 2)
	features[0] = "DEFAULT"
	features[1] = "ENTITLEMENTS"
	//features[2] = "AUTHORIZATIONS"
	//features[3] = "CRM"
	dirInput := &btpaccounts.CreateDirectoryInput{
		DisplayName: "data_terraform",
		//Subdomain:         "terraform",
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

func TestGetEntitlements(t *testing.T) {
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

func TestUpdateEntitlements(t *testing.T) {
	svc := btpentitlements.New(sess)

	assignmentInfo := btpentitlements.AssignmentInfo{
		Amount: sap.Uint(2),
		//Enable:         sap.Bool(true),
		SubAccountGuid: "a1754d1f-a9da-4e6b-8989-356359b84d5b",
	}
	plan := btpentitlements.SubAccountServicePlan{
		ServiceName:     "enterprise-messaging",
		ServicePlanName: "default",
		AssignmentInfo:  []btpentitlements.AssignmentInfo{assignmentInfo},
	}

	input := &btpentitlements.UpdateSubAccountServicePlanInput{
		SubAccountServicePlans: []btpentitlements.SubAccountServicePlan{plan},
	}
	if out, err := svc.UpdateSubAccountServicePlan(context.Background(), input); err != nil {
		t.Error(err)
	} else {
		if data, err := json.Marshal(out); err != nil {
			t.Error(err)
		} else {
			t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
		}

		jobInput := &btpentitlements.GetJobStatusInput{
			JobId: sap.StringValue(out.JobStatusId),
		}
		if jobOut, err := svc.GetJobStatus(context.Background(), jobInput); err != nil {
			t.Error(err)
		} else {
			if data, err := json.Marshal(jobOut); err != nil {
				t.Error(err)
			} else {
				t.Logf("\nSuccess StatusAndBodyFromResponse: %s\n", string(data))
			}
		}

	}
}
