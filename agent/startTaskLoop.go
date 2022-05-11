package agent

import (
	"context"
	"io"
	"log"
	"time"

	"gitflic.ru/vodolaz095/control/pb"
)

func StartTaskLoop(tag string, client pb.ControlClient, logger *log.Logger) {
	reportSink, err := client.ReportTaskUpdate(context.Background())
	if err != nil {
		logger.Fatalf("%s : while opening sink for task reports")
	}
	tasks, err := client.SubscribeToOrders(context.Background(), &pb.StringRequest{Data: tag})
	if err != nil {
		logger.Fatalf("%s : while opening sink for task reports")
	}
	for {
		var errL error
		order, errL := tasks.Recv()
		if errL == io.EOF {
			break
		}
		if err != nil {
			logger.Printf("%s : while receiving message", errL)
			continue
		}
		if order != nil {
			logger.Printf("Task received %s", order.TaskName)
			err = reportSink.Send(&pb.TaskUpdate{
				Timestamp: time.Now().Unix(),
				Nanos:     int32(time.Now().Nanosecond()),
				TaskID:    order.TaskName,
				Data:      "Accepted",
				Finished:  false,
				ExitCode:  0,
			})
			if err != nil {
				logger.Printf("%s : while sending message", errL)
			}
		} else {
			time.Sleep(20 * time.Millisecond)
		}
	}
}
