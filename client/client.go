/*
HTTP client for interacting with a Screenly server

    client, err := client.DefaultInstance()

*/
package client

import "net/http"
import "net/url"

type ScreenlyClient struct {
	httpClient *http.Client
	BaseUrl    *url.URL
}

// Return the current list of assets
func (sc *ScreenlyClient) GetAssets() (*http.Response, error) {
	return sc.get("assets")
}

// Private method for making GET requests
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
