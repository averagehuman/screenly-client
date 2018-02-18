/*
HTTP client for interacting with a Screenly server

    client, err := client.DefaultInstance()

*/
package client

import (
	"bytes"
	"encoding/json"
	//"io"
	"net/http"
	"net/url"
	//"os"
)

type ScreenlyClient struct {
	httpClient *http.Client
	BaseUrl    *url.URL
}

// Return the current Screenly list of assets as a PlayList object
func (sc *ScreenlyClient) GetAssets() *PlayList {
	playlist := &PlayList{}
	// The assets endpoint returns a JSON list not a JSON object, so the
	// response body can't be decoded directly to a PlayList. So we have
	// to unmarshal the asset list explicitly.
	response, err := sc.get("assets")

	if err == nil {
		// Create a buffer and read response body, eg. [{...}, {...}]
		b := new(bytes.Buffer)
		// (the first ignored parameter is the number of bytes read)
		_, err := b.ReadFrom(response.Body)

		if err == nil {
			// ...now unmarshal to the PlayList.Assets slice.
			err := json.Unmarshal(b.Bytes(), &playlist.Assets)
			if err == nil {
				return playlist
			}
		}
	}
	panic(err)
}

// Private method for making GET requests
func (sc *ScreenlyClient) get(path string) (*http.Response, error) {
	url, err := sc.BaseUrl.Parse(path)
	if err == nil {
		req, err := http.NewRequest("GET", url.String(), nil)
		if err == nil {
			return sc.httpClient.Do(req)
		}
	}
	return nil, err
}
