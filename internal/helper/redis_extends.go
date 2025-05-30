package helper

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisGetDel = redis.NewScript(`
-- this script has side-effects, so it requires replicate commands mode
redis.replicate_commands()
local x = redis.call('GET', KEYS[1]);
if not x then
	return x
end;
redis.call('DEL', KEYS[1]);
return x
`)

var redisIncByExpire = redis.NewScript(`
-- this script has side-effects, so it requires replicate commands mode
redis.replicate_commands()

local count = tonumber(ARGV[1])
local expire = tonumber(ARGV[2])
local key = KEYS[1]

local tat = redis.call('INCRBY', key, count);
if count > 0 then
    redis.call("expire", key, expire)
end
return tat
`)

var rpopCountScript = redis.NewScript(`
local result = redis.call('lrange',KEYS[1],ARGV[1],-1) -- 0 ~ -1 开始拿
redis.call('ltrim',KEYS[1],0,ARGV[2])
return result
`)

type RedisExtend struct {
	redis.UniversalClient
}

func (r RedisExtend) GetDel(ctx context.Context, key string) *redis.Cmd {
	var values []interface{}
	return redisGetDel.Run(ctx, r, []string{key}, values...)
}
func (r RedisExtend) IncByExpire(ctx context.Context, key string, value int, expire time.Duration) *redis.Cmd {
	var values = []interface{}{
		value, expire.Seconds(),
	}
	return redisIncByExpire.Run(ctx, r, []string{key}, values...)
}
func (r RedisExtend) LPopCount(ctx context.Context, key string, count int) *redis.Cmd {
	return rpopCountScript.Run(ctx, r, []string{key}, []interface{}{
		-count,
		-count - 1,
	})
}
