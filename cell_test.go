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

func TestCellNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cells.New(context.TODO(), cercago.CellNewParams{
		UserID: cercago.F("userId"),
		Configuration: cercago.F(cercago.CellConfigurationParam{
			ApprovalRequired:      cercago.F([]cercago.RuntimeToolName{cercago.RuntimeToolNameWebSearch}),
			ApprovalTimeoutMs:     cercago.F(0.000000),
			DefaultModel:          cercago.F("defaultModel"),
			ExcludedToolSourceIDs: cercago.F([]string{"string"}),
			Features:              cercago.F([]cercago.CellConfigurationFeature{cercago.CellConfigurationFeatureMemory}),
			Instructions:          cercago.F("instructions"),
			ToolApprovalOverrides: cercago.F(map[string]cercago.RuntimeToolApprovalOverrideMode{
				"foo": cercago.RuntimeToolApprovalOverrideModeAlways,
			}),
			ToolSourceMode: cercago.F(cercago.CellConfigurationToolSourceModeInherit),
			UserContext:    cercago.F("userContext"),
		}),
		EnvironmentID: cercago.F("environmentId"),
		Metadata: cercago.F(cercago.CellMetadataParam{
			"project": "alpha",
		}),
	})
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCellGet(t *testing.T) {
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
	_, err := client.Cells.Get(context.TODO(), "cell_abc123")
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCellUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cells.Update(
		context.TODO(),
		"cell_abc123",
		cercago.CellUpdateParams{
			Configuration: cercago.F(cercago.CellConfigurationParam{
				ApprovalRequired:      cercago.F([]cercago.RuntimeToolName{cercago.RuntimeToolNameWebSearch}),
				ApprovalTimeoutMs:     cercago.F(0.000000),
				DefaultModel:          cercago.F("defaultModel"),
				ExcludedToolSourceIDs: cercago.F([]string{"string"}),
				Features:              cercago.F([]cercago.CellConfigurationFeature{cercago.CellConfigurationFeatureMemory}),
				Instructions:          cercago.F("instructions"),
				ToolApprovalOverrides: cercago.F(map[string]cercago.RuntimeToolApprovalOverrideMode{
					"foo": cercago.RuntimeToolApprovalOverrideModeAlways,
				}),
				ToolSourceMode: cercago.F(cercago.CellConfigurationToolSourceModeInherit),
				UserContext:    cercago.F("userContext"),
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

func TestCellListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cells.List(context.TODO(), cercago.CellListParams{
		Active:        cercago.F(cercago.CellListParamsActiveTrue),
		Cursor:        cercago.F("cursor_abc123"),
		EnvironmentID: cercago.F("env_abc123"),
		Limit:         cercago.F("20"),
	})
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCellDelete(t *testing.T) {
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
	_, err := client.Cells.Delete(context.TODO(), "cell_abc123")
	if err != nil {
		var apierr *cercago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
