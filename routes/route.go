package routes

import (
	// "log"

	// "log"
	// "net/http"
	"trading/config"
	"trading/ws"

	// "github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	// "github.com/gorilla/mux"
	"trading/controllers"
	
	"github.com/gofiber/contrib/websocket"
)

// c *fiber.Ctx
func Routes() {
	app := fiber.New()
	
	app.Post("/sign-up",controllers.CreateUser)
	app.Get("/app", func(c *fiber.Ctx)error{
		return c.JSON("msg")
	})
	




	

	config.TradeServer.Route = app
}



// func WebSocketRoute(){
// 	r := mux.NewRouter()
// 	r.HandleFunc("/ws", ws.HandleConnection)
// 	ws.WebSocketRouteSet(r)

// }

func WebSocketRoute(){
	app:=fiber.New()
	app.Use("/ws",func (ctx *fiber.Ctx)error  {   //it checks the route have upgrade ws method or not ,its a middleware
		if websocket.IsWebSocketUpgrade(ctx){
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired

		
	})
	app.Get("/live", ws.UpgradeToWs)
	// app.Get("/", websocket.New(func(c *websocket.Conn){
	// 	log.Println(c.Locals("allowed"))  // true
	// 	log.Println(c.Params("id"))       // 123
	// 	log.Println(c.Query("v"))         // 1.0
	// 	log.Println(c.Cookies("session")) // ""
	// 	var (
	// 		mt  int
	// 		msg []byte
	// 		err error
	// 	)
	// 	for{
	// 		if mt,msg,err:=c.ReadMessage();err!=nil{
	// 			log.Println("error :",err)
	// 			break
	// 		}
	// 		if err=c.WriteMessage(mt, msg);err!=nil{
	// 			log.Println("error ",err)
	// 			break
	// 		}
	// 	}


	// }))
	config.TradeServer.WebSocketRoute=app
	go ws.StartWebSocketServer(app)
	
	

}