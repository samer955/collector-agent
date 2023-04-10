package repository

import (
	"database/sql"
	"github.com/samer955/collector-agent/metrics"
	"log"
)

type MetricRepositoryImpl struct {
	db *sql.DB
}

func NewMetricRepository(db *sql.DB) *MetricRepositoryImpl {
	return &MetricRepositoryImpl{db: db}
}

func (r *MetricRepositoryImpl) StoreCpu(cpu metrics.Cpu) error {

	query := "INSERT INTO `cpu`(`uuid`,`ip`,`model`,`utilization`,`time`) VALUES(?,?,?,?,?)"

	_, err := r.db.Exec(query,
		cpu.UUID,
		cpu.Ip,
		cpu.Model,
		cpu.Utilization,
		cpu.Time)

	if err != nil {
		return err
	}
	log.Println("STORED CPU METRIC: ", cpu)
	return nil

}

func (r *MetricRepositoryImpl) StoreSystem(sys metrics.System) error {

	log.Println("STORED SYSTEM METRIC: ", sys)
	return nil

}

func (r *MetricRepositoryImpl) StoreTcp(tcp metrics.Tcp) error {

	log.Println("STORED TCP METRIC: ", tcp)
	return nil

}

func (r *MetricRepositoryImpl) StoreMemory(mem metrics.Memory) error {

	log.Println("STORED MEMORY MEMORY")
	return nil

}
