package main

import (
	"github.com/k0kubun/pp/v3"
	"github.com/pelletier/go-toml"
)

type Postgres struct {
	User     string
	Password string
}
type Config struct {
	Postgres []Postgres
}

func main() {
	doc := []byte(`
[[Postgres]]
User = "pelletier"
Password = "mypassword"

[[Postgres]]
User = "pelletier"
Password = "mypassword"`)

	config := Config{}
	toml.Unmarshal(doc, &config)

	pp.Println(config)
}
