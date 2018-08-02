package types

type ConsulTlsConfig struct {
	CAFile   string `json:"cafile"`
	CAPath   string `json:"capath"`
	CertFile string `json:"certfile"`
	KeyFile  string `json:"keyfile"`
}

type ConsulConfig struct {
	Addr       string          `json:"addr" default:"localhost:8500"`
	Datacenter string          `json:"dc" default:"dc1"`
	Token      string          `json:"token"`
	Scheme     string          `json:"scheme" default:"http"`
	Tls        ConsulTlsConfig `json:"tls"`
}
