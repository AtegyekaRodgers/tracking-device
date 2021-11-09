package main


import (
    "fmt"
    "sync"
    "log"
	"flag"
	"strings"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/afenet-demo/afenet-server/db"
)

type Pubsub struct {
  mu     sync.RWMutex
  subs   map[string][]chan string
  closed bool
}

func NewPubsub() *Pubsub {
  ps := &Pubsub{}
  ps.subs = make(map[string][]chan string)
  return ps
}

var pubsubBroker *Pubsub

func (ps *Pubsub) Close() {
  ps.mu.Lock()
  defer ps.mu.Unlock()

  if !ps.closed {
    ps.closed = true
    for _, subs := range ps.subs {
      for _, ch := range subs {
        close(ch)
      }
    }
  }
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/login", userLogin)
	router.POST("/user/register", registeringUser)
	router.PUT("/user/rights", assignUserRights)
	router.POST("/policy/new", uploadingNewPolicy)
	router.GET("/policy/read", readPolicy)
	router.POST("/policy/agree", agreeToPolicy)
	router.POST("/policy/disagree", disagreeToPolicy)
	router.POST("/policy/violation", reportingPolicyViolation)
	router.POST("/alert/status", updateAlertStatus)
	router.POST("/notification/status", updateNotificationStatus)
	router.GET("/actions/monitor", monitoringActions)
	return router
}

var database *gorm.DB

func main() {
    //++++| os.Args |+++++
    wsEndPoint := ":4400" 
    addr := flag.String("addr", wsEndPoint, "AFENET API service address") 
    flag.Parse()
    //++++++++++++++++++++
    database, err := db.Connect()
      if err != nil {
        json.NewEncoder(w).Encode(struct{errors string}{ errors: err.Error()})
      }
    defer database.Close()
    
    fmt.Println("Server listening on port: "+(strings.Split(wsEndPoint,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}








