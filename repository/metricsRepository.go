package repository

import (
	"github.com/samer955/collector-agent/metrics"
	"log"
)

type MetricRepositoryImpl struct {
}

func (r *MetricRepositoryImpl) StoreCpu(cpu metrics.Cpu) error {

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
