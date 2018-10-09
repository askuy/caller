package gorm

import (
	"github.com/BurntSushi/toml"
	"github.com/godefault/caller/common"
	"github.com/jinzhu/gorm"
	"log"
)

var defaultCaller *callerStore

type callerStore struct {
	caller map[string]*GormClient
	cfg    Cfg
}

type GormClient struct {
	*gorm.DB
}

func New() common.Caller {
	defaultCaller = &callerStore{
		caller: make(map[string]*GormClient, 0),
	}
	return defaultCaller
}

func Caller(name string) *GormClient {
	return defaultCaller.caller[name]
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := parseConfig(cfg, &c.cfg); err != nil {
		return err
	}
	c.initGorm()
	return nil
}

func (c *callerStore) Get(key string) interface{} {
	return c.caller[key]
}

func (c *callerStore) Set(key string, val interface{}) {
	c.caller[key] = val.(*GormClient)
}

func (c *callerStore) initGorm() {
	for name, gormCfg := range c.cfg.CallerGorm {
		db, err := newGorm(gormCfg)
		if err != nil {
			if gormCfg.Level == "panic" {
				log.Panic("failed to connect mysql:" + ", error: " + err.Error())
			} else {
				log.Println("failed to connect mysql:" + ", error: " + err.Error())
			}
			continue
		}
		c.Set(name, db)
	}
}

func parseConfig(cfg []byte, value interface{}) error {
	var err error
	if err = toml.Unmarshal(cfg, value); err != nil {
		return err
	}
	return nil
}

func newGorm(cfg GormCfg) (resp *GormClient, err error) {
	var db *gorm.DB
	// dsn = "username:password@tcp(addr)/stt_config?charset=utf8&parseTime=True&loc=Local&readTimeout=1s&timeout=1s&writeTimeout=1s"
	db, err = gorm.Open(cfg.Dialect, cfg.Username+":"+cfg.Password+"@"+cfg.Network+"("+cfg.Addr+")/"+cfg.Db+
		"?charset="+cfg.Charset+"&parseTime="+cfg.ParseTime+"&loc="+cfg.Loc+
		"&timeout="+cfg.Timeout.Duration.String()+"&readTimeout="+cfg.ReadTimeout.Duration.String()+"&writeTimeout="+cfg.WriteTimeout.Duration.String())
	if err != nil {
		return
	}
	db.LogMode(cfg.Debug)
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	db.DB().SetConnMaxLifetime(cfg.ConnMaxLifetime.Duration)

	err = db.DB().Ping()
	if err != nil {
		return
	}
	resp = &GormClient{db}
	return
}
