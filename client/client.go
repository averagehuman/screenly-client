/*
HTTP client for interacting with a Screenly server

    client, err := client.DefaultInstance()

*/
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	//"os"
	"time"
)

type ScreenlyClient struct {
	httpClient *http.Client
	BaseUrl    *url.URL
}

// Add an asset of type 'webpage' to the screenly playlist
func (sc *ScreenlyClient) AddWebPage(name string, uri string, duration int64, expiry int64) (*Asset, error) {
	asset := new(AssetPayload)
	asset.Name = name
	asset.MimeType = "webpage"
	asset.Uri = uri
	asset.Duration = duration
	asset.Start = time.Now().UTC()
	asset.End = asset.Start.Add(time.Second * time.Duration(expiry))
	asset.IsEnabled = "1"
	asset.IsActive = true
	return sc.Post(asset)
}

// Return the current Screenly list of assets as a PlayList object
func (sc *ScreenlyClient) List() *PlayList {
	playlist := &PlayList{}
	// The assets endpoint returns a JSON list not a JSON object, so the
	// response body can't be decoded directly to a PlayList. So we have
	// to unmarshal to the PlayList.Assets field.
	response, err := sc.doHttp("GET", "assets", nil)

	if err == nil {
		//io.Copy(os.Stdout, response.Body)
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

// Return the asset with the given id
func (sc *ScreenlyClient) Get(id string) *Asset {
	asset := &Asset{}
	path := fmt.Sprintf("assets/%s", id)
	response, err := sc.doHttp("GET", path, nil)
	if err == nil {
		err = json.NewDecoder(response.Body).Decode(asset)
		if err == nil {
			return asset
		}
	}
	panic(err)
}

// Add an asset to the playlist
func (sc *ScreenlyClient) Post(payload *AssetPayload) (*Asset, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err == nil {
		path := "assets"
		response, err := sc.doHttp("POST", path, b)
		if err == nil {
			received := &Asset{}
			//io.Copy(os.Stdout, response.Body)
			err = json.NewDecoder(response.Body).Decode(received)
			if err == nil {
				return received, nil
			}
		}
	}
	return nil, err
}

// Update an existing asset
func (sc *ScreenlyClient) Put(id string, payload *AssetPayload) (*Asset, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err == nil {
		path := fmt.Sprintf("assets/%s", id)
		response, err := sc.doHttp("PUT", path, b)
		if err == nil {
			received := &Asset{}
			//io.Copy(os.Stdout, response.Body)
			err = json.NewDecoder(response.Body).Decode(received)
			if err == nil {
				return received, nil
			}
		}
	}
	return nil, err
}

// Delete an asset from the playlist
func (sc *ScreenlyClient) Delete(id string) error {
	path := fmt.Sprintf("assets/%s", id)
	_, err := sc.doHttp("DELETE", path, nil)
	return err
}

// Private method for making HTTP requests to the Screenly Server
func (sc *ScreenlyClient) doHttp(method string, path string, body io.Reader) (*http.Response, error) {
	url, err := sc.BaseUrl.Parse(path)
	if err == nil {
		req, err := http.NewRequest(method, url.String(), body)
		if err == nil {
			return sc.httpClient.Do(req)
		}
	}
	return nil, err
}
