package json_test

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/json"
	"github.com/cheebo/go-config/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

type Config struct {
	AMQP types.AMQPConfig `cfg:"amqp"`
}

func TestJsonFileSource(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	c := go_config.New()
	c.Use(json.FileSource("./fixtures/config.json"))
	c.Configure(&cfg)

	assert.Equal("localhost", cfg.AMQP.URL)
	assert.Equal("exch", cfg.AMQP.Exchange)
	assert.Equal("que", cfg.AMQP.Queue)
	assert.Equal("knd", cfg.AMQP.Kind)
	assert.Equal("k", cfg.AMQP.Key)
	assert.Equal(true, cfg.AMQP.Durable)
	assert.Equal(true, cfg.AMQP.AutoDelete)
	assert.Equal(2, int(cfg.AMQP.DeliveryMode))
}
