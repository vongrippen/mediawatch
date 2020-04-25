package lib

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Dirs []WatchDir `yaml:"dirs"`
}

type WatchDir struct {
	ContentType string `yaml:"content_type"`
	Path        string `yaml:"path"`
}

// GetWatchDirs - Get the directories to watch
func GetWatchDirs() []WatchDir {
	cfgPath := "./watch.yml"
	err := validateConfigPath(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := newConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	return cfg.Dirs
}

// NewConfig returns a new decoded Config struct
func newConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
