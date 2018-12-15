package backend

// TODO : add error checking

import (
	"log"
	"path/filepath"

	"github.com/go-redis/redis"
)

var client *redis.Client

// SetUpRedis sets up redis
func SetUpRedis() {
	setRedisClient(createRedisClient())
}

// createData creates a map
func createData(key string, mp map[string]string) {
	mapInterface := make(map[string]interface{})
	for k, v := range mp {
		mapInterface[k] = v
	}
	res := client.HMSet(key, mapInterface)
	logResponse(res)
}

// readFile
func readData(filename string) (map[string]string, error) {
	val, res := client.HGetAll(filename).Result()
	logResponse(res)
	if res != nil {
		log.Printf("Read data error: %v", res)
		return nil, res
	}
	return val, nil
}

// updateFile
func updateData(key string, mp map[string]interface{}) {
	res := client.HMSet(key, mp)
	logResponse(res)
}

// deleteFile
func deleteData(key string) {
	res := client.HDel(key)
	logResponse(res)
}

// createRedisClient creates a redis client
func createRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "rdb:6379",
		Password: "",
		DB:       0,
	})

}

// setRedisClient sets up the redis client
func setRedisClient(c *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "rdb:6379",
		Password: "",
		DB:       0,
	})
}

func getAbs(filename string) string {
	abName, res := filepath.Abs("data/" + filename)
	logResponse(res)
	return abName
}

func logResponse(res interface{}) {
	log.Println(res)
}
