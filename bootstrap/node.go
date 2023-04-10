package bootstrap

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"log"
)

type Node struct {
	Host         host.Host
	discoveryTag string
}

func InitializeNode(ctx context.Context, discoveryTag string) *Node {

	node := &Node{}
	node.discoveryTag = discoveryTag
	node.createLibp2pHost()
	go node.findPeers(ctx)
	return node

}

func (n *Node) createLibp2pHost() {

	libp2pHost, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		log.Println("Unable to create a Libp2p Host")
		panic(err)
	}
	n.Host = libp2pHost
	log.Printf("New node initialized with host-ID %s\n", n.Host.ID().ShortString())

}

func (n *Node) findPeers(ctx context.Context) {

	peerChan := initMDNS(n.Host, n.discoveryTag)
	for {
		peer := <-peerChan
		fmt.Println("Found peer:", peer.ID.ShortString(), ", connecting")

		if err := n.Host.Connect(ctx, peer); err != nil {
			fmt.Println("Connection failed:", err)
			continue
		}
	}
}
