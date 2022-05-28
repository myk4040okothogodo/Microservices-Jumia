package server

import (
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
    "io"
    "os"
    "strconv"
    "strings"
)



func  LoadCsvData()  []*protos.CustomerPurchaseDetails  {
    var allCustomersPurchases []*protos.CustomerPurchaseDetails 
    pwd, _ := os.Getwd()
    csvFile, _ := os.Open(pwd + "/data/test_file.csv")
    println(pwd + "/data/test_file.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        }else if error != nil {
            log.Fatal(error)
        }
        parcelweight,err  := strconv.ParseFloat(strings.TrimSpace(line[3]), 32)
        countryval, err      :=  strconv.ParseInt(strings.TrimSpace(line[2][0:4]), 10, 64)
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }
        
        allCustomersPurchases = append(allCustomersPurchases, &protos.CustomerPurchaseDetails{

            Email: strings.TrimSpace(line[1]),
            ParcelWeight: parcelweight,
            PhoneNumber:  strings.TrimSpace(line[2]),
            CountryCode:  countryval,
        })
    }
    return allCustomersPurchases
}
