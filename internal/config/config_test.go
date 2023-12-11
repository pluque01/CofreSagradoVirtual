package config

import (
	"os"
	"testing"

	"github.com/pluque01/CofreSagradoVirtual/internal/projectpath"
)

func TestNewConfig(t *testing.T) {
	var config *Config
	var err error
	// Test loading configuration from .env file
	// Create a temporary .env file for testing
	envFile := `
		log_folder=/path/to/logs
		etcd_endpoint=http://test.etcd.endpoint:1234
		etcd_timeout=1234
	`
	err = os.WriteFile(projectpath.Root+"/.env", []byte(envFile), 0644)
	if err != nil {
		t.Fatalf("Failed to create .env file: %v", err)
	}
	defer os.Remove(".env")

	config, err = NewConfig()
	if err != nil {
		t.Fatalf("Failed to load configuration from .env file: %v", err)
	}

	// Assert the values loaded from .env file
	if config.LogFolder != "/path/to/logs" {
		t.Errorf("Expected LogFolder to be '/path/to/logs', got '%s'", config.LogFolder)
	}
	if config.EtcdEndpoint != "http://test.etcd.endpoint:1234" {
		t.Errorf("Expected EtcdEndpoint to be 'http://test.etcd.endpoint:1234', got '%s'", config.EtcdEndpoint)
	}
	if config.EtcdTimeout != 1234 {
		t.Errorf("Expected EtcdTimeout to be 1234, got '%d'", config.EtcdTimeout)
	}

	envs := []struct {
		name  string
		value string
	}{
		{"CSV_LOG_FOLDER", "/path/to/logsfromenv"},
		{"CSV_ETCD_ENDPOINT", "http://test.etcd.endpoint:1234"},
	}
	// Set up test environment
	for _, env := range envs {
		os.Setenv(env.name, env.value)
	}
	// Clean up environment variables after the test
	defer func() {
		for _, env := range envs {
			os.Unsetenv(env.name)
		}
	}()

	// Test loading configuration from environment variables
	config, err = NewConfig()
	if err != nil {
		t.Fatalf("Failed to load configuration from environment variables: %v", err)
	}

	// Assert the values loaded from environment variables
	if config.LogFolder != envs[0].value {
		t.Errorf("Expected LogFolder to be %s, got '%s'", envs[0].value, config.LogFolder)
	}
	if config.EtcdEndpoint != envs[1].value {
		t.Errorf("Expected EtcdEndpoint to be %s, got '%s'", envs[1].value, config.EtcdEndpoint)
	}

	// Test loading configuration from flags
	// Save the original command-line arguments
	oldArgs := os.Args

	// Set the command-line arguments for testing
	os.Args = []string{"test", "--log_folder", "/path/to/logsfromflags"}

	// Create a new config
	config, err = NewConfig()
	if err != nil {
		t.Errorf("Error creating config: %v", err)
	} else {
		// Check the values of the config
		if config.LogFolder != "/path/to/logsfromflags" {
			t.Errorf("Expected LogFolder to be '/path/to/logsfromflags', but got '%s'", config.LogFolder)
		}
		if config.EtcdEndpoint != "http://test.etcd.endpoint:1234" {
			t.Errorf("Expected EtcdEndpoint to be 'http://test.etcd.endpoint:1234', but got '%s'", config.EtcdEndpoint)
		}
		if config.EtcdTimeout != 1234 {
			t.Errorf("Expected EtcdTimeout to be 1234, but got %d", config.EtcdTimeout)
		}
	}

	// Restore the original command-line arguments
	os.Args = oldArgs
}
