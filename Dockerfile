FROM golang AS builder

RUN apt-get update && apt-get install -y upx-ucl

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" cmd/server/main.go
RUN upx main

FROM gcr.io/distroless/base

WORKDIR /root
COPY --from=builder /app/main .

ENTRYPOINT [ "/root/main" ]
