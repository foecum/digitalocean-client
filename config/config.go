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
type Properties struct {
	Name        string `json:"name" xml:"name" yaml:"name" toml:"name" hcl:"name"`
	AccessToken string `json:"accessToken" xml:"accessToken" yaml:"accessToken" toml:"accessToken" hcl:"accessToken"`
}

// ReadConfig reads the config from the file provided
func ReadConfig(path string) (*Properties, error) {
	cfg := new(Properties)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(path) {
	case ".json":
		err := json.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	case ".xml":
		err := xml.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	case ".yml":
		err := yaml.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	case ".toml":
		err := toml.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	case ".hcl":
		obj, err := hcl.Parse(string(data))
		if err != nil {
			return nil, err
		}

		if err = hcl.DecodeObject(&cfg, obj); err != nil {
			return nil, err
		}
	default:
		return nil, ErrCfgUnsupported
	}

	return cfg, nil
}

func getEnvName(field string) string {
	camSplit := camelcase.Split(field)
	rst := strings.Join(camSplit, "_")
	return strings.ToUpper(rst)
}

// UseCustomEnvConfig updates configs with user environment configs
func (c *Properties) UseCustomEnvConfig() error {
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
