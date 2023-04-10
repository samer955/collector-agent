package service

import "github.com/samer955/collector-agent/metrics"

type MetricRepository interface {
	StoreCpu(cpu metrics.Cpu) error
	StoreSystem(sys metrics.System) error
	StoreTcp(tcp metrics.Tcp) error
	StoreMemory(mem metrics.Memory) error
}
