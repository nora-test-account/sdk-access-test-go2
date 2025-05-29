// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestprojectrepoaccess_test

import (
	"context"
	"os"
	"testing"

	"github.com/nora-test-account/sdk-access-test-go"
	"github.com/nora-test-account/sdk-access-test-go/internal/testutil"
	"github.com/nora-test-account/sdk-access-test-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := noratestprojectrepoaccess.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	foods, err := client.Fridge.ListItems(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", foods)
}
