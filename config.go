package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var (
	file []byte
)

func Load(path string) (err error) {
	file, err = os.ReadFile(path)
	return
}

func Bind(origin interface{}) (err error) {
	return yaml.Unmarshal(file, origin)
}
