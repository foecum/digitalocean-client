package config

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fatih/camelcase"
	"github.com/hashicorp/hcl"
	yaml "gopkg.in/yaml.v2"
)

// ErrCfgUnsupported ...
var ErrCfgUnsupported = errors.New("config file format not supported. Supported formats are json, xml, yaml, toml, hcl")

// Properties for the application
type Config struct {
	Name        string `json:"name" xml:"name" yaml:"name" toml:"name" hcl:"name"`
	AccessToken string `json:"accessToken" xml:"accessToken" yaml:"accessToken" toml:"accessToken" hcl:"accessToken"`
}

// ReadConfig reads the config from the file provided
func (c *Config) ReadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	switch filepath.Ext(path) {
	case ".json":
		err := json.Unmarshal(data, c)
		if err != nil {
			return err
		}
	case ".xml":
		err := xml.Unmarshal(data, c)
		if err != nil {
			return err
		}
	case ".yml":
		err := yaml.Unmarshal(data, c)
		if err != nil {
			return err
		}
	case ".toml":
		err := toml.Unmarshal(data, c)
		if err != nil {
			return err
		}
	case ".hcl":
		obj, err := hcl.Parse(string(data))
		if err != nil {
			return err
		}

		if err = hcl.DecodeObject(c, obj); err != nil {
			return err
		}
	default:
		return ErrCfgUnsupported
	}

	return nil
}

func getEnvName(field string) string {
	camSplit := camelcase.Split(field)
	rst := strings.Join(camSplit, "_")
	return strings.ToUpper(rst)
}

// UseCustomEnvConfig updates configs with user environment configs
func (c *Config) UseCustomEnvConfig() error {
	cfg := reflect.ValueOf(c).Elem()
	cTyp := cfg.Type()

	for i := range make([]struct{}, cTyp.NumField()) {
		field := cTyp.Field(i)

		cm := getEnvName(field.Name)
		env := os.Getenv(strings.ToUpper(cm))

		if env == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			cfg.FieldByName(field.Name).SetString(env)
		}
	}
	return nil
}
