# GRPC-demo
# build command : docker build . -t grpc-demo/server
# run command : docker run -it grpc-demo/server
FROM golang:latest

# Install grpc
RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin
WORKDIR /go/src/grpc-server
# Copy the project to be executed
RUN mkdir -p .
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY . .
RUN go mod download

EXPOSE 4040

ENTRYPOINT cd /go/src/grpc-server && go run main.go