# **UWSCR::act-gram RPA連携＆トランスパイラ実装仕様書 (Go)**

## **1\. 概要**

本モジュールは、無改造の本家 uwscr.exe（Rustコア）を子プロセスとして制御し、拡張構文を標準UWSCRスクリプトへ事前変換（トランスパイル）する「仮想拡張エンジン」である。

AI（ローカルLLM）は、Goの os/exec および正規表現を用いた変換処理を実装すること。

## **2\. マクロ・トランスパイラ (Virtual Extension)**

ユーザーが記述した .uws スクリプト内に存在する独自のAI関数を、Goのローカルサーバー（ポート8080等で待機）を呼び出す標準のUWSCRコードへ置換する。

* **変換仕様:**  
  * 変換前（ユーザー記述）:  
    Dim res \= AI\_EVAL("この伝票の金額は？", GetScreenCapture())  
  * 変換後（uwscr.exeへ流し込むコード）:  
    Dim \_img \= GetScreenCapture()  
    Dim res \= POWERSHELL("curl \-s \-X POST http://127.0.0.1:8080/ai\_eval \-d ...", true)  
* **実装要件:** Goの regexp パッケージを用いてスクリプト全体をパースし、安全に文字列置換を行ってから一時ファイル（temp\_exec.uws）として出力する。

## **3\. オーケストレーター層 (子プロセス制御)**

* Goの os/exec.Command("uwscr.exe", "temp\_exec.uws") を用いてスクリプトを実行。  
* StdoutPipe と StderrPipe を接続し、UWSCRからの実行ログやエラーをリアルタイムでGo側（コントロールセンターのUI）にストリーミング表示する仕組みを実装する。

## **4\. ローカルAPIサーバー機能**

トランスパイルされたUWSCRコードからのリクエスト（上記 curl 等）を受け止めるため、Go内で軽量なHTTPサーバー（net/http）をバックグラウンド稼働させる。

* エンドポイント /ai\_eval: 質問と画像を受け取り、LLMルーター（02\_backend\_llm\_router.md参照）へ渡して推論結果をテキストで返す。