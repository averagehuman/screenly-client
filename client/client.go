package client

import "time"
import "net/http"
import "net/url"

type Config struct {
	URL     string
	TIMEOUT uint
}

type ScreenlyClient struct {
	httpClient *http.Client
	BaseUrl    *url.URL
}

var defaultConfig = Config{
	URL:     "http://127.0.0.1:8080",
	TIMEOUT: 30,
}

func (sc *ScreenlyClient) get(path string) (*http.Response, error) {
	url, err := sc.BaseUrl.Parse(path)
	if err != nil {
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			return sc.httpClient.Do(req)
		}
	}
	return nil, err
}

func (sc *ScreenlyClient) GetAssets() (*http.Response, error) {
	return sc.get("assets")
}

func Instance(config Config) (*ScreenlyClient, error) {
	url, err := url.Parse(config.URL)
	if err != nil {
		return nil, err
	}
	timeout := time.Duration(config.TIMEOUT) * time.Second
	httpClient := &http.Client{Timeout: timeout}
	return &ScreenlyClient{httpClient: httpClient, BaseUrl: url}, nil
}

func DefaultInstance() (*ScreenlyClient, error) {
	client, err := Instance(defaultConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
