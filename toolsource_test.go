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

func TestToolSourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSources.New(
		context.TODO(),
		"fleet_abc123",
		cercago.ToolSourceNewParams{
			Source: cercago.ToolSourceNewParamsSourceCreateHTTPToolSourceRequest{
				Auth: cercago.F(cercago.ToolSourceAuthParam{
					"kind": "bar",
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{{
					"name":        "bar",
					"description": "bar",
					"inputSchema": "bar",
					"endpoint":    "bar",
				}}),
				Type:     cercago.F(cercago.ToolSourceNewParamsSourceCreateHTTPToolSourceRequestTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
					"timeoutMs":   "bar",
					"maxAttempts": "bar",
					"retryMode":   "bar",
				}),
			},
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

func TestToolSourceGet(t *testing.T) {
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
	_, err := client.ToolSources.Get(
		context.TODO(),
		"fleet_abc123",
		"toolsrc_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSources.Update(
		context.TODO(),
		"fleet_abc123",
		"toolsrc_abc123",
		cercago.ToolSourceUpdateParams{
			Source: cercago.ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequest{
				Auth: cercago.F(cercago.ToolSourceAuthParam{
					"kind": "bar",
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{{
					"name":        "bar",
					"description": "bar",
					"inputSchema": "bar",
					"endpoint":    "bar",
				}}),
				Type:     cercago.F(cercago.ToolSourceUpdateParamsSourceUpdateHTTPToolSourceRequestTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
					"timeoutMs":   "bar",
					"maxAttempts": "bar",
					"retryMode":   "bar",
				}),
			},
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

func TestToolSourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSources.List(
		context.TODO(),
		"fleet_abc123",
		cercago.ToolSourceListParams{
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

func TestToolSourceDelete(t *testing.T) {
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
	err := client.ToolSources.Delete(
		context.TODO(),
		"fleet_abc123",
		"toolsrc_abc123",
	)
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSourceListForAgent(t *testing.T) {
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
	_, err := client.ToolSources.ListForAgent(context.TODO(), "agent_abc123")
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
