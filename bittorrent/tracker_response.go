package bittorrent

//TrackerResponse A struct the params of a tracker response
type TrackerResponse struct {
	failureReason string
	interval      int
	minInterval   int
	trackerID     string
	complete      int
	incomplete    int
	peers         string
}
