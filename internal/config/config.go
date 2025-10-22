package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// make the path to the config file
	path := filepath.Join(homeDir,configFileName)
	return path, nil
}

func Read() (Config, error) {
	// get the path to the file 
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	// close the file at the end
	defer file.Close()
	// make an empty config
	var cfg Config
	decoder := json.NewDecoder()
	// populate the config struct
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil 
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}