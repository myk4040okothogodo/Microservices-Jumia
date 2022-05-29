package data4romServiceA


import (
    "log"
    "context"
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/orders"
    database "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/pkg/db/migrations/mysql" 

  )



func saveData(l *log.Logger, sc *protos.ServiceAClient) {
    //GetCsv data
    dr := protos.DataRequest{
      Token: "762346288fdgddghbddg",
    }
    resp,err := &sc.GetCsvData(context.Background(), dr)
    if err != nil {
        l.Println("[Error] error getting csv data ", err)
        return
    }

    stmt, err := database.Db.Prepare("INSERT INTO Orders(ParcelWeight, Country, Email, Phone) VALUES(?, ?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }

    for _, csvorder := range resp.PurchaseDetails {
        var order orders.Order
        order.Parcelweight = csvorder.ParcelWeight
        order.Country      = matchCodeToCountry(csvorder.Country)
        order.Email        = csvorder.Email
        order.Phone        = csvorder.Phone
        
        res, err := stmt.Exec(order.Parcelweight, order.Country, order.Email, order.Phone)
        if err != nil {
            log.Fatal(err)
        }
        log.Print("Inserted ", csvorder.Email, res)
    }

}
