package metrics

import (
	"fmt"
	"time"
)

type Cpu struct {
	UUID        string    `json:"uuid"`
	Hostname    string    `json:"hostname"`
	Model       string    `json:"model"`
	Utilization float64   `json:"utilization"`
	Time        time.Time `json:"time"`
}

func (c Cpu) String() string {
	return fmt.Sprintf("Hostname: %s, Model: %s, Utilization: %v%%, Time: %s",
		c.Hostname, c.Model, c.Utilization, c.Time)
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

func (s System) String() string {
	return fmt.Sprintf("Hostname: %s, Ip: %s, Os: %s, Arch: %s, Platform: %s, Version: %s, Online-Users: %v, Latency: %v ms, Time: %s",
		s.Hostname, s.Ip, s.Os, s.Architecture, s.Platform, s.Version, s.OnlineUsers, s.Latency, s.Time)
}

type Tcp struct {
	UUID             string    `json:"uuid"`
	Hostname         string    `json:"hostname"`
	QueueSize        int       `json:"queue_size"`
	SegmentsReceived int       `json:"segments_received"`
	SegmentsSent     int       `json:"segments_sent"`
	Time             time.Time `json:"time"`
}

func (t Tcp) String() string {
	return fmt.Sprintf("Hostname: %s, QueueSize: %v, Segments-Sent: %v, Segments-Received: %v, Time: %s",
		t.Hostname, t.QueueSize, t.SegmentsSent, t.SegmentsReceived, t.Time)
}

type Memory struct {
	UUID        string    `json:"uuid"`
	Hostname    string    `json:"hostname"`
	Total       float64   `json:"total"`
	Utilization float64   `json:"utilization,omitempty"`
	Time        time.Time `json:"time"`
}

func (m Memory) String() string {
	return fmt.Sprintf("Hostname: %s, Total: %v, Utilization %v%%, Time: %s",
		m.Hostname, m.Total, m.Utilization, m.Time)
}
