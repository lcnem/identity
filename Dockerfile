FROM golang:alpine AS build-env

# Set working directory for the build
WORKDIR /go/src/github.com/lcnem/identity

# Add source files
COPY . .

RUN go install ./cmd/identityd
RUN go install ./cmd/identitycli

# Final image
FROM ubuntu:latest

RUN set -x && \
  apt update && \
  apt upgrade -y && \
  apt install -y --no-install-recommends && \
  apt -y clean && \
  apt autoremove --purge -y

WORKDIR /root

COPY scripts/genesis.json genesis.json
COPY scripts/init.sh init.sh

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/identityd /usr/local/bin/identityd
COPY --from=build-env /go/bin/identitycli /usr/local/bin/identitycli

# Run identityd by default, omit entrypoint to ease using container with identitycli
CMD ["identityd"]
