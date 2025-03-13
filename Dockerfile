FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY ./be-inventory ./be-inventory

WORKDIR /app/be-inventory
RUN go mod download

RUN mkdir -p /usr/bin && CGO_ENABLED=0 go build -o /usr/bin/be-inventory ./cmd

FROM alpine

# Install tzdata to manage timezone settings
RUN apk add --no-cache tzdata

# Set the timezone to Asia/Jakarta
ENV TZ=Asia/Jakarta
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /usr/bin/be-inventory /usr/bin/be-inventory

CMD ["be-inventory"]