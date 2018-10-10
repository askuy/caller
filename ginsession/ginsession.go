package ginsession

import (
	"github.com/BurntSushi/toml"
	"github.com/godefault/caller/common"
	"github.com/gin-contrib/sessions/redis"
	"fmt"
)

var defaultCaller *callerStore

type callerStore struct {
	caller redis.Store
	cfg    Cfg
}


func New() common.Caller {
	defaultCaller = &callerStore{}
	return defaultCaller
}

func Caller() redis.Store {
	return defaultCaller.caller
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := parseConfig(cfg, &c.cfg); err != nil {
		return err
	}
	c.initCaller()
	return nil
}

func (c *callerStore) Get(key string) interface{} {
	return c.caller
}

func (c *callerStore) Set(key string, val interface{}) {
	c.caller = val.(redis.Store)
}

func (c *callerStore) initCaller() {
	caller,err := provider(c.cfg.CallerGinSession)
	if err != nil {
		panic(err.Error())
	}
	c.Set("", caller)
}

func parseConfig(cfg []byte, value interface{}) error {
	var err error
	if err = toml.Unmarshal(cfg, value); err != nil {
		return err
	}
	return nil
}

func provider(cfg CallerCfg) (store redis.Store,err error) {
	fmt.Println(cfg)
	store, err = redis.NewStore(cfg.Size, cfg.Network, cfg.Addr, cfg.Pwd, []byte(cfg.Keypairs))
	return
}
