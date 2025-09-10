from flask import Flask
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.resources import Resource
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.flask import FlaskInstrumentor


# 配置 Tracer
trace.set_tracer_provider(
  TracerProvider(resource=Resource.create({"service.name": "service-b"}))
)
otlp_exporter = OTLPSpanExporter(endpoint="http://localhost:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(otlp_exporter))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)

@app.route("/process")
def process():
  tracer = trace.get_tracer(__name__)
  with tracer.start_as_current_span("process-handler"):
    return "Service B processed request!"

if __name__ == "__main__":
  app.run(port=5001)

