from flask import Flask
from opentelemetry import trace
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
# from opentelemetry.exporter.jaeger.thrift import JaegerExporter
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.flask import FlaskInstrumentor

# Initialize tracing provider with service name
trace.set_tracer_provider(
  TracerProvider(
    resource=Resource.create({SERVICE_NAME: "my-python-service"})
  )
)

# TODO: old use
# Configure Jaeger exporter (Jaeger must be running)
# jaeger_exporter = JaegerExporter(
#   agent_host_name="localhost",  # Jaeger agent host
#   agent_port=6831,              # UDP port for Jaeger agent
# )

# Add exporter to trace provider
# trace.get_tracer_provider().add_span_processor(
#   BatchSpanProcessor(jaeger_exporter)
# )
# -----------------------------------------


# TODO: new use
otlp_exporter = OTLPSpanExporter(
  endpoint="http://localhost:4318/v1/traces",  # Jaeger all-in-one default OTLP/HTTP port
)

trace.get_tracer_provider().add_span_processor(
  BatchSpanProcessor(otlp_exporter)
)
# -----------------------------------------




# Create Flask app
app = Flask(__name__)

# Auto-instrument Flask routes
FlaskInstrumentor().instrument_app(app)

@app.route("/")
def hello():
  tracer = trace.get_tracer(__name__)
  with tracer.start_as_current_span("hello-handler"):
    return "Hello, Jaeger!"

@app.route("/work")
def do_work():
  tracer = trace.get_tracer(__name__)
  with tracer.start_as_current_span("work-handler"):
    # Simulate some work
    import time
    time.sleep(0.3)
    return "Work done!"

if __name__ == "__main__":
  app.run(port=5000)
