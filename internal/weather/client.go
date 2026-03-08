package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
	apiKey  = "will use my api key"
	units   = "metric"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FetchCurrentWeather(city string) (*apiResponse, error) {
	endpoint := fmt.Sprintf(
		"%s?q=%s&appid=%s&units=%s",
		baseURL,
		url.QueryEscape(city),
		apiKey,
		units,
	)

	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to call weather API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status: %s", resp.Status)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &data, nil
}
