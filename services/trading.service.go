package services

import (

	// "github.com/gofiber/fiber/v2"
	"context"
	"log"
	"trading/config"
	"trading/db"
	// "sync"
	// "strconv"
	sqlcdb "trading/db"
	"trading/ws"
	"encoding/json"
	// "encoding/json"
	// "container/heap"
	
)

func GetUserById(id uint64)(db.Users,error){
	db:=config.OpenMysqlConnectionQuery()
	defer db.Close()
	res,err:=db.GetUserByID(context.Background(),int64(id))
	if err!=nil{
		log.Println("error in service ",err)
		return res,err
	}
	
	return res,nil
	

}
func GetUserByUserName(){
	

}


func CreateTrade(trade db.Trades)(interface{},error){
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	params:=db.InsertTradeParams{
			BuyOrderID: trade.BuyOrderID,
			SellOrderID :trade.SellOrderID,
            Asset: trade.Asset,      
            Price:trade.Price,       
            Quantity :trade.Quantity,  
		}
	data,err:=dB.InsertTrade(context.Background(),params)
	if err!=nil{
		return nil,err
	}
		
	

	// totalOpenOrders,err:=dB.GetOpenOrdersByAsset(context.Background(),trade.Asset)
	// if err!=nil{
	// 	return nil,err
	// }
	for _,con:=range ws.WebsocketConnections{
		res:=map[string]interface{}{
			"type":"trade",
			"trade_id":data.ID,
			"data":map[string]interface{}{
				"asset":data.Asset,
				"price":data.Price,
				"value":data.Quantity,

			},
		}
		btData,err:=json.Marshal(res)
		if err!=nil{
			return nil,err

		}
		go func(){
			con.Send<-btData
		}()
		
			


	}
	
	// msg:=producer.SendMessage{
	// 	Event: "Trade",
	// 	UserId: trade.SellOrderID,
	// 	Response: totalOpenOrders,

	// }
	// data,err:=json.Marshal(msg)
	// if err!=nil{
	// 	log.Println(err)
	// 	return nil,err
	// }
	// producer.KafkaSendEvents(data)
	return nil,nil


}
func GetUserTradeByUserId(userId uint64)(interface{},error){
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	res,err:=dB.GetTradesByUserID(context.Background(),int64(userId))
	if err!=nil{
		return nil,err
	}
	return map[string]interface{}{
		"res":res,
	},nil

}

func GetUserPortFolioByUserId(userId int)(interface{},error){
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	res,err:=dB.GetPortfolioByUserID(context.Background(),int64(userId))
	if err!=nil{
		return nil,err
	}
	return map[string]interface{}{
		"res":res,
	},nil

}

func UpdateUserPortFolio(userId uint64,asset string,balance int64)error{
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	params:=db.UpdatePortfolioBalanceParams{
		Quantity: balance, 
        UserID: int64(userId),  
        Asset:asset,   
	}
	data,err:=dB.UpdatePortfolioBalance(context.Background(),params)
	if err!=nil{
		return err
	}
	for _,con:=range ws.WebsocketConnections{
		res:=map[string]interface{}{
			"type":"portfolio",
			"trade_id":data.ID,
			"data":map[string]interface{}{
				"asset":data.Asset,
				"value":data.Quantity,

			},
		}
		btData,err:=json.Marshal(res)
		if err!=nil{
			return err

		}
		go func(){
			con.Send<-btData

		}()
		
		
			


	}
	
	// uData,err:=dB.GetPortfolioByUserID(context.Background(),int64(userId))
	// dt:=ws.SendResponse{
	// 	Type: "portfolio",
	// 	Data: uData,
	// }
	// log.Println(dt)
	// if err!=nil{
	// 	log.Println(err)
	// }
	// else{
	// 	data, _ := json.Marshal(dt)
	// 	ws.GetWbsocketConnection().Send <-data

	// }
	
	return nil


}

func CreateUserPortfolio(userId uint64,asset string,balance int64)error{
	log.Println("enter in create portfolio",balance)
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	params:=db.InsertPortfolioParams{
        UserID: int64(userId),  
        Asset:asset,   
        Quantity:balance, 
	}
	data,err:=dB.InsertPortfolio(context.Background(),params)
	if err!=nil{
		log.Println("erro in portfolio ",err)
		// return err
	}
	for _,con:=range ws.WebsocketConnections{
		res:=map[string]interface{}{
			"type":"portfolio",
			"trade_id":data.ID,
			"data":map[string]interface{}{
				"asset":data.Asset,
				"value":data.Quantity,

			},
		}
		btData,err:=json.Marshal(res)
		if err!=nil{
			log.Println("error in creating portfolio :",err)
			// return err

		}

		go func(){
		    con.Send<-btData
		}()
		
			


	}
	log.Println("last in portfolio")	
	return nil

}


func CreateOrder(userId uint64,asset string,orderType string,price int64,quantity int64,remainingQuantity int64)error{
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	params:=db.CreateOrderParams{
		UserID:int64(userId),            
        Asset :asset,            
        Type :db.OrderType(orderType),            
        Price :price,            
        Quantity:quantity,          
        RemainingQuantity :remainingQuantity,
	}
	data,err:=dB.CreateOrder(context.Background(),params)
    
	// uData,err:=dB.GetOpenOrdersByAsset(context.Background(),asset)
	// dt:=ws.SendResponse{
	// 	Type: "orders",
	// 	Data: uData,

	// }
	for _,con:=range ws.WebsocketConnections{
		res:=map[string]interface{}{
			"type":"order",
			"order_id":data.ID,
			"data":map[string]interface{}{
				"asset":data.Asset,
				"type":data.Type,
				"price":data.Price,
				"value":data.Quantity,

			},
		}
		btData,err:=json.Marshal(res)
		if err!=nil{
			log.Println("error in creating order ",err)
			// return err

		}
		
		go func(){
			con.Send<-btData
		}()
		
			


	}

	if err!=nil{
		log.Println(err)
	}
	Match(&data)
	// log.Println(dt)
	// else{
	// 	data, _ := json.Marshal(dt)
	// 	ws.GetWbsocketConnection().Send <-data

	// }
	return nil

}
func ShowLiveOrder(asset string)(interface{},error){
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	res,err:=dB.GetOpenOrdersByAsset(context.Background(),asset)
	if err!=nil{
		return nil,err
	}
	return map[string]interface{}{
		"res":res,
	},nil
	

}
func UpdateOrderStatus(id uint64,remainingQuantity int64,status string)error{
	dB:=config.OpenMysqlConnectionQuery()
	defer dB.Close()
	params:=db.UpdateOrderStatusParams{
		RemainingQuantity:remainingQuantity, 
        Status:db.OrderStatus(status),            
        ID:int64(id),                

	}
	data,err:=dB.UpdateOrderStatus(context.Background(),params)
	if err!=nil{
		return err
	}
	for _,con:=range ws.WebsocketConnections{
		res:=map[string]interface{}{
			"type":"order",
			"trade_id":data.ID,
			"data":map[string]interface{}{
				"asset":data.Asset,
				"price":data.Price,
				"status":data.Status,
				"value":data.Quantity,

			},
		}
		btData,err:=json.Marshal(res)
		if err!=nil{
			return err

		}
		go func(){
			con.Send<-btData
		}()
		
			


	}

	return nil


}
// var mu sync.Mutex

func TradeLogic(
	buy *sqlcdb.Orders,
	sell *sqlcdb.Orders,
	tradeQty int,
	tradePrice int,
)error{

	db := config.OpenMysqlConnectionQuery()
	defer db.Close()

	ctx := context.Background()

	tx, err := config.TradeServer.PostgreSQL.PostgreSQL.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		// return
	}

	qtx := db.WithTx(tx)

	// defer func() {
	// 	if err != nil {
	// 		err=tx.Rollback()
	// 		if err!=nil{
	// 			return
	// 		}
	// 	}
	// }()

	total := int64(tradeQty * tradePrice)

	// update buyer order
	// err = qtx.UpdateOrderStatus(ctx,
	// 	sqlcdb.UpdateOrderFilledParams{
	// 		ID:                int64(buy.ID),
	// 		FilledQuantity:    tradeQty,
	// 		RemainingQuantity: buy.Qty - tradeQty,
	// 		Status:            getStatus(buy.Qty-tradeQty),
	// 	})
	var buystatus string="OPEN"
	if buy.Quantity- int64(tradeQty) <=0{
		buystatus="CLOSE"

	}
	_,err=qtx.UpdateOrderStatus(ctx, sqlcdb.UpdateOrderStatusParams{
		ID: int64(buy.ID),
		RemainingQuantity: buy.Quantity- int64(tradeQty),
		Status: sqlcdb.OrderStatus(buystatus),

	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
		// return
	}

	// pdate seller order
	// err = qtx.UpdateOrderFilled(ctx,
	// 	sqlcdb.UpdateOrderFilledParams{
	// 		ID:                int64(sell.ID),
	// 		FilledQuantity:    tradeQty,
	// 		RemainingQuantity: sell.Qty - tradeQty,
	// 		Status:            getStatus(sell.Qty-tradeQty),
	// 	})
	_,err=qtx.UpdateOrderStatus(ctx, sqlcdb.UpdateOrderStatusParams{
		ID: int64(buy.ID),
		RemainingQuantity: sell.Quantity- int64(tradeQty),
		Status: sqlcdb.OrderStatus(buystatus),

	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	//  buyer portfolio (+)
	// err = qtx.AddPortfolio(ctx,
	// 	sqlcdb.AddPortfolioParams{
	// 		UserID:   buy.UserID,
	// 		Asset:    buy.Asset,
	// 		Quantity: tradeQty,
	// 	})
	_,err=qtx.UpdatePortfolioCreditQuantity(ctx, sqlcdb.UpdatePortfolioCreditQuantityParams{
		UserID: buy.UserID,
		Asset: buy.Asset,
		Quantity: int64(tradeQty),
	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	// seller portfolio (-)
	// err = qtx.SubPortfolio(ctx,
	// 	sqlcdb.SubPortfolioParams{
	// 		UserID:   sell.UserID,
	// 		Asset:    sell.Asset,
	// 		Quantity: tradeQty,
	// 	})
	_,err=qtx.UpdatePortfolioDebitQuantity(ctx, sqlcdb.UpdatePortfolioDebitQuantityParams{
		Quantity: int64(tradeQty),
		Asset: sell.Asset,
	    UserID: sell.UserID,

	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	// buyer balance (-)
	// err = qtx.SubUserBalance(ctx,
	// 	sqlcdb.SubUserBalanceParams{
	// 		UserID: buy.UserID,
	// 		Amount: total,
	// 	})
	_,err=qtx.DebitUserBalance(ctx, sqlcdb.DebitUserBalanceParams{
		ID: buy.UserID,
		Balance: total,


	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	//  seller balance (+)
	// err = qtx.AddUserBalance(ctx,
	// 	sqlcdb.AddUserBalanceParams{
	// 		UserID: sell.UserID,
	// 		Amount: total,
	// 	})
	_,err=qtx.UpdateUserBalanceDelta(ctx, sqlcdb.UpdateUserBalanceDeltaParams{
		ID: sell.UserID,
		Balance: total,
	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	err=tx.Commit()
	if err!=nil{
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}
	log.Printf("TRADE EXECUTED: BUY %d SELL %d QTY %d PRICE %d\n",
		buy.ID, sell.ID, tradeQty, tradePrice)
	return nil
}
