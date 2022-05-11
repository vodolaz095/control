package agent

import (
	"context"
	"log"
	"time"

	"gitflic.ru/vodolaz095/control/pb"
)

const SamplingInterval = 5 * time.Second

// StartReportingTelemetry starts reporting telemetry
func StartReportingTelemetry(client pb.ControlClient, logger *log.Logger) {
	var report pb.Telemetry
	sink, err := client.ReportTelemetry(context.Background())
	if err != nil {
		logger.Fatalf("%s : while opening sink for reporting telemetry events")
	}
	ticker := time.NewTicker(SamplingInterval)
	defer ticker.Stop()
	for range ticker.C {
		report.Load1 = 1
		report.Load2 = 1
		report.Load3 = 1
		report.MemoryFree = 1
		report.MemoryUsed = 100
		err = sink.Send(&report)
		if err != nil {
			logger.Printf("%s : while sending telemetry report", err)
			continue
		}
	}
}
