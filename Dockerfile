FROM golang AS builder

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build cmd/server/main.go

FROM gcr.io/distroless/base

WORKDIR /root
COPY --from=builder /app/main .

ENTRYPOINT [ "/root/main" ]
