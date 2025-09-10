
# run java service
`opentelemetry-javaagent.jar` in the same folder with `eapi-0.0.1-SNAPSHOT.jar`  

```bash
java \
  -javaagent:opentelemetry-javaagent.jar \
  -Dotel.service.name=spring-boot-service \
  -Dotel.traces.exporter=otlp \
  -Dotel.exporter.otlp.protocol=grpc \
  -Dotel.exporter.otlp.endpoint=http://localhost:4317 \
  -Dotel.metrics.exporter=none \
  -jar eapi-0.0.1-SNAPSHOT.jar

```