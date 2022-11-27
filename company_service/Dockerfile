FROM golang:1.16 as builder

RUN mkdir -p company_service
WORKDIR company_service
COPY . ./


RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/company_service /

FROM alpine
COPY --from=builder company_service .

ENTRYPOINT ["/company_service"]
