// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/matrices/cerca-go"
	"github.com/matrices/cerca-go/internal/testutil"
	"github.com/matrices/cerca-go/option"
)

func TestConnectionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.New(context.TODO(), cercago.ConnectionNewParams{
		APIKey: cercago.F("sk_live_..."),
		Owner: cercago.F[cercago.ConnectionOwnerUnionParam](cercago.ConnectionOwnerOrganizationConnectionOwnerParam{
			Type: cercago.F(cercago.ConnectionOwnerOrganizationConnectionOwnerTypeOrganization),
		}),
		Provider:     cercago.F("custom"),
		AccountLabel: cercago.F("primary"),
	})
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.List(context.TODO(), cercago.ConnectionListParams{
		OwnerType: cercago.F(cercago.ConnectionListParamsOwnerTypeFleet),
		Cursor:    cercago.F("cursor_abc123"),
		FleetID:   cercago.F("fleet_abc123"),
		Limit:     cercago.F("20"),
	})
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.Delete(
		context.TODO(),
		"9f063c59-2775-4614-8a68-22e5f90f92f3",
		cercago.ConnectionDeleteParams{
			OwnerType: cercago.F(cercago.ConnectionDeleteParamsOwnerTypeFleet),
			FleetID:   cercago.F("fleet_abc123"),
		},
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionAttachWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.Attach(
		context.TODO(),
		"agent_abc123",
		cercago.ConnectionAttachParams{
			ConnectionID: cercago.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: cercago.F(cercago.ToolConnectionMetadataParam{
				"foo": "string",
			}),
		},
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionDetach(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.Detach(
		context.TODO(),
		"agent_abc123",
		"9f063c59-2775-4614-8a68-22e5f90f92f3",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionListForAgent(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cercago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Connections.ListForAgent(context.TODO(), "agent_abc123")
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
