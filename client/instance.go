package client

import "time"
import "net/url"
import "net/http"

// Utility factory function for creating a ScreenlyClient
func Instance(config Config) (*ScreenlyClient, error) {
	url, err := url.Parse(config.ENDPOINT)
	if err != nil {
		return nil, err
	}
	timeout := time.Duration(config.TIMEOUT) * time.Second
	httpClient := &http.Client{Timeout: timeout}
	return &ScreenlyClient{httpClient: httpClient, BaseUrl: url}, nil
}

// Utility factory function for creating a ScreenlyClient with default configuration
func DefaultInstance() (*ScreenlyClient, error) {
	client, err := Instance(defaultConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
