package Models

import	(
	"github.com/go-redis/redis"
	"fmt"
	"os"
)


func ExampleNewClient() {
	fmt.Println("Running up new Redis client")
	client := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_URL", "redis:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)
	pong, err := client.Ping().Result()
	fmt.Println(pong,err)
	if err==nil{
		fmt.Println("Connected Succesfully to Redis Cache Server")
	}
	// Output: PONG <nil>
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}