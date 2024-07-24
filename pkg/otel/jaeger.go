package otel

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

type JaegerConfig struct {
	Server      string `mapstructure:"server"`
	ServiceName string `mapstructure:"serviceName"`
	TracerName  string `mapstructure:"tracerName"`
}

func TracerJaeger(ctx context.Context, cfg *JaegerConfig, log *logrus.Logger) (trace.Tracer, error) {
	// Create the exporter
	client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint(cfg.Server), otlptracehttp.WithInsecure(), otlptracehttp.WithCompression(otlptracehttp.NoCompression))
	traceExporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, err
	}

	env := os.Getenv("APP_ENV")

	if env != "production" {
		env = "development"
	}

	// Create the resource to be traced
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(traceExporter),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
			attribute.String("environment", env),
		)),
	)

	go func() {
		select {
		case <-ctx.Done():
			err = tp.Shutdown(ctx)
			log.Info("open-telemetry exited properly")
			if err != nil {
				return
			}
		}
	}()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	t := tp.Tracer(cfg.TracerName)

	return t, nil
}
