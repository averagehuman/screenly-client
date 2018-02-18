package client

import "time"

type Asset struct {
	Id           string    `json:"asset_id"`
	Name         string    `json:"name"`
	Uri          string    `json:"uri"`
	Start        time.Time `json:"start_date"`
	End          time.Time `json:"end_date"`
	Duration     uint      `json:"duration"`
	MimeType     string    `json:"mimetype"`
	IsEnabled    bool      `json:"is_enabled,omitempty"`
	IsProcessing bool      `json:"is_processing,omitempty"`
	NoCache      bool      `json:"nocache,omitempty"`
	PlayOrder    uint      `json:"play_order,omitempty"`
}

type PlayList struct {
	Assets []Asset
}

func (p *PlayList) IsEmpty() bool {
	return len(p.Assets) == 0
}
