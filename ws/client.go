package ws

import (
	// "log"

	// "github.com/gorilla/websocket"
	"github.com/gofiber/contrib/websocket"

)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	// UserId int
}
type SendResponse struct{
	Type string  `json:"type"`
	Data interface{}  `json:"data"`
}
// var websocketConnection *Client
var  WebsocketConnections= make(map[int]*Client)  //user id key
var isConnectionAlliveMap= make(map[int]bool) //userid key 
var Rooms = make(map[string][]*websocket.Conn)  //making rooms , here one room example "btc" stores multiple connection
var RoomName="btc"
// func GetWbsocketConnection()*Client{
// 	return websocketConnection

// }
func SetMessageToClientQueue(userId int64,msg []byte){
	value:=WebsocketConnections[int(userId)]
	value.Send<-msg
	


}
func GetWebSocketCon(userid int64)*websocket.Conn{
	if isActive:=isConnectionAlliveMap[int(userid)];!isActive{
		return nil

	}
	con:=WebsocketConnections[int(userid)]
	return con.Conn
	

}

// func NewClient(conn *websocket.Conn) *Client {
// 	websocketConnection=&Client{
// 		Conn: conn,
// 		Send: make(chan []byte),
// 	}
// 	return &Client{
// 		Conn: conn,
// 		Send: make(chan []byte),
// 	}
// }

// func (c *Client) ReadMessages() {
// 	defer c.Conn.Close()

// 	for {
// 		_, msg, err := c.Conn.ReadMessage()
// 		if err != nil {
// 			log.Println("read error:", err)
// 			break
// 		}
// 		log.Printf("Recm map[Type]Type1eived from client: %s", msg)

// 		// (optional) process or forward msg to gRPC or Redis
// 	}
// }

// func (c *Client) WriteMessages() {
// 	defer c.Conn.Close()

// 	for msg := range c.Send {
// 		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
// 			log.Println("write error:", err)
// 			break
// 		}
// 	}
// }
