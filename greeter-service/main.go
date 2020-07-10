package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	ot "github.com/opentracing/opentracing-go"
	"greeter/handler"
	"greeter/internal/monitor"
	"greeter/internal/trace"
	greeter "greeter/proto/greeter"
	"io"
)

func main() {
	// New Service
	serviceName := "greeter"
	prometheusReporterAddress := "localhost:6831"

	go monitor.StartMonitor()

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Flags(
			&cli.StringFlag{
				Name:    "prometheus_reporter_address",
				Usage:   "Set the prometheus reporter address",
				EnvVars: []string{"PROMETHEUS_REPORTER_ADDRESS"},
				Value:   prometheusReporterAddress,
			},
		),
	)

	var tracer ot.Tracer
	var closer io.Closer

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			if f := c.String("prometheus_reporter_address"); len(f) > 0 {
				prometheusReporterAddress = f
			}
			tracer, closer = trace.NewTracer(serviceName, prometheusReporterAddress)
			return nil
		}),
	)

	service.Init(
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

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
