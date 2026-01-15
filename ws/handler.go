package ws

import (
	
	"log"
	// "net/http"

	// "github.com/gorilla/mux"
	// "github.com/gorilla/websocket"
	"sync"
	// "trading/config"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// upgrader upgrades HTTP to WebSocket
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }
var PubSubSystem = NewPubSub()
// HandleConnection upgrades HTTP → WS and registers the client
// func HandleConnection(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("Upgrade error:", err)
// 		return
// 	}

// 	client := NewClient(conn)
// 	RegisterClient(client)
// 	sub := PubSubSystem.Subscribe("portfolio_updates")

// 	for msg := range sub {
// 		conn.WriteJSON(msg)
// 	}

// 	go client.ReadMessages()
// 	go client.WriteMessages()
// }
func RegisterClient(client *Client,id int) {
	var mu sync.Mutex

	// manager.mu.Lock()
	// defer manager.mu.Unlock()

	// manager.clients[client] = true
	mu.Lock()
	defer mu.Unlock()
    isConnectionAlliveMap[id]=true
	WebsocketConnections[id]=client
	Rooms[RoomName]=append(Rooms[RoomName],client.Conn)

}

func UnregisterClient(client *Client,id int) {
	// manager.mu.Lock()
	// defer manager.mu.Unlock()

	// if _, ok := manager.clients[client]; ok {
	// 	delete(manager.clients, client)
	// 	close(client.Send)
	// }
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	
	for index,item:=range Rooms[RoomName]{
		if item==WebsocketConnections[id].Conn{
			newConn:=append(Rooms[RoomName][:index],Rooms[RoomName][:index]...)
			Rooms[RoomName]=newConn

		}
		
	}
	delete(isConnectionAlliveMap, id)
	delete(WebsocketConnections, id)
	close(client.Send)
	
}

func ReadClientMessage(c *websocket.Conn){
	// app:=config.TradeServer.WebSocketRoute
	var(
			dataType int
			msg []byte
			err error
			

		)
	for{
		if dataType,msg,err=c.ReadMessage();err!=nil{
			log.Println("error in read ",err)
			break
		}
		log.Println(dataType)
		log.Println(string(msg))

	}
		

}
func WriteClientMessage(c *websocket.Conn){
	// app:=config.TradeServer.WebSocketRoute
	var(
			dataType int
			msg []byte
			

		)
		for{
			if err:=c.WriteMessage(dataType, msg);err!=nil{
			    log.Println("errro in sending msg ",err)
				break
			
		    }

		}



	

}
func UpgradeToWs(ctx *fiber.Ctx)error{
	// app:=config.TradeServer.WebSocketRoute
	websocket.New(func(c *websocket.Conn){
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""
		userId,err:=strconv.Atoi(c.Params("id"))   //need to convert it to int
		if err!=nil{
			log.Println(" error in parsing userid :",err)
			return
		}
		RegisterClient(&Client{
			Conn: c,
			Send: make(chan []byte),

		},userId)
		go WriteClientMessage(c)
		go ReadClientMessage(c)
		
		
	})

	return nil

}
// StartWebSocketServer runs a separate Mux router on a different port
func StartWebSocketServer(f *fiber.App) {
    
	port := ":8080" // WebSocket service port
	log.Println("WebSocket server running on", port)
	if err := f.Listen(port); err != nil {
		log.Fatal("WebSocket server failed:", err)
	}
}
