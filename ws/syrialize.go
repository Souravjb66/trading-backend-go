package ws

import(
	"net/http"
	"encoding/json"
	"log"
)

type DataForCreateOrder struct{
	Type string `json:"type"`
	Order_Id int `json:"order_id"`
	Data  map[string]interface{} `json:"data"`

}

type DataForCreateTrade struct{
	Type string `json:"type"`
	Trade_Id int `json:"trade_id"`
	Data  map[string]interface{} `json:"data"`

}

type DataForGetOrders struct{
	Type string `json:"type"`
	Order_Id int `json:"order_id"`
	Data  map[string]interface{} `json:"data"`

}

type DataForUpdatedPortfolio struct{
	Type string `json:"type"`
	Portfolio_Id int `json:"portfolio_id"`
	Data  map[string]interface{} `json:"data"`

}
type DataForGetPortfolio struct{
	Type string `json:"type"`
	Portfolio_Id int `json:"portfolio_id"`
	Data  map[string]interface{} `json:"data"`

}


type BuyOrderReq struct{
	UserId int64 `json:"userId"`
	Asset  string `json:"asset"`
	Type   string  `json:"type"`
	Price  int64  `json:"price"`
	Quantity int64  `json:"quantity"`
	Remaining_quantity  int64  `json:"remaining_quantity"`

}

const(
	ALL_CLOSE_TRADE = "CLOSE_TRADES"
	ALL_OPEN_TRADE = "OPEN_TRADES"
	USER_PROFILE="PROFILE"

)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Failed to encode JSON: %v", err)
    }
}