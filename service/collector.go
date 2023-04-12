package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/samer955/collector-agent/bootstrap"
	"github.com/samer955/collector-agent/config"
	"github.com/samer955/collector-agent/consumer"
	"github.com/samer955/collector-agent/metrics"
	"github.com/samer955/collector-agent/repository"
	"github.com/samer955/collector-agent/utils"
	"log"
)

type Collector struct {
	Context       context.Context
	PubSubService *consumer.PubSubService
	Repository    MetricRepository
	Config        config.CollectorConfig
}

// MetricData is a help struct where the name must match the topic name. Payload is the message received from other agent
type MetricData struct {
	Name    string
	Payload []byte
}

func NewCollector() *Collector {

	ctx := context.TODO()
	cfg := config.GetConfig()
	dbCfg := config.GetDbConfig()
	node := bootstrap.InitializeNode(ctx, cfg.DiscoveryTag())
	ps := consumer.NewPubSubService(ctx, node.Host)
	repo := repository.NewMetricRepository(dbCfg.Connection)

	return &Collector{
		PubSubService: ps,
		Repository:    repo,
		Context:       ctx,
		Config:        cfg,
	}
}

func (c *Collector) Start() {

	c.subscribeToTopics()

	//new channel to receive data
	metricDataChan := make(chan MetricData)

	for _, topic := range c.Config.Topics() {
		go c.ReadFromTopic(metricDataChan, topic)
	}

	for metricData := range metricDataChan {
		log.Println("New Message received: ", metricData.Name, string(metricData.Payload))
		c.ConsumeMessages(metricData)
	}

}

func (c *Collector) subscribeToTopics() {

	for _, top := range c.Config.Topics() {
		topic, err := c.PubSubService.JoinTopic(top)
		if err != nil {
			panic(err)
		}

		_, suberr := c.PubSubService.Subscribe(topic)
		if suberr != nil {
			panic(err)
		}

	}
}

// ReadFromTopic reads from a topic and write the messages in a channel in order to be consumed
func (c *Collector) ReadFromTopic(msgChan chan MetricData, topic string) {

	subscr, err := c.PubSubService.GetSubscription(topic)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := subscr.Next(c.Context)
		if err != nil {
			log.Println("Unable to read from topic:" + " " + subscr.Topic())
			continue
		} else {
			msgChan <- MetricData{Name: topic, Payload: msg.Data}
		}
	}

}

// ConsumeMessages function filter the messages received in the channel by the topic name
func (c *Collector) ConsumeMessages(metricData MetricData) {

	switch metricData.Name {

	case "CPU":
		c.handleCpuMetric(metricData)
	case "SYSTEM":
		c.handleSystemMetric(metricData)
	case "TCP":
		c.handleTcpMetric(metricData)
	case "MEMORY":
		c.handleMemoryMetric(metricData)

	default:
		log.Printf("Unknown metric name: %s\n", metricData.Name)
	}

}

func (c *Collector) handleSystemMetric(metricData MetricData) {

	var sys metrics.System
	if err := utils.FromBytesToStruct(metricData.Payload, &sys); err != nil {
		log.Println(err)
		return
	}
	sys.UUID = uuid.New().String()

	if err := c.Repository.StoreSystem(sys); err != nil {
		log.Println(err)
		return
	}
	log.Println("STORED SYSTEM METRIC: ", sys)

}

func (c *Collector) handleCpuMetric(metricData MetricData) {

	var cpu metrics.Cpu
	if err := utils.FromBytesToStruct(metricData.Payload, &cpu); err != nil {
		log.Println(err)
		return
	}
	cpu.UUID = uuid.New().String()

	if err := c.Repository.StoreCpu(cpu); err != nil {
		log.Println(err)
		return
	}
	log.Println("STORED CPU METRIC: ", cpu)

}

func (c *Collector) handleTcpMetric(metricData MetricData) {

	var tcp metrics.Tcp
	if err := utils.FromBytesToStruct(metricData.Payload, &tcp); err != nil {
		log.Println(err)
		return
	}
	tcp.UUID = uuid.New().String()

	if err := c.Repository.StoreTcp(tcp); err != nil {
		log.Println(err)
		return
	}
	log.Println("STORED TCP METRIC: ", tcp)

}

func (c *Collector) handleMemoryMetric(metricData MetricData) {

	var mem metrics.Memory
	if err := utils.FromBytesToStruct(metricData.Payload, &mem); err != nil {
		log.Println(err)
		return
	}
	mem.UUID = uuid.New().String()

	if err := c.Repository.StoreMemory(mem); err != nil {
		log.Println(err)
		return
	}
	log.Println("STORED MEMORY METRIC: ", mem)

}
