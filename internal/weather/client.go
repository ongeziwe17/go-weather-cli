package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
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
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENWEATHER_API_KEY environment variable not set")
	}

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
		return nil, fmt.Errorf("could not fetch weather for %s: %s", city, resp.Status)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &data, nil
}
