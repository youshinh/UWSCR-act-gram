# **UWSCR AIネイティブRPA 開発エージェント指示書 (agent.md)**

## **1. プロジェクト概要**
本プロジェクトは、UWSC互換のスクリプトエンジンである「UWSCR」の仮想拡張ラッパー（サイドカー）を開発します。
バックエンドを **Go (Wails)**、フロントエンドを **Svelte (TypeScript)** で構築し、Windows上で動作する軽量かつ高機能な自律型RPAエージェントプラットフォームを実現します。

---

## **2. 開発エージェントのミッション**
1. **評価駆動開発 (EDD: Evaluation Driven Development) の徹底**
   * 仕様書（`docs/AIネイティブRPA 統合アーキテクチャ設計書 V4.md`）の要件を完全に満たす実装を行います。
   * Vibes（感覚）による修正は行わず、仕様ベースの評価基準を定義し、検証を行います。
2. **検証とE2Eテスト**
   * Wailsの開発モード（`http://localhost:3000/`）またはビルドされたアプリで動作確認を行います。
3. **日本語でのコミュニケーション**
   * 報告や説明はすべて日本語で行います。

---

## **3. 技術スタック**
* **Language (Backend)**: Go (1.26+)
* **Language (Frontend)**: Svelte, TypeScript, HTML/CSS (Vanilla CSSを基本とし、クリーンで美しいモダンUIを構築)
* **Framework**: Wails v2 (Go & WebUIブリッジ)
* **API Key Management**: Windows Credential Manager (`github.com/zalando/go-keyring`)
* **UWSCR Integration**: 本家UWSCR本体（Rust）とはプロセス起動・標準入出力等を介して疎結合で連携

---

## **4. 開発・設計の3大重要要件**

### **① Wailsバインディングの厳格な遵守**
Svelte（フロントエンド）からGo（バックエンド）の関数を呼び出す際、`fetch` や `axios` などのHTTPクライアントを**絶対に使用しないでください**。
必ず、Wailsのビルド・開発時に自動生成される `frontend/wailsjs/go/...` 配下のバインディング用JavaScript/TypeScriptモジュールをインポートして呼び出すこと。

### **② LLMレスポンス構造の抽象化**
各LLM（Gemini, Claude, Ollama等）はAPIレスポンスのJSON構造が異なります。依存ライブラリを増やさず `net/http` と `encoding/json` を使った直叩きHTTPクライアントを実装するため、バックエンド側でレスポンスを共通の抽象化構造体にパースしてください。
* 共通構造体の例: `LLMResponse { Text string, Error error }`

### **③ 非同期（Goroutine）によるメインスレッド保護**
`uwscr.exe` の実行制御や、バックグラウンドでのローカルAPIサーバー（デフォルトポート `31415`）の待受など、I/Oや重い処理を行う際は、必ず **Goroutine** を用いて非同期で実行してください。WailsのUIメインスレッド（UIの描画等）をブロックしてフリーズさせないこと。

---

## **5. 開発プロセス & 評価基準**
1. **仕様定義 (Observe & Red)**
   * 機能追加や修正を行う前に、`docs/` 配下の仕様書から Eval Criteria（評価基準）を定義します。
2. **進捗管理 (task.md の保守)**
   * 実装するタスクは一気にやらず、`task.md` で管理して1機能（コンポーネント）ずつ実装・テストを完了させてから次に進むこと。
3. **実装と適合 (Green)**
   * WailsブリッジおよびGo/Svelteのコードを実装します。
   * ローカル環境（`http://localhost:3000/` またはアプリウィンドウ）で動作させ、挙動が評価基準をクリアしているか検証します。
4. **リファクタ (Refactor)**
   * パフォーマンス最適化、セキュアコードの徹底、不要なコードのクリーンアップを行います。
