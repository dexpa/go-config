package types

type SMTPConfig struct {
	Host      string `json:"host"`
	Port      uint   `json:"port"`
	User      string `json:"username"`
	Password  string `json:"password"`
	SSL       bool   `json:"ssl"`
	LocalName string `json:"local_name"`
}
