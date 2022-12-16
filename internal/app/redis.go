package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli"
)

//nolint:gochecknoglobals
var (
	redisMasterNameFlag = &cli.StringFlag{
		Name:   "redis-master-name",
		EnvVar: "REDIS_MASTER_NAME",
		Value:  "",
		Usage:  "Master name for redis sentinel",
	}
	redisAddrsFlag = &cli.StringSliceFlag{
		Name:   "redis-addrs",
		EnvVar: "REDIS_ADDRS",
		Value:  &cli.StringSlice{"localhost:6379"},
		Usage:  "A list of address for connecting to redis. Default: localhost:6379",
	}
	redisDBFlag = &cli.IntFlag{
		Name:   "redis-db",
		EnvVar: "REDIS_DB",
		Value:  0,
		Usage:  "Redis database index",
	}
	redisUsernameFlag = &cli.StringFlag{
		Name:   "redis-username",
		EnvVar: "REDIS_USERNAME",
		Value:  "",
		Usage:  "Username for authenticating with redis server",
	}
	redisPasswordFlag = &cli.StringFlag{
		Name:   "redis-password",
		EnvVar: "REDIS_PASSWORD",
		Value:  "",
		Usage:  "Password for authenticating with redis server",
	}
	redisKeyPrefixFlag = &cli.StringFlag{
		Name:   "redis-key-prefix",
		EnvVar: "REDIS_KEY_PREFIX",
		Value:  "",
		Usage:  "Prefix of key for redis",
	}
	redisReadTimeoutFlag = &cli.DurationFlag{
		Name:   "redis-read-timeout",
		EnvVar: "REDIS_READ_TIMEOUT",
		Value:  0,
		Usage:  "Timeout for redis read operation",
	}
	redisWriteTimeoutFlag = &cli.DurationFlag{
		Name:   "redis-write-timeout",
		EnvVar: "REDIS_WRITE_TIMEOUT",
		Value:  0,
		Usage:  "Timeout for redis write operation",
	}
)

// NewRedisFlags returns flags for evmlistener.
func RedisFlags() []cli.Flag {
	return []cli.Flag{
		redisMasterNameFlag, redisAddrsFlag, redisDBFlag,
		redisUsernameFlag, redisPasswordFlag, redisKeyPrefixFlag,
		redisReadTimeoutFlag, redisWriteTimeoutFlag,
	}
}

func RedisOptionFromFlags(c *cli.Context) *redis.UniversalOptions {
	var (
		masterName                                             = c.String(redisMasterNameFlag.Name)
		sentinelUsername, sentinelPassword, username, password string
	)
	if masterName != "" {
		sentinelUsername = c.String(redisUsernameFlag.Name)
		sentinelPassword = c.String(redisPasswordFlag.Name)
	} else {
		username = c.String(redisUsernameFlag.Name)
		password = c.String(redisPasswordFlag.Name)
	}
	return &redis.UniversalOptions{
		MasterName:       c.String(redisMasterNameFlag.Name),
		Addrs:            c.StringSlice(redisAddrsFlag.Name),
		DB:               c.Int(redisDBFlag.Name),
		Username:         username,
		Password:         password,
		SentinelUsername: sentinelUsername,
		SentinelPassword: sentinelPassword,
		ReadTimeout:      c.Duration(redisReadTimeoutFlag.Name),
		WriteTimeout:     c.Duration(redisWriteTimeoutFlag.Name),
	}
}
