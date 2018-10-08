package gorm

import (
	"github.com/godefault/caller/common"
)

type Cfg struct {
	Gorm map[string]GormCfg
}

type GormCfg struct {
	Debug bool

	Network      string
	Dialect      string
	Addr         string
	Username     string
	Password     string
	Db           string
	Charset      string
	ParseTime    string
	Loc          string
	Timeout      common.Duration
	ReadTimeout  common.Duration
	WriteTimeout common.Duration

	Level           string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime common.Duration
}
