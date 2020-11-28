package libs

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func GrpcGetNetListener(port string) net.Listener {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}
	return listener
}

func GrpcNewServer() *grpc.Server {
	return grpc.NewServer()
}

func GrpcListenerServer(server *grpc.Server, con net.Listener) {
	println("Running...")
	err := server.Serve(con)
	if err != nil {
		panic(err)
	}
}
