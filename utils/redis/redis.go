package redis
import (
	"fmt"
	"github.com/go-redis/redis"
)
var redisDB *redis.Client
//redisDB的初始化链接
func RedisNewClient() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisDB.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}



