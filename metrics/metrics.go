package metrics

import "time"

type Cpu struct {
	UUID        string    `json:"uuid"`
	Ip          string    `json:"ip"`
	Model       string    `json:"model"`
	Utilization float64   `json:"utilization"`
	Time        time.Time `json:"time"`
}

type System struct {
	UUID         string    `json:"uuid"`
	Ip           string    `json:"ip"`
	Hostname     string    `json:"hostname"`
	Os           string    `json:"os"`
	Architecture string    `json:"architecture"`
	Platform     string    `json:"platform"`
	Version      string    `json:"version"`
	Time         time.Time `json:"time"`
}

type Tcp struct {
	UUID             string    `json:"uuid"`
	Ip               string    `json:"ip"`
	QueueSize        int       `json:"tcp_queue_size"`
	SegmentsReceived int       `json:"segments_received"`
	SegmentsSent     int       `json:"segments_sent"`
	Time             time.Time `json:"time"`
}

type Memory struct {
	UUID        string    `json:"uuid"`
	Ip          string    `json:"ip"`
	Total       float64   `json:"total"`
	Utilization float64   `json:"utilization,omitempty"`
	Time        time.Time `json:"time"`
}
