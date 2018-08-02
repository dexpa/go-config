package types

type RedisConfig struct {
	Addr          string `json:"addr"`
	Password      string `json:"password"`
	Database      uint   `json:"database"`
	PoolSize      uint   `json:"pool_size"`
	MasterName    string `json:"master_name"`
	SlaveReadOnly bool   `json:"slave_read_only"`
}
