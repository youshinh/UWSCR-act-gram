# **UWSCR::act-gram バックエンド・LLMルーター実装仕様書 (Go)**

## **1\. 概要**

本モジュールは、Wailsのバックエンド（Go）として稼働し、APIキーの安全な管理と、マルチLLM（Gemini, Anthropic, Ollama等）へのルーティング、およびモデル一覧の動的取得を行う。

AI（ローカルLLM）は、この設計に準拠してGoの構造体、インターフェース、およびAPI通信処理を実装すること。

## **2\. セキュア・クレデンシャル管理**

* パッケージ: github.com/zalando/go-keyring を使用。  
* 要件: APIキーは平文の設定ファイル (config.yaml) には絶対に書き込まず、OSの資格情報マネージャー（Windows Credential Manager）に保存・読み出しを行うラップ関数を実装する。  
  * SetAPIKey(provider string, key string) error  
  * GetAPIKey(provider string) (string, error)

## **3\. LLM Router (Factory Pattern)**

異なるLLMプロバイダーへのリクエストを抽象化するインターフェースを定義すること。

type LLMProvider interface {  
    GetAvailableModels() (\[\]string, error)  
    GenerateText(prompt string, options LLMOptions) (string, error)  
}

* **実装対象プロバイダー:**  
  1. GeminiProvider: GET /v1beta/models を叩いてリストを取得。  
  2. OllamaProvider: GET http://localhost:11434/api/tags を叩いてローカルのリストを取得。  
  3. AnthropicProvider: 現状は固定リスト（claude-3-7-sonnet-20250219等）を返す実装とする。

## **4\. 動的モデル解決戦略 (Latest Strategy)**

ルーティングの際、ユーザーが latest エイリアスを選択した場合はそのままプロバイダーへ投げ、固定バージョンを選択した場合は厳格にそのバージョンで推論を行うロジックを担保すること。