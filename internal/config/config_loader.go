// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	DefaultPort            = 9090
	DefaultHostname        = "localhost"
	DefaultInitialDataPath = "/etc/reading-list/initial_data.json"
)

var (
	EnvName         = "ENV"
	Hostname        = "HOSTNAME"
	Port            = "PORT"
	initialDataPath = "INIT_DATA_PATH"
)

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig() (*Config, error) {
	config = Config{
		Hostname:        getEnvString(Hostname, DefaultHostname),
		Port:            getEnvInt(Port, DefaultPort),
		Env:             os.Getenv(EnvName),
		InitialDataPath: resolveInitialDataPath(os.Getenv(initialDataPath), DefaultInitialDataPath, os.Stat),
	}
	return &config, nil
}

func LoadInitialData() (data InitialData) {
	if config.InitialDataPath == "" {
		return
	}
	contents, err := os.ReadFile(config.InitialDataPath)
	if err != nil {
		log.Fatalf("failed to read initial data at [%s]: %s", config.InitialDataPath, err)
	}
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatalf("failed to unmarshal initial data at [%s]: %s", config.InitialDataPath, err)
	}
	log.Printf("loaded initial data from [%s] with %d books", config.InitialDataPath, len(data.Books))
	return
}

func getEnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func getEnvString(key string, defaultVal string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	return s
}

func resolveInitialDataPath(envPath string, defaultPath string, stat func(string) (os.FileInfo, error)) string {
	if envPath != "" {
		return envPath
	}
	if _, err := stat(defaultPath); err == nil {
		return defaultPath
	}
	return ""
}
