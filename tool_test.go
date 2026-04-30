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

func TestToolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.New(
		context.TODO(),
		"fleet_abc123",
		cercago.ToolNewParams{
			Tool: cercago.ToolNewParamsToolCreateHTTPToolSourceRequest{
				Auth: cercago.F[cercago.ToolSourceAuthUnionParam](cercago.ToolSourceAuthNoToolSourceAuthParam{
					Kind: cercago.F(cercago.ToolSourceAuthNoToolSourceAuthKindNone),
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{{
					Description: cercago.F("Search documents"),
					Endpoint: cercago.F(cercago.HTTPEndpointParam{
						Method: cercago.F(cercago.HTTPEndpointMethodGet),
						URL:    cercago.F("https://docs.example.com/search"),
						Body:   cercago.F(cercago.HTTPEndpointBodyJsonParams),
						Headers: cercago.F(map[string]string{
							"foo": "string",
						}),
						Path: cercago.F(map[string]string{
							"foo": "params.query",
						}),
						Query: cercago.F(map[string]string{
							"foo": "params.query",
						}),
					}),
					InputSchema: cercago.F(map[string]interface{}{
						"type": "bar",
					}),
					Name:     cercago.F("search"),
					Approval: cercago.F(cercago.ToolApprovalModeAlways),
					Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
						IdempotencyKeyHeader: cercago.F("idempotencyKeyHeader"),
						MaxAttempts:          cercago.F(3.000000),
						RetryMode:            cercago.F(cercago.HTTPToolExecutionPolicyRetryModeSafeOnly),
						TimeoutMs:            cercago.F(10000.000000),
					}),
					Response: cercago.F(cercago.ResponseNormalizationHintParam{
						Mode: cercago.F(cercago.ResponseNormalizationHintModeAuto),
					}),
				}}),
				Type:     cercago.F(cercago.ToolNewParamsToolCreateHTTPToolSourceRequestTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
					IdempotencyKeyHeader: cercago.F("idempotencyKeyHeader"),
					MaxAttempts:          cercago.F(3.000000),
					RetryMode:            cercago.F(cercago.HTTPToolExecutionPolicyRetryModeSafeOnly),
					TimeoutMs:            cercago.F(10000.000000),
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

func TestToolGet(t *testing.T) {
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
	_, err := client.Tools.Get(
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

func TestToolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Update(
		context.TODO(),
		"fleet_abc123",
		"toolsrc_abc123",
		cercago.ToolUpdateParams{
			Tool: cercago.ToolUpdateParamsToolUpdateHTTPToolSourceRequest{
				Auth: cercago.F[cercago.ToolSourceAuthUnionParam](cercago.ToolSourceAuthNoToolSourceAuthParam{
					Kind: cercago.F(cercago.ToolSourceAuthNoToolSourceAuthKindNone),
				}),
				Namespace: cercago.F("docs"),
				Tools: cercago.F([]cercago.HTTPToolDefinitionParam{{
					Description: cercago.F("Search documents"),
					Endpoint: cercago.F(cercago.HTTPEndpointParam{
						Method: cercago.F(cercago.HTTPEndpointMethodGet),
						URL:    cercago.F("https://docs.example.com/search"),
						Body:   cercago.F(cercago.HTTPEndpointBodyJsonParams),
						Headers: cercago.F(map[string]string{
							"foo": "string",
						}),
						Path: cercago.F(map[string]string{
							"foo": "params.query",
						}),
						Query: cercago.F(map[string]string{
							"foo": "params.query",
						}),
					}),
					InputSchema: cercago.F(map[string]interface{}{
						"type": "bar",
					}),
					Name:     cercago.F("search"),
					Approval: cercago.F(cercago.ToolApprovalModeAlways),
					Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
						IdempotencyKeyHeader: cercago.F("idempotencyKeyHeader"),
						MaxAttempts:          cercago.F(3.000000),
						RetryMode:            cercago.F(cercago.HTTPToolExecutionPolicyRetryModeSafeOnly),
						TimeoutMs:            cercago.F(10000.000000),
					}),
					Response: cercago.F(cercago.ResponseNormalizationHintParam{
						Mode: cercago.F(cercago.ResponseNormalizationHintModeAuto),
					}),
				}}),
				Type:     cercago.F(cercago.ToolUpdateParamsToolUpdateHTTPToolSourceRequestTypeHTTP),
				Approval: cercago.F(cercago.ToolApprovalModeAlways),
				Enabled:  cercago.F(true),
				Execution: cercago.F(cercago.HTTPToolExecutionPolicyParam{
					IdempotencyKeyHeader: cercago.F("idempotencyKeyHeader"),
					MaxAttempts:          cercago.F(3.000000),
					RetryMode:            cercago.F(cercago.HTTPToolExecutionPolicyRetryModeSafeOnly),
					TimeoutMs:            cercago.F(10000.000000),
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

func TestToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.List(
		context.TODO(),
		"fleet_abc123",
		cercago.ToolListParams{
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

func TestToolDelete(t *testing.T) {
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
	err := client.Tools.Delete(
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
