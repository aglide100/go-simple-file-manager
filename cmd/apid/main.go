package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func realMain() error {
	gRPCWebAddr := flag.String("grpc.addr", "0.0.0.0:10112", "grpc address")
	usingTls := flag.Bool("grpc.tls", true, "using http2")
	serverCrt := flag.String("cert.crt", "keys/server.crt", "crt file location")
	serverKey := flag.String("cert.key", "keys/server.key", "ket file location")

	gRPCWebAddrL, err := net.Listen("tcp", *gRPCWebAddr)
	if err != nil {
		return err
	}
	defer gRPCWebAddrL.Close()

	var wait sync.WaitGroup
	wait.Add(1)

	var opts []grpc.ServerOption
	tls := *usingTls
	if tls {
		fmt.Println("Using tls keys")

		creds, err := credentials.NewServerTLSFromFile(*serverCrt, *serverKey)
		if err != nil {
			log.Fatalf("fail to load creds: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)

	wg, ctx := errgroup.WithContext(context.Background())
	_ = ctx

	wg.Go(func() error {
		wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
			// for test, TODO fix here
			return true
		}))

		handler := http.Handler(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if wrappedServer.IsGrpcWebRequest(req) || wrappedServer.IsAcceptableGrpcCorsRequest(req) {
				wrappedServer.ServeHTTP(resp, req)
			}
		}))

		log.Printf("Starting grpc server... %s", *gRPCWebAddr)
		var err error
		if *usingTls { 
			err = http.ServeTLS(gRPCWebAddrL, handler, *serverCrt , *serverKey)
		} else {
			log.Println("starting without tls....")
			err = http.Serve(gRPCWebAddrL, handler)
		}

		if err != nil {
			log.Fatalln("When error at serving grpc web... ", err)
		} 

		return err
	})

	return wg.Wait()
}