FROM golang:1.20 AS base
RUN apt update && apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
ENV PATH="${PATH}:/${GOPATH}/bin"
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.19.0/buf-linux-x86_64" -o "/usr/local/bin/buf" && chmod +x /usr/local/bin/buf
WORKDIR /protobufs
