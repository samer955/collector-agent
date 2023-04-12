package metrics

import "time"

type Cpu struct {
	UUID        string    `json:"uuid"`
	Hostname    string    `json:"hostname"`
	Model       string    `json:"model"`
	Utilization float64   `json:"utilization"`
	Time        time.Time `json:"time"`
}

// System Latency between sender and collector
type System struct {
	UUID         string    `json:"uuid"`
	Ip           string    `json:"ip"`
	Hostname     string    `json:"hostname"`
	Os           string    `json:"os"`
	Architecture string    `json:"architecture"`
	Platform     string    `json:"platform"`
	Version      string    `json:"version"`
	OnlineUsers  int       `json:"online_users"`
	Latency      int64     `json:"latency"`
	Time         time.Time `json:"time"`
}

type Tcp struct {
	UUID             string    `json:"uuid"`
	Hostname         string    `json:"hostname"`
	QueueSize        int       `json:"queue_size"`
	SegmentsReceived int       `json:"segments_received"`
	SegmentsSent     int       `json:"segments_sent"`
	Time             time.Time `json:"time"`
}

type Memory struct {
	UUID        string    `json:"uuid"`
	Hostname    string    `json:"hostname"`
	Total       float64   `json:"total"`
	Utilization float64   `json:"utilization,omitempty"`
	Time        time.Time `json:"time"`
}
