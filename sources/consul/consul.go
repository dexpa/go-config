package consul

import (
	"github.com/cheebo/consul-utils"
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/cheebo/go-config/utils"
	"github.com/hashicorp/consul/api"
	"reflect"
	"strings"
)

type consul struct {
	prefix string
	config types.ConsulConfig
	data   map[string]interface{}
	consul map[string]interface{}
	values map[string]*go_config.Variable
}

func Source(prefix string, config types.ConsulConfig) go_config.Source {
	return &consul{
		prefix: prefix,
		config: config,
		data:   make(map[string]interface{}),
		consul: make(map[string]interface{}),
	}
}

func (self *consul) Init(vals map[string]*go_config.Variable) error {
	self.values = vals

	config := &api.Config{Address: self.config.Addr, Scheme: self.config.Scheme, Token: self.config.Token}
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	for name, val := range vals {
		name = self.name(name)

		tag := val.Tag.Get("consul")
		opts := strings.Split(tag, ",")

		if len(opts[0]) > 0 {
			cc, err := consul_utils.GetKV(client, self.prefix+opts[0], consul_utils.QueryOptions{
				Datacenter: self.config.Datacenter,
				Token:      self.config.Token,
			})
			if err != nil {
				return err
			}

			switch val.Type.Kind() {
			case reflect.Struct:
				fallthrough
			case reflect.Slice:
				m, err := utils.JsonParse([]byte(cc))
				if err != nil {
					return err
				}
				for n, v := range m {
					self.data[name+"."+n] = v
				}
			default:
				self.data[name] = cc
			}
		}
	}
	return nil
}

func (self *consul) Int(name string) (int, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	return int(val.(float64)), nil
}

func (self *consul) Float(name string) (float64, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}

	return float64(val.(float64)), nil
}

func (self *consul) UInt(name string) (uint, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	return uint(val.(float64)), nil
}

func (self *consul) String(name string) (string, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return "", nil
	}
	return val.(string), nil
}

func (self *consul) Bool(name string) (bool, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return false, nil
	}
	b, ok := val.(bool)
	if !ok {
		return false, nil
	}
	return b, nil
}

func (self *consul) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (self *consul) Export(opt ...go_config.SourceOpt) ([]byte, error) {
	return []byte{}, nil
}

func (self *consul) name(name string) string {
	return strings.ToLower(name)
}
