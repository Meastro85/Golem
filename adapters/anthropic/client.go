package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiUrl = "https://api.anthropic.com/v1/messages"
const apiVersion = "2023-06-01"

type Client struct {
	APIKey     string
	Model      string
	HTTPClient *http.Client
}

func NewClient(apiKey, model string) *Client {
	return &Client{
		APIKey:     apiKey,
		Model:      model,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) send(ctx context.Context, req apiRequest) (*apiResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("anthropic: marshaling request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("anthropic: creating request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-Key", c.APIKey)
	httpReq.Header.Set("Anthropic-Version", apiVersion)

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("anthropic: request failed: %w", err)
	}
	defer httpResp.Body.Close()

	raw, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("anthropic: reading response body: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("anthropic: API error (status %d): %s", httpResp.StatusCode, raw)
	}

	var resp apiResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("anthropic: unmarshaling response: %w", err)
	}

	return &resp, nil
}
