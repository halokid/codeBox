package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initTracer() func() {
	ctx := context.Background()

	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-gin-service"),
		)),
	)
	otel.SetTracerProvider(tp)

	// 🔑 设置全局 Propagator，用于解析 HTTP Header (traceparent)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() { _ = tp.Shutdown(ctx) }
}

func main() {
	shutdown := initTracer()
	defer shutdown()

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// 🔑 从 HTTP 请求中提取 traceparent
		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

		tr := otel.Tracer("go-gin-handler")
		_, span := tr.Start(ctx, "say-hello")
		defer span.End()

		// 获取所有 Header（map[string][]string 类型）
    allHeaders := c.Request.Header
    c.JSON(http.StatusOK, gin.H{
        "all_headers": allHeaders,
    })
		//c.JSON(http.StatusOK, gin.H{"message": "Hello from Go Gin"})


	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}



