package server

import (
	"io"

	"gitflic.ru/vodolaz095/control/pb"
)

// Good read: https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc

// ReportTelemetry gathers telemetry reports
func (rpc *RPC) ReportTelemetry(server pb.Control_ReportTelemetryServer) error {
	for {
		telemetry, err := server.Recv()
		if err == io.EOF {
			err = server.SendAndClose(&pb.StringResponse{Data: "Farewell"})
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		rpc.Logger.Printf("Telemetry received %v %v %v", telemetry.Load1, telemetry.Load2, telemetry.Load3)
	}
}
