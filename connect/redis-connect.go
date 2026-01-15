package connect

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)
type RedisStruct struct{
	Redis  *redis.Client
    RedisPubSubForBuyOrder *redis.PubSub
	RedisPubSubForCreateOrder *redis.PubSub

}
var redisInstance RedisStruct
func RedisConfigSetUp(){
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // default Redis port
	})
	redisInstance.Redis=rdb


}
func GetRedisInstance()*RedisStruct{
	return &redisInstance
}
const buyChannel string="buy_channel"
const createOrderChannel string="sell_channel"
func SubScribe(){
	sub1:=redisInstance.Redis.Subscribe(context.Background(),buyChannel)   //used with orders when order buy
	sub2:=redisInstance.Redis.Subscribe(context.Background(),createOrderChannel)  //when create an order by seller
    redisInstance.RedisPubSubForBuyOrder=sub1
	redisInstance.RedisPubSubForCreateOrder=sub2

}
type PublishStruct struct{
	Response string

}
func PublishOrderBuy(msg PublishStruct){
	err := redisInstance.Redis.Publish(context.Background(), buyChannel, msg).Err()
		if err != nil {
			log.Println("publish error:", err)
		}


}
func PublishOrderCreate(msg PublishStruct){
	err := redisInstance.Redis.Publish(context.Background(), createOrderChannel, msg).Err()
	if err != nil {
		log.Println("publish error:", err)
	}

}