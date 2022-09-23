# Testing

## HTTP/GRPC requests

The `test/requests` folder and any subfolders required for organization store `.http` files that can be used to test
the services in this repository.

## Unit testing

All unit tests are in files sitting in the same package folder as the units under test.

To unit test this repository, run the dependencies as described in [RUNNING](RUNNING.md).

All unit tests should be created with the following build tag:

```go
//go:build unit && !integration && !smoke

```

Then you can run unit tests using Go Test:

```shell
go test -v --tags unit ./...
```

You can also set the `unit` build tag in your IDE.
