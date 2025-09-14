use opentelemetry::global;
use opentelemetry::sdk::trace as sdktrace;
use opentelemetry_jaeger::PipelineBuilder;
use reqwest::Client;
use tracing::{info, instrument};
use tracing_subscriber;

#[instrument]
async fn call_gin_service(client: &Client) -> Result<String, reqwest::Error> {
    let response = client
        .get("http://localhost:8080/your-endpoint")
        .send()
        .await?;

    let body = response.text().await?;
    Ok(body)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // 初始化 Jaeger 导出器
    let tracer = PipelineBuilder::default()
        .with_service_name("rust-client")
        .install_simple()?;

    global::set_tracer_provider(sdktrace::TracerProvider::builder().with_simple_exporter(tracer).build());

    // 设置 tracing_subscriber
    tracing_subscriber::fmt()
        .with_env_filter("info")
        .init();

    // 创建 reqwest 客户端
    let client = Client::new();

    // 调用 Gin 服务
    match call_gin_service(&client).await {
        Ok(response) => info!("Received response: {}", response),
        Err(e) => eprintln!("Error calling Gin service: {}", e),
    }

    global::shutdown_tracer_provider(); // flush

    Ok(())
}


