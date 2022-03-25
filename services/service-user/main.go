package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-protobuf/common/config"
	model "grpc-protobuf/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage model.UserList

func init() {
	localStorage := new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct{}

func (UsersServer) Register(ctx context.Context, params *model.User) (*empty.Empty, error) {
	fmt.Println(params.String())
	localStorage.List = append(localStorage.List, params)

	log.Println("Registering user", params.String())
	return new(empty.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return &localStorage, nil
}

func main() {
	srv := grpc.NewServer()

	var usersrv UsersServer
	model.RegisterUsersServer(srv, usersrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	listen, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(listen))
}
