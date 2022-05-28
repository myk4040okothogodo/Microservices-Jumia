package data

import (
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
    "io"
    "os"
    "strconv"
)


func  LoadCsvData() []*protos.CustomerPurchaseDetails {
    var allCustomersPurchases []*protos.CustomerPurchaseDetails; 
    pwd, _ := os.Getwd()
    csvFile, _ := os.Open(pwd + "../data/test_file.csv")
    println(pwd + "../data/test_file.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    for {
        line, errorr := reader.Read()
        if error == io.EOF {
            break
        }else if error != nil {
            log.Fatal(error)
        }
        parcelweight,err  := strconv.ParseFloat(line[3])
        PhoneNumber   := line[2]
        country      :=  "KENYA"
        
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

        allCustomersPurchases = append(allCustomersPurchases, *protos.CustomerPurchaseDetails{

            Email: line[1],
            ParcelWeight: parcelweight,
            PhoneNumber:  line[2],
            Country:  country  
        })
    }
    return allCustomerPurchases
}
