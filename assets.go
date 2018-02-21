package client

import "time"

// Represents a Screenly Asset. Use this when unmarshalling JSON received from the server.
// We can't use the AssetPayload type for sending and receiving because the server can
// return integers and empty strings for some boolean fields and this will break the
// unmarshaller which requires definite types in general. Any json field that doesn't
// have an associated struct field will be ignored by the unmarshaller.
type Asset struct {
	Id       string    `json:"asset_id"`
	Name     string    `json:"name"`
	Uri      string    `json:"uri"`
	Start    time.Time `json:"start_date"`
	End      time.Time `json:"end_date"`
	MimeType string    `json:"mimetype"`
	IsActive bool      `json:"is_active,omitempty"`
}

// Screenly Asset payload. Use this when sending JSON to the server.
type AssetPayload struct {
	Asset
	Duration  int64  `json:"duration"`
	IsEnabled string `json:"is_enabled,omitempty"`
	NoCache   string `json:"nocache,omitempty"`
}

// A list of Asset summaries.
type PlayList struct {
	Assets []Asset
}

func (p *PlayList) IsEmpty() bool {
	return len(p.Assets) == 0
}

func (p *PlayList) Size() int {
	return len(p.Assets)
}
