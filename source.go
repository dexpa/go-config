package go_config

import (
	"errors"
	"reflect"
)

const (
	NoVarInitialised = "no variables initialised"
)

var (
	NoVariablesInitialised = errors.New(NoVarInitialised)
)

type SourceOpt func(interface{})

// Source: implement this interface to get configurations from sources like env, flag, file, kv-store etc
type Source interface {
	Init(variables map[string]*Variable) error
	Int(name string) (int, error)
	Float(name string) (float64, error)
	UInt(name string) (uint, error)
	String(name string) (string, error)
	Bool(name string) (bool, error)
	Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error)
	Export(...SourceOpt) ([]byte, error)
}
