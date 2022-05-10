package server

import (
	"context"

	"gitflic.ru/vodolaz095/control/pb"
)

func (rpc *RPC) GetLine(ctx context.Context, in *pb.StringRequest) (resp *pb.StringResponse, err error) {
	rv := in.Data
	name, addr, err := rpc.getClientInfo(ctx)
	if err != nil {
		return
	}
	rpc.Logger.Printf("Client %s from %s send: %v", name, addr.String(), rv)
	return &pb.StringResponse{Data: rv}, nil
}
