package server

import (
	"gitflic.ru/vodolaz095/control/pb"
	"log"
)

func (rpc *RPC) ReportTelemetry(server pb.Control_ReportTelemetryServer) (err error) {
	telemetry, err := server.Recv()
	if err != nil {
		log.Printf("%s : while receiving telemetry", err)
		return
	}
	log.Printf("Telemetry received %v %v %v", telemetry.Load1, telemetry.Load2, telemetry.Load3)
	return
}
