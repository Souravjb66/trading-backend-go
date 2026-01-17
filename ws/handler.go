package ws

import (
	"encoding/json"
	"log"
	// "net/http"

	// "github.com/gorilla/mux"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	// "trading/services"

	// "trading/config"

	// "github.com/gofiber/contrib/websocket"
	// "github.com/gofiber/fiber/v2"
	"strconv"
)

// upgrader upgrades HTTP to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var PubSubSystem = NewPubSub()
// HandleConnection upgrades HTTP → WS and registers the client
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("in the ws con")
	q:=r.URL.Query()
	id:=q.Get("id")
	userId,err:=strconv.Atoi(id)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("userId not valid"))
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

    
	client := Client{
			Conn: conn,
			Send: make(chan []byte),

		}
	RegisterClient(&client,userId)
	// sub := PubSubSystem.Subscribe("portfolio_updates")

	// for msg := range sub {
	// 	conn.WriteJSON(msg)
	// }

	go ReadClientMessage(conn, userId,r)
	go WriteClientMessage(conn, userId)
}
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

func ReadClientMessage(c *websocket.Conn,userId int,r *http.Request){
	// app:=config.TradeServer.WebSocketRoute
	log.Println("userid",userId)
	var(
			// dataType int
			msg []byte
			err error
			msgFormat MsgFormat
			

		)
	_,ok:=isConnectionAlliveMap[userId]
	if !ok{
		log.Println("error in map")
		return 

	}
	
	for{

		log.Println("in read msg")
		if _,msg,err=c.ReadMessage();err!=nil{
			log.Println("error in read ",err)
			cl:=WebsocketConnections[userId]
            UnregisterClient(cl,userId)
			break
		}
	    err=json.Unmarshal(msg, &msgFormat)
		if err!=nil{
			log.Println("errror in unmarsal msg")
			return
		}
		// log.Println(dataType)
		// log.Println(string(msg))
		log.Println(msgFormat.Data["uu"])
		switch msgFormat.Type{
		case ALL_CLOSE_TRADE:

		case ALL_OPEN_TRADE:
		case USER_PROFILE:
		case OPEN_ORDERS:
			// a:=msgFormat.Data["asset"].(string)
	
			// if err!=nil{
			// 	log.Println("error in geting data")

			// }else{
			// 	cl:=WebsocketConnections[userId]
			// 	dt,err:=json.Marshal(res)
			// 	if err!=nil{
			// 	    log.Println("error in geting data")

			//     }
			// 	go func(){
			// 		cl.Send<-dt
			// 	}()
			// }

		
		}

	}
		

}
func WriteClientMessage(c *websocket.Conn , userId int){
	// app:=config.TradeServer.WebSocketRoute
	    // var ok bool
        // var msg []byte
		client,ok:=WebsocketConnections[int(userId)]
		if !ok{
			log.Println("value not present in the map")
			return
		}
		//ignore the warning
		for{
			
			select{
			case msg,ok:= <-client.Send:
				if !ok{
					log.Println("channel is closed")
					// break
					return
				}
				if err:=c.WriteMessage(websocket.TextMessage, msg);err!=nil{
					log.Println("error in send msg :",err)
					// continue
				}
				



			}
		

		}
		

}

// StartWebSocketServer runs a separate Mux router on a different port
func StartWebSocketServer(f *chi.Mux) {
    
	port := ":8082" // WebSocket service port
	log.Println("WebSocket server running on", port)
	if err :=http.ListenAndServe(":8080", f); err != nil {
		log.Fatal("WebSocket server failed:", err)
	}
}
