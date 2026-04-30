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

func TestUsage(t *testing.T) {
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
	thread, err := client.Threads.New(
		context.TODO(),
		"agent_abc123",
		cercago.ThreadNewParams{
			Message: cercago.F("What's on my calendar today?"),
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", thread.ID)
}
