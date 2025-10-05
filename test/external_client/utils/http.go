package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vimian/testing-playground/pkg/types"
)

func getWithContext(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

// APIStatusOk checks if the API at the given endpoint returns a 200 OK status.
func APIStatusOk(endpoint string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var response types.HTTPResponse
	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout reached without status ok")
		default:
			time.Sleep(1 * time.Second)
			resp, err := getWithContext(ctx, endpoint)
			if err != nil {
				fmt.Println("request error:", err)
				continue
			}
			defer resp.Body.Close()
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				fmt.Println("decode error:", err)
				continue
			}
			if response.Status == "ok" {
				return nil
			}
		}
	}
}

// PostJSON sends a POST request with a JSON body and decodes the JSON response.
func PostJSON(endpoint string, body []byte, result interface{}) error {
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}
