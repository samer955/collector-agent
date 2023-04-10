package consumer

import (
	"context"
	"errors"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"log"
)

type PubSubService struct {
	Psub          *pubsub.PubSub
	Subscriptions map[string]*pubsub.Subscription
}

// NewPubSubService return a new PubSubService Service using the GossipSub Service
func NewPubSubService(ctx context.Context, host host.Host) *PubSubService {
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		log.Println("unable to create the producer service")
		return nil
	}
	subscriptions := make(map[string]*pubsub.Subscription)
	return &PubSubService{Psub: ps, Subscriptions: subscriptions}
}

// JoinTopic allow the Peers to join a Topic on Pubsub
func (p *PubSubService) JoinTopic(room string) (*pubsub.Topic, error) {

	topic, err := p.Psub.Join(room)
	if err != nil {
		log.Println("Error while joining the topic: ", room)
		return nil, err
	}
	log.Println("Joined topic:", room)

	return topic, nil

}

// Subscribe returns a new Subscription for the topic.
func (p *PubSubService) Subscribe(topic *pubsub.Topic) (*pubsub.Subscription, error) {

	subscription, err := topic.Subscribe()

	if err != nil {
		log.Println("Cannot subscribe to: ", topic.String())
		return nil, err
	}
	log.Println("Subscribed to topic: " + subscription.Topic())

	p.Subscriptions[topic.String()] = subscription

	return subscription, nil
}

func (p *PubSubService) GetSubscription(topic string) (*pubsub.Subscription, error) {

	subscr, ok := p.Subscriptions[topic]
	if ok {
		return subscr, nil
	}
	return nil, errors.New("subscription to topic" + " " + topic + " " + "not found")
}
