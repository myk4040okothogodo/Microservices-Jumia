package orders

import (
    database "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/pkg/db/migrations/mysql"
    "log"
)


type  Order struct {
    ID           int
    Parcelweight float64
    Country      string
    Email        string
    Phone        string
    
} 


func (order Order) Save() int64 {
    stmt, err := database.Db.Prepare("INSERT INTO Orders(Parcelweight, Country, Email, Phone) VALUES(?, ?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }

    res, err := stmt.Exec(order.Parcelweight, order.Country, order.Email, order.Phone)
    if err != nil {
        log.Fatal(err)
    }

    id, err := res.LastInsertId()
    if err != nil {
        log.Fatal("Error:", err.Error())
    }

    log.Print("Row inserted")
    return id
}

func GetAll() []Order{
  stmt,err := database.Db.Prepare("Select id, parcelweight, country, email, phone from Orders")
  if err != nil {
      log.Fatal(err)
  }

  defer stmt.Close()
  rows, err := stmt.Query()
  if err != nil {
      log.Fatal(err)
  }

  defer rows.Close()
  var orders []Order
  for rows.Next() {
      var order Order
      err := rows.Scan(&order.ID, &order.Parcelweight, &order.Country, &order.Email, &order.Phone)
      if err != nil {
         log.Fatal(err)
      }
      orders = append(orders, order)
  }
  if err = rows.Err(); err != nil {
      log.Fatal(err)
  }
  return orders
}
