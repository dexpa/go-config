package flag

import (
	"flag"
	"os"
	"strings"

	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/utils"
	"reflect"
)

type flags struct {
	fs     *flag.FlagSet
	values map[string]interface{}
}

func Source() go_config.Source {
	return &flags{
		fs:     flag.NewFlagSet("", flag.ContinueOnError),
		values: make(map[string]interface{}),
	}
}

func (self *flags) Init(vals map[string]*go_config.Variable) error {
	for name, val := range vals {
		name = self.name(name)
		switch val.Type.Kind() {
		case reflect.String:
			self.values[name] = self.fs.String(name, val.Def.Interface().(string), val.Description)

		case reflect.Slice:
			defVal := ""
			self.values[name] = self.fs.String(name, defVal, val.Description)

		case reflect.Int:
			i, ok := val.Def.Interface().(int)
			if !ok {
				continue
			}
			self.values[name] = self.fs.Int(name, i, val.Description)

		case reflect.Uint:
			i, ok := val.Def.Interface().(uint)
			if !ok {
				continue
			}
			self.values[name] = self.fs.Uint(name, i, val.Description)

		case reflect.Float64:
			i, ok := val.Def.Interface().(float64)
			if !ok {
				continue
			}
			self.values[name] = self.fs.Float64(name, i, val.Description)

		case reflect.Bool:
			self.values[name] = self.fs.Bool(name, val.Def.Interface().(bool), val.Description)
		}
	}
	var args []string
	for _, f := range os.Args[:] {
		if f[0] == '-' {
			args = append(args, f)
		}
	}
	return self.fs.Parse(args)
}

func (self *flags) Int(name string) (int, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return 0, nil
	}
	i, ok := val.(*int)
	return *i, nil
}

func (self *flags) Float(name string) (float64, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return 0, nil
	}

	return float64(*val.(*float64)), nil
}

func (self *flags) UInt(name string) (uint, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return 0, nil
	}
	return uint(*val.(*uint)), nil
}

func (self *flags) String(name string) (string, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return "", nil
	}
	return *(val.(*string)), nil
}

func (self *flags) Bool(name string) (bool, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return false, nil
	}
	b, ok := val.(*bool)
	if !ok {
		return false, nil
	}
	return *b, nil
}

func (self *flags) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	val, ok := self.values[self.name(name)]
	if !ok {
		return []interface{}{}, nil
	}

	src := *(val.(*string))

	return utils.ParseSlice(src, delimiter, kind)
}

func (self *flags) Export(opt ...go_config.SourceOpt) ([]byte, error) {

	return []byte{}, nil
}

func (self *flags) name(name string) string {
	return strings.ToLower(name)
}
