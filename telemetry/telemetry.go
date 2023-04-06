package telemetry

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
	"os"
)

func New(serviceName, version, env string) (*trace.TracerProvider, error) {
	l := log.New(os.Stdout, "", 0)

	file, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer file.Close()

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(file),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
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
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	), nil
}
