package data4romServiceA

import (
    "log"
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
)



type Csvdata struct {
    l *log.Logger
    sc protos.ServiceAClient

}


func NewCsvdata(l *log.Logger, sc protos.ServiceAClient) *Csvdata {
    return &Csvdata{l, sc}
}


type GenericError struct{
    Message string `json:"message"`
}

type ValidationError struct {
    Messages []string `json: "messages"`
}
