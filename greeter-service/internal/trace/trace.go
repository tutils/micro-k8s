package trace

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func NewTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
            //SamplingServerURL: fmt.Sprintf("http://%s:%d/sampling", "jaegertracing-agent", jaeger.DefaultSamplingServerPort),
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
            LocalAgentHostPort: fmt.Sprintf("%s:%d", "jaegertracing-agent", jaeger.DefaultUDPSpanServerPort),
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
