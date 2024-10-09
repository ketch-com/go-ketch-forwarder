FROM --platform=$BUILDPLATFORM  golang:1.22-alpine AS builder
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache ca-certificates git

ENV PATH=/usr/local/go/bin:$PATH
ENV GO111MODULE="on" GOPRIVATE="go.ketch.com" CGO_ENABLED="0"
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

WORKDIR /dockerdev

COPY go.* .
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg --mount=type=secret,id=netrc,target=/root/.netrc,required go mod download

COPY . .
RUN --mount=target=. --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg --mount=type=secret,id=netrc,target=/root/.netrc,required go build -trimpath -buildmode=default -buildvcs=false -ldflags="-s -w" -o /docker-entrypoint ./cmd/ketch-event-forwarder/main.go


FROM scratch as runtime
COPY --from=builder /docker-entrypoint /docker-entrypoint
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/docker-entrypoint"]
