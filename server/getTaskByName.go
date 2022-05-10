package server

import (
	"context"

	"gitflic.ru/vodolaz095/control/pb"
)

func (rpc *RPC) GetTaskByName(ctx context.Context, in *pb.StringRequest) (resp *pb.Task, err error) {
	rv := in.Data
	name, addr, err := rpc.getClientInfo(ctx)
	if err != nil {
		return
	}
	rpc.Logger.Printf("Client %s from %s asks for task %s", name, addr.String(), rv)
	resp = new(pb.Task)
	resp.Code = pb.TaskStatusCode_TASK_STATUS_NOT_FOUND
	return
}
