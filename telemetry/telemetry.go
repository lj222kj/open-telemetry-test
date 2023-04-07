package telemetry

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func New(serviceName, version, env string) (*trace.TracerProvider, error) {
	jgr, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(version),
			attribute.String("environment", env),
		),
	)

	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(jgr),
		trace.WithResource(resource),
	), nil
}
