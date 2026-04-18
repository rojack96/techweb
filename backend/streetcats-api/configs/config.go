package configs

import (
	// golang import
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	// project import
	// external import
	"gopkg.in/yaml.v3"
)

func Config() (ConfigModel, error) {
	// retrieve file config.yml from source tree first
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configPath := filepath.Join(dir, "config.yml")

	if _, err := os.Stat(configPath); err != nil {
		// if the source path does not exist in the runtime container, fall back to executable-relative config path
		execPath, execErr := os.Executable()
		if execErr == nil {
			execDir := filepath.Dir(execPath)
			altPath := filepath.Join(execDir, "configs", "config.yml")
			if _, altErr := os.Stat(altPath); altErr == nil {
				configPath = altPath
			}
		}
	}

	config, err := defaultConfig(configPath)
	if err != nil {
		return ConfigModel{}, fmt.Errorf("error loading config: %w", err)
	}

	return config, nil
}

func defaultConfig(path string) (ConfigModel, error) {
	var (
		cfg  ConfigModel
		data []byte
		err  error
	)

	if data, err = os.ReadFile(path); err != nil {
		return ConfigModel{}, fmt.Errorf("error reading config file: %w", err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return ConfigModel{}, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	return cfg, nil
}
