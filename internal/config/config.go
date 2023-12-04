package config

import (
	"fmt"
	"log"
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

func NewConfig() (*Config, error) {
	k := koanf.New(".")

	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.String("log_folder", projectpath.Root+"/logs", "define the path for the log folder")
	f.String("etcd_endpoint", "http://localhost:2379", "define the endpoint for the etcd server")
	f.Int("etcd_timeout", 5000, "define the timeout for the etcd server")

	f.Parse(os.Args[1:])

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
