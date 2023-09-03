package jaeger

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

type Cfg struct {
	JaegerEnable bool
	ServiceName  string
	SamplerType  string
	SamplerParam float64
	JaegerHost   string
	JaegerPort   string
}

func NewTracer(jCfg Cfg) (opentracing.Tracer, io.Closer, error) {
	if !jCfg.JaegerEnable {
		return opentracing.NoopTracer{}, nil, nil
	}

	cfg := &config.Configuration{
		ServiceName: jCfg.ServiceName,

		Sampler: &config.SamplerConfig{
			Type:  jCfg.SamplerType,
			Param: jCfg.SamplerParam,
		},

		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jCfg.JaegerHost + ":" + jCfg.JaegerPort,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil
}
