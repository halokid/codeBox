import requests
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.requests import RequestsInstrumentor

# 设置 Trace Provider
trace.set_tracer_provider(
  TracerProvider(resource=Resource.create({"service.name": "python-client"}))
)

# 配置 OTLP Exporter，把 trace 发到本地 Jaeger
otlp_exporter = OTLPSpanExporter(endpoint="http://localhost:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(otlp_exporter))

# 自动追踪 requests 库的调用
RequestsInstrumentor().instrument()

tracer = trace.get_tracer(__name__)

def call_springboot():
  with tracer.start_as_current_span("call-springboot"):
    # resp = requests.get("http://localhost:8080/hello")  # Spring Boot 的接口
    resp = requests.get("http://localhost:19577/ping")  # Spring Boot 的接口
    print("Response:", resp.text)

if __name__ == "__main__":
  call_springboot()


