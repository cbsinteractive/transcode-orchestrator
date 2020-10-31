package db

import (
	"encoding/json"
	"errors"
	"log"
	"net"

	"github.com/go-redis/redis"
)

var (
	ErrJobNotFound = errors.New("job not found")
)

type Options struct {
	Addr string
	DB   int
}

func NewClient(opt *Options) (*Client, error) {
	if opt == nil {
		opt = &Options{}
	}
	if opt.Addr == "" {
		opt.Addr = "localhost:6379"
	}
	_, _, err := net.SplitHostPort(opt.Addr)
	if err != nil {
		opt.Addr = net.JoinHostPort(opt.Addr, "6379")
	}
	f := &Client{
		rc: redis.NewClient(&redis.Options{
			Addr:     opt.Addr,
			DB:       opt.DB,
			Password: "",
		}),
	}
	return f, nil
}

type Client struct {
	rc *redis.Client
}

func (c *Client) Get(key string, dst interface{}) (err error) {
	log.Printf("get key: %q", key)
	defer func() { log.Printf("get key: %q err: %v", key, err) }()
	val, err := c.rc.Get(key).Result()
	if err == redis.Nil {
		return ErrJobNotFound
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dst)
}

func (c *Client) Put(key string, val interface{}) (err error) {
	log.Printf("put key: %q", key)
	data, _ := json.Marshal(val)
	defer func() { log.Printf("put key: %q data: %q err: %v", key, data, err) }()
	return c.rc.Set(key, string(data), 0).Err()
}
