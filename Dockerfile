FROM golang:1.13.3-alpine AS build_base

RUN apk add bash ca-certificates git

WORKDIR /vta

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"' -o vta

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app/

RUN apk add --no-cache tzdata

COPY --from=builder vta .

RUN adduser -D vta
USER vta

ENV HTTP_PORT=8080
ENV LOG_LEVEL=debug

EXPOSE 8080

ENTRYPOINT ["./vta"]
