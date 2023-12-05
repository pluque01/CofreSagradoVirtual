package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/etcd/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/pluque01/CofreSagradoVirtual/internal/projectpath"
	flag "github.com/spf13/pflag"
)

type Config struct {
	LogFolder    string `koanf:"log_folder"`
	EtcdEndpoint string `koanf:"etcd_endpoint"`
	EtcdTimeout  int    `koanf:"etcd_timeout"`
}

var DefaultConfig *Config

func init() {
	c, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	DefaultConfig = c
}

func NewConfig() (*Config, error) {
	k := koanf.New(".")

	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.String("log_folder", projectpath.Root+"/logs", "define the path for the log folder")
	f.String("etcd_endpoint", "http://localhost:2379", "define the endpoint for the etcd server")
	f.Int("etcd_timeout", 5000, "define the timeout for the etcd server")

	f.Parse(os.Args[1:])

	if isFlagPassed("etcd_endpoint", f) {
		endpoint, _ := f.GetString("etcd_endpoint")
		var timeout int
		if isFlagPassed("etcd_timeout", f) {
			timeout, _ = f.GetInt("etc_timeout")
		}
		// Load configuration from etcd3
		etcdProvider, err := etcd.Provider(etcd.Config{
			Endpoints: []string{endpoint},

			DialTimeout: time.Duration(timeout) * time.Millisecond,

			// Key only readable from env var
			Key: os.Getenv("CSV_ETCD3_ACCESS_KEY"),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to load etcd provider: %w", err)
		}

		if err := k.Load(etcdProvider, nil); err != nil {
			fmt.Println("Error loading etcd config")
		}
	}

	// Load configuration from .env file
	// check if .env file exists
	if _, err := os.Stat(projectpath.Root + "/.env"); os.IsNotExist(err) {
		fmt.Println(".env file not found")
	} else {
		if err := k.Load(file.Provider(projectpath.Root+"/.env"), dotenv.Parser()); err != nil {
			return nil, fmt.Errorf("failed to load configuration from .env file: %w", err)
		}
	}

	// Load configuration from environment variables. The keys must be prefixed with CSV_ and
	// use _ as the separator
	k.Load(env.Provider("CSV_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "CSV_"))
	}), nil)

	// Load configuration from flags, this will override any values set before if passed as flags.
	// Non set values will get the default values specified in the flags definition.
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Extract configuration from koanf and save into Config struct
	config := &Config{}
	if err := k.Unmarshal("", config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return config, nil
}

func isFlagPassed(name string, f *flag.FlagSet) bool {
	found := false
	f.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
