package controllers

import (
	"encoding/json"
	"log"
	"trading/db"
	"trading/services"
    "strconv"
	"github.com/go-chi/chi/v5"
	"net/http"

)


func SignUpUserController(w http.ResponseWriter,r *http.Request){
	// bd:=ctx.Body()
	var user db.Users
    json.NewDecoder(r.Body).Decode(&user)
	// err:=json.Unmarshal(bd, &user)
	// if err!=nil{
	// 	return err
	// }

	err:=services.SignUp(user.Username, user.Email, user.PasswordHash)
	w.Header().Set("Content-Type", "application/json")

    // 2. Set the status code
    
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in creating user"))

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfull"))

}
func LoginUserController(w http.ResponseWriter,r *http.Request){
	// bd:=ctx.Body()
	var user db.Users
	// err:=json.Unmarshal(bd, &user)
	// if err!=nil{
	// 	return err
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	ok,value:=services.Login(user.Email, user.PasswordHash)
	if !ok{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(value))
		
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfull"))

}

func CreateOrderController(w http.ResponseWriter,r *http.Request){
	// bd:=ctx.Body()
	w.Header().Set("Content-Type", "application/json")
	var order db.Orders
	// err:=json.Unmarshal(bd, &order)
	// if err!=nil{
	// 	return err
	// }
	json.NewDecoder(r.Body).Decode(&order)
	err:=services.CreateOrder(uint64(order.UserID), order.Asset, string(order.Type), order.Price, order.Quantity, order.RemainingQuantity)
	if err!=nil{
		log.Println("error in creating order:",err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in creatig order"))

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully order created"))


}

func CreatePortfolioController(w http.ResponseWriter,r *http.Request){
	// bd:=ctx.Body()
	w.Header().Set("Content-Type", "application/json")
	var portfolio db.Portfolio
	// err:=json.Unmarshal(bd, &portfolio)
	// if err!=nil{
	// 	return err
	// }
	json.NewDecoder(r.Body).Decode(&portfolio)
	err:=services.CreateUserPortfolio(uint64(portfolio.UserID), portfolio.Asset, portfolio.Quantity)
	if err!=nil{
		log.Println("error in creating portfolior:",err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in create portfolio"))

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully portfolio created"))

}
func GetTradesController(w http.ResponseWriter,r *http.Request){
	// param:=ctx.Params("userId")
	w.Header().Set("Content-Type", "application/json")
	param := chi.URLParam(r, "userID")
    userId,ok:=strconv.Atoi(param)
	if ok!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in get userid"))

	}
	trades,err:=services.GetUserTradeByUserId(uint64(userId))
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in get trades"))
	}
    json.NewEncoder(w).Encode(trades)
}
func GetOrdersController(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	asset:=chi.URLParam(r, "asset")
	res,err:=services.ShowLiveOrder(asset)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in geting data"))
	}
	json.NewEncoder(w).Encode(res)
	
	

}