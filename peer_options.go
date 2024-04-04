package p2p

import (
	"fmt"
	"net"
	"time"
)

type PeerOptions func(p *Peer) error

func WithDialer(dial func(network, address string) (net.Conn, error)) PeerOptions {
	return func(p *Peer) error {
		p.dial = dial
		return nil
	}
}

func WithBatchDelay(batchDelay time.Duration) PeerOptions {
	return func(p *Peer) error {
		p.batchDelay = batchDelay
		return nil
	}
}

func WithIncomingConnection(conn net.Conn) PeerOptions {
	return func(p *Peer) error {
		p.incomingConn = conn
		return nil
	}
}

func WithMaximumMessageSize(maximumMessageSize int64) PeerOptions {
	return func(p *Peer) error {
		p.maximumMessageSize = maximumMessageSize
		return nil
	}
}

func WithUserAgent(userAgentName string, userAgentVersion string) PeerOptions {
	return func(p *Peer) error {
		if userAgentName == "" || userAgentVersion == "" {
			return fmt.Errorf("both user agent name '%s' and version '%s' must not be empty strings", userAgentName, userAgentVersion)
		}

		p.userAgentName = &userAgentName
		p.userAgentVersion = &userAgentVersion
		return nil
	}
}

func WithRetryReadWriteMessageInterval(d time.Duration) PeerOptions {
	return func(p *Peer) error {
		p.retryReadWriteMessageInterval = d
		return nil
	}
}
