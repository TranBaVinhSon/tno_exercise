package main

import (
	"log"
	"net"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/tnakade/tno_exercise/app/proto/services"
	"github.com/tnakade/tno_exercise/app/servers"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", listenAddressPort()) // nolint: gas
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer()

	client, err := getClient()
	if err != nil {
		log.Fatalln(err)
		return
	}

	service := servers.NewWallet(client)

	services.RegisterWalletServer(srv, service)

	if err := srv.Serve(listenPort); err != nil {
		log.Fatalln(err)
	}
}

func listenAddressPort() string {
	return ":1080"
}

func getClient() (*rpcclient.Client, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "35.187.215.246:18332",
		User:         "foo",
		Pass:         "qDDZdeQ5vw9XXFeVnXT4PZ--tGN2xNjjR4nrtyszZx0=",
	}, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
