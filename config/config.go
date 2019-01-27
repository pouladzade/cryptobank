package config

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
)

const Config_File string = "config.toml"

type RPCConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	Type string `toml:"type"`
}

func DefaultRPCConfig() *RPCConfig {
	return &RPCConfig{
		Host: "127.0.0.1",
		Port: "1362",
		Type: "tcp",
	}
}

type Config struct {
	RPC *RPCConfig `toml:"rpc"`
}

func (conf *Config) ToTOML() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	err := encoder.Encode(conf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (conf *Config) SaveToFile(file string) error {
	toml, err := conf.ToTOML()
	if err != nil {
		return err
	}
	if err := WriteFile(file, toml); err != nil {
		return err
	}

	return nil
}

func DefaultConfig() *Config {
	return &Config{
		RPC: DefaultRPCConfig(),
	}
}

func FromTOML(t string) (*Config, error) {
	conf := DefaultConfig()

	if _, err := toml.Decode(t, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func LoadFromFile(file string) (*Config, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return FromTOML(string(dat))
}

func WriteFile(filename string, data []byte) error {
	if err := Mkdir(filepath.Dir(filename)); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, data, 0777); err != nil {
		return fmt.Errorf("Failed to write to %s: %v", filename, err)
	}
	return nil
}

func Mkdir(dir string) error {
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("Could not create directory %s", dir)
	}
	return nil
}

func LoadCreateConfig() *Config {
	conf, err := LoadFromFile(Config_File)

	// if there is no config file, Load the default config and create the default config file
	if err != nil {
		fmt.Println(err)
		conf = DefaultConfig()
		conf.SaveToFile(Config_File)
	}
	return conf
}
