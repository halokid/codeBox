import requests
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.resources import Resource
from opentelemetry.instrumentation.requests import RequestsInstrumentor

# 设置 OTLP HTTP (走 4318)
trace.set_tracer_provider(
    TracerProvider(
        resource=Resource.create({"service.name": "python-client"})
    )
)
exporter = OTLPSpanExporter(endpoint="http://localhost:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(exporter))

RequestsInstrumentor().instrument()

tracer = trace.get_tracer("python-client")

with tracer.start_as_current_span("call-go-service"):
    resp = requests.get("http://localhost:8080/hello")
    print(resp.json())



