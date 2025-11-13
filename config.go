package aup_config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type service struct {
	Endpoint struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"Endpoint"`
	LogLevel string `json:"LogLevel"`
}

type WatcherService struct {
	service      `json:",inline"`
	RetryTimeout int `json:"RetryTimeout"`
}

type Config struct {
	Services struct {
		Watcher    WatcherService `json:"WatcherService"`
		Validator  service        `json:"ValidatorService"`
		DbWriter   service        `json:"DbWriterService"`
		SMTHandler service        `json:"SMTHandler"`
	} `json:"GrpcServer"`
	Directories struct {
		LogDir []string `json:"LogDir"`
		AupDir []string `json:"AupDir"`
	} `json:"Directories"`
	WorkDir string
}

func NewConfig() (*Config, error) {
	executable, err := os.Executable()
	workDir := filepath.Dir(executable)
	data, err := os.ReadFile(filepath.Join(workDir, "..", "sgs.json"))
	if err != nil {
		return nil, errors.New("ошибка открытия конфигурационного файла")
	}

	var config Config
	config.WorkDir = filepath.Join(workDir, "..")

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.New("ошибка преобразования конфигурационного файла из json")
	}

	return &config, nil
}

func (c *Config) GetLogDir() string {
	logDir := filepath.Join(c.WorkDir, c.Directories.LogDir[2], "/")
	return logDir
}

func (c *Config) GetArcInDir() string {
	arcInDir := filepath.Join(c.WorkDir, c.Directories.AupDir[2], c.Directories.AupDir[3], "/")
	return arcInDir
}

func (c *Config) GetArcOutDir() string {
	arcOutDir := filepath.Join(c.WorkDir, c.Directories.AupDir[2], "Out", "", "/")
	return arcOutDir
}

func (c *Config) GetOkDir() string {
	okDir := filepath.Join(c.WorkDir, c.Directories.AupDir[2], "Out", "OK", "/")
	return okDir
}

func (c *Config) GetOwDir() string {
	owDir := filepath.Join(c.WorkDir, c.Directories.AupDir[2], "Out", "OW", "/")
	return owDir
}

func (c *Config) GetRespDir() string {
	respDir := filepath.Join(c.WorkDir, c.Directories.AupDir[2], "Resp", "/")
	return respDir
}
