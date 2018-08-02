package go_config

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrNotAStructPtr = errors.New("expects pointer to a struct")
)

type Config interface {
	Use(sources ...Source)
	Configure(v interface{}) error
	Usage() map[string]string
}

type config struct {
	sources   []Source
	variables map[string]*Variable
}

func New() Config {
	return &config{
		sources:   []Source{},
		variables: make(map[string]*Variable),
	}
}

func (self *config) Use(sources ...Source) {
	self.sources = append(self.sources, sources...)
}

func (self *config) Usage() map[string]string {
	usage := make(map[string]string)
	for k, v := range self.variables {
		if v.Description == "-" {
			continue
		}
		usage[strings.ToLower(k)] = v.Description
	}
	return usage
}

func (self *config) Configure(v interface{}) error {
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return ErrNotAStructPtr
	}
	ref := ptr.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotAStructPtr
	}

	self.setup(v, "")

	for _, src := range self.sources {
		err := src.Init(self.variables)
		if err != nil {
			return err
		}
	}

	return self.fillData()
}

func (self *config) setup(v interface{}, parent string) error {
	refVal := reflect.ValueOf(v)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	if refVal.Kind() != reflect.Struct {
		return nil
	}

	refType := reflect.TypeOf(refVal.Interface())

	for i := 0; i < refVal.NumField(); i++ {
		field := refType.Field(i)
		refField := refVal.Field(i)

		name := field.Name
		tagName, _ := parseTag(field.Tag.Get("cfg"))
		if len(tagName) > 0 {
			name = tagName
		}
		if len(parent) > 0 {
			name = parent + "." + name
		}

		if refField.Kind() == reflect.Ptr {
			if refField.IsNil() {
				refField = reflect.New(refField.Type().Elem())
				refVal.Field(i).Set(refField)
				refField = refField.Elem()
			} else {
				refField = refField.Elem()
			}
		}

		if refField.Kind() == reflect.Struct {
			self.setup(refField.Addr().Interface(), name)
			self.variables[name] = &Variable{
				Name:        name,
				Description: field.Tag.Get("description"),
				Tag:         field.Tag,
				Type:        refField.Type(),
			}
			continue
		}

		if !refField.CanSet() {
			continue
		}

		tagDefault := parseDefault(field.Tag.Get("default"))
		defVal := reflect.Zero(refField.Type())
		if len(tagDefault) > 0 {
			switch refField.Kind() {
			case reflect.Int:
				i64, err := strconv.ParseInt(tagDefault, 10, 64)
				i := int(i64)
				if err == nil {
					defVal = reflect.ValueOf(i)
				}
			case reflect.Uint:
				ui64, err := strconv.ParseUint(tagDefault, 10, 64)
				ui := uint(ui64)
				if err == nil {
					defVal = reflect.ValueOf(ui)
				}
			case reflect.Float64:
				f, err := strconv.ParseFloat(tagDefault, 64)
				if err == nil {
					defVal = reflect.ValueOf(f)
				}
			case reflect.Bool:
				b, err := strconv.ParseBool(tagDefault)
				if err == nil {
					defVal = reflect.ValueOf(b)
				}
			case reflect.String:
				defVal = reflect.ValueOf(tagDefault)
			case reflect.Slice:
				// void: no default values for slices
			}
		}
		self.variables[name] = &Variable{
			Name:        name,
			Description: field.Tag.Get("description"),
			Def:         defVal,
			Set:         refField.Set,
			Tag:         field.Tag,
			Type:        refField.Type(),
		}

	}
	return nil
}

func (self *config) fillData() error {
	for _, val := range self.variables {
		changed := false

		for _, src := range self.sources {

			switch val.Type.Kind() {
			case reflect.Int:
				s, err := src.Int(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}
				if s == val.Def.Interface().(int) {
					continue
				}

				val.set(s)

			case reflect.Uint:
				s, err := src.UInt(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}
				if s == val.Def.Interface().(uint) {
					continue
				}

				val.set(s)

			case reflect.Float64:
				s, err := src.Float(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}
				if s == val.Def.Interface().(float64) {
					continue
				}

				val.set(s)

			case reflect.String:
				s, err := src.String(val.Name)
				if err != nil {
					continue
				}
				if s == "" || s == val.Def.Interface().(string) {
					continue
				}

				val.set(s)

			case reflect.Bool:
				s, err := src.Bool(val.Name)
				if err != nil {
					continue
				}
				if s == val.Def.Interface().(bool) {
					continue
				}

				val.set(s)

			case reflect.Slice:
				delimiter := val.Tag.Get("delimiter")
				if len(delimiter) == 0 {
					delimiter = ","
				}
				s, err := src.Slice(val.Name, delimiter, val.Type.Elem().Kind())

				if err != nil {
					continue
				}
				if len(s) == 0 {
					continue
				}
				// todo default values: check that slice does not equal to default slice

				val.set(s)
			}

			changed = true
		}
		if !changed {
			val.set(val.Def.Interface())
		}
	}
	return nil
}

func parseTag(tag string) (string, []string) {
	opts := strings.Split(tag, ",")
	return opts[0], opts[1:]
}

func parseDefault(tag string) string {
	return tag
}
