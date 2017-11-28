package test

import (
	"digitalocean/config"
	"os"
	"testing"
)

type TestConfig struct {
	Name        string
	AccessToken string
	ConfigFile  string
}

func TestGoodConfig(t *testing.T) {

	envValues := []string{"NAME", "BASE_URL", "DRIVER"}

	for i := range envValues {
		os.Unsetenv(envValues[i])
	}

	testCasesGood := []TestConfig{
		{
			Name:        "Test Config HCL",
			AccessToken: "accessToken",
			ConfigFile:  "fixtures/good/test-config.hcl",
		},
		{
			Name:        "Test Config TOML",
			AccessToken: "accessToken",
			ConfigFile:  "fixtures/good/test-config.toml",
		},
		{
			Name:        "Test Config XML",
			AccessToken: "accessToken",
			ConfigFile:  "fixtures/good/test-config.xml",
		},
		{
			Name:        "Test Config YAML",
			AccessToken: "accessToken",
			ConfigFile:  "fixtures/good/test-config.yml",
		},
		{
			Name:        "Test Config JSON",
			AccessToken: "accessToken",
			ConfigFile:  "fixtures/good/test-config.json",
		},
	}

	for i := range testCasesGood {
		cfg, err := config.ReadConfig(testCasesGood[i].ConfigFile)
		if err != nil {
			t.Fatalf("could not read config file %s: %v", testCasesGood[i].ConfigFile, err)
		}

		if cfg.AccessToken != testCasesGood[i].AccessToken {
			t.Errorf("Expected %s, Got %s", testCasesGood[i].AccessToken, cfg.AccessToken)
		}

		if cfg.Name != testCasesGood[i].Name {
			t.Errorf("Expected %s, Got %s", testCasesGood[i].Name, cfg.Name)
		}
	}
}

func TestBadConfig(t *testing.T) {

	envValues := []string{"NAME", "ACCESSTOKEN", "DRIVER"}

	for i := range envValues {
		os.Unsetenv(envValues[i])
	}

	testCasesBad := []string{
		"fixtures/bad/test-config.hcl",
		"fixtures/bad/test-config.toml",
		"fixtures/bad/test-config.xml",
		"fixtures/bad/test-config.yml",
		"fixtures/bad/test-config.json",
	}

	for i := range testCasesBad {
		_, err := config.ReadConfig(testCasesBad[i])
		if err == nil {
			t.Fatalf("Expected an error")
		}
	}
}

func TestEnvironmentConfig(t *testing.T) {

	envKeys := []string{"NAME", "ACCESS_TOKEN"}
	envValues := []string{
		"Testing Environment variables",
		"accessToken",
		"",
	}

	for i := range envKeys {
		os.Setenv(envKeys[i], envValues[i])
	}

	cfg := &config.Properties{}

	err := cfg.UseCustomEnvConfig()

	if err != nil {
		t.Fatalf("could not read enviroment variable: %v", err)
	}

	if cfg.Name != envValues[0] {
		t.Errorf("Expected %s, Got %s", envValues[0], cfg.Name)
	}

	if cfg.AccessToken != envValues[1] {
		t.Errorf("Expected %s, Got %s", envValues[1], cfg.AccessToken)
	}

	for i := range envKeys {
		os.Unsetenv(envKeys[i])
	}
}

func TestNoConfigFile(t *testing.T) {

	envValues := []string{"NAME", "BASE_URL", "DRIVER"}

	for i := range envValues {
		os.Unsetenv(envValues[i])
	}

	testCasesBad := []string{
		"fixtures/bad/test-config.hcl1",
	}

	for i := range testCasesBad {
		_, err := config.ReadConfig(testCasesBad[i])
		if err == nil {
			t.Fatalf("Expected an error")
		}
	}
}

func TestInvalidConfigFileExtension(t *testing.T) {

	envValues := []string{"NAME", "BASE_URL", "DRIVER"}

	for i := range envValues {
		os.Unsetenv(envValues[i])
	}

	testCasesBad := []string{
		"fixtures/bad/test-config.hc",
	}

	for i := range testCasesBad {
		_, err := config.ReadConfig(testCasesBad[i])
		if err.Error() != config.ErrCfgUnsupported.Error() {
			t.Errorf("Expected \"%v\" but got \"%v\"", config.ErrCfgUnsupported, err)
		}
	}
}
