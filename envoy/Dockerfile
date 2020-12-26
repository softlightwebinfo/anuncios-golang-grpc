FROM envoyproxy/envoy-alpine:v1.16-latest

ADD ./pem/crt /etc/
ADD ./pem/key /etc/
ADD envoy.yaml /etc/

CMD ["/usr/local/bin/envoy", "-c", "/etc/envoy.yaml", "-l", "debug"]