package main

import (
	cfg "github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/env"
	"github.com/cheebo/go-config/sources/flag"
	"github.com/cheebo/go-config/types"
	"github.com/davecgh/go-spew/spew"
)

type Master struct {
	AMQP *types.AMQPConfig `consul:"amqp"`
}

type Config struct {
	Master Master `description:"-"`

	Name string `description:"user's name'"`
	Pass string `cfg:"password" description:"user's password'"`

	GasPeerTx float64 `default:"10.11" description:"gas per transaction"`

	Timeout        uint `default:"101" description:"transaction timeout"`
	PricePerAction int  `default:"price per action"`

	AllowRegistration bool `default:"true" default:"allow new user registration"`

	Ips []string `description:"-"`
}

func main() {
	c := Config{}
	cfgr := cfg.New()
	eSrc := env.Source()
	cfgr.Use(eSrc)
	cfgr.Use(flag.Source())
	//cfgr.Use(cfg.ConsulSource("/example/config", types.ConsulConfig{
	//	Addr:   "localhost:8500",
	//	Scheme: "http",
	//}))
	cfgr.Configure(&c)

	//spew.Dump(c)
	//spew.Dump(cfgr.Usage())
	spew.Dump(eSrc.Export())
}
