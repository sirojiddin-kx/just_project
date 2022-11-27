FROM golang:1.16 as builder

RUN mkdir -p position_service
WORKDIR position_service
COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/position_service /

FROM alpine
COPY --from=builder position_service .

ENTRYPOINT ["/position_service"]
