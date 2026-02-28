package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CommonEtcd       discov.EtcdConf `json:",optional"`
	UserEtcd         discov.EtcdConf `json:",optional"`
	AuthEtcd         discov.EtcdConf `json:",optional"`
	SystemEtcd       discov.EtcdConf `json:",optional"`
	ApiLogEtcd       discov.EtcdConf `json:",optional"`
	ConfigCenterEtcd discov.EtcdConf `json:",optional"`
	LogExcludePaths  []string        `json:",optional"`
	Auth             struct {
		AccessSecret string `json:",optional"`
		AccessExpire int64  `json:",optional"`
	}
	Static struct {
		Dir    string
		Prefix string
		Index  string
	}
}
