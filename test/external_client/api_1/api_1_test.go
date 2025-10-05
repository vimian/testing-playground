package api1

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/vimian/testing-playground/pkg/types"
	"github.com/vimian/testing-playground/test/external_client/utils"
)

var (
	api1Url        string
	api1HealthPath string
)

func TestInit(t *testing.T) {
	api1Url = os.Getenv("API_1_URL")
	api1HealthPath = os.Getenv("API_1_HEALTH_PATH")
	endpoint := api1Url + api1HealthPath
	if err := utils.APIStatusOk(endpoint); err != nil {
		t.Fatalf("API 1 health check failed: %v", err)
	}
}

func TestEndpointMin(t *testing.T) {
	numbers := []float64{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	body, err := json.Marshal(numbers)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}
	var resp types.HTTPResponse

	endpoint := api1Url + "/min"
	if err := utils.PostJSON(endpoint, body, &resp); err != nil {
		t.Fatalf("failed to call /min: %v", err)
	}
	fmt.Print(resp.Result)
	if resp.Result != float64(2) {
		t.Errorf("expected min 2, got %v", resp.Result)
	}
}
