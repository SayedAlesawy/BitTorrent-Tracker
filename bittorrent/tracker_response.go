package bittorrent

//TrackerResponse A struct the params of a tracker response
type TrackerResponse struct {
	FailureReason string `json:"infoHash,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	MinInterval   int    `json:"minInterval,omitempty"`
	TrackerID     string `json:"trackerID,omitempty"`
	Complete      int    `json:"complete,omitempty"`
	Incomplete    int    `json:"incomplete,omitempty"`
	Peers         string `json:"peers,omitempty"`
}
