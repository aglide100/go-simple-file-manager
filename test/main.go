package main

import (
	"context"
	"log"
	"time"

	pb_svc "github.com/aglide100/go-simple-file-manager/pb/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "127.0.0.1:10112"
)


func main() {
	opts := grpc.WithInsecure()

	rootCACert := "keys/ca.cert"
	creds, err := credentials.NewClientTLSFromFile(rootCACert, "")
	if err != nil {
		log.Fatalf("failt to load CA trust certificate: %v", err)
	}
	opts = grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial(address, opts)
	if err != nil {
		log.Fatalf("Can't connect grpc server %v", err)
	}

	defer conn.Close()
	client := pb_svc.NewDirClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// str := &dir.Dir{
	// 	Spec: {},
	// }

	res, err := client.GetDir(ctx, &pb_svc.GetDirReq{Path: "./"})
	if err != nil {
		log.Fatalf("Can't receive anything! %v", err)
	}

	log.Println("res %v", res)
}
