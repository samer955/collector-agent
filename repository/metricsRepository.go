package repository

import (
	"database/sql"
	"github.com/samer955/collector-agent/metrics"
)

type MetricRepositoryImpl struct {
	db *sql.DB
}

func NewMetricRepository(db *sql.DB) *MetricRepositoryImpl {
	return &MetricRepositoryImpl{db: db}
}

func (r *MetricRepositoryImpl) StoreCpu(cpu metrics.Cpu) error {

	query := "INSERT INTO `cpu`(`uuid`,`hostname`,`model`,`utilization`,`time`) VALUES(?,?,?,?,?)"

	_, err := r.db.Exec(query,
		cpu.UUID,
		cpu.Hostname,
		cpu.Model,
		cpu.Utilization,
		cpu.Time)

	if err != nil {
		return err
	}

	return nil

}

func (r *MetricRepositoryImpl) StoreSystem(sys metrics.System) error {

	query := "INSERT INTO `system`(`uuid`,`ip`,`hostname`,`os`, `architecture`, `platform`, `version`, `latency`,`online_users`,`time`) VALUES(?,?,?,?,?,?,?,?,?,?)"

	_, err := r.db.Exec(query,
		sys.UUID,
		sys.Ip,
		sys.Hostname,
		sys.Os,
		sys.Architecture,
		sys.Platform,
		sys.Version,
		sys.Latency,
		sys.OnlineUsers,
		sys.Time)

	if err != nil {
		return err
	}

	return nil

}

func (r *MetricRepositoryImpl) StoreTcp(tcp metrics.Tcp) error {

	query := "INSERT INTO `tcp`(`uuid`,`hostname`,`queue_size`,`segments_sent`, `segments_received`, `time`) VALUES(?,?,?,?,?,?)"

	_, err := r.db.Exec(query,
		tcp.UUID,
		tcp.Hostname,
		tcp.QueueSize,
		tcp.SegmentsSent,
		tcp.SegmentsReceived,
		tcp.Time)

	if err != nil {
		return err
	}

	return nil

}

func (r *MetricRepositoryImpl) StoreMemory(mem metrics.Memory) error {

	query := "INSERT INTO `memory`(`uuid`,`hostname`,`total`,`utilization`,`time`) VALUES(?,?,?,?,?)"

	_, err := r.db.Exec(query,
		mem.UUID,
		mem.Hostname,
		mem.Total,
		mem.Utilization,
		mem.Time)

	if err != nil {
		return err
	}

	return nil

}
