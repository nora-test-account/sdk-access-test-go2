// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestprojectrepoaccess_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nora-test-project-repo-access-go"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/internal/testutil"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/option"
)

func TestFridgeAddItem(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fridge.AddItem(context.TODO(), noratestprojectrepoaccess.FridgeAddItemParams{
		Food: noratestprojectrepoaccess.FoodParam{
			Name:     "name",
			Quantity: "quantity",
		},
	})
	if err != nil {
		var apierr *noratestprojectrepoaccess.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeDeleteItem(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fridge.DeleteItem(context.TODO(), "name")
	if err != nil {
		var apierr *noratestprojectrepoaccess.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeListItems(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fridge.ListItems(context.TODO())
	if err != nil {
		var apierr *noratestprojectrepoaccess.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeGetItem(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fridge.GetItem(context.TODO(), "name")
	if err != nil {
		var apierr *noratestprojectrepoaccess.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeUpdateItem(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fridge.UpdateItem(
		context.TODO(),
		"name",
		noratestprojectrepoaccess.FridgeUpdateItemParams{
			Food: noratestprojectrepoaccess.FoodParam{
				Name:     "name",
				Quantity: "quantity",
			},
		},
	)
	if err != nil {
		var apierr *noratestprojectrepoaccess.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
