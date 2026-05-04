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

func TestWebhookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.New(
		context.TODO(),
		"fleet_abc123",
		cercago.WebhookNewParams{
			URL:    cercago.F("https://example.com/webhook"),
			Events: cercago.F([]cercago.WebhookSubscriptionEventType{cercago.WebhookSubscriptionEventTypeAgentCreated}),
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

func TestWebhookGet(t *testing.T) {
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
	_, err := client.Webhooks.Get(
		context.TODO(),
		"fleet_abc123",
		"webhook_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.Update(
		context.TODO(),
		"fleet_abc123",
		"webhook_abc123",
		cercago.WebhookUpdateParams{
			Enabled: cercago.F(true),
			Events:  cercago.F([]cercago.WebhookSubscriptionEventType{cercago.WebhookSubscriptionEventTypeAgentCreated}),
			URL:     cercago.F("https://example.com/webhook"),
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

func TestWebhookListWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.List(
		context.TODO(),
		"fleet_abc123",
		cercago.WebhookListParams{
			Cursor: cercago.F("cursor_abc123"),
			Limit:  cercago.F("20"),
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

func TestWebhookDelete(t *testing.T) {
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
	err := client.Webhooks.Delete(
		context.TODO(),
		"fleet_abc123",
		"webhook_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookRotate(t *testing.T) {
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
	_, err := client.Webhooks.Rotate(
		context.TODO(),
		"fleet_abc123",
		"webhook_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookTest(t *testing.T) {
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
	_, err := client.Webhooks.Test(
		context.TODO(),
		"fleet_abc123",
		"webhook_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
