# **UWSCR AIネイティブRPA 統合アーキテクチャ設計書 V4**

～ 動的モデル検出（Dynamic Discovery）と環境適応・ベンダーフリー戦略 ～

## **1\. プロジェクト・ビジョン (First Principles)**

「ベンダーロックインの完全な排除」と「実行環境への適応力」、そして「運用保守のゼロ化」を備えた自律型Windows自動化プラットフォーム。

GoとWailsで記述された仮想拡張ラッパー（サイドカー）が、クラウド/ローカルのLLMを動的かつグラフィカルに管理し、非IT部門のユーザーへ安全・爆速なRPA体験を提供する。

## **2\. コア・アーキテクチャ：仮想拡張・サイドカーモデル**

本家UWSCR本体（Rust）は無改造（疎結合）とし、Go製のエージェント（agent.exe）がUI（Wails）とAIルーティングを担う。

独自関数（AI\_EVAL 等）は、実行直前に標準のUWSCRコードへトランスパイル（展開）され、ネイティブ拡張と同等のDXを提供する。

## **3\. 動的モデル検出 & ベンダーフリー戦略 (Dynamic Model Discovery)**

ユーザー（情シス部門）が設定ファイルを直接編集する際のエラーを防ぐため、Goエージェントは各プロバイダーのAPIから「利用可能な最新モデル一覧」をリアルタイムに取得（Fetch）し、WailsのUIで選択させる仕組みを実装する。

### **3.1 各プロバイダーの動的取得アプローチ**

WailsのコントロールパネルでAPIキーを入力（またはローカルエンドポイントを指定）した瞬間に、Goが以下のAPIを叩いてリストを生成する。

* **Google (Gemini)**  
  * エンドポイント: GET /v1beta/models?key={API\_KEY}  
  * 動作: 権限のある最新のGeminiモデル一覧（gemini-2.5-flash-lite, gemini-1.5-pro等）を動的にリスト化。  
* **ローカルLLM (Ollama)**  
  * エンドポイント: GET http://localhost:11434/api/tags  
  * 動作: ユーザーのローカルPCに現在ダウンロード（pull）されているモデル群（qwen2.5-coder, llama3.2-vision等）をリストアップ。  
* **OpenAI / 互換サーバー (LM Studio, vLLM 等)**  
  * エンドポイント: GET /v1/models  
* **Anthropic (Claude)**  
  * ※現状公式APIにモデル一覧取得エンドポイントがないため、ここだけはGo側で最新の定数リスト（claude-3-7-sonnet-20250219, claude-3-5-haiku-latest等）を保持しつつ、APIがアップデートされ次第取得ロジックへ切り替える。

### **3.2 レイヤー別アサインメントのUI統合**

ユーザーはWailsの設定画面（コントロールセンター）を開き、プルダウンから各層の脳を選ぶだけで config.yaml が自動生成される。

* **🧠 Brain層 (コード生成・計画) \-\> 賢さと「安定性」優先**  
  * プルダウンから claude-3-7-sonnet-20250219（企業向け安定板）などを選択。  
* **👁️ Eye層 (画面解析・要素特定) \-\> 視覚能力優先**  
  * プルダウンから gemini-2.5-flash または ローカルの llama-3.2-vision を選択。  
* **🛠️ Utility / Hand層 (データ変換・正規表現代替) \-\> 速度・コスト最優先**  
  * プルダウンから gemini-2.5-flash-lite を選択。

## **4\. セキュリティと認証情報管理 (Secure Credential Management)**

APIキーの平文保存による漏洩リスクをゼロにするため、GoエージェントはOSのネイティブ・セキュアストレージ（Windows Credential Manager）を活用して鍵を管理する。

### **4.1 実装メカニズム (Go keyring パッケージの活用)**

* **登録フェーズ:** Wails UI画面にAPIキーを入力すると、Go側で github.com/zalando/go-keyring を使用し、Windowsの「資格情報マネージャー」に暗号化保存。  
* **実行フェーズ:** API呼び出し時に、GoがOSのセキュア領域からキーを取得し、メモリ上でのみリクエストヘッダーに付与する（config.yamlにはAPIキーを記述しない）。

## **5\. アクション・ロードマップ (Implementation Steps)**

* **Phase 1: 基礎インフラとセキュア管理 (Go \+ Wails)**  
  * GoでWailsを用いた軽量ウィジェットUIを作成。  
  * go-keyring を実装し、各種APIキーの安全な登録・読み出し機構を確立。  
* **Phase 2: 動的モデルルーターの構築 (Dynamic Discovery)**  
  * 各プロバイダー（Gemini, OpenAI, Ollama）のモデル一覧取得APIを叩く関数をGoで実装。  
  * 取得したリストをWails UIのプルダウンに連携し、GUIで各レイヤーのモデルを選択・保存（config.yaml自動更新）できる仕組みを完成させる。  
* **Phase 3: トランスパイラ（仮想拡張機能）の実装**  
  * AI\_EVAL などの独自関数を、Goローカルサーバーを叩く標準UWSCRコマンドへ置換するマクロエンジンをGo内に構築。  
* **Phase 4: レコーダー統合と自律マニュアル生成**  
  * Windows OSフックによる操作ログ取得と画面キャプチャを統合し、「記録 ➔ 変換 ➔ 実行」のサイクルを確立する。