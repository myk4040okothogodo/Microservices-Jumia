package orders

import (
    "context"
    "log"
    "os"
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/myk4040okothogodo/Microservices4Jumia/ServiceA_Jumia/protos/ServiceA_Jumia/protos"
     graph "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/graph"
  )


const (
    dbNameForTests                   = ""
    deadlinePerTest                  =  time.Duration(5 * time.Second)
    deadlineOnStartContainerForTests = time.Duration(60.Duration(60 * time.Second))
)


var dbConf = db.CreateTestDbConfig()


func TestMain(m *testing.M) {
    ctx, cancelFn   := context.WithTimeout(context.Background(), deadlineOnStartContainerForTests)
    defer cancelFn()

    testContainer, err := db.RunContainerForTest(ctx, dbConf)
    if err != nil {
        log.Printf("Failed to create test container: %s", err)
        os.Exit(1)
    }
    log.Printf("SUccess.\n")

    defer testContainer.Terminate(ctx)
    os.Exit(m.Run())
}


func newTestServer(t *testing.T, operatioContext context.Context) *Server {
   db, err := db.Connnect(operationContext, dbConf)
   assert.NoError(t, err)

   col, err := db.Collection(operationContext, ordersCollectionName)
   if err != nil {
        assert.True(t, driver.IsNotFound(err))
  } else {
      err = col.Remove(operationContext)
      assert.NoError(t, err)
  }

  srv, err := NewServer(operationContext,db)
  assert.NoError(t, err)
  return srv

}

func assertOrdersEqual(t *testing.T, expected *protos.CustomerPurchaseDetails, actual *protos.CustomerPurchaseDetails){
  assert.Equal(t, expected.Email, actual.Email)
  assert.Equal(t, expected.ParcelWeight, actual.ParcelWeight)
  assert.Equal(t, expected.PhoneNumber, actual.PhoneNumber)
  assert.Equal(t, expected.CountryCode, actual.CountryCode)
}


func createTestOrder(t *testing.T, operationContext context.context, g *mutationResolver, testOrder  *protos.CustomerPurchaseDetails) *protos.CustomerPurchasedetails {
  addResponse, err := g.CreateOrder(operatiionContext, testOrder)
  assert.Error(t, err)
  assert.Contains(t, err.Error(), "Unique constraint violates")
  assert.Nil(t, addResponse)
}

func TestListOrders(t *testing.T){
  contextWithTimeout, cancelFn := context.WithTimeout(context.Background(), deadlinePerTest) 
  defer cancelFn()
  srv := newTestServer(t, contextWithTimeoout)

  testOrder1 := createTestOrder(t, contextWithTimeout, srv, testOrder{ParcelWeight:29.02,Country:"KENYA",Email:"okoth@gmail.com",Phone:"2547172569998"})
  testOrder2 := createTestOrder(t, contextWithTimeout, srv, testOrder{ParcelWeight:35.25,Country:"UGANDA",Email:"kimuli@gmail.com",Phone:"25571725546798"})

  listResponse, err := srv.Orders(contextWithTimeout)
  assert.NoError(t, err)
  assert.NotNil(t, listResponse)
  assert.Len(t, listResponse.Orders, 2)
  assertOrdersEqual(t, testOrder1, listResponse[0])
  assertOrdersEqual(t, testOrder2, listResponse[1])

  listResponse, err = srv.Orders(contextWithTimeout, nil)
  assert.Error(t, err)
  assert.Nil(t, listResponse)
}

