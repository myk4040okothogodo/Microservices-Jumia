package server

import (
    "context"
    "github.com/hashicorp/go-hclog"
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    //"github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/data"
)



type ServiceA struct {
    log hclog.Logger
}


func NewServiceA(l hclog.Logger ) *ServiceA {
    return &ServiceA{l}
}

 func (s *ServiceA)GetCsvData(ctx context.Context, dr *protos.DataRequest) (*protos.DataResponse, error){
     s.log.Info("Handle GetCsvData","token", dr.GetToken())
     
     CsvData := LoadCsvData()
     return &protos.DataResponse{PurchaseDetails : CsvData},nil     

 }
