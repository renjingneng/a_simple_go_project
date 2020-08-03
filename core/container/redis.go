package container

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/renjingneng/a_simple_go_project/core/config"
)

var redisContainer map[string]*redis.Client

//GetEntityFromRedisContainer is
func GetEntityFromRedisContainer(database string, mode string) *redis.Client {
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
	db := redis.NewClient(&redis.Options{
		Addr:     config.DatabaseMap[dbname],
		Password: "",
		DB:       0,
	})
	redisContainer[dbname] = db
	return db
}
func init() {
	if redisContainer == nil {
		redisContainer = make(map[string]*redis.Client)
	}
}
