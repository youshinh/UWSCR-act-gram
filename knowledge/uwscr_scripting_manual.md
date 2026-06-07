# UWSCR 1.1.9 スクリプト生成・開発マニュアル (UWSCR Scripting Manual)

本マニュアルは、UWSCR (UWSC互換スクリプトエンジン) において、AIおよび開発者が構文エラーや実行時エラーを起こさず、高品質で正確な自動化スクリプトを生成・実装するための仕様ガイドラインです。

---

## 1. 開発環境と文字コードのルール

UWSCRのスクリプトファイル（`.uws`）を作成・編集する際は、以下のルールを厳守してください。

- **ファイル文字コード**: **UTF-8** （BOMなしを推奨、BOMありも対応）
  - UTF-8以外のエンコーディング（Shift_JIS等）で保存されたファイルは、日本語コメントや文字列リテラルが文字化けし、構文エラーを引き起こします。
- **大文字・小文字の区別**: 識別子、変数名、関数名、定数名において、英字の**大文字小文字は区別されません**。
  - 例: `getid` と `GETID`、`st_title` と `ST_TITLE` は同一とみなされます。可読性のため大文字統一、あるいはスネークケース等の表記法を統一してください。

---

## 2. 変数定義とデータ型

### 2.1 変数定義
変数定義には以下のキーワードを使用します。

- **ローカル変数**: `dim 変数名` または `var 変数名`
  - 代入による暗黙的な定義も可能ですが、明示的に `dim` や `var` を使用することで、バグを未然に防ぐことができます。
- **グローバル変数**: `public 変数名`
- **定数**: `const 定数名 = 値`

### 2.2 配列と連想配列
UWSCRでは、配列や連想配列の表記が非常に柔軟です。

- **配列リテラル**: `arr = [1, 2, 3]`
  - `dim arr[3]` のように宣言してから個別に代入する旧方式も動作しますが、リテラル表記 `[val1, val2, ...]` を推奨します。
- **連想配列（ハッシュ）**: `hashtbl` 宣言が**不要**になり、代入だけで初期化できます。
  - キーの大文字小文字の区別の設定、並び替えなどのオプションは、ハッシュの定義時に指定可能です。
  - 例:
    ```uwscr
    // 空の連想配列を作成
    hash_data = []  // または hash_data = hashtbl()
    hash_data["key1"] = "value1"
    
    // オプション指定付きの初期化
    hash with_option = HASH_CASECARE or HASH_SORT // キーの大文字小文字を区別し、自動ソートする
    ```

### 2.3 UObject (JSON / YAML 互換オブジェクト)
JSONやYAML形式の構造化データを、スクリプト内で直感的に扱えるUWSCR独自のデータ型です。

- **UObjectリテラル**: JSONの記述を `@` と `@` で括ります。
  ```uwscr
  obj = @{
    "foo": "fooooo",
    "bar": {
      "baz": true
    },
    "qux": [
      { "quux": 1 },
      { "quux": 2 }
    ]
  }@
  ```
- **値の取得と変更**: ドット演算子 `.` または配列の添字でアクセス可能です。
  ```uwscr
  print obj.foo           // "fooooo"
  print obj.bar.baz       // True
  obj.foo = "new_value"   // 変更可能
  print obj.qux[1].quux   // 2
  ```
- **注意点 (制限)**: UObjectのキーやメンバを**後から動的に追加することはできません**。
  - 例: `obj.new_key = 1` はエラーになります。オブジェクトを丸ごと作成して再代入することは可能です。
- **変数展開**: UObjectリテラルは変数展開が可能です。二重引用符内に `<#変数名>` を記述します。
  ```uwscr
  name = "John"
  user = @{
    "name": "<#name>"
  }@
  ```

---

## 3. 主要なビルトイン関数とシグネチャ

### 3.1 ウィンドウ制御

#### ① `GETID` (ウィンドウIDの取得)
```uwscr
id = getid(タイトル [, クラス名=EMPTY, 待ち時間=1])
// または
id = getid(定数)
```
- **パラメータ**:
  - `タイトル`: ウィンドウタイトル（部分一致）。
  - `クラス名`: ウィンドウクラス名（部分一致、省略可）。
  - `待ち時間`: ウィンドウが出現するまでの最大タイムアウト（秒、省略可）。見つからなかった場合は `-1` を返します。
- **定数指定**:
  - `GET_ACTIVE_WIN`: アクティブウィンドウ
  - `GET_FROMPOINT_WIN`: マウスカーソル下のウィンドウ
  - `GET_FROMPOINT_OBJ`: マウスカーソル下の子ウィンドウ/コントロール
  - `GET_CONSOLE_WIN`: UWSCRを実行しているコンソールウィンドウのID
  - `GET_LOGPRINT_WIN`: Printウィンドウ

#### ② `GETALLWIN` (全ウィンドウIDの取得)
```uwscr
ids = getallwin([ID=EMPTY])
```
- **注意 (UWSCとの絶対的差異)**:
  - UWSCとは異なり、戻り値は**ウィンドウIDの配列**です。見つかった個数ではありません。
  - 特殊変数 `ALL_WIN_ID` は**廃止**されました。
  - 子ウィンドウの一覧を取得したい場合は、親ウィンドウの `ID` を引数に渡します。

#### ③ `CTRLWIN` (ウィンドウ操作命令)
```uwscr
ctrlwin(ID, コマンド定数)
```
- **コマンド定数**:
  - `CLOSE` : ウィンドウを閉じる
  - `CLOSE2` : 強制終了
  - `ACTIVATE` : アクティブにする
  - `HIDE` : 非表示にする
  - `SHOW` : 表示する
  - `MIN` : 最小化
  - `MAX` : 最大化
  - `NORMAL` : 通常サイズに戻す
  - `TOPMOST` : 最前面に固定
  - `NOTOPMOST` : 最前面固定を解除

#### ④ `STATUS` (ウィンドウ状態の取得)
```uwscr
value = status(ID, ST定数)
// 複数指定時は連想配列が返ります
stat = status(ID, ST_TITLE, ST_X, ST_Y)
print stat[ST_TITLE]
```
- **【超重要】引数に指定する定数は `ST_` から始まる定数です。**
  - **絶対に `STATUS_X` や `STATUS_Y` と記述しないでください。** UWSCRには存在しない定数のため、変数未定義エラーとなります。
  - 主要な `ST_` 定数:
    - `ST_TITLE` : タイトル
    - `ST_CLASS` : クラス名
    - `ST_X` / `ST_Y` : 座標
    - `ST_WIDTH` / `ST_HEIGHT` : ウィンドウ幅・高さ
    - `ST_VISIBLE` : 可視状態（True/False）
    - `ST_ACTIVE` : アクティブ状態（True/False）
    - `ST_PATH` : 実行プログラムのフルパス
    - `ST_PROCESS` : プロセスID

---

### 3.2 画面操作とUI Automation (UIA)

#### ① `CLKITEM` (項目・ボタンのクリック)
```uwscr
clkitem(ID, アイテム名 [, CLK定数=0, チェック指定=TRUE, n番目=1])
```
- **CLK定数 (アイテム種別)**:
  - `CLK_BTN` : ボタン、チェックボックス、ラジオボタン
  - `CLK_LIST` : リストボックス、コンボボックス
  - `CLK_TAB` : タブ
  - `CLK_MENU` : メニュー
  - `CLK_TREEVIEW` : ツリービュー
  - `CLK_LISTVIEW` : リストビュー
  - `CLK_LINK` : ハイパーリンク
- **CLK定数 (クリック方式・オプション)**:
  - `CLK_API` : Win32 APIによるクリック（高速、バックグラウンド対応）
  - `CLK_ACC` : Microsoft Active Accessibility (MSAA) による操作
  - `CLK_UIA` : **UI Automation (UIA) による検索とクリック**
  - `CLK_BACK` : ウィンドウをアクティブにせずに裏で実行（バックグラウンド）
  - `CLK_SHORT` : アイテム名の部分一致で検索
- **ヒント (複数指定)**:
  - 定数は `OR` 演算子で連結して複数指定が可能です。
  - 例: `clkitem(id, "OK", CLK_BTN or CLK_UIA)`

#### ② `SENDSTR` (文字列の送信・入力)
```uwscr
sendstr(ID, 文字列 [, n番目=0, 送信モード=FALSE, ACC指定=FALSE])
```
- **注意 (UWSCとの絶対的差異)**:
  - 関数名は `SENDSSTR` ではなく **`SENDSTR`** です。
  - `ID` を `0` に指定した場合は、**クリップボードに文字列を送信**（コピー）します。
  - `送信モード`: `FALSE` / `0` (追記)、`TRUE` / `1` (置き換え)、`2` (一文字ずつ送信)。
  - `ACC指定 (クリック・検索方式)`:
    - `FALSE` / `0` : Win32 APIによる送信。該当がなければ自動で UI Automation (UIA) にフォールバックします。
    - `TRUE` / `1` : MSAA (ACC) を使用して送信。
    - `STR_UIA` (定数: 6) : **UI Automation (UIA) を明示的に使用。** (UIA時は送信モードに関わらず常に「置き換え」になります)

---

### 3.3 低レベル操作 (マウス・キーボード)

#### ① `MMV` (マウスカーソル移動)
```uwscr
mmv(x, y [, ms=0])
```
- **パラメータ**:
  - `x`, `y`: 移動先座標。デフォルトは画面のスクリーン座標(0,0が左上)。
  - `ms`: 移動する前に待機するミリ秒数。

#### ② `BTN` (マウスボタン操作の送信)
```uwscr
btn(ボタン定数 [, 状態=CLICK, x=EMPTY, y=EMPTY, ms=0])
```
- **ボタン定数**:
  - `LEFT` (左クリック)、`RIGHT` (右クリック)、`MIDDLE` (中央クリック)、`WHEEL` (ホイール回転)、`TOUCH` (タッチ・スワイプ)
- **状態定数**:
  - `CLICK` : 押し下げて離す
  - `DOWN` : 押し下げる
  - `UP` : 離す
- **タッチ操作とスワイプ**:
  - `btn(TOUCH, DOWN, x1, y1)` を呼び出した後、別座標に対して `btn(TOUCH, UP, x2, y2, ms)` を実行すると、スワイプ操作になります。

#### ③ `KBD` (キーボード入力送信)
```uwscr
kbd(仮想キーまたは文字コード [, 状態=CLICK, ms=0])
// または
kbd(送信文字列 [, 状態=CLICK, ms=0])
```
- **仮想キーの例**:
  - `VK_RETURN` (Enterキー), `VK_TAB`, `VK_SPACE`, `VK_SHIFT`, `VK_CONTROL`, `VK_ALT`, `VK_BACK` (BackSpace) 等。
- **状態定数**:
  - `CLICK` (デフォルト), `DOWN`, `UP`

#### ④ `MOUSEORG` (座標基準の設定)
```uwscr
mouseorg(ID [, 基準=MORG_WINDOW, 画面取得=MORG_FORE, HWND=FALSE])
```
- マウス移動（`mmv`）、クリック（`btn`）、画像検索（`ChkImg`）の座標基準を、画面全体から特定のウィンドウ相対座標へと切り替えます。
- `ID = 0` を指定するとスクリーン全体基準に戻ります。

---

### 3.4 画像検索

#### ① `CHKIMG`
```uwscr
result_coords = chkimg([ファイル名={clipboard}, 探索方式=0, x1, y1, x2, y2, n番目=1, 色幅=0])
```
- **【超重要】UWSCとの絶対的差異**:
  - 戻り値は**見つかった座標の配列 `[x, y]`** です。見つからなかった場合は**空の配列 `[]`** が返ります。
  - **特殊変数 `G_IMG_X`, `G_IMG_Y` は完全に廃止されました。** 記述するとコンパイル・構文エラーになります。
  - `n番目` に `-1` を指定した場合、一致するすべての箇所を二次元配列 `[[x1,y1], [x2,y2], ...]` で返します。
  - ファイル名を省略するか `"{clipboard}"` を指定すると、クリップボード上の画像から検索します。

#### ② `SearchImage` (UWSCR専用・高速画像認識)
```uwscr
results = SearchImage(画像ファイルパス [, スコア=95, 最大検索数=5, left, top, right, bottom, オプション=0, モニタ番号=0])
```
- **戻り値**: 一致した座標と類似度スコアを含む二次元配列 **`[[X座標, Y座標, スコア], ...]`**。見つからない場合は空の配列。
- **オプション**:
  - `SCHIMG_NO_GRAY`: 画像をグレースケール化せず検索。
  - `SCHIMG_USE_WGCAPI`: GraphicsCaptureAPIにより通常キャプチャできないウィンドウからキャプチャ・検索。

---

## 4. 特殊機能とCOMオブジェクト制御

### 4.1 COMオブジェクトの制御
ExcelやWordなどの外部アプリケーション、WshShellなどのWindows機能をOLE/COM経由で操作します。

- **オブジェクト生成**: `obj = createoleobj("Excel.Application")`
- **COMエラーの抑制 (`COM_ERR_IGN` と `COM_ERR_RET`)**:
  COM操作中に発生するエラーでプログラムが中断するのを防ぎます。
  ```uwscr
  COM_ERR_IGN  // エラー抑制を開始 (特殊変数 COM_ERR_FLG は FALSE に初期化される)
  
  // 例: エラーが発生する可能性のある処理
  obj.ActiveSheet.Cells(1, 1).Value = "Test"
  
  if COM_ERR_FLG then
    print "COMエラーが発生しました"
  endif
  
  COM_ERR_RET  // エラー抑制を解除
  ```

### 4.2 WMI (Windows Management Instrumentation) クエリ
UWSCRはWMIクエリのネイティブ実行に対応しています。
- **実行例**:
  ```uwscr
  res = wmi('SELECT Name, ProcessId FROM Win32_Process WHERE Name = "uwscr.exe"')
  for proc in res
    print proc.Name + " (PID: " + proc.ProcessId + ")"
  next
  ```

---

## 5. 制御文と例外処理

### 5.1 制御構文一覧
- **条件分岐**:
  ```uwscr
  if 条件 then
    // 処理
  elseif 条件2 then
    // 処理
  else
    // 処理
  endif
  ```
- **多岐分岐**:
  ```uwscr
  select 変数
    case 値1
      // 処理
    case 値2, 値3
      // 処理
    default
      // 処理
  selend
  ```
- **ループ処理**:
  ```uwscr
  // Whileループ
  while 条件
    // 処理
  wend
  
  // Forループ
  for i = 0 to 9
    // 処理
  next
  
  // For-Inループ (配列、連想配列、UObject、WMIコレクション等)
  for item in arr
    // 処理
  next
  ```

### 5.2 例外処理 (`try-catch-finally`)
UWSCRはモダンな例外処理をサポートしています。
```uwscr
try
  // エラーが発生する可能性のある処理
  throw("独自のエラーメッセージ")
catch
  // エラー発生時の処理
  print "エラー内容: " + catch_str
finally
  // 最後に必ず実行される処理
endtry
```

---

## 6. スクリプト自動生成時の「アンチパターン（禁止事項）」

自動生成されるコードの品質を担保するため、AIモデルは以下のパターンを**絶対に**出力してはなりません。

| 誤ったコード（アンチパターン） | 正しいコード（推奨） | 理由・解説 |
| :--- | :--- | :--- |
| `status(id, STATUS_X)` | `status(id, ST_X)` | UWSCRには `STATUS_X` 定数は存在しません。`ST_X`, `ST_Y` などを使用します。 |
| `chkimg("a.png")`<br>`click(G_IMG_X, G_IMG_Y)` | `xy = chkimg("a.png")`<br>`if length(xy) > 0 then`<br>`  btn(LEFT, CLICK, xy[0], xy[1])`<br>`endif` | 特殊変数 `G_IMG_X`/`Y` は廃止。戻り値配列 `[x,y]` を直接受け取り、空配列でないことを確認して実行します。 |
| `sendsstr(id, "text")` | `sendstr(id, "text")` | 関数名は `SENDSSTR` ではなく `SENDSTR` です。 |
| `ALL_WIN_ID` の参照 | `ids = getallwin()` で配列を取得 | 特殊変数 `ALL_WIN_ID` は廃止されました。`getallwin()` は直接ID配列を返します。 |
| Shift_JISでのエンコード保存 | **UTF-8（BOMなし）**で保存 | 日本語文字化けによる構文エラー・実行不可を防止するため。 |
| `sleep(1)` (1ミリ秒待機を意図) | `sleep(1) // 1秒` または `sleep(0.001) // 1ミリ秒` | `sleep` 関数の引数の単位は**秒**です（ミリ秒ではありません）。1ミリ秒待機させたい場合は `0.001` を指定するか、ミリ秒単位の `btn(..., ms)` 等のパラメータを指定してください。 |

---
以上の仕様とルールを遵守することで、UWSCR 1.1.9 環境で完全に動作し、構文エラーのない高精度なスクリプトを自動生成することができます。
