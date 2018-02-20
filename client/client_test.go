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
		t.Fatalf(
			"The initial PlayList is not empty (has %d items). "+
				"Make sure to restart the screenly docker container between each test run.", len(playlist.Assets))
	}

	// Add a webpage asset
	asset, err := client.AddWebPage("BBC Home", "http://bbc.co.uk", 10, 60)

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

	// The asset list endpoint should return the one asset just added.
	playlist = client.List()

	if playlist.Size() != 1 {
		t.Errorf("Expected PlayList to return one item but it returned %d.", playlist.Size())
	}
}
