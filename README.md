# **UWSCR::act-gram**

### **AI-Native RPA Sidecar Agent for UWSCR**

---

## **🌌 敬意の表明 (Respect & Acknowledgement)**

本プロジェクト **`UWSCR::act-gram`** は、Windows自動化における伝説的なスクリプト言語「UWSC」の精神を引き継ぎ、現代のRust言語を用いてゼロから爆速・高機能なエンジンとして構築された **[本家 UWSCR (Ultimate Windows Scripting Engine for Google Chrome / Chromium)](https://github.com/stuncloud/UWSCR)** の存在なしには成立しません。

素晴らしいコアエンジンを開発・公開し、Windowsの自動化エコシステムに計り知れない貢献をもたらしてくださっている作者の **`stuncloud`** 氏（およびコントリビューターの皆様）の卓越した技術力と情熱に、最大限の敬意と感謝を表します。

本システムは、本家UWSCRを無改造のまま子プロセスとして連携させ、AIモデルルーティングおよびセキュアな認証情報管理を提供する**仮想サイドカーエージェント（ラッパー）**です。UWSCRの可能性をさらに広げ、AIとのコラボレーションによる次世代の自動化体験（EDD: 評価駆動開発の導入）を目指しています。

---

## **🚀 特徴 (Core Features)**

1. **環境依存ゼロ（ゼロインフラ）**
   * ZIPを解凍して置くだけで即座に動作する設計。システム環境変数PATHの変更をユーザーに強要しません。
2. **マルチLLMルーター（ベンダーフリー）**
   * **Gemini**, **Claude**, **Ollama（ローカル）** へのルーティングを自前実装の軽量HTTPクライアントで高速処理。
   * 「思考（Brain）」「視覚（Eye）」「反射（Utility）」の3層で、最適なAI頭脳をGUIから自由に割り当てられます。
3. **セキュア認証キー管理**
   * APIキーは平文設定ファイルに一切保存されず、Windows資格情報マネージャー（Windows Credential Manager）のセキュア領域に暗号化保存されます。
4. **マクロ・トランスパイラ**
   * `.uws` スクリプト内の `AI_EVAL("金額は？", GetScreenCapture())` などの拡張構文を、本家UWSCRが解釈できるローカルAPIサーバー連携コードへ瞬時に事前変換します。

---

## **🛠️ 開発環境とビルド (Development & Build)**

本アプリは **Go (Wails v2)** および **Svelte + TypeScript** で構築されています。

### **必須要件**
* Go 1.26+
* Node.js LTS (v24+)
* npm (v11+)
* Wails CLI

### **開発モードでの起動**
ファイル変更時に自動でホットリロードが走り、素早いUI調整が可能です。
```bash
wails dev
```

### **プロダクションビルド**
コンパイルされ、`build/bin/act_gram.exe` に単一バイナリが出力されます。
```bash
wails build
```
