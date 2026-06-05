package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LayerConfig は各レイヤー（Brain, Eye, Utility）のプロバイダーとモデルの設定を保持します。
type LayerConfig struct {
	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
}

// Config はアプリケーション全体の設定を保持します。
type Config struct {
	Port      int                    `yaml:"port"`
	UWSCRPath string                 `yaml:"uwscr_path"`
	Layers    map[string]LayerConfig `yaml:"layers"`
}

const (
	DefaultPort = 31415
	ConfigName  = "config.yaml"
)

// GetConfigPath は設定ファイルの保存先パスを返します。
// 実行ファイルのディレクトリに保存するようにします。
func GetConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ConfigName
	}
	return filepath.Join(filepath.Dir(exePath), ConfigName)
}

// LoadConfig は設定ファイルから設定を読み込みます。ファイルがない場合はデフォルト値を返します。
func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()

	// デフォルト値
	config := &Config{
		Port:      DefaultPort,
		UWSCRPath: "",
		Layers: map[string]LayerConfig{
			"brain":   {Provider: "anthropic", Model: "claude-3-7-sonnet-20250219"},
			"eye":     {Provider: "google", Model: "gemini-2.5-flash"},
			"utility": {Provider: "google", Model: "gemini-2.5-flash-lite"},
		},
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// ファイルが存在しない場合は初期デフォルト設定で保存して返す
			err = SaveConfig(config)
			return config, err
		}
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	// ポートが未指定か不正な場合はデフォルトに補正
	if config.Port <= 0 {
		config.Port = DefaultPort
	}
	if config.Layers == nil {
		config.Layers = make(map[string]LayerConfig)
	}

	return config, nil
}

// SaveConfig は設定ファイルを設定オブジェクトに基づいて保存します。
func SaveConfig(config *Config) error {
	configPath := GetConfigPath()

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
