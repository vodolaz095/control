package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

func (rpc *RPC) getClientInfo(ctx context.Context) (name string, address net.Addr, err error) {
	info, ok := peer.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("unauthorized")
		return
	}
	tlsInfo := info.AuthInfo.(credentials.TLSInfo)
	if len(tlsInfo.State.PeerCertificates) == 0 {
		err = fmt.Errorf("unauthorized")
		return
	}
	name = tlsInfo.State.PeerCertificates[0].Subject.CommonName
	address = info.Addr
	return
}
