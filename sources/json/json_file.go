package json

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/utils"
	"io/ioutil"
	"reflect"
	"strings"
)

type jsonFile struct {
	path   string
	data   map[string]interface{}
	values map[string]*go_config.Variable
}

func FileSource(path string) go_config.Source {
	return &jsonFile{
		path: path,
		data: map[string]interface{}{},
	}
}

func (self *jsonFile) Init(vals map[string]*go_config.Variable) error {
	self.values = vals
	data, err := ioutil.ReadFile(self.path)
	if err != nil {
		return err
	}
	self.data, err = utils.JsonParse(data)
	return err
}

func (self *jsonFile) Int(name string) (int, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case int:
		return val.(int), nil
	case float64:
		return int(val.(float64)), nil
	}
	return 0, nil
}

func (self *jsonFile) Float(name string) (float64, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case float64:
		return val.(float64), nil
	}
	return 0, nil
}

func (self *jsonFile) UInt(name string) (uint, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case uint:
		return val.(uint), nil
	case float64:
		return uint(val.(float64)), nil
	}
	return 0, nil
}

func (self *jsonFile) String(name string) (string, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return "", nil
	}
	switch val.(type) {
	case string:
		return val.(string), nil
	}
	return "", nil
}

func (self *jsonFile) Bool(name string) (bool, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return false, nil
	}
	switch val.(type) {
	case bool:
		return val.(bool), nil
	}
	return false, nil
}

func (self *jsonFile) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return []interface{}{}, nil
	}
	switch val.(type) {
	case string:
		src := val.(string)
		return utils.ParseSlice(src, delimiter, kind)
	default:
		return []interface{}{}, nil
	}
}

func (self *jsonFile) Export(opt ...go_config.SourceOpt) ([]byte, error) {

	return []byte{}, nil
}

func (self *jsonFile) name(name string) string {
	return strings.ToLower(name)

}
