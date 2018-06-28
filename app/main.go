package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/tnakade/tno_exercise/app/servers"
	"github.com/tnakade/tno_exercise/app/proto/services"
)

func main() {
	listenPort, err := net.Listen("tcp", listenAddressPort()) // nolint: gas
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer()
	service := servers.Wallet{}

	services.RegisterWalletServer(srv, &service)

	if err := srv.Serve(listenPort); err != nil {
		log.Fatalln(err)
	}
}

func listenAddressPort() string {
	return ":1080"
}
