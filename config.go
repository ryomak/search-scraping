package search

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type Word struct {
	Name string
}

type Config struct {
	MaxNum int
	MyURL  string
  All bool
	Word   []Word
}

var configure Config

func LoadConfig(filePath string) *Config {
	if _, err := toml.DecodeFile(filePath, &configure); err != nil {
		panic(err)
	}
	if configure.MaxNum == 0 {
		panic(errors.New("max Num undefined"))
	}
	if len(configure.Word) == 0 {
		panic(errors.New("Word undefined"))
	}
	if configure.MyURL == "" {
		panic(errors.New("MyURL undefined"))
	}
  return &configure
}
