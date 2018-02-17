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

}
