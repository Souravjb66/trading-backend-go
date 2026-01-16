package config
import (
	// "database/sql"
	"fmt"
	"net/http"
	"trading/connect"
    // "github.com/gofiber/fiber/v2"
	"github.com/go-chi/chi/v5"
	sqlcdb "trading/db"
	// "github.com/gorilla/mux"
    // "github.com/gorilla/websocket"


	
    
	// "firebase.google.com/go/v4/messaging"
)

var TradeServer Server


type Server struct {
	// FireBase *firebase.FireBaseInstance
	// Mongodb  *mongodb.MongoDbStorage
	// MySql    *connect.MySqlDbInstance
	PostgreSQL *connect.PostgreSQLDbInstance
	HttpConnection *http.Server
	// RedisClient *redisConnector.Redis
	// Route   *http.ServeMux
	// WorkerPool  chan bool
	Route     *chi.Mux
	Redis      *connect.RedisStruct
	WebSocketRoute *chi.Mux
	
}

func GetServerInstance() *Server {
	return &TradeServer
}
func CreateServer(mysqlConfig connect.PostgreSQLRequired) {
	connect.RedisConfigSetUp()
	
	// mongodbClient := mongodb.Connect(mongodbConfig)
	TradeServer = Server{
		// Mongodb:  mongodb.NewMongoStorage(mongodbClient),
		// FireBase: firebase.GetFireBaseInstance(),
		// MySql:    connect.ConnectMysql(mysqlConfig),
		PostgreSQL: connect.ConnectMysql(mysqlConfig),
		// RedisClient: redisConnector.RedisConnect(redisConfig),
		Redis: connect.GetRedisInstance(),
	}

}
func StartServer(serverPort string) {
	// port := fmt.Sprintf(":%s", serverPort)
	// server:=http.Server{
	// 	Addr: fmt.Sprintf(":%s", serverPort),
	// 	Handler: TradeServer.Route,
	// }
	if err:=http.ListenAndServe(":8081", TradeServer.Route);err!=nil{
		fmt.Printf("failed to start the rest api server : ERROR : %s \n", err)
		panic(err)

	}else{
		fmt.Println("server start at port : 80")
	}
	
		
	
	
	

}

func OpenMysqlConnectionQuery()*sqlcdb.Queries{
	db:=sqlcdb.New(TradeServer.PostgreSQL.PostgreSQL)
	return db

}


type OrderBook struct{
	Id int 
	// SellerId int
	Asset string
	Price string
	Quantity string
}

