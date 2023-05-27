# Protobufs Central Repository template

## Features of the generated library

Define the services and messages in ./protos/

Run `make` to run a dockerized build producing:

1. Go protobuf message definitions
2. Go gRPC clients and server stubs 
3. Go grpc-gateway reverse-proxy
4. OpenAPI/Swagger documents. Includes a minimal go package that embeds the *.json documents so it's easy to serve via HTTP.

## How to use the template:

1. Install cookiecutter (`pip3 install cookiecutter` will do it if you already have pip installed)
2. Instantiate the template
	
	cookiecutter .

3. Answer the prompts to initialize a repo
4. Copy the results into your own git repo for your project.

