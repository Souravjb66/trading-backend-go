package services

import (
	"container/heap"
	"context"
	"log"
	// "strconv"
	"trading/config"
	"trading/db"
)

// highest bid first , highest buy price first trade match, for sell order price is less expensive to expensive ,
// buy price upper to low,sell price low to upper, so highest buyer price have get the match chance with sell order from low price to high

// type BookOrder []*db.Orders
type BuyerHeap []*db.Orders

// Len implements [heap.Interface].
func (b *BuyerHeap) Len() int {
	return len(*b)
	
}

// Len implements [heap.Interface].
// func (b *BuyerHeap) Len() int {

// }

type SellerHeap []*db.Orders

func (b BuyerHeap) Length() int {
	return len(b)

}

// MAX heap → higher price has higher priority
func (h BuyerHeap) Less(i, j int) bool {

	return h[i].Price > h[j].Price
}
func (h BuyerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *BuyerHeap) Push(x interface{}) {
	*h = append(*h, x.(*db.Orders))
}

func (h *BuyerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

//

func (h SellerHeap) Len() int { return len(h) }

// MIN heap → lower price has higher priority
func (h SellerHeap) Less(i, j int) bool {

	return h[i].Price < h[j].Price
}

func (h SellerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SellerHeap) Push(x interface{}) {
	*h = append(*h, x.(*db.Orders))
}

func (h *SellerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// buyHeap := &BuyHeap{}
// sellHeap := &SellHeap{}
// var buyHeap *BuyerHeap
// var sellHeap *SellerHeap
var buyHeap *BuyerHeap
var sellHeap *SellerHeap

func HeapInit() {
	buyHeap = &BuyerHeap{}
	sellHeap = &SellerHeap{}
	heap.Init(buyHeap)
	heap.Init(sellHeap)

}

func Match(order *db.Orders) {
	
	// value:=config.BookMap[order.Asset]
	for buyHeap.Len() > 0 && sellHeap.Len() > 0 {
		bestBuy := (*buyHeap)[0]
		bestSell := (*sellHeap)[0]

		// Trade condition
		if bestBuy.Price < bestSell.Price {
			break

		}
		// Execute trade
		tradeQty := min(bestBuy.Quantity, bestSell.Quantity)

		tb:= bestBuy.Quantity
		

		ts:= bestSell.Quantity
	
		tradeQt:=tradeQty
		
		tb -= tradeQt
		ts -= tradeQt

			// Remove filled orders
		if tb == 0 {
			heap.Pop(buyHeap)
				
				//
		}
		if ts == 0 {
			heap.Pop(sellHeap)
				//
		}
		
	}

}

func AddDbBuySellDataToHeap() {
	dB := config.OpenMysqlConnectionQuery()
	defer dB.Close()
	data, err := dB.GetOpenBuyOrders(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	sell_data, err := dB.GetOpenSellOrders(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, value := range data {
		heap.Push(buyHeap, value)
	}
	for _, value := range sell_data {
		heap.Push(sellHeap, value)
	}

}
