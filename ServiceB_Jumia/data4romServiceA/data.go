package data4romServiceA


import (
    "context"
    protos "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
    "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/orders"
    database "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/pkg/db/migrations/mysql" 

  )



func (c *Csvdata)saveData( ) {
    //GetCsv data
    dr := &protos.DataRequest{
      Token: "762346288fdgddghbddg",
    }
    resp,err := c.sc.GetCsvData(context.Background(), dr)
    if err != nil {
        c.l.Println("[Error] error getting csv data ", err)
        return
    }

    stmt, err := database.Db.Prepare("INSERT INTO Orders(ParcelWeight, Country, Email, Phone) VALUES(?, ?, ?, ?)")
    if err != nil {
        c.l.Fatal(err)
    }

    for _, csvorder := range resp.PurchaseDetails {
        var order orders.Order
        order.Parcelweight = csvorder.ParcelWeight
        order.Country      = matchCodeToCountry(csvorder.CountryCode)
        order.Email        = csvorder.Email
        order.Phone        = csvorder.PhoneNumber
        
        res, err := stmt.Exec(order.Parcelweight, order.Country, order.Email, order.Phone)
        if err != nil {
            c.l.Fatal(err)
        }
        c.l.Println("Inserted ", csvorder.Email, res)
    }

}
