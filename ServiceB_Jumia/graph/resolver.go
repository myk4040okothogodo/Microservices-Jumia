package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
    apiHolder "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/api_holder"

)
type Resolver struct{
    apiHolder apiHolder.Service
}


func NewResolver(ah apiHolder.Service) *Resolver {
    return &Resolver{apiHolder: ah}
}
