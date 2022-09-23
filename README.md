# go-ketch-forwarder

This repository contain a Go reference implementation of the Ketch Forwarder.

## Getting

To get the source code, use the following:

```shell
git clone git@github.com:ketch-com/go-ketch-forwarder.git
cd go-ketch-forwarder
```

## Building

To build the code, use the following:

```shell
./scripts/build.sh
```

## Testing

There are several sample requests in `test/requests`. You can run those requests (in IntelliJ) to simulate valid and
invalid incoming requests.

## Running in Docker Compose

To run locally in Docker Compose, use the following:

```shell
./scripts/build.sh linux
docker compose up --build
```

## Distributing

To build the production docker container, use the following:

```shell
./scripts/build.sh linux
docker build -f docker/ketch-event-forwarder/Dockerfile --tag ketch-event-forwarder:latest .
```

Now, you can run the container:

```shell
docker run -d -p 5000:5000 -v $PWD/certs:/tls -e KETCH_USER_NAME=user1 -e KETCH_USER_PASSWORD=password1 ketch-event-forwarder:latest
```

You will now have a running event forwarder listening on port 5000.

## Configuring

To change the configuration of the docker container, there are several environment variables you can set:

| Variable              | Default           | Description                                                                                                              |
|-----------------------|-------------------|--------------------------------------------------------------------------------------------------------------------------|
| `KETCH_USER_NAME`     | None              | Username required for basic authentication. If not supplied, the server starts unauthenticated for development purposes. |
| `KETCH_USER_PASSWORD` | None              | Password required for basic authentication. If not supplied, the server starts unauthenticated for development purposes. |
| `KETCH_LISTEN`        | `5000`            | Port to listen on.                                                                                                       |
| `KETCH_TLS_CERT_FILE` | `/tls/server.crt` | Location of the TLS certificate file.                                                                                    |
| `KETCH_TLS_KEY_FILE`  | `/tls/server.key` | Location of the TLS private key file.                                                                                    |

## Extending

Given this is a reference implementation, the only "implementation" provided is logging the incoming request. We recommend
you update files in `pkg/handler` and provide implementations of the types of request you want to handle.
