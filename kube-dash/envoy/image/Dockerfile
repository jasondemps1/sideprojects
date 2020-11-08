FROM envoyproxy/envoy:v1.12.2

COPY . .

RUN chmod +x ./entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh", "/etc/envoy/envoy.yaml" ]