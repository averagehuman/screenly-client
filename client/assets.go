package client

import "time"

type Asset struct {
	Id        string    `json:"asset_id"`
	Name      string    `json:"name"`
	Uri       string    `json:"uri"`
	Start     time.Time `json:"start_date"`
	End       time.Time `json:"end_date"`
	Duration  int64     `json:"duration"`
	MimeType  string    `json:"mimetype"`
	IsActive  bool      `json:"is_active,omitempty"`
	IsEnabled string    `json:"is_enabled,omitempty"`
	NoCache   string    `json:"nocache,omitempty"`
}

type AssetSummary struct {
	Id       string    `json:"asset_id"`
	Name     string    `json:"name"`
	Uri      string    `json:"uri"`
	Start    time.Time `json:"start_date"`
	End      time.Time `json:"end_date"`
	MimeType string    `json:"mimetype"`
	IsActive bool      `json:"is_active,omitempty"`
}
type PlayList struct {
	Assets []AssetSummary
}

func (p *PlayList) IsEmpty() bool {
	return len(p.Assets) == 0
}

func (p *PlayList) Size() int {
	return len(p.Assets)
}
