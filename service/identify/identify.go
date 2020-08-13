// @Description
// @Author  renjingneng
// @CreateTime  2020/8/13 21:20
package identify

import (
	"encoding/json"
	"time"

	"github.com/renjingneng/a_simple_go_project/data/mysql"
	"github.com/renjingneng/a_simple_go_project/data/redis"
)

type Identify struct {
	prefix        string
	tableUserInfo *mysql.Open
	redis         *redis.Open
}

func NewIdentify() *Identify {
	thisService := &Identify{
		prefix:        "open_identify_",
		tableUserInfo: mysql.NewOpen(),
		redis:         redis.NewOpen(),
	}
	thisService.tableUserInfo.SetTablename("op_identify_user_info")
	return thisService
}

func (thisService *Identify) GetInfoByToken(token string) (map[string]string, error) {
	var res map[string]string
	str, err := thisService.redis.Get(thisService.prefix + token)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(str), &res); err != nil {
		return res, err
	}
	return res, nil
}

func (thisService *Identify) StoreInfoByToken(token string, info map[string]string, ttl time.Duration) error {
	str, _ := json.Marshal(info)
	thisService.redis.Set(thisService.prefix+token, string(str), ttl)
	return nil
}
