import requests
from flask import Flask
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.resources import Resource
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor


#配置 Tracer
trace.set_tracer_provider(
  TracerProvider(resource=Resource.create({"service.name": "service-a"}))
)
otlp_exporter = OTLPSpanExporter(endpoint="http://localhost:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(otlp_exporter))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()  # 自动追踪 requests 库调用

@app.route("/call-b")
def call_b():
  tracer = trace.get_tracer(__name__)
  with tracer.start_as_current_span("call-b-handler"):
    resp = requests.get("http://localhost:5001/process")
    return f"Service A got response: {resp.text}"

if __name__ == "__main__":
  app.run(port=5000)



