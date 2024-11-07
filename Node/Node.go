package main

import (
	cc "Client"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ip string
var targetPort string

type server struct {
	cc.UnimplementedClientServer
	msg string
	key int32
}

func newServer() *server {
	s := &server{msg: "", key: 0}
	return s
}

func passAlong() {



	//Making client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	var serverAddress = fmt.Sprintf("%s:%s",ip,targetPort)
	conn, err := grpc.NewClient(serverAddress, opts...)

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	// close connection when function terminates
	defer conn.Close()

	// create client
	client := cc.NewClientClient(conn)

}

func main() {

	//Making server
	ip := "localhost"

	var port string
	flag.StringVar(&port, "p", "5050", "Sets the port of the node")

	var targetPort string
	flag.StringVar(&targetPort, "tp", "5051", "Sets the port of the target node")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Now listening to port: %d", port)
	}

	

	grpcServer := grpc.NewServer()
	cc.RegisterClientServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}