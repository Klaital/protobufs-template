# Protobufs Central Repository

Define the services and messages in ./protos/

Run `make` to run a dockerized build producing:

1. Go protobuf message definitions
2. Go gRPC clients and server stubs 
3. Go grpc-gateway reverse-proxy
4. OpenAPI/Swagger documents. Includes a minimal go package that embeds the *.json documents so it's easy to serve via HTTP.

## Building

The code generation is dockerized. Assuming you have `docker` and `docker-compose` available, just run

	make

to generate the various output files. I recommend setting this up as a precommit hook in your source control, so that the generated files are always in sync with the protobuf definitions.

