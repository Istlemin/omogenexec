package main

import (
	"flag"
	"io/ioutil"
	"net"

	"github.com/google/logger"
	"google.golang.org/grpc"

	fileservice "github.com/jsannemo/omogenjudge/filehandler/service"
	toolservice "github.com/jsannemo/omogenjudge/problemtools/service"
	"github.com/jsannemo/omogenjudge/runner/language"
	runnerservice "github.com/jsannemo/omogenjudge/runner/service"
)

var (
	listenAddr = flag.String("localjudge_listen_addr", "127.0.0.1:61811", "The local judge server address to listen to")
)

func main() {
	flag.Parse()
	defer logger.Init("localjudge", true, false, ioutil.Discard).Close()
	if err := language.Init(); err != nil {
		logger.Fatalf("failed initializing languages: %v", err)
	}
	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err != nil {
		logger.Fatalf("failed to create server: %v", err)
	}
	if err := toolservice.Register(grpcServer); err != nil {
		logger.Fatalf("failed registering tool service :%v")
	}
	if err := runnerservice.Register(grpcServer); err != nil {
		logger.Fatalf("failed registering runner service :%v")
	}
	fileservice.Register(grpcServer)
	logger.Infof("serving on %v", *listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("could not listen: %v", err)
	}
}
