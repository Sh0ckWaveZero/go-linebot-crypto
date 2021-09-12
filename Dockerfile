# get a base image from the docker hub
FROM golang:1.17-buster as builder

LABEL "xyz.yadom.vendor"="Sh0ckWaveZero"
LABEL version="1.0"
LABEL description="Golang Linebot"

# Create and change to the app directory
WORKDIR /app

# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image
COPY . ./

# Build the binary.
RUN go build -v -o server
EXPOSE 8080

# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
  ca-certificates && \
  rm -rf /var/lib/apt/lists/*

# SET Timezone (Asia/Bangkok GTM+07:00)
ENV TZ Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# Run the web service on container startup.
CMD ["/app/server"]