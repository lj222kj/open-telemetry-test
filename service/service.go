package service

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
)

type service struct {
	tel *trace.TracerProvider
}

func (s *service) Hello(ctx context.Context) string {
	ctx, span := s.tel.Tracer("service").Start(ctx, "Hello")
	defer span.End()
	return "hello"
}

func New(tel *trace.TracerProvider) Service {

	return &service{tel: tel}
}
