# gRPC / REST Gateway sample app

This is a demo app for building a service in Go, then exporting it with a REST and gRPC API.

# Quickstart

`make` will start the server. You can then either:
* navigate to http://localhost:5000/swagger-ui for the REST docs & interacter
* `cd ruby && bundle exec bin/rest_client.rb localhost:5000` to run the ruby client against the REST API
* `cd ruby && bundle exec bin/grpc_client.rb <grpc host>` to run the ruby client against the gRPC API

# How the app is built

### Service definition

[releases.proto](releases.proto)

This contains the main definition of the types, and APIs in use.

### Server

[go/server/main.go](go/server/main.go)

Actually starts the servers

### Service implementation

[go/server/releaser/service.go](go/server/releaser/service.go)

Implementation of the service

### Ruby gRPC client

[ruby/bin/grpc_client.rb](ruby/bin/grpc_client.rb)

Consumes the generated gRPC client in [ruby/lib](ruby/lib)

### Ruby REST client

[ruby/bin/rest_client.rb](ruby/bin/rest_client.rb)

Consumes the generated REST/Swagger client gem in [ruby/releases_client](ruby/releases_client)