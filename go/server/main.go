package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/lstoll/grpc-rest-example/go/server/releaser"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT env var must be set")
		os.Exit(1)
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	listenAddr := host + ":" + port

	// Start the gRPC server locally on a random port. In theory grpc and rest could
	// co-exist on the same http server, but only if we use SSL. and if we use
	// SSL, we can't run on Heroku.
	grpcServer := grpc.NewServer()
	releaser.RegisterReleasesServer(grpcServer, &releaser.ReleaseService{})
	grpcLis, err := net.Listen("tcp", "localhost:")
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Printf("Starting gRPC server on %s\n", grpcLis.Addr().String())
		err := grpcServer.Serve(grpcLis)
		if err != nil {
			panic(err)
		}
	}()

	mux := http.NewServeMux()
	// Start the REST gateway service, connecting to the gRPC server
	gwmux := runtime.NewServeMux()
	ctx := context.Background()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	err = releaser.RegisterReleasesHandlerFromEndpoint(ctx, gwmux, grpcLis.Addr().String(), dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}
	mux.Handle("/", gwmux)

	// Serve up a UI
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "go/server/releaser/service.swagger.json")
	})
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui/dist"))))
	mux.HandleFunc("/swagger-ui", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, fmt.Sprintf("/swagger-ui/?url=http://%s/swagger.json#/Releases", listenAddr), 302)
	})

	httpLis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	fmt.Printf("Starting http server on %s\n", listenAddr)
	err = srv.Serve(httpLis)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
