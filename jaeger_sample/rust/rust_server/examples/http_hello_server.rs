#[macro_use]
extern crate trackable;

use bytecodec::bytes::Utf8Encoder;
use bytecodec::null::NullDecoder;
use cf_rustracing::sampler::AllSampler;
use cf_rustracing_jaeger::reporter::JaegerCompactReporter;
use cf_rustracing_jaeger::span::SpanContext;
use cf_rustracing_jaeger::Tracer;
use fibers_http_server::{HandleRequest, Reply, Req, Res, ServerBuilder, Status};
use futures::future::ok;
use httpcodec::{BodyDecoder, BodyEncoder};
use std::collections::HashMap;
use std::net::Ipv4Addr;

struct Hello {
    tracer: Tracer,
}
impl HandleRequest for Hello {
    const METHOD: &'static str = "GET";
    const PATH: &'static str = "/**";

    type ReqBody = ();
    type ResBody = String;
    type Decoder = BodyDecoder<NullDecoder>;
    type Encoder = BodyEncoder<Utf8Encoder>;
    type Reply = Reply<Self::ResBody>;

    fn handle_request(&self, req: Req<Self::ReqBody>) -> Self::Reply {
        let mut carrier = HashMap::new();
        let header = req.header();
        for field in header.fields() {
            carrier.insert(field.name(), field.value());
        }

        let context = track_try_unwrap!(SpanContext::extract_from_http_header(&carrier));
        let _span = self
            .tracer
            .span("Hello::handle_request")
            .child_of(&context)
            .start();
        let body = format!("Hello: {}\n", req.url().path());
        Box::new(ok(Res::new(Status::Ok, body)))
    }
}

#[tokio::main]
async fn main() -> trackable::result::MainResult {
    let (tracer, mut span_rx) = Tracer::new(AllSampler);
    let handler = Hello { tracer };

    tokio::spawn(async move {
        let reporter = track_try_unwrap!(
            JaegerCompactReporter::new(
                "http_hello_server",
                (Ipv4Addr::LOCALHOST, 6831).into(),
                (Ipv4Addr::LOCALHOST, 0).into()
            )
            .await
        );

        while let Some(span) = span_rx.recv().await {
            track_try_unwrap!(reporter.report(&[span]).await);
        }
    });

    let mut builder = ServerBuilder::new(track_any_err!("127.0.0.1:8081".parse())?);
    builder.add_handler(handler)?;
    let server = builder.finish(fibers_global::handle());
    fibers_global::execute(server)?;
    Ok(())
}


