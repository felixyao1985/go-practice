package redisx

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Nil reply Redis returns when key does not exist.
const Nil = redis.Nil

var redisStore = &sync.Map{}

// Conf redis
type Conf struct {
	DSN string
}

// Register initialize redis
// dsn format -> redis://:password@url/dbNum[optional,default 0]
func Register(name string, conf Conf) *redis.Client {
	if s := os.Getenv("REDIS_DSN_" + strings.ToUpper(name)); s != "" {
		conf.DSN = s
	}

	opt, err := redis.ParseURL(conf.DSN)
	if err != nil {
		panic(err)
	}

	// TODO opt set using config
	client := redis.NewClient(opt)
	act, ok := redisStore.LoadOrStore(name, client)
	if ok {
		// already exist
		client.Close()
		return act.(*redis.Client)
	}

	return client
}

// Client return redis client
func Client(name string) (*redis.Client, error) {
	v, ok := redisStore.Load(name)
	if !ok {
		return nil, fmt.Errorf("redis %q not registered", name)
	}

	return v.(*redis.Client), nil
}

// Use return redis client with given name or nil if not exist
func Use(name string) *redis.Client {
	cli, _ := Client(name)
	return cli
}

// HealthCheck ping rds
func HealthCheck() error {
	errs := make(map[string]error)

	redisStore.Range(func(k, v interface{}) bool {
		err := v.(*redis.Client).Ping().Err()
		if err != nil {
			errs[k.(string)] = err
		}
		return true
	})

	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}

	return nil
}

// Close close all redis conn
func Close() error {
	redisStore.Range(func(k, v interface{}) bool {
		v.(*redis.Client).Close()
		return true
	})
	return nil
}

// Duration is helper func to gen time.Second
func Duration(sec int64) time.Duration {
	return time.Duration(sec) * time.Second
}

// PushAlias is helper func to get jpush alias that app registered.
func PushAlias(cli *redis.Client, coder interface{ Code() int32 }, uId string) string {
	key := fmt.Sprintf("%d_%s", coder.Code(), uId)

	if jti, err := cli.Get("u:sso:" + key).Result(); err == nil {
		return key + "_" + jti
	}

	// no need to push
	return ""
}

// UserJTI
func UserJTI(cli *redis.Client, coder interface{ Code() int32 }, uId string) string {
	key := fmt.Sprintf("%d_%s", coder.Code(), uId)

	if jti, err := cli.Get("u:sso:" + key).Result(); err == nil {
		return jti
	}

	return ""
}

func CheckFormId(cli *redis.Client, uId string) bool {
	if uId == "" {
		return true
	}
	if _, err := cli.Get("api:formid:" + uId).Result(); err == nil {
		return false
	} else {
		cli.SetNX("api:formid:"+uId, uId, time.Minute*15)
	}

	return true
}
