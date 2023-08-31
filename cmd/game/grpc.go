package main

import (
	"context"
	character "game-service/gen/character"
	characterpb "game-service/gen/grpc/character/pb"
	charactersvr "game-service/gen/grpc/character/server"
	inventorypb "game-service/gen/grpc/inventory/pb"
	inventorysvr "game-service/gen/grpc/inventory/server"
	itempb "game-service/gen/grpc/item/pb"
	itemsvr "game-service/gen/grpc/item/server"
	inventory "game-service/gen/inventory"
	item "game-service/gen/item"
	"log"
	"net"
	"net/url"
	"sync"

	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, characterEndpoints *character.Endpoints, inventoryEndpoints *inventory.Endpoints, itemEndpoints *item.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		characterServer *charactersvr.Server
		inventoryServer *inventorysvr.Server
		itemServer      *itemsvr.Server
	)
	{
		characterServer = charactersvr.New(characterEndpoints, nil)
		inventoryServer = inventorysvr.New(inventoryEndpoints, nil)
		itemServer = itemsvr.New(itemEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	characterpb.RegisterCharacterServer(srv, characterServer)
	inventorypb.RegisterInventoryServer(srv, inventoryServer)
	itempb.RegisterItemServer(srv, itemServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
