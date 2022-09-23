# Environment

The environment variables available for this repository are documented here, by the container name.

## ketch-event-forwarder

| Variable                                       | Default   | Description                               |
|------------------------------------------------|-----------|-------------------------------------------|
| `KETCH_EVENT_FORWARDER_SERVER_BIND`            | `0.0.0.0` | IP address to bind to                     |
| `KETCH_EVENT_FORWARDER_SERVER_LISTEN`          | `5000`    | Port to listen to                         |
| `KETCH_EVENT_FORWARDER_SERVER_TLS_CERT_FILE`   | N/A       | Location of the TLS certificate file      |
| `KETCH_EVENT_FORWARDER_SERVER_TLS_KEY_FILE`    | N/A       | Location of the TLS private key file      |
| `KETCH_EVENT_FORWARDER_SERVER_TLS_ROOTCA_FILE` | N/A       | Location of the TLS root certificate file |
| `KETCH_EVENT_FORWARDER_USER_NAME`              | N/A       | Username for basic authentication         |
| `KETCH_EVENT_FORWARDER_USER_PASSWORD`          | N/A       | Password for basic authentication         |
