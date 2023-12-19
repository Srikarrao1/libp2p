package main

import (
	"bufio"
	"context"
	"github.com/libp2p/go-libp2p"
	"fmt"
	"log"
	
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2p "github.com/libp2p/go-libp2p"
	peer "github.com/libp2p/go-libp2p-peer"
	ipfsaddr "github.com/multiformats/go-multiaddr"
)

func main() {

	ctx := context.Background()
    host, err := libp2p.New(context.Background())
    if err != nil {
	   log.Fatal(err)
    }
}

host.SetStreamHandler("/chat/1.1.0", handleStream)

func handleStream(stream net.Stream) {
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)
}

dht, err := dht.New(ctx, host)

for _, addr := range bootstrapPeers {
	iaddr, _ := ipfsaddr.ParseString(addr)
	peerInfo, _ := peer.InfoFromP2pAddr(iaddr.Multiaddr())
	if err := host.Connect(ctx, *peerInfo); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection established with bootstrap node: ", *peerInfo)
	}

}


routingDiscovery := routingDiscovery.FindPeers(ctx, config.RendezvousString)

go func() {
	for peer := range peerChan {
		if peer.ID == host.ID() {
			continue
		} 
		fmt.Println("Found peer:", peer)

		fmt.Println("Connecting to:", peer)
		stream, err := host.NewStream(ctx, peer.ID, protocol.ID(config.ProtocolID))

		if err != nil {
			fmt.Println("Connection failed", err)
			continue

		} else {
			rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

            go writeData(rw)
			go readData(rw)
		}
		fmt.Println("Connected to:", peer)
	}
}()