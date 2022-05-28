package main

import (
    "net"
    "os"
    "github.com/hashicorp/go-hclog"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/server"
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"

  )


func main() {
    log := hclog.Default()

    gs := grpc.NewServer( )
    ss := server.NewServiceA(log)

    protos.RegisterServiceAServer(gs, ss)

    reflection.Register(gs)

    l, err := net.Listen("tcp", ":9092")
    if err != nil {
      log.Error("Unable to Listen", "error", err)
      os.Exit(1)
    }
    gs.Serve(l)
}
