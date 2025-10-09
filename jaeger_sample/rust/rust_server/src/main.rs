use cf_rustracing::sampler::AllSampler;
use cf_rustracing_jaeger::Tracer;
use cf_rustracing_jaeger::reporter::JaegerCompactReporter;
use reqwest::header::{HeaderMap, HeaderValue};
use std::net::Ipv4Addr;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let (tracer, mut span_rx) = Tracer::new(AllSampler);

    {
        let span = tracer.span("sample_op").start();

        if let Some(ctx) = span.context() {
            let state = ctx.state();

            // trace_id 已经实现了 Display，可以直接 format
            let trace_id = format!("{:032}", state.trace_id()); // 补齐成 32 字符
            let span_id  = format!("{:016x}", state.span_id());
            let flags    = format!("{:02x}", state.flags());

            let traceparent = format!("00-{}-{}-{}", trace_id, span_id, flags);
            println!("constructed traceparent = {}", traceparent);

            let mut headers = HeaderMap::new();
            headers.insert("traceparent", HeaderValue::from_str(&traceparent)?);

            let client = reqwest::Client::new();
            let resp = client
              .get("http://localhost:8080/hello")
              .headers(headers)
              .send()
              .await?
              .text()
              .await?;

            println!("resp = {}", resp);
        }

        println!("job done!");
    }

    let span = span_rx.recv().await.unwrap();

    let reporter = JaegerCompactReporter::new(
        "sample_service",
        (Ipv4Addr::LOCALHOST, 6831).into(),
        (Ipv4Addr::LOCALHOST, 0).into(),
    )
      .await
      .unwrap();

    reporter.report(&[span]).await.unwrap();
    Ok(())
}

