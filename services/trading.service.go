package services

import (

	// "github.com/gofiber/fiber/v2"
	"context"
	// "database/sql"
	"log"
	// "time"
	"trading/config"
	"trading/db"

	// "sync"
	// "strconv"
	"encoding/json"
	sqlcdb "trading/db"
	"trading/ws"
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
			log.Println("error in creating portfolio :",err)
			return err

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
			return err

		}
		
		go func(){
			con.Send<-btData
		}()
		
			


	}

	if err!=nil{
		log.Println(err)
	}
	AddDbBuySellDataToHeap()
	Match()
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
	tradeQty int,  //trade quantity
	tradePrice int,  //sell price
	isBuyOrderClose bool,
	isSellOrderCLose bool,
	buyerRemainingQuantity int64,
	sellerRemainingQuantity int64,
)error{
    log.Printf("TRADE STARTED BUY %d SELL %d QTY %d PRICE %d BUYER ORDER STATUS %v SELLER ORDER STATUS %v BUYER REMAINING QUANTITY %d SELLER REMAINING QUANTITY %d\n",buy.ID, sell.ID, tradeQty, tradePrice,isBuyOrderClose,isSellOrderCLose,buyerRemainingQuantity,sellerRemainingQuantity)
	db := config.OpenMysqlConnectionQuery()
	defer db.Close()

	ctx := context.Background()

	tx, err := config.TradeServer.PostgreSQL.PostgreSQL.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		// return
	}

	qtx := db.WithTx(tx)

	

	total := int64(tradeQty * tradePrice)

	var buystatus string="OPEN"
	if isBuyOrderClose{
		buystatus="FILLED"

	}
	var sellstatus = "OPEN"
	if isSellOrderCLose{
		sellstatus="FILLED"
	}
	//buy
	_,err=qtx.UpdateOrderStatus(ctx, sqlcdb.UpdateOrderStatusParams{
		ID: int64(buy.ID),
		RemainingQuantity: buyerRemainingQuantity,
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

	
	//seller
	_,err=qtx.UpdateOrderStatus(ctx, sqlcdb.UpdateOrderStatusParams{
		ID: int64(sell.ID),
		RemainingQuantity: sellerRemainingQuantity,
		Status: sqlcdb.OrderStatus(sellstatus),

	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	//  buyer portfolio (+)
    buyerPort,err:=qtx.GetPortfolioByUserID(ctx, buy.UserID)
	if err!=nil{
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}


	}
	
	//buyer
	_,err=qtx.UpdatePortfolioCreditQuantity(ctx, sqlcdb.UpdatePortfolioCreditQuantityParams{
		UserID: buy.UserID,
		Asset: buy.Asset,
		Quantity: buyerPort.Quantity+int64(tradeQty),
	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	// seller portfolio (-)
	sellerPort,err:=qtx.GetPortfolioByUserID(ctx, sell.UserID)
	if err!=nil{
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}


	}

	//seller
	_,err=qtx.UpdatePortfolioDebitQuantity(ctx, sqlcdb.UpdatePortfolioDebitQuantityParams{
		Quantity: sellerPort.Quantity-int64(tradeQty),
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
    buyerUser,err:=qtx.GetUserByID(ctx,buy.UserID)
    if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}
	//buyer
	_,err=qtx.DebitUserBalance(ctx, sqlcdb.DebitUserBalanceParams{
		ID: buy.UserID,
		Balance: buyerUser.Balance-total,


	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	//  seller balance (+)
	sellerUser,err:=qtx.GetUserByID(ctx,buy.UserID)
    if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}
	_,err=qtx.UpdateUserBalanceDelta(ctx, sqlcdb.UpdateUserBalanceDeltaParams{
		ID: sell.UserID,
		Balance: sellerUser.Balance+total,
	})
	if err != nil {
		log.Println(err)
		err=tx.Rollback()
			if err!=nil{
				return err
			}
	}

	

	_,err=qtx.InsertTrade(ctx, sqlcdb.InsertTradeParams{
		BuyOrderID: buy.ID,
		SellOrderID: sell.ID,
	    Asset: buy.Asset,
		Price: total,
		Quantity: int64(tradeQty),
	

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

	// go func(){

	// }()
	return nil
}
