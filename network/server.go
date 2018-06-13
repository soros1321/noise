package network

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/perlin-network/noise/actor"
	"github.com/perlin-network/noise/crypto"
	"github.com/perlin-network/noise/dht"
	"github.com/perlin-network/noise/log"
	"github.com/perlin-network/noise/peer"
	"github.com/perlin-network/noise/protobuf"
	"io"
)

type Server struct {
	network *Network
	routes  *dht.RoutingTable
	actors  []actor.Actor
}

func createServer(network *Network, actors ...actor.Actor) *Server {
	return &Server{
		network: network,
		routes:  dht.CreateRoutingTable(peer.CreateID(network.Host(), network.Keys.PublicKey)),
		actors:  actors,
	}
}

func (s Server) Stream(server protobuf.Noise_StreamServer) error {
	var id *peer.ID
	var client protobuf.Noise_StreamClient

	for {
		raw, err := server.Recv()

		if err == io.EOF || err != nil {
			if id != nil {
				s.routes.RemovePeer(*id)
				log.Info("Peer " + id.Address + " has disconnected.")
			}
			return nil
		}

		if raw.Message == nil || raw.Sender == nil || raw.Sender.PublicKey == nil || len(raw.Sender.Address) == 0 || raw.Signature == nil {
			log.Debug("Received an invalid message (either no message, no sender, or no signature) from a peer.")
			continue
		}

		if !crypto.Verify(raw.Sender.PublicKey, raw.Message.Value, raw.Signature) {
			continue
		}

		val := peer.ID(*raw.Sender)
		id = &val

		var ptr ptypes.DynamicAny
		if err := ptypes.UnmarshalAny(raw.Message, &ptr); err != nil {
			continue
		}

		msg := ptr.Message

		switch msg.(type) {
		case *protobuf.HandshakeRequest:
			// Update routing table w/ peer's ID.
			s.routes.Update(*id)

			if client == nil {
				// Dial and send handshake response to peer.
				client, err = s.network.Dial(raw.Sender.Address)
				if err != nil {
					continue
				}
				err = s.network.Tell(client, &protobuf.HandshakeResponse{})
				if err != nil {
					continue
				}

				log.Info("Peer " + raw.Sender.Address + " has connected to you.")
			}

			continue
		case *protobuf.HandshakeResponse:
			// Update routing table w/ peer's ID.
			s.routes.Update(*id)

			log.Info("Successfully bootstrapped w/ peer " + raw.Sender.Address + ".")

			continue
		case *protobuf.LookupNodeRequest:
			if client != nil {
				response := &protobuf.LookupNodeResponse{Peers: []*protobuf.ID{}}
				msg := msg.(*protobuf.LookupNodeRequest)

				// Update routing table w/ peer's ID.
				s.routes.Update(*id)

				// Respond back with closest peers to a provided target.
				for _, id := range s.routes.FindClosestPeers(peer.ID(*msg.Target), dht.BucketSize) {
					id := protobuf.ID(id)
					response.Peers = append(response.Peers, &id)
				}

				s.network.Tell(client, response)
			}
		}

		if client != nil {
			for _, actor := range s.actors {
				actor.Receive(client, *id, msg)
			}
		}
	}
	return nil
}
