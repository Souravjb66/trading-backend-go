package routes

import (
	// "log"

	// "log"
	// "net/http"
	// "log"
	// "strconv"
	"trading/config"
	"trading/ws"

	// "github.com/gofiber/adaptor/v2"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gorilla/mux"
	"github.com/go-chi/chi/v5"
	"trading/controllers"
    // "time"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/websocket"
	// "github.com/gofiber/contrib/websocket"
	"github.com/go-chi/chi/v5/middleware"
)

// c *fiber.Ctx
func Routes() {
	app := chi.NewRouter()
	app.Use(middleware.RequestID)
    app.Use(middleware.RealIP)
    app.Use(middleware.Logger)
    app.Use(middleware.Recoverer)
	
	//
	app.HandleFunc("/ws/live", ws.HandleConnection)
	config.TradeServer.WebSocketRoute=app
	go ws.StartWebSocketServer(app)
	//

	app.Post("/sign-up",controllers.SignUpUserController)
	app.Post("/login",controllers.LoginUserController)
	app.Post("/create-portfolio",controllers.CreatePortfolioController)
	app.Post("/create-order",controllers.CreateOrderController)
	app.Get("/orders",controllers.GetOrdersController)
	app.Get("/trades",controllers.GetTradesController)
	




	

	config.TradeServer.Route = app
}



// func WebSocketRoute(){
// 	r := mux.NewRouter()
// 	r.HandleFunc("/ws", ws.HandleConnection)
// 	ws.WebSocketRouteSet(r)
	
// 	config.TradeServer.WebSocketRoute=r
// 	go ws.StartWebSocketServer(r)

// }

// func WebSocketRoute(){
// 	app:=
// 	app.Use("/ws",func (ctx *fiber.Ctx)error  {   //it checks the route have upgrade ws method or not ,its a middleware
// 		if websocket.IsWebSocketUpgrade(ctx){
// 			ctx.Locals("allowed", true)
// 			return ctx.Next()
// 		}
// 		return fiber.ErrUpgradeRequired

		
// 	})
// 	// app.Get("/ws/live", ws.UpgradeToWs)
// 	app.Get("/ws/live", websocket.New(func(c *websocket.Conn){
// 		log.Println("enter in ws")
		
// 		log.Println("enter in ws uppp")
// 		log.Println(c.Locals("allowed"))  // true
// 		log.Println(c.Query("id"))       // 123
// 		log.Println(c.Query("v"))         // 1.0
// 		log.Println(c.Cookies("session")) // ""
// 		userId,err:=strconv.Atoi(c.Query("id"))   //need to convert it to int
// 		if err!=nil{
// 			log.Println(" error in parsing userid :",err)
// 			return
// 		}
// 		ws.RegisterClient(&ws.Client{
// 			Conn: c,
// 			Send: make(chan []byte),

// 		},userId)
// 		time.Sleep(1* time.Second)
// 		log.Println("conn")
// 		go ws.ReadClientMessage(c,userId)
// 		go ws.WriteClientMessage(c,userId)
		
		
		
	 


// 	}))


// 	config.TradeServer.WebSocketRoute=app
// 	go ws.StartWebSocketServer(app)
	
	

// }