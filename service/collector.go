package service

import (
	"container/list"
	"context"
	"github.com/google/uuid"
	"github.com/samer955/collector-agent/bootstrap"
	"github.com/samer955/collector-agent/config"
	"github.com/samer955/collector-agent/consumer"
	"github.com/samer955/collector-agent/metrics"
	"github.com/samer955/collector-agent/repository"
	"github.com/samer955/collector-agent/utils"
	"log"
	"sync"
	"time"
)

type Collector struct {
	Context       context.Context
	PubSubService *consumer.PubSubService
	Repository    MetricRepository
	Config        config.CollectorConfig
	Mutex         sync.Mutex
}

// MetricData is a help struct used to wrap the messaged received
// The Name represents the topic name.
// Payload is the message byte-content received.
// From is the sender.
type MetricData struct {
	Name    string
	Payload []byte
	From    string
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
		Mutex:         sync.Mutex{},
	}
}

// Start subscribes to the topics and listen on each topic for incoming messages.
func (c *Collector) Start() {

	c.subscribeToTopics()

	//new linked list to receive data
	metricDataList := list.New()

	for _, topic := range c.Config.Topics() {
		go c.ReadFromTopic(metricDataList, topic)
	}

	for {
		c.Mutex.Lock()
		switch metricDataList.Len() > 0 {
		case true:
			metricData := metricDataList.Remove(metricDataList.Front()).(MetricData)
			c.Mutex.Unlock()
			log.Printf(">>New Message received from %s. Metric Name: %s. Content: %s", metricData.From, metricData.Name, string(metricData.Payload))
			c.ConsumeMessages(metricData)
		default:
			c.Mutex.Unlock()
		}
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

// ReadFromTopic reads messages from a topic and write them in a linked list to be consumed
func (c *Collector) ReadFromTopic(msgList *list.List, topic string) {

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
			c.Mutex.Lock()
			msgList.PushBack(MetricData{Name: topic, Payload: msg.Data, From: msg.ReceivedFrom.ShortString()})
			c.Mutex.Unlock()
		}
	}

}

// ConsumeMessages function filters the messages received in the list by the topic name
// the HandlePanicError() avoids the program to stop if a panic is thrown by the database e.g. duplicate primary key, database not connected

func (c *Collector) ConsumeMessages(metricData MetricData) {

	defer utils.HandlePanicError()

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
	sys.Latency = utils.LatencyCalc(time.Now(), sys.Time)

	if err := c.Repository.StoreSystem(sys); err != nil {
		log.Println(err)
		return
	}
	log.Println("STORED SYSTEM METRIC:", sys.String())

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
	log.Println("STORED CPU METRIC:", cpu.String())

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
	log.Println("STORED TCP METRIC:", tcp.String())

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
	log.Println("STORED MEMORY METRIC:", mem.String())

}
