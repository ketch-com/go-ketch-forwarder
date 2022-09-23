FROM alpine
ARG TARGETOS TARGETARCH
RUN apk update && apk --no-cache upgrade && mkdir /docker-entrypoint-initdb.d /app
COPY docker/ketch-event-forwarder/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
COPY .build/$TARGETOS-$TARGETARCH/ketch-event-forwarder /app/ketch-event-forwarder
WORKDIR /app
ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
