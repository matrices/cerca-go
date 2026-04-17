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

func TestEnvironmentToolSourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.ToolSources.New(
		context.TODO(),
		"env_abc123",
		cercago.EnvironmentToolSourceNewParams{
			Body: cercago.EnvironmentToolSourceNewParamsBodyObject{
				Auth: cercago.F[any](map[string]interface{}{
					"kind": "none",
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{map[string]interface{}{
					"name":        "search",
					"description": "Search documents",
					"inputSchema": map[string]interface{}{
						"type": "object",
					},
					"endpoint": map[string]interface{}{
						"method": "GET",
						"url":    "https://docs.example.com/search",
					},
				}}),
				Type:     cercago.F(cercago.EnvironmentToolSourceNewParamsBodyObjectTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F[any](map[string]interface{}{
					"timeoutMs":   10000,
					"maxAttempts": 3,
					"retryMode":   "safe_only",
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

func TestEnvironmentToolSourceGet(t *testing.T) {
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
	_, err := client.Environments.ToolSources.Get(
		context.TODO(),
		"env_abc123",
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

func TestEnvironmentToolSourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.ToolSources.Update(
		context.TODO(),
		"env_abc123",
		"toolsrc_abc123",
		cercago.EnvironmentToolSourceUpdateParams{
			Body: cercago.EnvironmentToolSourceUpdateParamsBodyObject{
				Auth: cercago.F[any](map[string]interface{}{
					"kind": "none",
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{map[string]interface{}{
					"name":        "search",
					"description": "Search documents",
					"inputSchema": map[string]interface{}{
						"type": "object",
					},
					"endpoint": map[string]interface{}{
						"method": "GET",
						"url":    "https://docs.example.com/search",
					},
				}}),
				Type:     cercago.F(cercago.EnvironmentToolSourceUpdateParamsBodyObjectTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F[any](map[string]interface{}{
					"timeoutMs":   10000,
					"maxAttempts": 3,
					"retryMode":   "safe_only",
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

func TestEnvironmentToolSourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.ToolSources.List(
		context.TODO(),
		"env_abc123",
		cercago.EnvironmentToolSourceListParams{
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

func TestEnvironmentToolSourceDelete(t *testing.T) {
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
	err := client.Environments.ToolSources.Delete(
		context.TODO(),
		"env_abc123",
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
