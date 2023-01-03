package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_service_user_login/user_login/user_login"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := user_login.NewUserLoginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserLogin(ctx, &user_login.Request{UserName: "reation", Password: "19861029"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("ID: %d", r.Id)
}
