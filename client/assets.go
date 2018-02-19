package client

import "time"

type Asset struct {
	Id           string    `json:"asset_id"`
	Name         string    `json:"name"`
	Uri          string    `json:"uri"`
	Start        time.Time `json:"start_date"`
	End          time.Time `json:"end_date"`
	Duration     int64     `json:"duration"`
	MimeType     string    `json:"mimetype"`
	IsEnabled    int       `json:"is_enabled,omitempty"`
	IsProcessing int       `json:"is_processing,omitempty"`
	NoCache      int       `json:"nocache,omitempty"`
	PlayOrder    int       `json:"play_order,omitempty"`
}

type PlayList struct {
	Assets []Asset
}

func (p *PlayList) IsEmpty() bool {
	return len(p.Assets) == 0
}

func (p *PlayList) Add(asset Asset) {
	p.Assets = append(p.Assets, asset)
}
