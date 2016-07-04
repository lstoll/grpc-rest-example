# gRPC / REST Gateway sample app

This is a demo app for building a service in Go, then exporting it with a REST and gRPC API.

# Quickstart

`make` will start the server. You can then either:
* navigate to http://localhost:5000/swagger-ui for the REST docs & interacter
* `cd ruby && bundle exec bin/rest_client.rb localhost:5000` to run the ruby client against the REST API
* `cd ruby && bundle exec bin/grpc_client.rb <grpc host>` to run the ruby client against the gRPC API