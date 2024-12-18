package helper

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	Client  *http.Client
	BaseURL string
	APIKey  string
}

func NewHttpClient(baseURL, APIKey string) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Timeout: 10 * time.Second, // Set a reasonable timeout
		},
		BaseURL: baseURL,
		APIKey:  APIKey,
	}
}
func (hc *HttpClient) Get(endpoint string) ([]byte, error) {
	url := hc.BaseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("x-api-key", hc.APIKey)

	resp, err := hc.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))

	// Check for non-OK status codes
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // Read error details from the response
		return nil, fmt.Errorf("API responded with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
