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
	Port            int                    `yaml:"port"`
	UWSCRPath       string                 `yaml:"uwscr_path"`
	KnowledgeDir    string                 `yaml:"knowledge_dir"`
	UWSCRDocURL     string                 `yaml:"uwscr_doc_url"`
	CustomBaseURL   string                 `yaml:"custom_base_url"`
	LocalLLMType    string                 `yaml:"local_llm_type"`
	LocalLLMURL     string                 `yaml:"local_llm_url"`
	TestTimeout     int                    `yaml:"test_timeout"`
	Layers          map[string]LayerConfig `yaml:"layers"`
	UseUnifiedModel bool                   `yaml:"use_unified_model"`
}

const (
	DefaultPort = 31415
	ConfigName  = "config.yaml"
)

// GetConfigPath は設定ファイルの保存先パスを返します。
// 実行ファイルのディレクトリに保存するようにします。
// ... (後続の処理)
func GetConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ConfigName
	}
	return filepath.Join(filepath.Dir(exePath), ConfigName)
}

// LoadConfig は設定ファイルから設定を読み込みます。ファイルがない場合はデフォルト値を返します。
// 返却される設定オブジェクトに空の部分があれば自動補正します。
func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()

	// デフォルト値
	config := &Config{
		Port:            DefaultPort,
		UWSCRPath:       "",
		UWSCRDocURL:     "https://stuncloud.github.io/UWSCR/",
		UseUnifiedModel: true,
		LocalLLMType:    "ollama",
		LocalLLMURL:     "",
		TestTimeout:     60,
		Layers: map[string]LayerConfig{
			"brain":   {Provider: "google", Model: "gemini-flash-lite-latest"},
			"eye":     {Provider: "google", Model: "gemini-flash-lite-latest"},
			"utility": {Provider: "google", Model: "gemini-flash-lite-latest"},
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
		// 設定ファイルが破損している場合は削除してデフォルトで初期化・保存する
		_ = os.Remove(configPath)
		_ = SaveConfig(config)
		return config, nil
	}

	// ポートが未指定か不正な場合はデフォルトに補正
	if config.Port <= 0 {
		config.Port = DefaultPort
	}
	if config.Layers == nil {
		config.Layers = make(map[string]LayerConfig)
	}
	if config.KnowledgeDir == "" {
		config.KnowledgeDir = filepath.Join(filepath.Dir(configPath), "knowledge")
	}
	if config.UWSCRDocURL == "" {
		config.UWSCRDocURL = "https://stuncloud.github.io/UWSCR/"
	}
	if config.CustomBaseURL == "" {
		config.CustomBaseURL = "http://localhost:8080/v1"
	}
	if config.LocalLLMType == "" {
		config.LocalLLMType = "ollama"
	}
	if config.TestTimeout <= 0 {
		config.TestTimeout = 30 // 未指定やマイナス値なら30秒に自動補正
	}

	// "ollama" プロバイダーを "local" プロバイダーに自動移行
	for k, v := range config.Layers {
		if v.Provider == "ollama" {
			config.Layers[k] = LayerConfig{
				Provider: "local",
				Model:    v.Model,
			}
		}
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
