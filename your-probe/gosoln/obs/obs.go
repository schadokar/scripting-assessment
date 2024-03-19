package main

import (
	"context"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	metric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func InitOtel() error {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("shop-backend"),
			semconv.ServiceVersion("v1.0"),
			semconv.ServiceInstanceID("http://localhost:4317")),
	)
	if err != nil {
		slog.Error("Merging resources", "err", err)
		return err
	}
	driver := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
		otlptracegrpc.WithCompressor("gzip"),
		otlptracegrpc.WithEndpoint("http://localhost:4317"),
	)
	exp, err := otlptrace.New(context.TODO(), driver)
	if err != nil {
		slog.Error("Making new otlptrace", "err", err)
		return err
	}
	pr := trace.NewBatchSpanProcessor(exp)
	tp := trace.NewTracerProvider(
		trace.WithSpanProcessor(pr),
		trace.WithResource(r),
	)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tp)

	mexp, err := otlpmetricgrpc.New(context.TODO(),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithDialOption(grpc.WithBlock()),
		otlpmetricgrpc.WithEndpoint("http://localhost:4317"))
	if err != nil {
		slog.Error("Creating otlpmetricgrpc", "err", err)
		return err
	}
	mp := metric.NewMeterProvider(metric.WithResource(r), metric.WithReader(metric.NewPeriodicReader(mexp, metric.WithInterval(time.Second*60))))
	otel.SetMeterProvider(mp)

	return nil
}

func get_credit() int64 {
	credit := int64(rand.Intn(100))

	return credit
}

func main() {
	InitOtel()

	ctx := context.Background()

	mtr := otel.Meter("Example meter name")

	successCounter, err := mtr.Int64Counter("success_counter")
	if err != nil {
		slog.Error("successCounter", "err", err)
	}

	failuerCounter, err := mtr.Int64Counter("failuer_counter")
	if err != nil {
		slog.Error("failuerCounter", "err", err)
	}

	mtr.Int64ObservableGauge("credit_gauge", otelmetric.WithInt64Callback(func(ctx context.Context, io otelmetric.Int64Observer) error {

		credit := get_credit()
		//You can use any function here

		io.Observe(credit)
		return nil
	}), otelmetric.WithUnit("Dollar"))

	successCounter.Add(ctx, 1)
	failuerCounter.Add(ctx, 1)
}
