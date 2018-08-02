# Go Config

Package provides routines that loads configuration into provided structure from provided sources.


Example

```go
type Config struct {
	Redis *types.RedisConfig

	Name  string `description:"user's name'"`
	Pass  string `cfg:"password" description:"user's password'"`

	GasPeerTx           float64

	Timeout             uint
	PricePerAction      int

	AllowRegistration   bool
}

func main() {
	c := Config{}
	cfgr := configurer.New()
	
	// parse env variables
	cfgr.Use(env.Source())
	
	// parse flags
	cfgr.Use(flag.Source())
	
	// parse json file
	cfgr.Use(json.FileSource("./config/configurations.json"))
	
	// parse consul KV values
	cfgr.Use(consul.Source("/project/name"))
	
	cfgr.Configure(&c)
}
```

Supported field tags:
- cfg:"param_name"
- description:"variable description"
- default:"default_value"
- consul:"/kv/path/to/json/config"

Supported sources:
- env variables
- flags
- json file
- consul (you can define relative path in structure's tag and basepath in consul source creation)