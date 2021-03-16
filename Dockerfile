FROM golang:1.16-alpine AS builder
WORKDIR /vta
RUN apk add ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o vta

FROM alpine:latest
WORKDIR /app
RUN adduser -D vta
RUN apk add --no-cache tzdata
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder vta .
USER vta
ENV HTTP_PORT=8080
ENV LOG_LEVEL=debug
EXPOSE 8080
ENTRYPOINT ["./vta"]
