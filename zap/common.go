package zap

type Cfg struct {
	CallerZap map[string]ZapCfg
}

type ZapCfg struct {
	Debug bool
	Level string
	Path  string
}
