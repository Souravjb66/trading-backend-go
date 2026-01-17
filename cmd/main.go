
package main

import (
	"fmt"
	"log"


	mysqlConnect "trading/connect"
	server "trading/config"
	
	"os"

	"path/filepath"

    "trading/routes"
	"github.com/joho/godotenv"
	"trading/services"
	// "net/http"
	// "trading/ws"
)

func main() {
	fmt.Println("Starting trading Service...")
	// logger.InitiateLogger()
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../env/local.env")
	fmt.Println(envPath)
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("%v ended", err)
		return
	}
	// mongoUrl := os.Getenv("MONGO_DB_URL")
	// mongoName := os.Getenv("MONGO_DB_NAME")
	mysqlPort := os.Getenv("POSTGRES_DB_PORT")
	mysqlDbUsername := os.Getenv("POSTGRES_DB_USERNAME")
	mysqlDbName := os.Getenv("POSTGRES_DB_NAME")
	mysqlHost := os.Getenv("POSTGRES_DB_HOST")
	mysqlDbPassword := os.Getenv("POSTGRES_DB_PASSWORD")

	appPort := os.Getenv("SERVER_PORT")

	server.CreateServer( mysqlConnect.PostgreSQLRequired{Url: mysqlHost, UserName: mysqlDbUsername, Password: mysqlDbPassword, Port: mysqlPort, DbName: mysqlDbName})
    
	routes.Routes()
    services.HeapInit()
	services.AddDbBuySellDataToHeap()
	services.Match()

	log.Println("server started at port:", appPort)
	server.StartServer(appPort)
	

}