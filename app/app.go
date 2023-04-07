package app

import (
	"go.opentelemetry.io/otel"
	"net/http"
	"open-telemetry-test/service"
	"open-telemetry-test/telemetry"
)

const name = "app"

func New() error {

	mux := &http.ServeMux{}
	tel, err := telemetry.New("open-telemetry-test", "v0.0.1", "dev")
	if err != nil {
		return err
	}

	svc := service.New(tel)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tel.Tracer(name).Start(r.Context(), "/hello")
		defer span.End()
		hello := svc.Hello(ctx)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hello))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	otel.SetTracerProvider(tel)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
