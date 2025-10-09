use cf_rustracing::sampler::AllSampler;
use cf_rustracing_jaeger::Tracer;
use cf_rustracing_jaeger::Span;
use cf_rustracing_jaeger::reporter::JaegerCompactReporter;
use std::net::Ipv4Addr;
use std::collections::HashMap;

/// Modified main function with trace context propagation
#[tokio::main]
async fn main() {
  let (tracer, mut span_rx) = Tracer::new(AllSampler);

  {
    let span = tracer.span("sample_op").start();

    // Create HTTP client
    let http_client = reqwest::Client::new();

    // Create a request builder with trace context injection
    let mut request = http_client.get("http://localhost:8080/hello");

    // Inject trace context into request headers
    // This requires implementing a propagator (e.g., Jaeger or W3C format)
    inject_trace_context(&span, &mut request);

    // Send request with trace context
    let resp = request.send().await.unwrap().text().await;

    println!("{resp:#?}");
    println!("{}", "job done!");
  }

  let span = span_rx.recv().await.unwrap();
  assert_eq!(span.operation_name(), "sample_op");

  let reporter = JaegerCompactReporter::new(
    "sample_service",
    (Ipv4Addr::LOCALHOST, 6831).into(),
    (Ipv4Addr::LOCALHOST, 0).into(),
  )
    .await
    .unwrap();

  reporter.report(&[span]).await.unwrap();
}

/// Function to inject trace context into HTTP request headers
fn inject_trace_context(span: &Span, request: &mut reqwest::RequestBuilder) {
  // Example for Jaeger propagation format
  // You'll need to implement the actual propagation logic based on your tracer library
  let trace_id = span.context().trace_id().to_string();
  let span_id = span.context().span_id().to_string();

  // Add headers for trace context propagation
  // Format depends on your propagation standard (Jaeger, W3C, etc.)
  request.header("uber-trace-id", format!("{}:{}:0:1", trace_id, span_id));
}


