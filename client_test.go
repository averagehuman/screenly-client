package client

import "fmt"
import "testing"

func TestDefaultClient(t *testing.T) {
	client, err := DefaultClient()
	if err != nil {
		fmt.Errorf("Failed to create default ScreenlyClient: %s", err)
	}
	if client.BaseUrl.Host != "127.0.0.1" {
		fmt.Errorf("Unexpected default host. Expected 127.0.0.1, got %s", client.BaseUrl.Host)
	}
	if client.BaseUrl.Port() != "8080" {
		fmt.Errorf("Unexpected default port. Expected 8080, got %s", client.BaseUrl.Host)
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

	// The item returned in the list should be identical to the asset as saved.
	item := playlist.Assets[0]

	if item.Id != asset.Id {
		t.Error("Asset Id of list item doesn't match the Id of the object created.")
	}
	if item.MimeType != asset.MimeType {
		t.Error("Asset MimeType of list item doesn't match the MimeType of the object created.")
	}
	if item.Uri != asset.Uri {
		t.Error("Asset Uri of list item doesn't match the Uri of the object created.")
	}

	// Delete the item
	err = client.Delete(asset.Id)
	if err != nil {
		t.Errorf("The delete method returned an error: %s", err)
	}

	// The playlist should now be empty.
	playlist = client.List()
	if playlist.Size() != 0 {
		t.Errorf("Expected PlayList to be empty but it has size %d.", playlist.Size())
	}

}
