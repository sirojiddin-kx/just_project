FROM golang:1.16 as builder

RUN mkdir -p example_api_gateway
WORKDIR  example_api_gateway
COPY . ./


RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/example_api_gateway /

FROM alpine
COPY --from=builder  example_api_gateway .

ENTRYPOINT ["/example_api_gateway"]