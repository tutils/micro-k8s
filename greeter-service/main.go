package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
    "github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
    "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"greeter/handler"
    "greeter/internal/monitor"
    "greeter/internal/trace"

	greeter "greeter/proto/greeter"
)

func main() {
	// New Service
    serviceName := "greeter"

    tracer, closer := trace.NewTracer(serviceName)

    go monitor.StartMonitor()

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
        micro.WrapClient(
			opentracing.NewClientWrapper(tracer),
			prometheus.NewClientWrapper(),
		),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(tracer),
			prometheus.NewHandlerWrapper(),
		),
        micro.AfterStop(func() error {
			return closer.Close()
		}),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
