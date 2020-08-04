package container

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/renjingneng/a_simple_go_project/core/config"
	"strings"
)

var redisContainer map[string]interface{}

//GetEntityFromRedisContainer is
func GetEntityFromRedisContainer(database string, mode string) interface{} {
	if database == "" || mode == "" {
		return nil
	}
	dbname := database + mode
	if db, ok := redisContainer[dbname]; ok {
		return db
	}
	if _, ok := config.DatabaseMap[dbname]; !ok {
		return nil
	}
	var db interface{}
	if mode == "Single" {
		db = newClient(config.DatabaseMap[dbname])
	} else if mode == "Cluster" {
		addrs := strings.Split(config.DatabaseMap[dbname], ",")
		db = newClusterClient(addrs)
	}
	redisContainer[dbname] = db
	return db
}
func init() {
	if redisContainer == nil {
		redisContainer = make(map[string]interface{})
	}
}
func newClient(addr string) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return db
}
func newClusterClient(addrs []string) *redis.ClusterClient {
	db := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})
	return db
}
