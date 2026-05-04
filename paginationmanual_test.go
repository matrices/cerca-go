// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago_test

import (
	"context"
	"os"
	"testing"

	"github.com/matrices/cerca-go"
	"github.com/matrices/cerca-go/internal/testutil"
	"github.com/matrices/cerca-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Agents.List(context.TODO(), cercago.AgentListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, agent := range page.Agents {
		t.Logf("%+v\n", agent.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, agent := range page.Agents {
			t.Logf("%+v\n", agent.ID)
		}
	}
}
