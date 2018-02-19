package client

import "fmt"
import "testing"

func TestDefaultClient(t *testing.T) {
	client, err := DefaultInstance()
	if err != nil {
		fmt.Errorf("Failed to create default ScreenlyClient: %s", err)
	}
	if client.BaseUrl.Host != "127.0.0.1" {
		fmt.Errorf("%s", err)
	}

	playlist := client.List()

	if !playlist.IsEmpty() {
		t.Errorf("Expected initial PlayList to be empty but it has %d items", len(playlist.Assets))
	}

	// Add a webpage asset
	asset, err := client.AddWebPage("BBC Home", "http://bbc.co.uk", 10, 60)
	fmt.Println(asset.MimeType)

	if err != nil {
		t.Errorf("Failed to add webpage asset - %s", err)
	}

	if len(asset.Id) == 0 {
		t.Error("Got null asset id for posted asset")
	}
	if asset.MimeType != "webpage" {
		t.Errorf("Unexpected mimetype. Expected 'webpage', got '%s'", asset.MimeType)
	}
	if asset.Uri != "http://bbc.co.uk" {
		t.Errorf("Unexpected uri. Expected 'http://bbc.co.uk', got '%s'", asset.Uri)
	}

	playlist = client.List()

	if playlist.Size() != 2 {
		t.Errorf("Expected PlayList to return one item but it returned %d", playlist.Size())
	}
}
