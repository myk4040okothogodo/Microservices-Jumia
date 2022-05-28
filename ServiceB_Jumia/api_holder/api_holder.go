package api_holder

import (
     serviceA "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    "io"
    "log"
    "google.golang.org/grpc"
)


//type ServiceConfig struct {
//    csvData string
//}

type service  struct {
   io.Closer
   service_AClientConn  *grpc.ClientConn
   service_AClient      serviceA.ServiceAClient
}


type Service interface {
  service_A()  serviceA.ServiceAClient
}


func NewServiceHolder() (Service, error) {
    log.Printf("Connection to serviceA .......")
    serviceAConnection, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
    if err != nil {
        return nil,err
    }

    ah := &service {
        service_AClientConn: serviceAConnection,
        service_AClient:  serviceA.NewServiceAClient(serviceAConnection),     
    }
    return ah, nil
}

func (ah *service) service_A() serviceA.ServiceAClient{
    return ah.service_AClient
}


func (ah *service) Close() error {
    err := ah.service_AClientConn.Close()
    if err != nil {
        log.Printf("An error occured while closing connection on ServiceA: %s ", err)
    }
    return nil
}
