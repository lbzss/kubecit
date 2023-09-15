FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates libprotobuf-dev protobuf-compiler
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

RUN make init && make all && make build

FROM debian:stable-slim



COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
