
# UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/index.html

## UWSCR 1.1.9#
UWSCR 1.1.9 オンラインヘルプ
## 使い方#
使い方
インストール方法
zipファイルダウンロード
wingetコマンド
実行方法
コマンドラインオプション
スクリプトファイルのエンコーディング
注意
設定ファイル
ファイルの所在
設定ファイルに関するコマンド
設定の優先順位
設定ファイルの構成
設定ファイルを読み取れない場合
json schemaについて
ビルド方法
UWSCR
ドキュメント
Language Server機能について
Language Clientとの通信方法
Language Serverが提供する機能
サンプルコード集
新機能を使ったサンプルコード
TIPS
## スクリプト構文#
スクリプト構文
スクリプト構文
識別子
変数定義
enum
関数定義
UObject
評価の順序
スコープ
数値
文字列
ホワイトスペース
演算子
条件式の判定
コメント
行結合
マルチステートメント
組み込み定数
16進数
起動時パラメータ
OPTION
def_dll
構造体
スレッド
タスク
with
textblock
call
例外処理
制御文
COMオブジェクト
リストの改行表記
値型
特殊変数
## ビルトイン関数#
ビルトイン関数
引数の型について
文字列
真偽値
文字列操作関数
コピー
copy()
betweenstr()
token()
置換
replace()
chgmoj()
サイズ
length()
lengthb()
lengthu()
lengths()
lengthw()
正規表現
NewRE()
regex()
TestRE()
Match()
利用可能な正規表現
JSON
FromJson()
ToJson()
YAML
FromYaml()
ToYaml()
検索
pos()
変換系
chknum()
val()
trim()
chr()
chrb()
asc()
ascb()
isunicode()
strconv()
format()
encode()
decode()
ウィンドウ操作関数
ID取得
getid()
getallwin()
idtohnd()
hndtoid()
getctlhnd()
ID0について
ウィンドウ操作
clkitem()
ctrlwin()
sckey()
setslider()
sendstr()
mouseorg()
chkmorg()
ウィンドウ情報取得
status()
monitor()
posacc()
muscur()
peekcolor()
getslider()
chkbtn()
getstr()
getitem()
getslctlst()
chkclr()
画像検索
ChkImg()
SearchImage()
saveimg()
低レベル関数
mmv()
btn()
kbd()
acw()
仮想キーコード一覧
スクリプト制御
実行をブロック
sleep()
動的評価
eval()
エラー発生
raise()
assert_equal()
タスク
Task()
WaitTask()
型チェック
type_of()
システム関数
システム情報
kindofos()
env()
setenv()
wmi()
cpuuserate()
sensor()
プロセス実行
exec()
shexec()
CUIシェル
doscmd()
powershell()
pwsh()
入力制御
lockhard()
lockhardex()
音声出力
sound()
beep()
キー入力
getkeystate()
sethotkey()
仮想キーコード一覧
システム制御
poff()
日時
gettime()
音声
speak()
recostate()
dictate()
ファイル操作関数
テキストファイル
fopen()
fget()
fput()
fdelline()
fclose()
CSVファイル
csvopen()
csvclose()
csvread()
csvwrite()
iniファイル
readini()
writeini()
deleteini()
INI関数のファイルID利用について
その他のファイル操作
deletefile()
getdir()
dropfile()
ZIPファイル
zip()
unzip()
zipitems()
GUI
ダイアログ
msgbox()
input()
slctbox()
popupmenu()
メッセージ表示
balloon()
fukidasi()
logprint()
HTMLフォーム
createform()
Form情報
Formオブジェクト
WebViewRemoteObject
組み込みウィンドウのクラス名一覧
配列操作関数
配列の変更
qsort()
reverse()
resize()
setclear()
shiftarray()
配列長を得る
Length()
配列要素を使う
slice()
calcarray()
文字列との相互変換
join()
split()
数学関数
isnan()
random()
abs()
zcut()
int()
ceil()
round()
sqrt()
power()
exp()
ln()
logn()
sin()
cos()
tan()
arcsin()
arccos()
arctan()
COMオブジェクト
COMオブジェクトの作成・取得
createoleobj()
getactiveoleobj()
コレクション
getoleitem()
VARIANT
vartype()
VAR定数
非推奨関数
safearray()
Excel
xlopen()
xlclose()
xlactivate()
xlsheet()
xlgetdata()
xlsetdata()
ウェブ関連
ブラウザ操作
BrowserControl()
Browserbuilder()
RemoteObjectType()
IE関数互換
BrowserBuilderオブジェクト
Browserオブジェクト
TabWindowオブジェクト
RemoteObject
ブラウザ操作サンプル
ダウンロード先やその方法の制御について
HTTPリクエスト
Webrequest()
WebRequestBuilder()
WebRequestオブジェクト
WebResponseオブジェクト
HTTPパーサー
ParseHTML()
HtmlNodeオブジェクト
コレクションのインデックスアクセス
ソケット通信
共通
sclose()
UDP通信
UdpClient()
UdpSend()
UdpRecv()
TCP通信
TcpSend()
TcpListener()
WebSocket
WebSocket()
WsSend()
WsRecv()
## ライセンス#
UWSCRのソースコードは MIT でライセンスされています
依存crateのライセンスは サードパーティライセンス を参照ください
## コミュニティ#
Discord
## 開発支援#
UWSCR開発支援 - CAMPFIRE (キャンプファイヤー)

---

# インストール方法 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/installation.html

## インストール方法#
## zipファイルダウンロード#
githubからダウンロードできます
## 最新版#
githubの 最新リリース を開きます
そのリリースのAssetsから UWSCRx64.zip または UWSCRx86.zip をダウンロードします
## 任意のバージョン#
githubの リリース一覧 からダウンロードしたいバージョンのリリースを探します
そのリリースのAssetsから UWSCRx64.zip または UWSCRx86.zip をダウンロードします
## zipファイルからのインストール#
任意のフォルダにzipファイルを展開します
展開されるファイル
以下のファイルが展開されます
uwscr.exe
必要に応じてファイルを展開したフォルダをPATH環境変数に登録します
## wingetコマンド#
wingetコマンドによるインストールに対応しています
wingetのバージョン1.4.11071以降からご利用いただけます
コマンドプロンプト、またはPowerShellで以下のいずれかのコマンドを実行します
winget install UWSCR
# または
winget install - -id stuncloud . uwscr
初回インストールの場合は uwscr.exe がインストールされたパスが環境変数PATHに登録されます
必要に応じてコマンドプロンプトまたはPowerShellを再起動します
UWSCRが実行できることを確認します
# UWSCRのバージョンを表示
cmd / c uwscr - -version
wingetコマンドが使えない場合
以下の方法でwingetをインストールします
Microsoft Storeアプリを開きます
アプリ インストーラー を検索します ( アプリインストーラー や app installer で見つかります)
アプリ インストーラー をインストールしてください
コマンドプロンプトやPowerShellで winget コマンドが使えることを確認します
# wingetのバージョンを表示
winget - -version
最新リリースがwingetで公開されるタイミングについて
wingetリポジトリへの登録申請はgithubでのリリース後に行われます
登録には数日を要すため、githubでのリリース直後はwingetによるインストールが行えません
また、何かしらの理由により登録申請自体を行わない場合がありえます
登録申請が通らなかった場合
登録申請を行えなかった場合

---

# 実行方法 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/how_to_run.html

## 実行方法#
UWSCRはコンソールアプリケーションです
コマンドプロンプトやPowerShell上で実行してください
Explorer等から実行した場合はコンソールウィンドウが表示されます
## コマンドラインオプション#
## スクリプトの実行#
スクリプトパス [PARAM_STR...] #
スクリプトを実行します
スクリプトパス
実行するスクリプトファイルのパス
PARAM_STR
スクリプトに渡すパラメータ
半角スペース区切りで複数のパラメータを指定可能
渡されたパラメータは PARAM_STR 特殊変数に格納されています
実行例
uwscr hoge.uws foo bar baz
print PARAM_STR // [foo, bar, baz]
ハイフンから始まる文字を渡す場合
- や -- から始まる文字列はコマンドラインオプションと見なされます
PARAM_STRの文字列としてそれらを渡す場合は -- の後に記述してください
uwscr 123 456 - 78 # -78 がオプションだと見なされエラーになる
uwscr 123 456 &quot;-78&quot; # &quot;&quot;で括っても同様
# -- の後に書けばOK
uwscr -- 123 456 - 78 # [&quot;123&quot;, &quot;456&quot;, &quot;-78&quot;] として渡る
batファイルでパラメータを渡す場合の注意
batファイル (またはcmdファイル) からUWSCRを呼ぶ場合にスクリプトのパラメータとして特定の記述をした場合に正しく PARAM_STR へと反映されません
パラメータを &quot;&quot; で括っている
かつその末尾が \ である
かつその後に別のパラメータが続く
以上を満たすパラメータがある場合 PARAM_STR が意図しない値となります
以下の例では [&quot;aaa\&quot;, &quot;bbb&quot;] という二つの文字列の出力が期待されますが、単一の文字列として出力されてしまいます
// test.uws
print PARAM_STR
rem test.bat
rem PARAM_STRがおかしくなる記述
uwscr test.uws &quot;aaa\&quot; &quot;bbb&quot;
# 出力
&gt; test . bat
[ &quot;aaa\&quot; bbb &quot;]
batファイルの記述方法を変更することで回避可能です
rem test2.bat
rem ダブルクォートで括るのを止める
uwscr test.uws aaa\ &quot;bbb&quot;
rem ダブルクォートで括る場合は \ を \ でエスケープする
uwscr test.uws &quot;aaa\\&quot; &quot;bbb&quot;
# 出力
&gt; test2 . bat
[ &quot;aaa\&quot; , &quot;bbb&quot; ]
[ &quot;aaa\&quot; , &quot;bbb&quot; ]
-w , --window #
コンソールから実行された場合にwindowモードでの起動を強制します
スクリプトパスが指定されていない場合使えません
-a , --ast #
スクリプトの構文木を出力します
スクリプトパスが指定されていない場合使えません
--continue #
構文木の出力後にスクリプトを実行する場合に指定
--ast が指定されていない場合使えません
-p , --prettify #
出力される構文木を見やすくします
--ast が指定されていない場合使えません
## REPLモード#
モジュールパス [PARAM_STR...] #
REPL起動前に読み込ませるモジュールファイルのパス
PARAM_STRを渡すこともできる
PS&gt; uwscr hoge.uws foo bar baz --repl
uwscr&gt; PARAM_STR
[ &quot;foo&quot; , &quot;bar&quot; , &quot;baz&quot; ]
-r , --repl #
Replを起動します
Replの使い方
プロンプトに式や文を入力しEnterキーを押すと実行されます
変数への代入などは次の入力にも引き継がれます
スクリプトを読み込ませることで事前に定義した関数等も使用できます
Tabキーで以下の補完が行なえます、いずれも小文字のみにマッチします
ビルトイン関数
ビルトイン定数
キーワードの一部
Alt+Enterで改行します
ブロック構文の入力や複数行の一括実行が行なえます
ヒント
コマンドライン引数がない場合もREPLモードで起動します
## UWSCRライブラリ(uwsl)ファイル出力#
スクリプトパス #
uwslの変換するスクリプトのパス
-l , --lib #
スクリプトのあるディレクトリに スクリプト名.uwsl ファイルを出力します
## コード実行#
-c , --code &lt;CODE&gt; #
渡された文字列を評価して実行します
CODE
UWSCRで評価可能な式または文を示す文字列
半角スペースを含む場合は &quot;&quot; で括ってください
実行例
uwscr -c &quot;msgbox(&#39;hello world!&#39;)&quot;
## 操作記録#
--record [&lt;FILE&gt;] #
実行した操作の低レベル記録を行います。
FILEには記録した操作を保存するパスを指定します
FILEを省略した場合はクリップボードに保存します
## 設定ファイル#
-s , --settings [&lt;OPTION&gt;] #
設定ファイル( settings.json )を開きます
設定ファイルは %APPDATA%\UWSCR\settings.json に出力されます
OPTION
設定ファイルがすでに存在する場合にどのように開くかのオプションを指定します
設定ファイルが存在しない場合これらのオプションは無視され、設定ファイルが新規に作成されます
省略時
設定ファイルが存在していればそれを開きます
init
設定ファイルが存在する場合はそれを破棄し、新たな設定ファイルを出力します
merge
古いバージョンの設定ファイルの内容を可能な限りマージした新しいバージョンの設定ファイルを出力します
--schema [&lt;DIR&gt;] #
設定ファイル用のjson schemaファイル( uwscr-settings-schema.json )を出力します
DIR
出力先ディレクトリのパスを指定
省略した場合はuwscr.exeと同じディレクトリに出力されます
## オンラインヘルプ#
-o , --online-help #
オンラインヘルプをブラウザで表示します
--license #
サードパーティライセンスをブラウザで表示します
## 情報表示#
-h , --help #
ヘルプを表示します
-V , --version #
UWSCRのバージョンを表示します
## スクリプトファイルのエンコーディング#
以下に対応しています
UTF-8
UTF-16 (BE, LE)
Shift-JIS
## 注意#
## ANSIコードポイントについて#
UWSCRではOSのANSIコードポイントが932であることを想定しています
65001(UTF8)等に変更している場合の動作保証はありません

---

# 設定ファイル - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/settings.html

## 設定ファイル#
## ファイルの所在#
uwscr - -settings
を実行することで %APPDATA%\UWSCR\settings.json に出力されます
## 設定ファイルに関するコマンド#
# 現在の設定ファイルを標準のエディタで開く
uwscr - -settings
# 設定ファイルを初期化する
uwscr - -settings init
# 旧バージョンの設定を保持しつつ、新しい設定項目を追加する
uwscr - -settings merge
# いずれの場合も設定ファイルが存在しない場合は新規に作成します
## 設定の優先順位#
スクリプト実行時の設定の優先順位が以下の通りです (上にあるほど優先される)
OPTION設定
設定ファイル(settings.json)
デフォルト設定
## 設定ファイルの構成#
ヒント
設定ファイルのデータ型について
bool
true または false を指定
string (文字列)
文字列、設定が不要な場合は null を指定する
number (数値)
数値
{
&quot;options&quot; : {
// bool : fially部を必ず実行するかどうか
&quot;opt_finally&quot; : false ,
// bool : 変数初期化にdim宣言を必須とするかどうか
&quot;explicit&quot; : false ,
// string: ダイアログのタイトル
&quot;dlg_title&quot; : null ,
// number: ログファイルの出力方法
// 0: 通常のログ出力
// 1: ログ出力なし
// 2: 日時出力なし
// 3: 通常のログ出力 (標準で秒を含むため0と同じ)
// 4: 以前のログを破棄
// それ以外: ログ出力なし
&quot;log_file&quot; : 1 ,
// number: ログファイルの行数
&quot;log_lines&quot; : 400 ,
// string: ログファイルの出力先ディレクトリ
// nullの場合はスクリプトファイルと同じ場所
&quot;log_path&quot; : null ,
// (未対応)
&quot;position&quot; : {
// number: ウィンドウ左上のx座標
&quot;left&quot; : 0 ,
// number: ウィンドウ左上のy座標
&quot;top&quot; : 0
},
// ダイアログなどのフォント
&quot;default_font&quot; : {
// string: フォント名
&quot;name&quot; : &quot;Yu Gothic UI&quot; ,
// number: フォントサイズ
&quot;size&quot; : 15
},
// bool : 仮想デスクトップにも吹き出しを表示するかどうか
&quot;fix_balloon&quot; : false ,
// bool : (未対応)
&quot;no_stop_hot_key&quot; : false ,
// bool : 短絡評価を行うかどうか
&quot;short_circuit&quot; : true ,
// bool : publicの重複宣言を禁止するかどうか
&quot;opt_public&quot; : false ,
// bool : 文字列比較で大文字小文字を区別するかどうか
&quot;same_str&quot; : false ,
// bool : print文をGUIに出力するかどうか
&quot;gui_print&quot; : false
// bool : if文などの条件式をTRUEかFALSEに限定する
&quot;force_bool&quot; : false
// bool : if文などの条件式の判定方法をUWSCと同じにする
&quot;cond_uwsc&quot; : false
},
// BrowserControl設定
&quot;browser&quot; : {
// string: 操作するGoogle Chromeのパス (nullなら自動取得)
&quot;chrome&quot; : null ,
// string: Microsoft Edgeのパス (nullなら自動取得)
&quot;msedge&quot; : null
},
// chkimg設定
&quot;chkimg&quot; : {
// bool : chkimg実行時のスクリーン画像を保存する(chkimg_ss.png)
&quot;save_ss&quot; : false
},
// print窓のフォント
&quot;logfont&quot; : {
// string: フォント名
&quot;name&quot; : &quot;MS Gothic&quot; ,
// number: フォントサイズ
&quot;size&quot; : 15
},
// json schemaのurl: x.x.xはリリースバージョン
&quot;$schema&quot; : &quot;https://github.com/stuncloud/UWSCR/releases/download/x.x.x/uwscr-settings-schema.json&quot;
}
## 設定ファイルを読み取れない場合#
書式が不正な場合は設定ファイルの内容は読み取られません
その場合はデフォルト設定が適用されます
また、エラー(読み取れなかった理由)がコンソールに出力されます
## json schemaについて#
設定ファイルの $schema は設定ファイルに対応したjson schemaのURLです
Visual Studio Code等でjsonファイルを編集する際に補完機能が使えるようになります
## json schemaのオフライン利用#
schemaファイルをローカルに出力することでオフライン環境でもjson schemaが利用できます
以下のコマンドを実行すると指定パスに uwscr-settings-schema.json が出力されます
uwscr - -schema [ パス ]
このファイルのパスをurlに変換し設定ファイルの $schema に指定します
ヒント
ファイルパス→URL変換方法
C:\\uwscr\\uwscr-settings-schema.json であれば file:///C:/uwscr/uwscr-settings-schema.json のように変換する必要があります
PowerShellで簡単に変換できます
PS &gt; ( [uri] &#39;C:\\uwscr\\uwscr-settings-schema.json&#39; ). AbsoluteUri
file :/// C :// uwscr // uwscr-settings-schema . json

---

# ビルド方法 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/how_to_build.html

## ビルド方法#
## UWSCR#
ソースからuwscr.exeをビルドする方法
## Rust開発環境の準備#
Windows 10 x64環境での手順
## Visual C++ Build Tools のインストール#
Visual Studio 2019のツール からBuild Tools for Visual Studio 2019のインストーラをダウンロード
インストーラからVisual C++ Build Toolsのインストールを行う
## Rustのインストール#
Rust をインストール - Rustプログラミング言語 から rustup-init.exe をダウンロード
PowerShellなどから rustup-init.exe を実行
プロンプトに従いインストールを完了する
rustup --version や cargo --version が正常に実行できればOK
ヒント
実行できない場合は一旦PowerShellなどを再起動してみてください
rustup target install i686-pc-windows-msvc を実行してx86版もビルドできるようにする
rustup show を実行し以下のような出力になっていればOK
Default host: x86_64-pc-windows-msvc
rustup home: C:\Users\(your name)\.rustup
installed toolchains
--------------------
stable-i686-pc-windows-msvc
stable-x86_64-pc-windows-msvc (default)
installed targets for active toolchain
--------------------------------------
i686-pc-windows-msvc
x86_64-pc-windows-msvc
active toolchain
----------------
stable-x86_64-pc-windows-msvc (default)
rustc 1.62.0 (a8314ef7d 2022-06-27)
## OpenCV#
chkimgを含める場合は事前にOpenCVをインストール、または静的リンクライブラリをビルドしておく必要があります
インストールした場合はuwscr.exe実行時にopencv_worldXXX.dllを参照できるようにする必要があります
ライブラリをビルドした場合はdllは不要です
chkimgを含めずビルドする場合はこの項目をスキップしてください
## OpenCVのインストール#
OpenCV で opencv-X.Y.Z-vc14_vc15.exe をダウンロード
opencv-X.Y.Z-vc14_vc15.exe を実行し、任意のフォルダに展開
## OpenCVのビルド#
## 準備#
OpenCV のソースコードをダウンロードして任意のフォルダに展開
Cmake をダウンロードしインストール
## cmake#
ヒント
以下ではOpenCVの展開先を C:\tools\opencv としています
生成されるファイルの出力先を C:\tools\opencv64 または C:\tools\opencv86 としています
msvcビルドツールは Visual Studio 16 2019 としています
いずれも環境に合わせて読み替えてください
スタートメニューからcmake-guiを起動
Where is the source code (ソース) に C:\tools\opencv を指定
Where to build the binaries (出力先) に C:\tools\opencv64 を指定 (x86はopencv86)
Configure ボタンを押す (出力先フォルダが存在しない場合はダイアログで確認されるので作成してもらう)
ダイアログが表示されたら
generatorは Visual Studio 16 2019 を選択
platformは x64 を選択
x86なら Win32 にする
toolsetは空欄
Finish ボタンを押してしばらく待つ
リストが表示されるので変更を加える
BUILD_SHARED_LIBS のチェックを外す
BUILD_opencv_* 系は以下のみチェックし、ほかは外す
BUILD_opencv_core
BUILD_opencv_imgcodecs
BUILD_opencv_imgprocs
*_TESTS 系のチェック外す
BUILD_JAVA のチェック外す
WITH_ADE のチェック外す
WITH_QUIRC のチェック外す
WITH_OPENEXR のチェック外す
VC++ランタイムを静的リンクしない場合のみ
BUILD_WITH_STATIC_CRT のチェックを外す
再度 Configure ボタンを押ししばらく待つ
BUILD_FAT_JAVA_LIB が赤くなるけど無視
リストが赤くなっていればなくなるまで Configure ボタンを押す
Generate ボタンを押す
Tip
スクリプトによる実行方法
UWSCRリポジトリにある CmakeOpencv.ps1 で上記と同等のことができます
.\ CmakeOpencv . ps1 -Source C :\ tools \ opencv \ -OutDir C :\ tools \ opencv64 \ -Architecture x64 -WithStaticCrt
## msbuild#
ヒント
msvcビルドツールは Visual Studio 16 2019 がインストールされているものとします
また、cmakeの出力先が C:\tools\opencv64 または C:\tools\opencv86 であるものとします
環境に合わせて適宜読み替えてください
スタートメニューから x64 Native Tools Command Prompt for VS 2019 を起動
以下を実行
x64
cd /d c:\tools\opencv64
chcp 65001
msbuild -p:Configuration=Release;Platform=x64;CodePage=65001 INSTALL.vcxproj
x86
cd /d c:\tools\opencv86
chcp 65001
msbuild -p:Configuration=Release;Platform=Win32;CodePage=65001 INSTALL.vcxproj
C:\tools\opencv64\install (x86なら C:\tools\opencv86\install ) に出力される
## ビルド#
重要
Rustのバージョンについて
UWSCR0.8.1よりCargo.tomlで rust-version が指定されています
このバージョン未満のRustではビルドができません
重要
VC++ランタイムライブラリについて
以下のコマンドでそのままビルドした場合は実行時にVC++ランタイムライブラリが必要になります
exe単体で動作させる(ライブラリを静的リンクする)ためには事前に以下の環境変数をセットしてください
$env:RUSTFLAGS = &#39;-C target-feature=+crt-static&#39;
UWSCRを git clone し、PowerShellでそのディレクトリへ移動
以下のコマンドを実行
## デバッグビルド#
# x64
cargo build
注釈
.\target\debug\ に出力されます
# x86
cargo build - -target = i686-pc-windows-msvc
注釈
.\target\i686-pc-windows-msvc\debug\ に出力されます
## リリースビルド#
# x64
cargo build - -release
注釈
.\target\release\ に出力されます
# x86
cargo build - -target = i686-pc-windows-msvc - -release
注釈
.\target\i686-pc-windows-msvc\release\ に出力されます
## chkimgを含める場合#
## 準備#
LLVM で LLVM-X.Y.Z-win64.exe をダウンロードしてインストール
以下の環境変数を設定する必要があります
具体的な設定値については後述します
OPENCV_INCLUDE_PATHS
includeフォルダのパス
OPENCV_LINK_PATHS
libファイルのあるパス
OPENCV_LINK_LIBS
読み込むlibファイル
## ビルド#
cargo 実行時に --features chkimg を指定しchkimgが含まれるようにします
## OpenCVをインストールした場合#
ヒント
OpenCVのインストール を実行している必要があります
opencvの展開先は C:\tools\opencv\ としています、環境に合わせて適宜読み替えてください
注意
この方法ではx86版はビルドできません
以下は環境変数を設定しつつcargoによるビルドを行うPowerShellスクリプトのサンプルです
# includeフォルダ
$env:OPENCV_INCLUDE_PATHS = &#39;C:\tools\opencv\build\include\&#39;
# libファイルのあるフォルダ
$env:OPENCV_LINK_PATHS = &#39;C:\tools\opencv\build\x64\vc15\lib&#39;
# 読み込むlibファイル
$env:OPENCV_LINK_LIBS = &#39;opencv_worldXXX&#39;
# XXXの部分はopencvのバージョンにより変わります (バージョン4.6.0→460)
cargo build - -features chkimg
重要
この方法でビルドしたuwscr.exeは opencv_worldXXX.dll が参照できないと実行できません
以下のいずれかの方法でdllを参照できるようにしてください
C:\tools\opencv\build\x64\vc15\bin にPATHを通す
C:\tools\opencv\build\x64\vc15\bin\opencv_worldXXX.dll をuwscr.exeと同じフォルダにコピーする
## OpenCVをビルドした場合#
ヒント
OpenCVのビルド を実行している必要があります
msbuildの出力先は C:\tools\opencv64\install\ ( C:\tools\opencv86\install\ ) としています、環境に合わせて適宜読み替えてください
重要
BUILD_WITH_STATIC_CRTについて
VC++ランタイムライブラリを静的リンクしてビルドする場合はopencvビルド時に BUILD_WITH_STATIC_CRT をオンにします
VC++ランタイムライブラリを静的リンクしない場合はopencvビルド時に BUILD_WITH_STATIC_CRT をオフにします
以下は環境変数を設定しつつcargoによるビルドを行うPowerShellスクリプトのサンプルです
x64
# includeフォルダ
$env:OPENCV_INCLUDE_PATHS = &#39;C:\tools\opencv64\install\include&#39;
# libファイルのあるフォルダ
$env:OPENCV_LINK_PATHS = &#39;C:\tools\opencv64\install\x64\vc16\staticlib&#39;
# 読み込むlibファイル
# 複数ある場合は , で連結する
$env:OPENCV_LINK_LIBS = @(
&#39;opencv_coreXXX&#39;
&#39;opencv_imgcodecsXXX&#39;
&#39;opencv_imgprocXXX&#39;
&#39;ippiw&#39;
&#39;ittnotify&#39;
&#39;ippicvmt&#39;
&#39;liblibjpeg-turbo&#39;
&#39;liblibopenjp2&#39;
&#39;liblibpng&#39;
&#39;liblibtiff&#39;
&#39;liblibwebp&#39;
&#39;zlib&#39;
) -join &#39;,&#39;
# XXXの部分はopencvのバージョンにより変わります (バージョン4.6.0→460)
# libから始まるファイルは先頭にlibを追加する必要があります (libpng→liblibpng)
cargo build - -features chkimg
x86
$env:OPENCV_INCLUDE_PATHS = &#39;C:\tools\opencv86\install\include&#39;
$env:OPENCV_LINK_PATHS = &#39;C:\tools\opencv86\install\x86\vc16\staticlib&#39;
$env:OPENCV_LINK_LIBS = @(
&#39;opencv_coreXXX&#39;
&#39;opencv_imgcodecsXXX&#39;
&#39;opencv_imgprocXXX&#39;
&#39;ippiw&#39;
&#39;ittnotify&#39;
&#39;ippicvmt&#39;
&#39;liblibjpeg-turbo&#39;
&#39;liblibopenjp2&#39;
&#39;liblibpng&#39;
&#39;liblibtiff&#39;
&#39;liblibwebp&#39;
&#39;zlib&#39;
) -join &#39;,&#39;
cargo build - -features chkimg - -target = i686-pc-windows-msvc
## GUIアプリケーションとしてビルド#
gui featureフラグを加えてビルドすることでUWSCRはGUIアプリケーションとして振る舞います
cargo build - -features gui
cargo build - -features gui - -release
# chkimgも加える
cargo build - -features gui , chkimg
このfeatureによるビルドは動作保証外です
gui featureは十分なテストが行われていません
このfeatureによるビルドを行った場合の動作については保証されません
## cargoによるテスト実行#
cargoを使ったuwscrのテスト実行方法
都度ビルド→実行を行います
# スクリプトの実行
cargo run -- C :\ uwscr \ test . uws
# x86
cargo run - -target = i686-pc-windows-msvc -- C :\ uwscr \ test . uws
# リリース版で実行
cargo run - -release -- C :\ uwscr \ test . uws
# repl
cargo run
cargo run -- - -repl
# 設定ファイルを開く
cargo run -- - -settings merge
# schemaファイルを出力
cargo run -- - -schema .\ schema
## ドキュメント#
重要
Python実行環境が必要です
## 準備#
pip 等で以下をインストール
Sphinx (ドキュメントのビルド)
furo (ドキュメントのテーマ)
pygments (サンプル構文のシンタックスハイライト)
sphinx-favicon (faviconの設定)
## ビルド#
.\documents\make.bat html を実行
ヒント
.\documents\build\html\ に出力されます

---

# Language Server機能について - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/language_server.html

## Language Server機能について#
UWSCRは Language Server Protocol 準拠のLanguage Serverを実装しています。任意のエディタでUWSCR用のLanguage Clientを実装することでLanguage ServerからUWSCRのコーティング支援を受けられます。
## Language Clientとの通信方法#
UWSCRのLanguage ServerはClientの子プロセスとして動作し、標準入出力によりLanguage Server Protocolでの通信を行います。
Clientからは以下のコマンドでLanguage Serverを起動します。
uwscr - -language-server
## Language Serverが提供する機能#
## Diagnostics#
構文解析を行い解析エラー情報をClientに送信します。UWSCRでは以下の通知(notification)にがあった場合に textDocument/publishDiagnostics 通知を送信します。
textDocument/didOpen : .uwsファイルが開かれたときにServerに送信される通知
textDocument/didSave : ファイルを保存したときにServerに送信される通知
## Completion#
コード補完機能です。Clientの textDocument/completion 通知に対して以下を返します。
組み込み定数
組み込み関数
コードスニペット
## Semantic Tokens#
キーワードのハイライト機能です。Clientからの textDocument/semanticTokens/full 通知に対して以下をSemantic Tokenとして返します。
組み込み定数名
組み込み関数名

---

# サンプルコード集 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/usage/example.html

## サンプルコード集#
UWSCRのサンプルコードを掲載していきます
思いつき次第拡充していく予定です
サンプルコードを見たい機能がある場合は オンラインヘルプ関連 のissueでリクエストしてください
## 新機能を使ったサンプルコード#
## 時間の計測 (クラス、クロージャ)#
実行時間計測用のタイマーを実装してみます
gettime を二度呼び出しその差分を得ることで経過時間を取得できますが、UWSCでは以下のような問題がありました
ミリ秒での計測が面倒
最初のgettimeの結果を隠蔽しにくい
UWSCRの gettime はミリ秒に対応しています
それと、新たに追加されたクラスやクロージャ機能を使ってこれらを解決していきます
この方法だと上記を解決するばかりか、複数の計測を並行するのも簡単になります
## クラスを使う#
Timerクラスを実装し、それにより経過時間を計測します
インスタンスを複数作ることで並行計測も可能です
class Timer
dim from
// コンストラクタで計測開始時の時間をセット
procedure Timer ()
this . from = this . now ()
fend
// 現在時刻をミリ秒で取得
dim now = function ()
result = gettime ( , , , TRUE )
fend
// 経過時間を返す
function elapsed ()
result = this . now () - this . from
fend
endclass
t = Timer ()
msgbox ( &quot;好きなタイミングでOKを押す&quot; )
msg = t . elapsed () + &quot; ミリ秒経過しました&quot;
msgbox ( msg )
## クロージャを使う#
クラスよりもスッキリ書きたい場合はこちら
関数の戻り値を無名関数にすることで、その中に値を閉じ込めておくことができます
この場合は計測開始時を戻り値の無名関数に持たせておくことで、その関数を実行すると経過時間が得られる仕組みです
// 経過時間を得る関数を返す関数 (エンクロージャ)
function timer ()
s = gettime ( ,,, TRUE )
// 経過時間を返す関数 (クロージャ)
// 計測開始時間(s)を保持している
result = function ()
result = gettime ( ,,, TRUE ) - s
fend
fend
elapsed = timer ()
msgbox ( &quot;好きなタイミングでOKを押す&quot; )
msg = elapsed () + &quot; ミリ秒経過しました&quot;
msgbox ( msg )
## UWSCでやるなら？#
モジュールを使えば一応要件を満たすことはできます
id指定により並行計測も可能としていますが、都度id指定が必要なのが不便ですね
Timer . Start ( 1 )
msgbox ( &quot;好きなタイミングでOKを押す&quot; )
msg = Timer . End ( 1 ) + &quot; ミリ秒経過しました&quot;
msgbox ( msg )
module Timer
hashtbl s
procedure Start ( id )
s [ id ] = GetTickCount ()
fend
function End ( id )
result = GetTickCount () - s [ id ]
fend
def_dll GetTickCount () : dword : kernel32 . dll
endmodule
## イテレータ#
classや無名関数を用いてイテレータっぽいものが作れます
## 実装#
class Iter
dim list
dim index = 0
dim type
procedure Iter ( list )
t = type_of ( list )
select t
case TYPE_ARRAY
this . list = list
this . type = t
case TYPE_HASHTBL
hashtbl cpy
for key in list
cpy [ key ] = list [ key ]
next
this . list = cpy
this . type = t
default
raise ( &quot;&lt;#t&gt;はイテレータにできません&quot; , &quot;Iter型エラー&quot; )
selend
fend
function to_list ()
result = this . list
fend
function next ()
if this . index &lt; length ( list ) then
select this . type
case TYPE_ARRAY
result = this . list [ this . index ]
case TYPE_HASHTBL
result = this . list [ this . index , HASH_VAL ]
selend
this . index += 1
else
result = EMPTY
endif
fend
function map ( f : func )
select this . type
case TYPE_ARRAY
for i = this . index to length ( this . list ) - 1
list [ i ] = f ( this . list [ i ])
next
case TYPE_HASHTBL
for i = this . index to length ( this . list ) - 1
key = this . list [ i , HASH_KEY ]
list [ key ] = f ( this . list [ key ])
next
selend
result = this
fend
function filter ( f : func )
select this . type
case TYPE_ARRAY
new = []
for i = this . index to length ( this . list ) - 1
if f ( this . list [ i ]) then
new += this . list [ i ]
endif
next
this . list = new
case TYPE_HASHTBL
for key in this . list
if ! f ( key , this . list [ key ]) then
|=&gt; this . list [ key , HASH_REMOVE ] | ()
endif
next
selend
result = this
fend
function find ( f : func )
select this . type
case TYPE_ARRAY
for i = this . index to length ( this . list ) - 1
if f ( this . list [ i ]) then
result = this . list [ i ]
exit
endif
next
case TYPE_HASHTBL
for key in this . list
if f ( key , this . list [ key ]) then
result = this . list [ key ]
exit
endif
next
selend
result = this
fend
function reduce ( f : func )
select this . type
case TYPE_ARRAY
result = this . list [ this . index ]
for i = this . index + 1 to length ( this . list ) - 1
result = f ( result , this . list [ i ])
next
case TYPE_HASHTBL
result = this . list [ this . index , HASH_VAL ]
for i = this . index + 1 to length ( this . list ) - 1
result = f ( result , this . list [ i , HASH_VAL ])
next
selend
fend
endclass
## 使い方#
a = [ 1 , 2 , 3 ]
hash h
&quot;a&quot; = 1
&quot;b&quot; = 2
&quot;c&quot; = 3
endhash
// map
f = | n =&gt; n * 2 |
print Iter ( a ) . map ( f ) . to_list () // [2, 4, 6]
print Iter ( h ) . map ( f ) . to_list () // {&quot;A&quot;: 2, &quot;B&quot;: 4, &quot;C&quot;: 6}
// filter
print Iter ( a ) . filter ( | n =&gt; n mod 2 == 1 | ) . to_list () // [1, 3]
// 連想配列はキーと値を受けてフィルタできる
print Iter ( h ) . filter ( | k , v =&gt; v mod 2 == 1 | ) . to_list () // {&quot;A&quot;: 1, &quot;C&quot;: 3}
// reduce
f = | x , y =&gt; x + y |
print Iter ( a ) . reduce ( f ) // 6
print Iter ( h ) . reduce ( f ) // 6
// find
print Iter ( a ) . find ( | n =&gt; n mod 2 == 0 | ) // 2
print Iter ( h ) . find ( | k , v =&gt; k == &quot;c&quot; | ) // 3
// 複合
print Iter ([ 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8 , 9 ]) _
. filter ( | n =&gt; n mod 2 == 0 | ) _ // 偶数
. map ( | n =&gt; n + 1 | ) _ // それぞれに+1
. reduce ( | m , n =&gt; m + n | ) // 合計を出す: 24
## TIPS#
## UWSCとUWSCRを判別#
GET_UWSC_PRO で判別できます
UWSCではPro版か否かでTRUEまたはFALSEを返していましたが、UWSCRではEMPTYを返します
select GET_UWSC_PRO
case TRUE
print &quot;UWSC Pro版です&quot;
case FALSE
print &quot;UWSC 無料版です&quot;
case EMPTY
print &quot;UWSCRです&quot;
selend

---

# スクリプト構文 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/syntax/statement.html

## スクリプト構文#
## 識別子#
識別子とは変数、定数、関数などの名前を示す文字列です
以下の文字の組み合わせで命名できます
英字 (大文字・小文字の区別はしません)
数字
記号
_
全角文字
## キーワード一覧#
## 識別子 (変数名、定数名、関数名など) に使用できないキーワード#
特殊構文キーワード
call
async
await
特殊な値を示すもの
null
empty
nothing
true, false
NaN
演算子
mod
and, andL, andB
or, orL, orB
xor, xorL, xorB
## 変数定義#
## dim#
ローカル変数を定義します
dim hoge // 変数 hoge を定義
dim fuga = 1 // 値の代入も同時に行える
piyo = 1 // 未宣言変数への代入式で新たな変数が定義される(dim省略)
OPTION EXPLICIT指定時の動作
OPTION EXPLICIT を指定した場合は未宣言の変数への代入がエラーとなります
未宣言変数への代入や複合代入は解析エラーとなります
OPTION EXPLICIT
dim foo = 1 // ok
foo = 2 // ok
bar = 3 // 解析エラー
bar += 4 // 解析エラー
baz := 5 // 解析エラー
foo = qux := 100 // 解析エラー
## public#
グローバル変数を定義します
public hoge = 1
fuga () // 1
hoge = 2
fuga () // 2
procedure fuga ()
print hoge
fend
## const#
定数を定義します
再代入ができません
const hoge = 1
hoge = 2 // エラー
## 一括定義#
, 区切りで変数を一括定義できます
dim a = 1 , b = 2 , c , d [ 3 ] , e [] = 1 , 2 , 3 , 4 , 5
public f = 1 , g = 2
const h = 1 , i = 2
UWSCではエラーになっていたconstの一括定義も可能
注意
配列定義に続けて記述するのはNG
dim foo [] = 1 , 2 , 3 , a = 1 // a = 1 は定義できない
## 配列#
配列の定義はdimを使った方法と、配列リテラル(新機能)を使う方法があります
配列の各要素には 配列変数[添字] という書式でアクセスできます
添字は数値で、n番目の要素に対してn-1を指定します
// 従来の配列定義
dim hoge [] = 1 , 2 , 3
print hoge [ 0 ] // 1
// 配列リテラル
fuga = [ 1 , 2 , 3 ]
print fuga [ 0 ] // 1
// 配列リテラルにインデックスを指定することも可能
print [ 4 , 5 , 6 ][ 0 ] // 4
リスト末尾のカンマについて
dim配列宣言ではリストの末尾がカンマだと解析エラーになります
dim hoge [] = 1 , 2 , 3 , // [1, 21] - リストの終端をカンマにはできません
配列リテラルであれば末尾カンマは有効です (無視されます)
fuga = [ 1 , 2 , 3 , ]
print fuga // [1, 2, 3]
piyo = [
4 ,
5 ,
6 ,
]
print piyo // [4, 5, 6]
## +演算子による要素の追加#
+ 演算子で配列の末尾に要素を追加できます
print [ 1 , 2 , 3 ] + 4
// [1, 2, 3, 4]
dim arr = [ 5 , 6 , 7 ]
arr += 8
print arr
// [5, 6, 7, 8]
## 多次元配列#
// 2 次元
dim 配列名 [ 要素数 ][ 要素数 ] = 値 , 値 , 値 , 値 ...
// 3 次元
dim 配列名 [ 要素数 ][ 要素数 ][ 要素数 ] = 値 , 値 , 値 , 値 ...
// 以下の書式も可能
dim 配列名 [ 要素数 , 要素数 ] = 値 , 値 , 値 , 値 ...
dim 配列名 [ 要素数 , 要素数 , 要素数 ] = 値 , 値 , 値 , 値 ...
// 一番左の要素数のみ省略可能
dim 配列名 [][ 要素数 ][ 要素数 ] = 値 , 値 , 値 , 値 ...
dim 配列名 [, 要素数 , 要素数 ] = 値 , 値 , 値 , 値 ...
// 呼び出しは次元数分だけ [] をつける
print 配列名 [ 0 ][ 0 ][ 0 ] // 3 次元配列の1つ目の要素
// 不足分はEMPTYで埋められる
dim sample1 [ 2 ][ 1 ] = 0 , 1 , 2 , 3
print sample1 // [[0, 1], [2, 3], [, ]]
// 超過分は捨てられる
dim sample2 [ 1 , 1 ] = 0 , 1 , 2 , 3 , 4 , 5
print sample2 // [[0, 1], [2, 3]]
// 要素数省略
dim sample3 [][ 1 ] = 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8
print sample3 // [[1,2] ,[3,4], [5,6], [7,8]]
// 一番左以外は省略不可
dim bad_sample [][][ 1 ] // エラー
配列リテラルを使って多次元配列を作ることもできます
dim sample4 [] = [ 1 , 2 ] , [ 3 , 4 ] , [ 5 , 6 ] , [ 7 , 8 ]
sample5 = [[ 1 , 2 ] , [ 3 , 4 ] , [ 5 , 6 ] , [ 7 , 8 ]]
## 連想配列#
hashtbl 連想配列変数 // 連想配列を宣言
hashtbl 連想配列変数 = HASH_CASECARE // キーの大文字小文字を区別
hashtbl 連想配列変数 = HASH_SORT // キーでソート(※1)
hashtbl 連想配列変数 = HASH_CASECARE or HASH_SORT // 大小文字区別かつソート
連想配列変数[キー] = 値 // 任意のキー名で値を代入、数値のキーは文字列に変換される
値 = 連想配列変数[キー] // キー名で値を読み出す、キーがない場合はEMPTY
真偽値 = 連想配列変数[キー, HASH_EXISTS] // キーが存在するかどうか ※2
真偽値 = 連想配列変数[キー, HASH_REMOVE] // キーを削除、成功時はTRUE
キー = 連想配列変数[i, HASH_KEY] // i番目の要素のキーを取得 ※3
値 = 連想配列変数[i, HASH_VAL] // i番目の要素の値を取得 ※3
連想配列変数 = HASH_REMOVEALL // 要素をすべて消す
// カンマ区切りで一括定義可能、オプションも指定できる
hashtbl 変数1, 変数2 = HASH_CASECARE, 変数3 = HASH_SORT
※1
HASH_SORT によるキーソート順はUWSCと異なる場合があります
※2
連想配列変数[キー, HASH_REMOVE] は式であるためこれ単体では実行できません (解析時エラーとなる)
有効な文の中に該当式を記述する必要があります (例: ダミー変数に代入)
// エラーを回避してHASH_REMOVEする方法の例
// 変数に代入
_dummy = hoge [ key , HASH_REMOVE ]
// 即時関数内で実行
| =&gt; hoge [ key , HASH_REMOVE ] | ()
ただし、replモードであればコンソールに戻り値を表示するため式のみで実行できます
uwscr&gt; hoge[key, HASH_REMOVE]
True
※3
iは0から
HASH_SORT がない場合は代入した順序
HASH_SORT がある場合はキーによりソートされた順序
hashtbl hoge
hoge [ &quot;foo&quot; ] = 100
print hoge [ &quot;foo&quot; ] // 100
hoge [ &quot;FOO&quot; ] = 200
print hoge [ &quot;foo&quot; ] // 200 大小文字区別がないため上書きされた
hoge [ &quot;bar&quot; ] = 400
hoge [ &quot;baz&quot; ] = 600
for i = 0 to length ( hoge ) - 1
print hoge [ i , HASH_KEY ] // foo, bar, baz の順で表示される
print hoge [ i , HASH_VAL ] // 200, 400, 600
next
print hoge [ &quot;bar&quot; , HASH_EXISTS ] // true
print hoge [ &quot;qux&quot; , HASH_EXISTS ] // false
hoge [ &quot;bar&quot; , HASH_REMOVE ] // 変数で受けなくてもOK
print hoge [ &quot;bar&quot; , HASH_EXISTS ] // false
hashtbl fuga = HASH_CASECARE
fuga [ &quot;foo&quot; ] = 1
fuga [ &quot;Foo&quot; ] = 2
fuga [ &quot;FOO&quot; ] = 3
print fuga // {&quot;foo&quot;: 1, &quot;Foo&quot;: 2, &quot;FOO&quot;: 3}
hashtbl piyo = HASH_SORT
piyo [ &quot;b&quot; ] = &quot;&quot;
piyo [ &quot;z&quot; ] = &quot;&quot;
piyo [ &quot;a&quot; ] = &quot;&quot;
print piyo // {&quot;A&quot;: , &quot;B&quot;: , &quot;Z&quot;: }
## 連想配列一括定義#
hash-endhash で連想配列を一括定義できます
hash [ public ] 変数名 [ = オプション ]
[ キー = 値 ]
endhash
public (省略可)
指定するとグローバル変数、省略時はローカル変数になる
オプション (省略可)
HASH_SORT と HASH_CASECARE を指定可能
省略時はオプションなし
キー = 値
キーと値の組み合わせを指定する
複数指定可
キーは文字列だが '' や &quot;&quot; は省略可能
一つも指定しない場合空の連想配列ができる
// 一括定義
hash foobar
&#39;foo&#39; = 1 // キー = 値形式で記述
bar = 2 // キーは文字列でなくても良い
endhash
// 以下と同じ
// hashtbl foobar
// foobar[&#39;foo&#39;] = 1
// foobar[&#39;bar&#39;] = 2
// グローバル変数にする
hash public pub
endhash
// 以下と同じ
// public hashtbl pub
// オプション指定
hash with_option = HASH_CASECARE or HASH_SORT
endhash
// 以下と同じ
// hashtbl with_option = HASH_CASECARE or HASH_SORT
## enum#
列挙体を定義します
グローバルスコープの定数として定義されます
// 定義
enum 定数名
メンバ名
メンバ名 [ = 数値 ]
endenum
// 呼び出し
定数名 . メンバ名
メンバには上から順に数値が割り当てられます (0から)
メンバ名 = 数値 とすることで任意の値を割り当てられます
ただし前のメンバより大きな値のみ有効です
// 0から順に割り当てられる
enum E
foo // 0
bar // 1
baz // 2
endenum
// 呼び出しは定数名.メンバ名
print E . foo // 0
print E . bar // 1
print E . baz // 2
// 数値を指定
enum E
foo = 10 // 10
bar = 20 // 20
endenum
// 一箇所指定するとそれ以降はその値から加算されていく
enum E
foo = 10 // 10
bar // 11 (上の10 に +1される)
baz // 12
endenum
// 途中も可
enum E
foo // 0
bar = 10 // 10
baz // 11
endenum
enum E
foo = 100 // 100
bar // 101
baz = 200 // 200
qux // 201
endenum
// 以下はNG
enum E
foo
foo // 同じ名前はダメ
endenum
// 前の数値より大きくないとダメ
enum E
foo // 0
bar // 1
baz = 1 // 2以上じゃないとダメ
endenum
enum E
foo = 50
bar = 1 // 51以上じゃないとダメ
endenum
## 関数定義#
関数名には英数字、一部記号、全角文字列が使えます
英字の大文字小文字は区別しません
## procedure#
## function#
procedure 関数名([引数, 引数, …])
処理
fend
function 関数名([引数, 引数, …])
[result = 戻り値]
fend
procedure
戻り値がありません
function
result 変数の値が戻り値となります
result (省略可)
初期値は EMPTY です
記述がない場合は EMPTY を返します
hoge ( 1 , 2 , 3 ) // 6
print fuga ( 1 , 2 , 3 ) // 6
procedure hoge ( a , b , c )
print a + b + c
fend
function fuga ( a , b , c )
result = a + b + c
fend
関数定義の入れ子はダメ
// エラーになる
procedure p ()
procedure q ()
fend
fend
## 特殊な引数#
## 参照渡し#
引数の前に var または ref キーワードをつけることで参照渡しが可能な引数になります
引数に変数を渡すとその変数に関数実行中の変更が反映されます
変数以外の式を渡した場合は通常の引数と同様に振る舞います
a = 2
print a // 2
p ( a )
print a // 6
q ( a )
print a // 16
procedure p ( ref r )
r *= 3
fend
procedure q ( var v )
v += 10
fend
## 配列表記#
引数[] 形式で記述します
互換性のため表記自体はできますが、動作は通常の引数と同様です
受けられる引数を配列や連想配列に限定したい場合は 引数の型チェック を使用してください
// 以下は同じ意味です
procedure p ( arr [])
procedure p ( arr )
## デフォルト値#
引数 = 値 とすることで引数のデフォルト値を指定できます
値を省略した場合は EMPTY がデフォルト値になります
呼び出し時に引数を渡さなかった場合デフォルト値が適用されます
print f ( 2 ) // 0
print f ( 2 , 3 ) // 6
function f ( n , m = 0 )
result = n * m
fend
// デフォルト値を省略した場合はEMPTYが入る
procedure p ( arg = )
print arg == EMPTY // True
fend
デフォルト値を持つ引数のあとに別の種類の引数は指定できません
procedure p ( a = 1 , b = 2 , c = 3 ) // ok
fend
procedure q ( a = 1 , b , c = 3 ) // エラー
fend
procedure r ( a , b = 2 , c = 3 ) // 前ならok
fend
## 可変長引数#
引数の前に args または prms キーワードをつけることで可変長の引数を受けられるようになります
関数内ではその引数が配列になります
可変長引数は最後の引数でなくてはいけません
print f ( 1 ) // 1
print f ( 1 , 2 , 3 , 4 , 5 ) // 5
function f ( args v )
result = length ( v )
fend
可変長引数のあとに引数があるとエラーになります
procedure p ( prms a , b ) // エラー
fend
procedure q ( a , b , prms c ) // ok
fend
## 特殊な引数の組み合わせ#
原則として組み合わせられません
配列表記の参照渡しのみOK
procedure p ( ref foo []) // これはOK
// こういうのはダメ
procedure p ( ref foo = 1 ) // 参照 + デフォルト値
procedure p ( ref params bar ) // 参照 + 可変長
procedure p ( params bar = 1 ) // 可変長 + デフォルト値
## 引数の型チェック#
function 関数名 ( 引数名 : 型 , var 引数名 : 型 , 引数名 : 型 = デフォルト値 )
通常の引数、参照渡し、デフォルト値を持つ引数であれば受ける型を指定できます
関数呼び出し時に指定した型が渡されなかった場合は実行時エラーになります
指定可能な型
string
文字列
number
数値
bool
真偽値 (TRUE/FALSE)
array
配列
hash
連想配列
func
関数 (ユーザー定義、無名関数)
uobject
UObject
クラス名
クラスオブジェクトのインスタンス
列挙体名
列挙体(enum)メンバの値 (該当する数値でも良い)
function f ( str : string )
result = str
fend
print f ( &quot;hoge&quot; ) // OK
print f ( 123 ) // 数値なのでエラー
// 列挙体名指定の場合
enum Hoge
foo
bar
baz
endenum
function f2 ( n : Hoge )
select n
case Hoge . foo
result = &#39;foo!&#39;
case Hoge . bar
result = &#39;bar!&#39;
case Hoge . baz
result = &#39;baz!&#39;
selend
fend
print f2 ( Hoge . foo ) // OK
print f2 ( &quot;Hoge&quot; ) // 文字列はエラー
print f2 ( 0 ) // OK ※Hoge.fooに一致するため
print f2 ( 100 ) // Hogeに含まれない値なのでエラー
## 無名関数#
名前を持たない関数です
変数に代入して使えます
変数 = function ([ 引数 , ... ])
[ result = 戻り値 ]
fend
変数 = procedure ([ 引数 , ... ])
fend
変数に関数を代入できます
hoge = function ( x , y )
result = x + y
fend
print hoge ( 2 , 3 ) // 5
無名関数の中でpublic/constを宣言した場合は実行時に初めて評価されます
print x // エラー
proc = procedure ()
public x = 5
fend
print x // エラー
proc ()
print x // 5
通常の関数と同様に特殊な引数も定義できます
f = function ( a , b [] , var c , d = 0 )
fend
p = procedure ( args e )
fend
## 簡易関数式#
無名関数を単行の式で記述できます
通常の無名関数と異なり処理部に文は書けません(式のみ)
その代わりに即時関数として利用できます
関数 = | 引数 [, 引数, …] =&gt; 式 [; 式; …] |
引数は,区切りで複数指定可能
result は省略可能です
func = | a , b =&gt; a + b |
print func ( 1 , 2 ) // 3
式は ; 区切りで複数書けます
この場合一番最後の式が戻り値となります
func = | a , b =&gt; a *= 2 ; b *= 3 ; a + b |
print func ( 1 , 2 ) // 8
## 即時関数#
print | n , m =&gt; n * m | ( 7 , 6 ) // 42
// 値だけ返す
print |=&gt; 42 | () // 42
// 関数の引数にする
function f ( fn )
result = fn ( &quot;world!&quot; )
fend
print f ( | s =&gt; &quot;hello &quot; + s | ) // hello world!
特殊な引数にも対応
print | args a =&gt; length ( a ) | ( 1 , 2 , 3 , 4 , 5 , 6 ) // 6
## 関数の特殊な使用例#
## 高階関数#
関数の引数に関数を指定できます
print Math ( 10 , 5 , Add ) // 15
print Math ( 10 , 5 , Multiply ) // 50
subtract = function ( n , m )
result = n - m
fend
print Math ( 10 , 5 , subtract ) // 5
function Math ( n , m , func )
result = func ( n , m )
fend
function Add ( n , m )
result = n + m
fend
function Multiply ( n , m )
result = n * m
fend
## クロージャ#
関数の戻り値として関数(クロージャ)を返すことができます
クロージャは元の関数内での値を保持します
hoge = test ( 5 ) // test関数内の変数nを5にする
// 関数hogeはn=5を保持している
print hoge ( 3 ) // 8 (5+3が行われる)
print hoge ( 7 ) // 12 (5+7が行われる)
print hoge ( &quot;あ&quot; ) // 5あ (5+&#39;あ&#39;が行われる)
function test ( n )
result = function ( m )
result = n + m
fend
fend
## エイリアス#
関数を変数に代入することでその関数を別の名前で呼び出せるようになります
function hoge ( n )
result = n
fend
h = hoge // 変数hにhoge関数を代入
print h ( &#39;hoge&#39; ) // hoge
// ビルトイン関数も代入できる
mb = msgbox
mb ( &#39;ほげほげ&#39; )
## module#
機能のモジュール化
モジュール名.メンバ名 で各機能を利用可能にします
module モジュール名
const 定数名 = 式 // モジュール名.定数名 で外部からアクセス可
public 変数名[ = 式] // モジュール名.変数名 で外部からアクセス可
dim 変数名[ = 式] // 外部からアクセス不可
procedure モジュール名 // コンストラクタ、module定義の評価直後に実行される
procedure 関数名() // モジュール名.関数名() で外部からアクセス可
function 関数名() // モジュール名.関数名() で外部からアクセス可
textblock 定数名 // モジュール名.定数名 で外部からアクセス可
endmodule
## module関数内でのみ使える特殊な書式#
this
自module内のメンバの呼び出しを明示する
global
グローバル変数および関数を呼び出す(ビルトイン含む)
(本家と異なり変数や定数も可)
module sample
dim d = 1
public p = 2
const c = 3
function f1 ()
// 各メンバーには以下のようにアクセス可能
print d
print this . d
print sample . d
print p
print this . p
print sample . p
print c
print this . c
print sample . c
print f2 ()
print this . f2 ()
print sample . f2 ()
fend
function f2 ()
result = 4
fend
function f3 ()
print this . f4 () // in メンバ関数が呼ばれる
print global . f4 () // out module外の関数が呼ばれる
print f4 () // in メンバ関数が呼ばれる
fend
function f4 ()
result = &quot;in&quot;
fend
endmodule
function f4 ()
result = &quot;out&quot;
fend
## プライベート関数#
無名関数を用いたプライベート関数の実装例
Sample . Private () // エラー
Sample . Func () // OK
module Sample
function Func ()
result = Private ()
fend
dim Private = function ()
result = &quot;OK&quot;
fend
endmodule
## class#
classを定義します
class名() を実行することによりインスタンスを作成します
注意
UWSCのclassとは互換性がありません
class class名
procedure class名 () // コンストラクタ ( 必須 )
procedure _class名_ () // デストラクタ ( オプション )
const 定数名 = 式 // classインスタンス . 定数名 で呼び出し可
public 変数名 [ = 式 ] // classインスタンス . 変数名 で呼び出し可
dim 変数名 [ = 式 ] // class内からのみ呼び出し可
procedure 関数名 () // classインスタンス . 関数名 () で呼び出し可
function 関数名 () // classインスタンス . 関数名 () で呼び出し可
textblock 定数名 // classインスタンス . 定数名 で呼び出し可
endclass
h1 = hoge ( 3 , 5 )
print h1 . Total () // 8
h2 = hoge ( 8 , 10 )
print h2 . Total () // 18
print hoge ( 11 , 22 ) . Total () // 33
class hoge
dim a = 1 , b = 2
procedure hoge ( a , b )
this . a = a
this . b = b
fend
function Total ()
result = this . a + this . b
fend
endclass
注意
moduleと異なりclass名から直接メンバにアクセスすることはできません
print hoge . p () // エラー
## デストラクタ#
デストラクタはインスタンスへの参照がなくなった際に実行される関数です
_class名_() で命名された関数がデストラクタとして定義されます
デストラクタに引数は指定できません
デストラクタが実行されるタイミング
すべての参照が失われたとき
いずれかのインスタンス変数に NOTHING を代入したとき (明示的に破棄する)
インスタンス変数は NOTHING になります
with``に渡す式でインスタンスを作成した場合で ``endwith に到達したとき
関数スコープを抜ける際に削除されるこローカルスコープ変数だった場合
スクリプト終了時に削除されるローカル・グローバル定数だった場合
class Sample
dim msg
procedure Sample ( msg )
this . msg = msg
fend
procedure _Sample_ ()
print msg
fend
endclass
obj1 = Sample ( &quot;すべての参照が失われた&quot; )
obj2 = obj1
obj3 = obj1
obj1 = 1
obj2 = 1
obj3 = 1 // すべての参照が失われた がprintされる
obj1 = Sample ( &quot;NOTHINGが代入された&quot; )
obj2 = obj1
obj3 = obj1
obj1 = NOTHING // NOTHINGが代入された がprintされる
print obj1 // NOTHING
print obj2 // NOTHING
print obj3 // NOTHING
with Sample ( &quot;withを抜けた&quot; )
endwith // withを抜けた がprintされる
procedure p ()
obj = Sample ( &quot;関数スコープを抜けた&quot; )
fend
p () // 関数スコープを抜けた がprintされる
## UObject#
json互換のオブジェクト
## オブジェクトの作成#
UObjectリテラル: jsonを &#64; で括る
FromJson 関数
obj = @{
&quot;foo&quot; : &quot;fooooo&quot; ,
&quot;bar&quot; : {
&quot;baz&quot; : true
} ,
&quot;qux&quot; : [
{ &quot;quux&quot; : 1 } ,
{ &quot;quux&quot; : 2 } ,
{ &quot;quux&quot; : 3 }
]
}@
arr = @[ 1 , 2 , 3 ]@
有効な値は
数値
文字列
真偽値
NULL
配列
オブジェクト
Tip
UObjectリテラル内での変数展開について
&#64; で括られたjson部分は文字列として扱われます
これは展開可能文字列であるため &quot;&lt;#変数名&gt;&quot; が利用可能です
foo = &#39;文字列を展開&#39;
bar = 123
textblock baz
,
&quot;baz&quot;:{
&quot;qux&quot;: &quot;jsonの一部を一気に書き込むことも可能&quot;
}
endtextblock
obj = @{
&quot;foo&quot; : &quot;&lt;#foo&gt;&quot; ,
&quot;bar&quot; : &lt;# bar &gt;
&lt;# baz &gt;
}@
print obj . foo // 文字列を展開
print obj . bar // 123
print obj . baz . qux // jsonの一部を一気に書き込むことも可能
## 値の呼び出し、変更#
print obj . foo // fooooo
obj . foo = &quot;FOOOOO&quot;
print obj . foo // FOOOOO
print obj [ &quot;foo&quot; ] // 配列の添字にしてもOK
print obj . bar . baz ? &quot;baz is true!&quot; : &quot;baz is fasle!&quot; // baz is true!
obj . qux [ 1 ] . quux = 5
print obj . qux [ 1 ] . quux // 5
obj . qux [ 2 ] = &quot;overwrite!&quot;
print obj . qux [ 2 ] // overwrite!
obj . corge = 1 // エラー、追加はできない
// オブジェクトを作って代入ならOK
obj . foo = fromjson ( &#39;{&quot;hoge&quot;: 1, &quot;fuga&quot;: 2}&#39; )
print obj . foo
## UObjectメソッド#
keys ( ) #
オブジェクト配下のキー名の一覧を取得します
UObjectがオブジェクトではない場合(配列等)は空の配列を返します
配列内の順序は元のjsonの順序と一致するとは限りません
parameter
:rtype: 配列(文字列)
:return: キー名の一覧
サンプルコード
obj = @{ &quot;foo&quot; : 1 , &quot;bar&quot; : 2 , &quot;baz&quot; : 3 }@
print obj . keys () // [bar, baz, foo]
values ( ) #
オブジェクト配下のキーが持つ値の一覧を取得します
UObjectがオブジェクトではない場合(配列等)は空の配列を返します
配列内の順序は元のjsonの順序と一致するとは限りませんが、インデックスはkeys()に対応します
parameter
:rtype: 配列(値)
:return: キーが持つ値一覧
サンプルコード
obj = @{ &quot;foo&quot; : 1 , &quot;bar&quot; : 2 , &quot;baz&quot; : 3 }@
print obj . values () // [2, 3, 1]
## 評価の順序#
グローバル変数や定数、関数定義は実行より先に評価されます
public, const, textblockを記述順に評価
function, procedure, moduleを記述順に評価
関数内で宣言されているpublicやconstも評価
残りの構文を評価/実行する
## スコープ#
スコープは大まかに分けると
スクリプト本文
関数内
という区分になっています
変数にはローカルとグローバルという区分があり、
スクリプト本文のローカル変数はスクリプト本文内でしかアクセスできない
関数のローカル変数は関数内でしかアクセスできない
グローバル変数はいずれからでもアクセスできる
という特徴があります
ローカル
dim宣言した変数
宣言省略した変数も含む
hashtbl宣言した連想配列
グローバル
public宣言した変数
public hashtbl
const宣言した定数
定義した関数 (変数ではないが扱いはグローバル)
public global1 = &quot;グローバル変数1&quot;
dim local = &quot;本文ローカル&quot;
print global1 // ok
print global2 // ok
print local // ok
print proc_local // ng
print func () // ok
procedure proc ()
public global2 = &quot;グローバル変数2&quot;
dim proc_local = &quot;関数ローカル&quot;
print global1 // ok
print global2 // ok
print local // ng
print proc_local // ok
print func () // ok
fend
function func ()
result = &quot;関数&quot;
fend
## 無名関数のスコープ#
無名関数の中はスコープが分かれていません
ローカル変数がそのまま使えます
dim local = 1
dim func = function ( n )
result = local + n
fend
print func ( 1 ) // 2
## moduleのスコープ#
moduleメンバに関しては独自のスコープを持ちます
module関数内で定義したpublic, const, function/procedureはグローバル空間には置かれず、
moduleメンバのみがアクセスできるmoduleローカル空間に配置されます
これらは module名.メンバ名 でアクセスできます
## 数値#
UWSCRにおける数値は倍精度浮動小数点数として扱われます
そのため、小数の演算が期待と異なる結果になる場合があります
print 0 . 1 + 0 . 2 // 0.30000000000000004
また、UWSCとは異なり負のゼロが存在します
print - 0 // -0
a = - 0
print &quot;&lt;#a&gt;&quot; // -0
print a + 0 // 0
print a == - 0 // True
print a == 0 // True
## 文字列#
文字列リテラルは &quot;&quot; または '' で括ります
&quot; で括った文字列では特殊文字が展開されます
' で括った文字列では特殊文字が展開されません
str = &quot;文字列&quot;
str = &#39;文字列&#39;
文字列の結合は + 演算子を使います
str = &quot;文字列&quot; + &quot;の&quot; + &quot;結合&quot;
print str // 文字列の結合
## 特殊文字の展開#
&quot;&quot; で括った文字列中にある以下の特殊文字は、それぞれ該当する別の文字に変換されます
&lt;#CR&gt; : 改行 (CRLF)
&lt;#TAB&gt; : タブ文字
&lt;#DBL&gt; : ダブルクォーテーション ( &quot; )
&lt;#NULL&gt; : NULL文字 ( chr(0) )
&lt;#変数名&gt; : 変数が存在する場合、その値
print &quot;hoge&lt;#CR&gt;fuga&lt;#CR&gt;piyo&quot;
// hoge
// fuga
// piyo
print &quot;hoge&lt;#TAB&gt;fuga&lt;#TAB&gt;piyo&quot;
// hoge fuga piyo
print &quot;&lt;#DBL&gt;hoge&lt;#DBL&gt;&quot;
// &quot;hoge&quot;
dim a = 123
print &quot;a is &lt;#a&gt;&quot;
// a is 123
print &quot;b is &lt;#b&gt;&quot; // 変数が存在しない場合は展開されない
// b is &lt;#b&gt;
print &quot;length of a is &lt;#length(a)&gt;&quot; // 式はダメ、変数のみ展開される
// length of a is &lt;#length(a)&gt;
print &#39;a is &lt;#a&gt;&#39; // シングルクォーテーション文字列は展開しない
// a is &lt;#a&gt;
## ホワイトスペース#
半角スペース
タブ文字
全角スペース
はホワイトスペース扱いです
式と式の区切りとして機能します
改行(CRLF、CR、LF)は行末扱いです
## 演算子#
+
数値の加算、文字列の結合、配列要素の追加
+=
数値の加算、文字列の結合、配列要素の追加をして代入
-
数値の減算
-=
減算して代入
*
数値の乗算、文字列の繰り返し
*=
乗算して代入
/
数値の除算 ※ 0で割ると0を返す
/=
除算して代入
mod
数値の剰余演算 (割った余りを返す)
!
論理否定
真偽値以外の値について
真偽値以外に対して ! 演算子を用いた場合、右辺に対して 条件式の判定 が行われます
OPTION FORCEBOOL や OPTION CONDUWSC の有無により異なる結果を返します
? :
三項演算子 b ? t : f
:=
代入 (代入した値を返す)
=
代入、等価演算
==
等価演算
&lt;&gt;
!=
不等価演算
and
数値のAND演算(ビット演算)
or
数値のOR演算(ビット演算)、真偽値の論理演算
xor
数値のXOR演算(ビット演算)
andL
論理演算 (両辺の真偽性評価を行う)
orL
論理演算 (両辺の真偽性評価を行う)
xorL
論理演算 (両辺の真偽性評価を行う)
andB
ビット演算 (両辺を数値とみなし評価を行う)
orB
ビット演算 (両辺を数値とみなし評価を行う)
xorB
ビット演算 (両辺を数値とみなし評価を行う)
&lt;
小なり
&lt;=
小なりイコール
&gt;
大なり
&gt;=
大なりイコール
.
moduleやオブジェクトのメンバへのアクセス
## 演算式の優先順位#
優先順位の高いものから先に演算を行います
( ) 内の式
.
!
* , / , mod
+ , -
= (等価比較), == , &lt;&gt; , !=
and (L,Bを含む)
or (L,Bを含む), xor (L,Bを含む)
? : (三項演算子)
:=
代入系の演算子は順位判定とは別に代入処理判定を行っています
代入演算子
=
複合代入演算子
+=
-=
*=
/=
// 2つ目の = は代入ではなく比較になるので a にはboolが入る
a = b = c
// こういうのはダメ、演算中に代入はしない
a + b + c += d
例外として := による代入があります
:= による代入は式であり、変数に代入された値を返します
print n := 1 // 1 (代入した値が返る)
print n // 1
print 1 + 2 + ( n := 3 ) + 4 // 10 (代入した値が返り、その値で計算が行われる)
print n // 3
// 一度に複数の変数に値を代入することもできる
a = b := c := 10
print a // 10
print b // 10
print c // 10
// a := b := c := 10 でも可
## 特殊な演算#
数値以外を含む演算には一部特殊な仕様があります
型に対して不適切な演算子が用いられた場合はエラーになります
数値 + 文字列
右辺の文字列が数値変換可能な場合は数値にします
print 1 + &#39;2&#39; // 3
右辺の文字列が数値変換できない場合は左辺の数値を文字列にします
print 1 + &#39;a&#39; // 1a
数値とEMPTYの演算
EMPTYは0として扱われます
print 3 * EMPTY // 0
数値と真偽値の演算
TRUEは1、FALSEは0として扱われます
print 3 + TRUE // 4
文字列 + 数値
右辺の数値を文字列にします
print &#39;a&#39; + 3 // a3
print &#39;1&#39; + 2 // 12
文字列 * 数値
左辺の文字列が数値変換可能な場合数値にします
print &#39;2&#39; * 3 // 6
print &#39;123&#39; * 2 // 246
左辺の文字列が数値に変換できない場合、文字列を数値分繰り返します
print &#39;a&#39; * 3 // aaa
print &#39;xyz&#39; * 3 // xyzxyzxyz
文字列と数値の演算 (+, * 以外)
左辺の文字列が数値変換可能な場合は数値にします
print &#39;15&#39; / 3 // 5
左辺の文字列が数値変換できない場合はエラーになります
print &#39;a&#39; / 3 // エラー
文字列 + NULL
null文字(chr(0))を付け加えます
hoge = &quot;HOGE&quot; + NULL
print hoge // HOGE
print length ( hoge ) // 5
文字列 + その他の値
上記例以外の値型はすべて文字列として扱われます
&#39;a&#39; + TRUE // aTrue
配列 + 値
配列の末尾に値を追加します
print [ 1 , 2 , 3 ] + 4 // [1,2,3,4]
NULL * 数値
数値分連続したnull文字を返します
hoge = NULL * 5
print hoge // (なにも表示されない)
print length ( hoge ) // 5
空文字 == EMPTY
空文字とEMPTYの等価比較は常にFALSEです
UWSCとの挙動の差異について
UWSCでは以下のような挙動でした
dim a = EMPTY
print &quot;&quot; = a // True
print &quot;&quot; = EMPTY // False
空文字に対して EMPTY である変数は等価になりますが、リテラルでは非等価になっていました
同一であるべき式が異なる結果を返すのは不正なのでUWSCRではいずれもFALSEを返します
## 三項演算子#
条件式 ? 真で返す式 : 偽で返す式
式を評価しその真偽により値を返します
単行のIF文に似ていますが、三項演算子は値を返します
また、IF文とは異なり文を書くことができません
条件式について
三項演算子の条件式はオプションにより異なる判定を行います
詳しくは 条件式の判定 を参照してください
a = FALSE
print a ? &quot;a is TRUE&quot; : &quot;a is FALSE&quot; // a is FALSE
// 入れ子もできる
// fizzbuzz
for i = 1 to 100
print i mod 15 ? i mod 5 ? i mod 3 ? i : &quot;fizz&quot; : &quot;buzz&quot; : &quot;fizzbuzz&quot;
next
// 三項演算子では中に式しか書けない
// 例: print文を書いた場合
hoge ? print &quot;hoge is truthy&quot; : print &quot;hoge is falsy&quot; // エラー
## ビット演算子、論理演算子#
AND、OR、XORは両辺の値型により論理演算またはビット演算のいずれかを行っていました
UWSCRでは演算子が追加され論理演算およびビット演算を明示的に行うことができます
## 論理演算子#
AndL, OrL, XorL
真偽値を返します
両辺に不適切な値型が含まれる場合はエラーになります
// 両辺の真偽性を評価してから演算を行う
print true andl false // false
print true andl NOTHING // false
print NULL andl &#39;a&#39; // true
print 1 xorl [ 1 , 2 ] // false
## ビット演算子#
AndB, OrB, XorB
数値を返します
両辺に不適切な値型が含まれる場合はエラーになります
// 両辺を数値として評価してから演算を行う
print 3 andb 5 // 1
print 3 orb 5 // 7
print 3 xorb 5 // 6
print 1 andb &#39;1&#39; // 1
print 1 andb true // 1
## 条件式の判定#
以下の式で条件判定が行われます
if文における if 及び elseif の式
while-wend文における while の式
repeat-until文における until の式
三項演算子における ? の左辺の式
! 演算子の右辺の式
条件判定はオプションにより以下の三種類の方法で行われます
## 真偽性判定 (デフォルト)#
デフォルトでは条件式において単に真偽値( TRUE , FALSE )であることだけではなく、式の真偽性が評価されます
式の評価結果が以下となる場合は 偽 と判定されます
FALSE
EMPTY
0
NOTHING
長さ0の文字列
長さ0の配列
これら以外の値を取る場合は 真 となります
print NOTHING ? &#39;真&#39; : &#39;偽&#39; // 偽
print &quot;&quot; ? &#39;真&#39; : &#39;偽&#39; // 偽
print &quot;空ではない文字列&quot; ? &#39;真&#39; : &#39;偽&#39; // 真
print [ 1 , 2 , 3 ] ? &#39;真&#39; : &#39;偽&#39; // 真
print [] ? &#39;真&#39; : &#39;偽&#39; // 偽
UWSCとの差異による注意点
UWSCでは条件式において、式の評価結果が文字列であった場合にそれを数値( VAR_DOUBLE 相当)へと変換し、変換成功時は0または0以外による判定を行っていました
そのため、以下のような記述をした場合にUWSCとUWSCRで異なる結果となります
if &quot;0&quot; then
print &quot;UWSCRでは長さ1以上の文字列であるため真と判定される&quot;
else
print &quot;UWSCでは数値の0に変換され偽と判定される&quot;
endif
このような場合は val 関数を使ってください
if val ( &quot;0&quot; ) then
print &quot;未到達&quot;
else
print &quot;0なので偽と判定される&quot;
endif
あるいは後述する CONDUWSC オプションをご利用ください
## 真偽値のみ#
OPTION FORCEBOOL が指定されている場合は真偽値( TRUE , FALSE )を返す式のみが有効となります
OPTION FORCEBOOL
// 真偽値を返す式のみ有効となる
print TRUE ? &#39;真&#39; : &#39;偽&#39; // 真
print FALSE ? &#39;真&#39; : &#39;偽&#39; // 偽
print 1 == 1 ? &#39;真&#39; : &#39;偽&#39; // 真
print 1 &gt; 2 ? &#39;真&#39; : &#39;偽&#39; // 偽
// 以下は真偽値を返す式ではないためエラーとなる
print NOTHING ? &#39;真&#39; : &#39;偽&#39; // エラー
print &quot;&quot; ? &#39;真&#39; : &#39;偽&#39; // エラー
print &quot;空ではない文字列&quot; ? &#39;真&#39; : &#39;偽&#39; // エラー
print [ 1 , 2 , 3 ] ? &#39;真&#39; : &#39;偽&#39; // エラー
print [] ? &#39;真&#39; : &#39;偽&#39; // エラー
## UWSC方式#
OPTION CONDUWSC が指定されている場合はUWSCと同等の判定を行います
FORCEBOOL との併用はできず、 いずれも有効の場合 FORCEBOOL が優先されます
OPTION CONDUWSC
print NOTHING ? &#39;真&#39; : &#39;偽&#39; // 真
print &quot;&quot; ? &#39;真&#39; : &#39;偽&#39; // エラー
print &quot;123&quot; ? &#39;真&#39; : &#39;偽&#39; // 真
print &quot;0&quot; ? &#39;真&#39; : &#39;偽&#39; // 偽
print &quot;hoge&quot; ? &#39;真&#39; : &#39;偽&#39; // エラー
## コメント#
// 以降は行末までコメントです (構文解析されない)
// があった時点で行末扱いになります
a = 1
// a = a + 1
print a // 1 が出力される
## ダミーコメント#
//- を記述することでUWSCでは以降をコメント扱いにしますが、UWSCRではこの部分は無視されます
これによりUWSCと併用するスクリプトでUWSCRのみで使える構文や関数を記述することが可能となります
print 1
//- print 2
print 3 //- + 5
// print 4
結果
# UWSC
1
3
# UWSCR
1
2
8
## 行結合#
行末に _ を記述することで次の行と結合させます
a = 1 + 2 + _
3 + 4
print a // 10
## マルチステートメント#
; をつけることで複数の文を1行に記述できます
a = 1 ; a = a + 1 ; print a // 2
// UWSCでは IF行のマルチ宣言はNG エラー
if 123 then ; print &quot;condition is truthy&quot; ; else ; print &quot;condition is falsy&quot; ; endif
## 組み込み定数#
TRUE
true または 1
FALSE
false または 0
NULL
振る舞い未実装
EMPTY
空文字
NOTHING
オブジェクトがない状態
NaN
Not a number
## NaNについて#
NaN は NaN 自身を含めあらゆる値と等価ではありません
また NaN との大小の比較結果も必ず偽です
print NaN == NaN // False
print n == NaN // False (nは何かしらの値)
print NaN != NaN // True
print NaN &lt; n // False
print NaN &lt;= n // False
print NaN &gt; n // False
print NaN &gt;= n // False
## 16進数#
16進数リテラル表記は $ を使います
print $FF // 255
## 起動時パラメータ#
スクリプトにパラメータを付与した場合にそれらが PARAM_STR[] に格納されます
uwscr hoge.uws foo bar baz
// hoge.uws
for p in PARAM_STR
print p
next
# 結果
foo
bar
baz
## OPTION#
OPTION 設定名 [ = 値 ]
値が真偽値指定の場合は省略可能で、省略時はtrueになります
各OPTIONのデフォルト値は設定ファイルからも変更可能です
設定ファイルについては 設定ファイルの構成 を参照してください
OPTION EXPLICIT // explicit設定をtrueにする
OPTION SHORTCIRCUIT = FALSE // デフォルトtrueなのでfalseにする
OPTION EXPLICIT[=bool]
trueの場合未宣言の変数への代入を許可しない (初期値:false)
未宣言の変数への代入および複合代入が行われる場合に解析エラーとなります
OPTION SAMESTR[=bool]
文字列の比較等で大文字小文字を区別するかどうか (初期値:false)
OPTION OPTPUBLIC[=bool]
public変数の重複宣言を禁止するかどうか (初期値:false)
以下の場合に解析エラーとなります
同名のグローバル変数宣言を行ったとき
OPTION OPTPUBLIC
public p = 1
public p = 2 // エラー
hoge = procedure ()
public p = 3 // エラー
fend
procedure p ()
public p = 4 //エラー
fend
同一モジュール内で同名のpublic変数を宣言したとき
OPTION OPTPUBLIC
module m
public p = 1
public p = 2 // エラー
procedure m
public p = 3 // エラー
fend
public x = procedure ()
public p = 4 // エラー
fend
endmodule
OPTION OPTFINALLY[=bool]
tryで強制終了時にfinally部を実行するかどうか (初期値:false)
OPTION SPECIALCHAR[=bool]
trueで特殊文字(&lt;#CR&gt;など)や変数展開が行われなくなる (初期値: false)
OPTION SHORTCIRCUIT[=bool]
論理演算で短絡評価を行うかどうか (初期値:true)
このOPTIONはデフォルト有効です
UWSCとは違いデフォルトで OPTION SHORTCIRCUIT が有効になっています。
無効にするには以下を実行してください
OPTION SHORTCIRCUIT = FALSE
短絡評価とは
論理演算において左辺の評価のみで結果が確定する場合に右辺の評価を行いません
論理和(OR)の場合、左辺が真なら右辺によらず真なので右辺を評価しない
論理積(AND)の場合、左辺が偽なら右辺によらず偽なので右辺を評価しない
短絡評価が行われるのは以下の状況です
サンプルコード内では事前に以下が実行されているものとします
function t ( n )
result = true
print n + &quot;: &quot; + result
fend
function f ( n )
result = false
print n + &quot;: &quot; + result
fend
AndL演算子の左辺が偽となる値を取る場合
// t(2) は評価されない
print f ( 1 ) AndL t ( 2 )
// 1: False
// False
OrL演算子の左辺が真となる値を取る場合
// f(2) は評価されない
print t ( 1 ) OrL f ( 2 )
// 1: True
// True
ifなどの条件式にて、AndまたはAndL演算子の左辺が偽となる値を取る場合
// t(2) は評価されない
print f ( 1 ) and t ( 2 ) ? true : false
// 1: False
// False
ifなどの条件式にて、OrまたはOrL演算子の左辺が真となる値を取る場合
// f(2) は評価されない
print t ( 1 ) or f ( 2 ) ? true : false
// 1: True
// True
短絡評価におけるUWSCとの差異
ANDとORの複合条件でUWSCでは短絡評価が行われないケースがありましたが、UWSCRでは適切に短絡評価を行います
評価結果に影響はありません
// UWSCで短絡評価が行われない例
if f ( 1 ) and t ( 2 ) or t ( 3 ) then
print true
else
print false
endif
// UWSCR
// 1: False … f(1) and t(2) で短絡評価されfalse
// 3: True … false or t(3) は短絡評価されないのでt(3)も評価される
// True
// UWSC
// 1: False
// 2: True … 評価されてしまう
// 3: True
// True
OPTION NOSTOPHOTKEY[=bool]
注意
この設定は無効です
OPTION TOPSTOPFORM[=bool]
注意
この設定は無効です
OPTION FIXBALLOON[=bool]
吹き出しを仮想デスクトップを跨いで表示するかどうか (初期値:false)
OPTION DEFAULTFONT=&quot;name,n&quot;
ダイアログ等のフォント指定 (初期値:&quot;Yu Gothic UI,20&quot;)
OPTION POSITION=x,y
注意
この設定は無効です
OPTION LOGPATH=&quot;path&quot;
ログ保存フォルダを指定 (初期値:スクリプトのあるフォルダ)
存在するディレクトリを指定するとそこに uwscr.log を出力します
それ以外はログファイルのパスとして扱われます
OPTION LOGLINES=n
ログファイルの最大行数を指定 (初期値:400)
OPTION LOGFILE=n
ログファイルの出力方法 (初期値:1)
0: 通常のログ出力
1: ログ出力なし
2: 日時出力なし
3: 通常のログ出力 (標準で秒を含むため0と同じ)
4: 以前のログを破棄
それ以外: ログ出力なし
UWSCRのログ出力について
UWSCRはデフォルトではログを出力しません
ログを出力するには0, 2, 3, 4のいずれかを指定してください
OPTION DLGTITLE=&quot;title&quot;
ダイアログのタイトルを指定します (初期値:&quot;UWSCR - スクリプト名&quot;)
OPTION GUIPRINT[=bool]
TRUEにした場合print文実行時にコンソールではなくGUIに出力します
uwscr --window で実行されている場合はこの設定が強制的にtrueになります
OPTION FORCEBOOL[=bool]
TRUEにした場合if文やwhile, repeatの条件式がTRUEまたはFALSEしか受け付けなくなります
CONDUWSC と競合した場合はこちらが優先されます
OPTION FORCEBOOL
if TRUE then
print &quot;OK&quot;
endif
if 1 then
print &quot;↑はエラーになります&quot;
endif
OPTION CONDUWSC[=bool]
TRUEにした場合if文やwhile, repeatの条件式の判定方法がUWSCと同等になります
FORCEBOOL が有効な場合は無視されます
## def_dll#
DLL関数 (Win32 APIなど) を呼び出せるようにします
32bit版UWSCRでは32bitのDLL、64bit版では64bitのDLLに対応します
呼び出す関数の名前、引数の型、戻り値の型、dllのパスを指定します
dllパスは拡張子(.dll)を省略できます
別名を指定して本来の関数名ではなく別名で呼び出せるようにもできます
def_dll 関数名 ( 型名 , 型名 , ... ): 型名 : DLLパス
// 戻り値がvoidの場合省略できる
def_dll 関数名 ( 型名 , 型名 , ... ): DLLパス
// 配列引数指定
def_dll 関数名 ( 型名 [] ): 型名 : DLLパス
// 配列サイズ指定
def_dll 関数名 ( 型名 [ サイズ ] ): 型名 : DLLパス
// 参照渡し
def_dll 関数名 ( var 型名 ): 型名 : DLLパス
def_dll 関数名 ( ref 型名 ): 型名 : DLLパス
// 構造体
def_dll 関数名 ( { 型名 , ... } ): 型名 : DLLパス
// 関数名エイリアス
// dll関数に呼び出すための別の名前をつける
def_dll 別名 : 関数名 ( 型名 , ... ): 型名 : DLLパス
## 使用可能な型名#
以下の型を指定できます
一部の型はx86/x64でサイズが変わります
一部の型は引数定義、または戻り値定義でのみ指定可能です
文字列型に EMPTY , NULL , NOTHING を渡した場合はNULL文字として扱われます
型名
サイズ
詳細
対応する値型
引数
戻り型
備考
int, long
4
符号あり32ビット整数
数値
可
可
bool
4
符号あり32ビット整数
真偽値
可
可
uint, dword
4
符号なし32ビット整数
数値
可
可
float
4
単精度浮動小数点数
数値(小数)
可
可
double
8
倍精度浮動小数点数
数値(小数)
可
可
word
2
符号なし16ビット整数
数値
可
可
wchar
2
符号なし16ビット整数
文字(列)
可
可
byte
1
符号なし8ビット整数
数値
可
可
char
1
符号なし8ビット整数
文字(列)
可
可
boolean
1
符号なし8ビット整数
真偽値
可
可
longlong
8
符号あり64ビット整数
数値
可
可
string
可変
ANSI文字列のポインタ
文字列
可
pchar
可変
ANSI文字列のポインタ
文字列
可
wstring
可変
ワイド文字列のポインタ
文字列
可
pwchar
可変
ワイド文字列のポインタ
文字列
可
hwnd
可変
ウィンドウハンドル
数値
可
可
handle
可変
各種ハンドル
数値
可
可
pointer
可変
ポインタを示す数値(符号なし)
数値
可
可
struct
可変
ユーザー定義構造体のポインタ
構造体
可
callback
可変
コールバック関数のポインタ
ユーザー定義関数
可
コールバック関数の型定義 を行う
safearray
可変
SAFEARRAYのポインタ
配列
可
void
1
型がないことを示す
可
可変サイズについて
一部の数値型はOSのアーキテクチャによりそのサイズが変わります
x86: 4
x64: 8
hwnd , handle , pointer , size にデータ上の区別はありません
## 配列引数#
型名[] と記述することでその型に該当する値型の配列を渡せるようになります
型名[サイズ] のように配列サイズを数値または定数でしていすることで、そのサイズの配列を受けることを明示します
サイズ指定時は異なるサイズの配列を渡した場合エラーになります
サイズ未指定時は渡す配列のサイズは可変ですが、十分なサイズを確保してください
## 参照渡し#
var 型名 または ref 型名 で参照渡しになります
引数として変数を渡した場合、関数実行後にその変数の値が更新されます
配列引数も参照渡しできます
## 構造体#
引数として構造体のポインタを受ける場合に {型名, 型名, ...} と記述することでその構造体として値の受け渡しができるようになります
関数呼び出し時に型名に該当する値を渡す必要があります
構造体の場合は参照渡しにしなくても変数に値が返ります
def_dll GetCursorPos ({ long , long }) : bool : user32 . dll
dim x , y
// 呼び出し時は{}内に書いた型名の分引数を渡す必要がある
GetCursorPos ( x , y )
// 参照渡しとして記述しなくても引数が更新される
print [ x , y ]
## ポインタではない構造体#
構造体のポインタではなく構造体そのものを受ける関数の場合は {} 表記が使えません
その場合は {} を使わずメンバーの方を引数として直接記述します
// MonitorFromPointは引数としてPOINT構造体とDWORDを受けます
// POINT構造体は2つのLONGで構成されているため、以下のように記述できます
def_dll MonitorFromPoint ( long , long , dword ) : dword : user32
## ネストした構造体#
メンバが構造体のポインタである場合は {型名, 型名, {型名, ...}, ...} のようにネスト構造で表記します
メンバが構造体そのものである場合は子構造体メンバの型名を展開して記述します
typedef struct tagWINDOWPLACEMENT {
UINT length ;
UINT flags ;
UINT showCmd ;
POINT ptMinPosition ; // POINT構造体は long, long
POINT ptMaxPosition ;
RECT rcNormalPosition ; // RECT構造体は long, long, long, long
} WINDOWPLACEMENT ;
def_dll GetWindowPlacement ( hwnd , { uint , uint , uint , long , long , long , long , long , long , long , long }) : bool : user32 . dll
dim len , flags , cmd , minx , miny , maxx , maxy , left , top , right , bottom
len = 44
h = hndtoid ( getid ( &quot;hoge&quot; ))
print GetWindowPlacement ( h , len , flags , cmd , minx , miny , maxx , maxy , left , top , right , bottom )
## コールバック#
以下の書式でコールバック関数の引数と戻り値の型を定義します
// callback(型名, 型名, ...):型名
def_dll hoge ( callback ( dword , dword ) : bool ) : hoge . dll
// 戻り値型は省略可能
def_dll fuga ( callback ( int ) ) : fuga . dll
dll関数呼び出し時に対応したユーザー定義関数を渡します
function hoge_callback ( foo , bar )
result = foo &gt; bar
fend
hoge ( hoge_callback )
コールバック定義に使える型は以下の通りです
型名
サイズ
詳細
対応する値型
引数
戻り型
int, long
4
符号あり32ビット整数
数値
可
可
bool
4
符号あり32ビット整数
真偽値
可
可
uint, dword
4
符号なし32ビット整数
数値
可
可
float
4
単精度浮動小数点数
数値(小数)
可
可
double
8
倍精度浮動小数点数
数値(小数)
可
可
word
2
符号なし16ビット整数
数値
可
可
byte
1
符号なし8ビット整数
数値
可
可
boolean
1
符号なし8ビット整数
真偽値
可
可
longlong
8
符号あり64ビット整数
数値
可
可
hwnd
可変
ウィンドウハンドル
数値
可
可
handle
可変
各種ハンドル
数値
可
可
pointer
可変
ポインタを示す数値(符号なし)
数値
可
可
void
1
型がないことを示す
可
コールバック実行例
// 1: デバイスコンテキストハンドル
// 2: RECT構造体のポインタ、今回は使わないのでstructではなくpointerを指定
// 3: コールバック関数
// 1. モニタハンドル
// 2. デバイスコンテキストハンドル
// 3. モニタのRECTのポインタ
// 4. LPARAM
// 4: LPARAM
def_dll EnumDisplayMonitors ( handle , pointer , callback ( handle , handle , pointer , pointer ) : bool , pointer ) : bool : user32 . dll
// lparamとして渡される構造体
struct UserData
// モニタハンドルを入れる配列
handles : handle [ 10 ]
// ハンドル数
count : uint
endstruct
// UserData構造体を初期化
data = UserData ()
// 構造体アドレスをLPARAMとして渡す
lparam = data . address ()
// callbackにはコールバック関数として呼ばれるユーザー定義関数を渡す
EnumDisplayMonitors ( null , null , MonitorEnumProc , lparam )
for i = 0 to data . count - 1
handle = data . handles [ i ]
print &quot;モニタ&lt;#i&gt;: &lt;#handle&gt;&quot;
next
function MonitorEnumProc ( hmonitor , hdc , prect , lparam )
// lparamからUserData構造体を得る
data = UserData ( lparam )
// モニタハンドルを配列に入れる
data . handles [ data . count ] = hmonitor
// カウントを進める
data . count += 1
if data . count == length ( data . handles ) then
// 取得上限を超えたら終了する
result = false
else
// trueを返して次に進む
result = true
endif
fend
## 文字列型について#
以下の型は引数として文字列を受けますが、それぞれ性質が異なります
戻り値の場合、可能な限り文字列として返します
ANSI文字列は日本語環境であれば主にCP932です
型名
型詳細
引数として渡された場合の処理
参照渡しの場合
char
ANSI文字を示す符号なし8ビット整数値
数値に変換され渡される
文字として返る
wchar
Unicode文字を示す符号なし16ビット整数値
char[]
ANSI文字列を示す符号なし8ビット整数の配列
数値配列に変換され渡される
文字列として返る
wchar[]
Unicode文字列を示す符号なし16ビット整数の配列
string
ANSI文字列を示す符号なし8ビット整数配列のポインタ
別途数値配列を作成しそのポインタを渡す
作成された配列は関数実行後開放される
最初のNULL文字までを文字列として返る
wstring
Unicode文字列を示す符号なし16ビット整数配列のポインタ
pchar
ANSI文字列を示す符号なし8ビット整数配列のポインタ
NULL文字も含めた文字列として返る
pwchar
Unicode文字列を示す符号なし16ビット整数配列のポインタ
## DLL関数定義およびその呼び出し方の例#
// Win32のA関数ではstringかpcharを使う
def_dll MessageBoxA ( hwnd , string , pchar , uint ) : int : user32 . dll
// Win32のW関数ではwstringかpwcharを使う
def_dll MessageBoxW ( hwnd , wstring , pwchar , uint ) : int : user32 . dll
// 呼び出す際は単に文字列を渡すだけで良い
print MessageBoxA ( 0 , &#39;メッセージ&#39; , &#39;タイトル&#39; , 0 )
print MessageBoxW ( 0 , &#39;メッセージ&#39; , &#39;タイトル&#39; , 0 )
// 構造体定義は{}
def_dll SetWindowPlacement ( hwnd , { uint , uint , uint , long , long , long , long , long , long , long , long }) : bool : user32 . dll
id = getid ( &quot;メモ帳&quot; )
h = idtohnd ( id )
// 構造体を渡すときは定義した型の数だけ値を並べる
SetWindowPlacement ( h , 44 , 0 , 1 , 0 , 0 , 0 , 0 , 200 , 200 , 600 , 600 )
// 参照渡し
path = GET_CUR_DIR + &quot;\test.ini&quot;
writeini ( &quot;foo&quot; , &quot;foo&quot; , &quot;foo&quot; , path )
writeini ( &quot;bar&quot; , &quot;bar&quot; , &quot;bar&quot; , path )
writeini ( &quot;baz&quot; , &quot;baz&quot; , &quot;baz&quot; , path )
print path
def_dll GetPrivateProfileStringA ( string , string , string , var pchar , dword , string ) : dword : kernel32
buffer = NULL * 100
// bufferがpcharなのでNULLを含んだ文字列が返ってくる
print GetPrivateProfileStringA ( NULL , NULL , NULL , buffer , length ( buffer ) , path )
print split ( buffer , NULL )
def_dll GetPrivateProfileStringA ( string , string , string , var string , dword , string ) : dword : kernel32
buffer = NULL * 100
// bufferをstringにすると最初のNULL以前の文字列のみ返ってくる
print GetPrivateProfileStringA ( NULL , NULL , null , buffer , length ( buffer ) , path )
print buffer
// 構造体で値を受ける
// varは不要
def_dll GetCursorPos ({ long , long }) : bool : user32 . dll
dim x , y
print GetCursorPos ( x , y )
print [ x , y ]
// 構造体はそのサイズに合う配列でも代用可能
// varで渡す
def_dll GetCursorPos ( var long []) : bool : user32 . dll
dim point = [ 0 , 0 ] // long, long
print GetCursorPos ( point )
print point
// サイズを明示するとより安全
// def_dll GetCursorPos(var long[2]):bool:user32.dll
## 別名による呼び出し例#
本来のDLL関数名とは異なる名前でそのDLL関数を呼び出すことができます
例: MessageBoxWをMessageBoxという名前で呼び出す
def_dll MessageBox : MessageBoxW ( hwnd , wstring , wstring , uint ) : int : user32 . dll
print MessageBox ( 0 , &quot;別名呼び出しサンプル&quot; , &quot;テスト&quot; , 0 )
Win32 APIの GetKeyState 関数を登録した場合、組み込み関数の getkeystate と競合してしまうという問題がありました
この場合も別名を登録することで関数の使い分けが可能になります
// GetKeyStateWin32という別名でGetKeyState関数を登録
def_dll GetKeyStateWin32 : GetKeyState ( int ) : word : user32
print GetKeyStateWin32 // GetKeyState(int):word:user32 as GetKeyStateWin32
print GetKeyStateWin32 ( VK_RETURN ) // Win32のGetKeyStateが呼ばれる
print getkeystate ( VK_RETURN ) // 組み込み関数が呼ばれる
## 構造体#
def_dllのstruct型に渡す構造体を定義します
## 構造体定義#
struct 構造体名
メンバ名: 型名
メンバ名: 型名[サイズ]
メンバ名: var 型名
︙
endstruct
id: int のようにメンバ名と型名を指定します
メンバが配列の場合は buffer: byte[260] のようにメンバ名、型名に加えてサイズを示す数値または定数を [] 内に記述します
型名の前に var または ref キーワードを記述した場合そのメンバは指定した型のポインタとなります
型名には以下が利用可能です
型名
サイズ
詳細
対応する値型
サイズ指定時
int, long
4
符号あり32ビット整数
数値
bool
4
符号あり32ビット整数
真偽値
uint, dword
4
符号なし32ビット整数
数値
float
4
単精度浮動小数点数
数値
double
8
倍精度浮動小数点数
数値
word
2
符号なし16ビット整数
数値
wchar
2
符号なし16ビット整数
文字
文字列
byte
1
符号なし8ビット整数
数値
char
1
符号なし8ビット整数
文字
文字列
boolean
1
符号なし8ビット整数
真偽値
longlong
8
符号あり64ビット整数
数値
string
可変
ANSI文字列(char配列)へのポインタ
文字列
文字列バッファのサイズ
pchar
可変
ANSI文字列(char配列)へのポインタ
文字列
文字列バッファのサイズ
wstring
可変
ワイド文字列(wchar配列)へのポインタ
文字列
文字列バッファのサイズ
pwchar
可変
ワイド文字列(wchar配列)へのポインタ
文字列
文字列バッファのサイズ
hwnd
可変
ウィンドウハンドル
数値
handle
可変
各種ハンドル
数値
pointer
可変
ポインタを示す数値(符号なし)
数値
size
可変
サイズ可変の符号なし整数
数値
var 型名
ref 型名
可変
型名のポインタ
型に対応する値型
可変サイズについて
一部の数値型はOSのアーキテクチャによりそのサイズが変わります
x86: 4
x64: 8
hwnd , handle , pointer , size にデータ上の区別はありません
文字列型について
string, wstring, pchar, pwcharはそれぞれの文字列を示す数値配列へのポインタとなります
文字列型メンバに代入された文字列は内部で数値配列(バッファ)に変換され、そのポインタが構造体にセットされます
構造体定義時にサイズを指定した場合はバッファサイズは固定となり、そのサイズを超える文字列の代入はできません
サイズ指定がない場合のバッファサイズは代入した文字列により可変です
実際のバッファサイズは bufsize() メソッドで取得できます
文字列型メンバにNULLが代入された場合はバッファが削除され構造体にNULLポインタがセットされます
struct Hoge
fuga : wstring
piyo : wstring [ 260 ]
endstruct
hoge = Hoge ()
// バッファサイズを確認する
// 代入前は0が返る
print hoge . bufsize ( &quot;fuga&quot; ) // 0
print hoge . bufsize ( &quot;piyo&quot; ) // 0
// 代入後はバッファサイズが得られる
// サイズ未指定時は代入した文字列による
// サイズ指定時はサイズ固定
hoge . fuga = &quot;fugafuga&quot;
hoge . piyo = &quot;piyopiyo&quot;
print hoge . bufsize ( &quot;fuga&quot; ) // 9
print hoge . bufsize ( &quot;piyo&quot; ) // 260
// サイズ指定時はサイズを越える文字列は代入不可
// hoge.piyo = &quot;p&quot; * 500 // エラー
// NULLを代入するとバッファが削除され、構造体にはNULLポインタがセットされる
hoge . piyo = NULL
print hoge . piyo // EMPTY
メンバが構造体の場合
メンバが構造体のポインタである場合
メンバの型名をpointerとしそのメンバに構造体のアドレスを代入するか、メンバの値から構造体を得ます
def_dll WNetGetUniversalNameW ( wstring , long , struct , var long ) : long : mpr
// WNetGetUniversalNameWに渡す構造体
struct BufferStruct
puni : pointer // UNIVERSAL_NAME_INFOWのポインタが返る
name : wchar [ 260 ] // 文字列バッファ
endstruct
struct UNIVERSAL_NAME_INFOW
lpUniversalName : wchar [ 260 ]
endstruct
// 関数に渡す構造体を初期化
buf = BufferStruct ()
if WNetGetUniversalNameW ( &quot;Z:\hoge&quot; , 1 , buf , 260 ) == 0 then
// 構造体で受けたポインタでUNIVERSAL_NAME_INFOWを得る
uni = UNIVERSAL_NAME_INFOW ( buf . pUni )
print uni . lpUniversalName
endif
メンバが構造体そのものである場合
メンバとなる構造体を定義し、型名として構造体名を記述します
このようなネスト構造の場合は parent.child.member のように . を連結してメンバにアクセスできます
// POINTとRECTは構造体
typedef struct tagWINDOWPLACEMENT {
UINT length ;
UINT flags ;
UINT showCmd ;
POINT ptMinPosition ;
POINT ptMaxPosition ;
RECT rcNormalPosition ;
RECT rcDevice ;
} WINDOWPLACEMENT ;
typedef struct tagPOINT {
LONG x ;
LONG y ;
} POINT , * PPOINT , * NPPOINT , * LPPOINT ;
typedef struct tagRECT {
LONG left ;
LONG top ;
LONG right ;
LONG bottom ;
} RECT , * PRECT , * NPRECT , * LPRECT ;
struct POINT
x : long
y : long
endstruct
struct RECT
left : long
top : long
right : long
bottom : long
endstruct
struct WINDOWPLACEMENT
length : uint
flags : uint
showCmd : uint
ptMinPosition : POINT
ptMaxPosition : POINT
rcNormalPosition : RECT
rcDevice : RECT
endstruct
wp = WINDOWPLACEMENT ()
wp . ptMinPosition . x = 100
wp . ptMinPosition . y = 100
## 構造体の利用方法#
## 構造体の初期化#
構造体名() で構造体を初期化します
各メンバは0で初期化されます
struct Point
x : long
y : long
endstruct
dim p = Point ()
## 構造体メンバへのアクセス#
構造体.メンバ名 でメンバへアクセスします
dim p = Point ()
print p . x // 0
print p . y // 0
p . x = 100
p . y = 200
print p . x // 100
print p . y // 200
## 構造体のメソッド#
構造体は以下のメソッドを持ちます
size : 構造体のサイズを得る
address : 構造体のアドレスを得る
bufSize(メンバ名) : 文字列型メンバのバッファサイズを得る、文字列型以外は0
struct Hoge
foo : dword
bar : dword
baz : wstring
qux : wstring [ 260 ]
endstruct
dim h = Hoge ()
print h . size () // 24
print h . address () // アドレスを返す
// 代入していない場合は文字列バッファがないので0
print h . bufSize ( &quot;baz&quot; ) // 0
print h . bufSize ( &quot;qux&quot; ) // 0
// 代入後はバッファのサイズが返る
h . baz = &quot;baz&quot;
h . qux = &quot;qux&quot;
print h . bufSize ( &quot;baz&quot; ) // 4
print h . bufSize ( &quot;qux&quot; ) // 260
print h . bufSize ( &quot;foo&quot; ) // 0 ※文字列型じゃない場合も0
## ポインタから構造体を得る#
DLL関数が返す構造体のポインタから構造体にアクセスできます
// 第四引数にWTS_SESSION_INFO_1Wのポインタが返る
def_dll WTSEnumerateSessionsExW ( handle , var dword , dword , var pointer , var dword ) : bool : Wtsapi32
def_dll WTSFreeMemoryExW ( dword , pointer , dword ) : bool : Wtsapi32
struct WTS_SESSION_INFO_1W
ExecEnvId : dword
State : int
SessionId : dword
pSessionName : wstring
pHostName : wstring
pUserName : wstring
pDomainName : wstring
pFarmName : wstring
endstruct
// 構造体のアドレスと個数を得るための変数
dim ptr , cnt
dim size = length ( WTS_SESSION_INFO_1W )
if WTSEnumerateSessionsExW ( null , 1 , 0 , ptr , cnt ) then
for i = 0 to cnt - 1
// 構造体は連続しているため、構造体サイズ分のオフセットを加える
addr = ptr + i * size
// アドレスから構造体を得る
wsi = WTS_SESSION_INFO_1W ( addr )
print &quot;addr: &lt;#addr&gt;&quot;
print &quot;ExecEnvId : &quot; + wsi . ExecEnvId
print &quot;State : &quot; + wsi . State
print &quot;SessionId : &quot; + wsi . SessionId
print &quot;pSessionName : &quot; + wsi . pSessionName
print &quot;pHostName : &quot; + wsi . pHostName
print &quot;pUserName : &quot; + wsi . pUserName
print &quot;pDomainName : &quot; + wsi . pDomainName
print &quot;pFarmName : &quot; + wsi . pFarmName
print
next
// WTS_SESSION_INFO_1W構造体をすべて開放する
WTSFreeMemoryExW ( 2 , ptr , cnt )
endif
## スレッド#
## thread#
関数を別のスレッドで実行します
thread func ()
スレッドスコープで実行されます
(その中でさらに関数スコープに入ります)
グローバルスコープへのアクセスは可能
public, const, function/procedure, module/class
呼び出した関数内でエラーが発生した場合スクリプトが終了します
## タスク#
関数を非同期実行します
threadとは異なり関数が完了し次第戻り値を受け取れます
タスク関数
Task
WaitTask
構文
async
await
## async#
タスクを返す関数を宣言します
// function宣言の前に async キーワードを付与
async function 関数名 ()
fend
async function MyFuncAsync ( n )
sleep ( n )
result = &quot;&lt;#n&gt;秒待ちました&quot;
fend
task = MyFuncAsync ( 5 ) // resultの値ではなくタスクを返す
// 以下と同じ結果になります
function MyFuncAsync ( n )
sleep ( n )
result = &quot;&lt;#n&gt;秒待ちました&quot;
fend
task = Task ( MyFuncAsync , 5 )
## await#
async宣言した関数の終了を待ち、resultの値を得ます
async function MyFuncAsync ( n )
sleep ( n )
result = &quot;&lt;#n&gt;秒待ちました&quot;
fend
// MyFuncAsync()の処理が終了するまで待つ
print await MyFuncAsync ( 5 ) // 5秒待ちました
## with#
. 演算子の左辺(module名やオブジェクト)を省略できます
module foo
public bar = &#39;bar&#39;
procedure baz ()
fend
endmodule
with foo
print . bar // foo.bar
. baz () // foo.baz()
endwith
// ネストも可
module m
public p = &quot;m.p&quot;
function f ()
result = m2
fend
endmodule
module m2
public p = &quot;m2.p&quot;
endmodule
with m
print . p // m.p
with . f () // m.f() のwithでネスト
print . p // m2.p
endwith
print . p // m.p
endwith
## textblock#
複数行文字列の定数を定義します
textblock内での改行は &lt;#CR&gt; と同様です
特殊文字( &lt;#CR&gt; , &lt;#DBL&gt; , &lt;#TAB&gt; )はtextblock文の評価時に展開されます
textblock [ 定数名 ]
( 複数行文字列 )
endtextblock
定数名が省略された場合は複数行コメントとなり、スクリプトの一部として扱われません
(構文木が作られない)
// 定数hogeが作られる
textblock hoge
foo
bar
baz
endtextblock
// 定数省略時はコメント扱い
// 値を呼び出すことができない
textblock
ここはコメントです
endtextblock
## textblockex#
変数展開が可能なtextblockです
textblockex変数の評価時に展開されます
textblockex hoge
&lt;#fuga&gt;
endtextblock
fuga = 123
print hoge // 123
fuga = 456
print hoge // 456
## call#
他のスクリプトを取り込みます
call hoge . uws // 実行するスクリプトからの相対パス
call hoge // 拡張子のないファイルもOK、見つからない場合は.uwsを付けて開く
call fuga . uws ( 1 , 2 , 3 ) // 引数を渡すと PARAM_STR にはboolが入る
// urlから読み込み
call url [ https://example.com/hoge.uws ] // url[ ] の中でurlを指定
call url [ https://example.com/hoge.uws ]( 1 , 2 , 3 ) // url[ ] の後に()をつけて引数を渡せる
グローバル定義はスクリプト実行前に処理されます
public
const
textblock
function
procedure
module
class
それ以外の処理部分はcall文が呼ばれる際に実行されます
呼び出し元とは異なるスコープで実行されます
呼び出し元の PARAM_STR にはアクセスできません (独自の PARAM_STR を持つため)
## 同一ファイルの複数箇所呼び出し#
同じファイルを複数箇所でcallした場合、グローバル定義 (関数やmodule等) および宣言 (定数やpublic) は初回のみ処理されます
実行可能部分はcall呼び出しの都度処理されます
複数箇所呼び出しの例:
ファイルAから複数箇所でファイルBがcallされる場合
ファイルAからファイルB、Cをcallし、BとCのそれぞれからファイルDがcallされる場合
ファイルAと、ファイルAからcallされたファイルBのそれぞれからファイルCがcallされる場合
## uwslファイルの読み込み#
uwslファイルをcallして使えます
call mylib . uwsl // 拡張子はuwslのみ (省略不可)
## uwslファイルについて#
構文木をバイナリとして保存したものです
以下のコマンドでバイナリファイルを生成できます
ファイルはスクリプトと同じディレクトリに作成されます
拡張子は .uwsl になります
uwscr - -lib path \ to \ module . uws # module.uwsl が出力される
callでの呼び出しにのみ対応しており、直接実行することはできません
uwscr module . uwsl // ng
## uwslファイル作成の流れ#
指定されたスクリプトを読み出す
構文解析を行い構文木を生成する
構文木をバイナリデータとしてファイルに書き出す
## 使用例#
多段callしているファイルをまとめてバイナリ化
ファイル構成例
mylib.uws (module1 ~ 3 をcall)
module1.uws
module2.uws (submodule1, 2 をcall)
submodule1.uws
submodule2.uws
module3.uws
uwscr -l mylib . uws # mylib.uwslが出力される
uwslファイルをcallして使う
call mylib . uwsl
MyLib . DoSomething ()
Module1 . DoSomethingElse ()
Module2 . DoSomethingWithSubmodule ( Submodule1 . DoSomething )
## 例外処理#
try-except-endtry
try-finally-endtry
try-except-finally-endtry
try部で発生した実行エラーを抑制し、以下の特殊変数にエラー情報を格納します
TRY_ERRMSG : エラーメッセージ
TRY_ERRLINE : エラー行
except部はtryでエラーが発生した場合のみ実行されます
finally部は必ず実行されます
finally部では continue , break , exit が使えません (構文解析エラーになる)
try-except-finally-endtry は
try
try
except
endtry
finally
endtry
と同等です
## except例#
try
print 1
raise ( &quot;エラー&quot; ) // ここでエラー
print 2 // 実行されない
except
print TRY_ERRMSG // 実行される
endtry
try
// エラーが発生しない場合
except
print 1 // 実行されない
endtry
## finally例#
try
print 1
raise ( &quot;エラー&quot; ) // ここでエラー
print 2 // 実行されない
finally
print TRY_ERRMSG // 実行される
endtry
try
// エラーが発生しない場合
finally
print 1 // 実行される
endtry
## except-finally例#
try
print 1
raise ( &quot;エラー&quot; ) // ここでエラー
print 2 // 実行されない
except
print TRY_ERRMSG // 実行される
finally
print TRY_ERRMSG // 実行される
endtry
try
// エラーが発生しない場合
except
print 1 // 実行されない
finally
print 2 // 実行される
endtry
## 制御文#
説明文中の 式 とは主に値を返す演算式や関数など
文 は制御文のことです
ブロック文``は ``文 が複数行ある状態です
## if#
注釈
if と ifb が区別されません
どちらも同じものとして扱われます
## 単行if#
if 条件式 then 文 [ else 文 ]
条件式について
if文の条件式はオプションにより異なる判定を行います
詳しくは 条件式の判定 を参照してください
if foo then bar // foo が真の場合 bar が実行され、偽の場合なにもしない
if foo then bar else baz // foo が真の場合 bar、偽の場合 baz が実行される
// UWSCとは異なり ifb でもエラーにはならない
ifb foo then bar
## 複数行if#
if 条件式 [ then ]
ブロック文
[ elseif 条件式 [ then ]]
ブロック文
[ else ]
ブロック文
endif
条件式について
if及びelseifの条件式はオプションにより異なる判定を行います
詳しくは 条件式の判定 を参照してください
注釈
elseif は複数回記述できる
if foo then
// fooが真なら実行され偽ならなにもしない
endif
if foo then
// fooが真なら実行される
else
// fooが偽なら実行される
endif
if foo then
// fooが真なら実行される
elseif bar then
// fooが偽かつbarが真なら実行される
elseif baz then
// fooが偽かつbazが真なら実行される
else
// foobarbazいずれも偽なら実行される
endif
## for#
for 変数 = 式1 to 式2 [ step 式3 ]
ブロック文
next
式1 ～ 式3 はいずれも数値を返す必要があります
step 式3 が省略された場合 step 1 として扱われます
小数が渡された場合は整数に丸められます (UWSCとは仕様が異なります)
変数 に 式1 を代入した状態で ブロック文 を処理
変数 の値に 式3 を加算したものを再代入し ブロック文 を処理
変数 に 式2 を超える値が代入されたら終了
終了後も変数の値は維持されます
for i = 0 to 2
print i // 順に 0 1 2 が出力される
next
print i // 3
for i = 0 to 5 step 2
print i // 順に 0 2 4 が出力される
next
print i // 6
// stepは減算も可能
for i = 5 to 0 step - 1
print i
next
// ループ変数に代入した場合
for i = 0 to 0
print i // 0
i = 10
print i // 10
next
print i // 1
// UWSCでは小数が利用可能でしたがUWSCRでは整数値に変換されます
for i = 0 . 1 to 1 . 9 step 0 . 1 // 0.1 -&gt; 0, 1.9 -&gt; 2 に丸められます
next
## for-in#
for 変数 [, 位置 , 最終周フラグ ] in 式
ブロック文
next
式 は以下を返す必要があります
配列
連想配列
文字列
COMのコレクション
Iteratorを実装するRemoteObject
式 が返す値をその種類に応じて分解し 変数 に代入していきます
位置 に識別子(変数)を入れた場合、その識別子に位置(インデックス)番号を代入します
最終周フラグ に識別子(変数)を入れた場合、最終周であればTRUE、そうでなければFALSEをその識別子に代入します
// 文字列は1文字ずつ分解
for char in &quot;あいうえお&quot;
print char // あ い う え お が順に出力される
next
// 配列は各要素
for value in [ &quot;あ&quot; , &quot;い&quot; , &quot;う&quot; , &quot;え&quot; , &quot;お&quot; ]
print value // あ い う え お が順に出力される
next
// 連想配列はキーを返す
hashtbl hoge = HASH_SORT
hoge [ &quot;b&quot; ] = 2
hoge [ &quot;a&quot; ] = 1
hoge [ &quot;d&quot; ] = 3
hoge [ &quot;c&quot; ] = 4
for key in hoge
print key // a b c d の順に出力される
print hoge [ key ] // 1 2 3 4 の順に出力される
next
// 位置を得る
for n , i in [ 1 , 3 , 5 ]
print n // 1, 3, 5 と出力される
print i // 0, 1, 2 と出力される
next
// インデックスおよび最終周フラグを得る
for n , i , l in [ 1 , 3 , 5 ]
print n // 1, 3, 5 と出力される
print i // 0, 1, 2 と出力される
print l // False, False, True と出力される
next
// インデックス得ず最終周フラグのみ得る
// 識別子を省略する
for n , , l in [ 1 , 3 , 5 ]
print n // 1, 3, 5 と出力される
print l // False, False, True と出力される
next
## for-else-endfor#
for i = a to b
ブロック文
else
ブロック文
endfor
for a in b
ブロック文
else
ブロック文
endfor
forループをbreakで抜けなかった場合にelse句以降が実行されます
for i = 0 to length ( items ) - 1
if items [ i ] == target then
// 要素のいずれかがtargetと一致した場合break
target_found ()
break
endif
else
// いずれの要素もtargetと一致しない場合はbreakしないのでこちらが実行される
target_not_found ()
endfor
// for-inにも対応
for item in items
if item == target then
target_found ()
break
endif
else
target_not_found ()
endfor
// ループ内の処理が行われない場合でもelseが実行される
for i = 0 to - 1
print 1 // 実行されないため表示もされない
else
print 2 // 2と表示される
endfor
for a in []
print 3 // 実行されない
else
print 4 // 4と表示される
endfor
## while#
while 条件式
ブロック文
wend
条件式 が真である限り ブロック文 を繰り返し処理します
(ループ中に条件式を偽にしない限り無限ループする)
条件式について
while文の条件式はオプションにより異なる判定を行います
詳しくは 条件式の判定 を参照してください
a = TRUE
while a
a = DoSomething () // 偽値を返せばループ終了
wend
while false
// 式が偽なら何も実行されない
wend
while TRUE
print &quot;無限ループ&quot;
wend
## repeat#
repeat
ブロック文
until 条件式
条件式 が偽である限り ブロック文 を繰り返し処理します
(ループ中に式を真にしない限り無限ループする)
条件式について
repeat文の条件式はオプションにより異なる判定を行います
詳しくは 条件式の判定 を参照してください
a = false
repeat
a = DoSomething () // 真値を返せばループ終了
until a
repeat
// 式が真でも一度は必ず実行される
until TRUE
repeat
print &quot;無限ループ&quot;
until FALSE
## continue#
continue [ 式 ]
ループ文(for, while, repet)にてループの先頭に戻ります
式 は正の整数を指定します
省略した場合 式 は 1 として扱われます
多重ループで複数のループをcontinueしたい場合 式 に2以上(ループの数分)を指定します
for i = 0 to 2
print &quot;3回出力される&quot;
continue
print &quot;出力されない&quot;
next
a = 1
b = 1
while a &lt; 5
while TRUE
a = a + 1
continue 2
b = b + 1
wend
wend
print a // 5
print b // 1
## break#
break [ 式 ]
ループ文(for, while, repet)にてループを抜けます
式 は正の整数を指定します
省略した場合 式 は 1 として扱われます
多重ループで複数のループをbreakしたい場合 式 に2以上(ループの数分)を指定します
for i = 0 to 2
print &quot;1回だけ出力される&quot;
break
print &quot;出力されない&quot;
next
a = 0
repeat
repeat
repeat
break 3
a = a + 1
until FALSE
a = a + 1
until FALSE
a = a + 1
until FALSE
print a // 0
## select#
select 式
case 式
ブロック文
[case 式, 式 …]
ブロック文
[default]
ブロック文
selend
select式を評価し、その結果とcase式が一致した場合にそのcase以下のブロック文が処理されます
caseに , 区切りで式を複数指定した場合、いずれかが一致すればそのブロック文が処理されます
select式を評価し結果を得る
case式を評価しselect式の結果と比較
一致した場合: その下のブロック文を処理しselectを終了する
不一致の場合: 次のcaseまたはdefaultに進む
defaultに到達した場合必ずその下のブロック文を処理する
defaultがなくいずれのcaseにも一致しない場合なにも行わない
select hoge
case 1
// hogeが1なら実行される
case 2 , 3
// hogeが2か3なら実行される
case 3
// hogeが3でも上のcaseが該当してるので実行されない
default
// hogeが1～3以外なら実行される
selend
select hoge
default
// 必ず実行される
selend
select 1
case 2
// なにも実行されない
selend
## exit#
スクリプト本文に記述した場合 スクリプト実行を終了します
関数内に記述した場合 関数を抜けます
REPLで実行した場合 REPLを終了します
hoge () // 2 は出力されない
procedure hoge ()
print 1
exit
print 2
fend
## exitexit#
exitexit [ 数値 ]
数値を指定した場合はUWSCRの終了コードになります
(省略時は終了コード0)
## print#
評価した式を文字列として出力します
またそれをログファイルに記録します
print 式
## print文の出力#
標準ではコンソールウィンドウに対して出力が行われます
以下の場合はprintウィンドウに出力します
OPTION GUIPRINT が有効の場合
ウィンドウモードでUWSCRを実行している場合 ( uwscr.exe -w )
GUIビルドのUWSCRを実行している場合
## ログファイルへの書き出し#
ログ出力が有効になっている場合はprintした内容をログファイルに書き出します
## COMオブジェクト#
createoleobj() , getactiveoleobj() により取得可能
またCOMオブジェクトのプロパティやメソッドが別のCOMオブジェクトを返す場合もあります
## プロパティの取得#
COMオブジェクト.プロパティ名 でプロパティの値を取得できます
## プロパティ取得#
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws . CurrentDirectory // 現在のワーキングフォルダのパスが表示される
## インデックス指定#
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws . Environment . item [ &quot;windir&quot; ] // %SystemRoot%
print ws . Environment . item ( &quot;windir&quot; ) // () も可
## コレクションに対するインデックス指定#
COMオブジェクトがコレクションの場合、インデックス指定で要素を得られる
Item(i)の糖衣構文として実装されているため、Itemメソッドを持たない場合はエラーになる
ws = createoleobj ( &quot;WScript.Shell&quot; )
// ws.SpecialFoldersはIWshCollectionというコレクション
print ws . SpecialFolders [ 0 ] // いずれかの特殊フォルダのパスが表示される
print ws . SpecialFolders ( 0 ) // ()でもOK
// これらは以下と同じ
print ws . SpecialFolders . Item ( 0 )
## プロパティの変更#
プロパティに対して値を代入することでプロパティを変更できます
## 代入#
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws . CurrentDirectory // 元々のカレントディレクトリ
ws . CurrentDirectory = &quot;D:\Hoge&quot;
print ws . CurrentDirectory // D:\Hoge
## インデックス指定による代入#
excel = createoleobj ( &quot;Excel.Application&quot; )
excel . visible = TRUE
excel . Workbooks . Add ()
range = excel . ActiveSheet . Range ( &quot;A1:A2&quot; )
// A1に値を代入
range [ 1 ] . Value = &quot;hoge&quot;
print range [ 1 ] . Value // hoge
// A2にA1を代入
range [ 2 ] = range [ 1 ]
print range [ 2 ] . Value // hoge
## メソッドの実行#
COMオブジェクト.メソッド名([引数, 引数, ...]) でメソッドを実行できます
通常の引数に加え、名前付き引数( 名前 := 値 )や参照渡し( ref 変数 )が利用可能です
メソッドの()なし実行について
UWSCでは引数のない(または全て省略可能な)メソッドは () を付けなくても実行できましたが、UWSCRではこれをメソッド扱いしません
// ()省略サンプル
excel = createoleobj ( &quot;Excel.Application&quot; )
excel . Quit
UWSCではQuitメソッドを実行していましたが、UWSCRではQuitプロパティへのアクセス扱いとなります
Excel.ApplicationにはQuitプロパティが存在しないためエラーになります
メソッドとして実行する場合は必ず () を付けてください
// こうすればUWSCでもUWSCRでも正常に動作する
excel . Quit ()
## 名前なし引数#
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws . Popup ( &quot;テキスト&quot; , 0 , &quot;タイトル&quot; )
## 名前付き引数#
名前を指定してメソッドに引数を渡すことができます
引数名 := 値 と記述します
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws . Popup ( Text := &quot;テキスト&quot; , Title := &quot;タイトル&quot; )
// 名前なし引数との併記
// 名前なし引数は正しい位置に書く必要がある
print ws . Popup ( &quot;テキスト&quot; , Title := &quot;タイトル&quot; )
// 名前付き引数のあとに名前なしは書けない
print ws . Popup ( Text := &quot;テキスト&quot; , 0 , &quot;タイトル&quot; ) // エラー
## 参照渡し#
ref または var キーワードで参照渡しになります
// uwscr x86でのみ動作
sc = createoleobj ( &quot;ScriptControl&quot; )
sc . language = &quot;VBScript&quot;
sc . ExecuteStatement ( script )
dim n = 50
print sc . CodeObject . Hoge ( ref n ) // 50
print n // 100
textblock script
Function Hoge(ByRef n)
Hoge = n &#39;引数をそのまま返す
n = 100 &#39;引数の値を更新する
End Function
endtextblock
## 一部のWMIオブジェクトのメソッドについて#
一部WMIメソッドの注意点
一部のWMIオブジェクトのメソッドは通常のCOMオブジェクトのようなメソッド実行ができません
このようなメソッドに対しては内部で自動的にWMIオブジェクトのメソッド実行処理に切り替わります
この場合以下の制限があります
名前付き引数が利用できません
該当するWMIオブジェクトは以下になります
ISWbemObject
ISWbemObjectEx
以下は実行例です
dim hDefKey = $80000002
dim sSubKeyName = &quot;SOFTWARE\Microsoft\Windows NT\CurrentVersion\&quot;
dim sValueName = &quot;CurrentVersion&quot;
locator = CreateOleObj ( &quot;Wbemscripting.SWbemLocator&quot; )
service = locator . ConnectServer ( &quot;&quot; , &quot;root\default&quot; )
stdRegProv = service . Get ( &quot;StdRegProv&quot; )
print stdRegProv // ComObject(ISWbemObjectEx)
// stdRegProv.GetStringValueは通常の実行ができないためWMIメソッド処理で実行される
dim sValue
print stdRegProv . GetStringValue ( hDefKey , sSubKeyName , sValueName , ref sValue )
print sValue
// 上記のメソッド実行は以下のコードと同等の処理を内部的に行っています
inparam = stdRegProv . Methods_ . Item ( &quot;GetStringValue&quot; ) . InParameters . SpawnInstance_ ()
inparam . hDefKey = hDefKey
inparam . sSubKeyName = sSubKeyName
inparam . sValueName = sValueName
out = stdRegProv . ExecMethod_ ( &quot;GetStringValue&quot; , inparam )
print out . ReturnValue
print out . sValue
## 型の確認#
COMオブジェクトをprintすることで型の名前を確認できます
オブジェクトがコレクションの場合は 型名[] と表示されます
ws = createoleobj ( &quot;WScript.Shell&quot; )
print ws // ComObject(IWshShell3)
print ws . specialfolders // ComObject(IWshCollection[])
## COM_ERR_IGN-COM_ERR_RET#
COMエラーの発生を無視して処理を続行させることができます
COM_ERR_IGN でCOMエラーを抑制します
COM_ERR_RET でCOMエラーの抑制を解除します
COM_ERR_IGN から COM_ERR_RET の間でCOMエラーが発生した場合
実行時エラーで終了することなく処理を続行します
その際に COM_ERR_FLG が TRUE になります
COM_ERR_FLG は COM_ERR_IGN を呼んだ際に FALSE に初期化されます
COM_ERR_RET を呼んだ場合は値がそのまま維持されます
COM_ERR_IGN によるCOMエラーの抑制はスレッド単位で有効です
// 通常はCOMエラーで動作停止する
obj = createoleobj ( &quot;Some.ComObject&quot; )
obj . FireError () // COMエラー！
// COMエラーを抑制するパターン
obj = createoleobj ( &quot;Some.ComObject&quot; )
// COMエラー抑制開始
COM_ERR_IGN
print COM_ERR_FLG // False
obj . FireError () // エラーになるがスクリプトは停止しない
print COM_ERR_FLG // Trueになる
// COMエラー抑制終了
COM_ERR_RET
print COM_ERR_FLG // True; COM_ERR_RETでは初期化されない
obj . FireError () // 抑制していないのでCOMエラー
## リストの改行表記#
一部リスト表記 ( , 区切りの式) で改行を含めることができます
従来では _ による行連結が必要だったところを簡潔に記述できるようになりました
// 従来
hoge = [ _
&quot;foo&quot; , _
&quot;bar&quot; , _
&quot;baz&quot; _
]
// 0.5.0以降
hoge = [
&quot;foo&quot; ,
&quot;bar&quot; ,
&quot;baz&quot;
]
改行を含めることができる構文
配列リテラル
関数呼び出し時の引数
関数定義の引数
def_dllの引数型指定時
改行を含めることができない構文
dim配列定義
select文case句の複数条件
// 配列リテラル
print [ // [ の後の改行
&quot;foo&quot; , // カンマの後の改行
&quot;bar&quot; // 式の後の改行
, &quot;baz&quot; // カンマを式の前に書いてもいい
] // ] 前の改行
// 関数呼び出し
print func ( // ( の後の改行
foo , // カンマの後の改行
, // 引数省略
bar // 式の後の改行
, baz // カンマを式の前に書いてもいい
) // ) 前の改行
// 関数定義
function hoge (
a ,
ref b ,
c : string ,
d = 1
)
b = do_something_with ( a )
do_something_with ( c , d )
fend
// def_dll
def_dll MessageBoxA (
hwnd ,
string ,
string ,
uint
) : int : user32
// 以下は対象外
// dim配列宣言で改行するとエラーになる
dim fuga [] = 1 ,
2 ,
3
// 必ず一行で書く
dim hoge [] = 1 , 2 , 3
// select文case句の複数条件で改行するとエラー
select fuga
case 1 ,
2 ,
3
print &quot;ng&quot;
selend
// 必ず一行で書く
select hoge
case 1 , 2 , 3
print &quot;ok&quot;
selend
## 値型#
UWSCRは動的型付け言語であり、その値は状況により様々な型を持ちます
以下に値が取りうる型とその詳細を記します
各項目は以下を示します
型: 型の呼称
解説: 型についての解説
文字列変換時: 暗黙の型変換で文字列に変換された場合
_name_ のような表記はプレースホルダです
種別: 値であるか参照であるか
型一覧 #
型
解説
種別
文字列変換時
数値
double (倍精度浮動小数点型)
値
123 → &quot;123&quot;
文字列
文字列
値
そのまま
真偽値
TRUE/FALSE
値
&quot;True&quot; / &quot;False&quot;
配列
要素として異なる値型を格納できる
値
[1,&quot;a&quot;] → [1, a]
連想配列
連想配列
参照
{&quot;KEY1&quot;: value1, &quot;KEY2&quot;: value2}
無名関数
名前なしで定義されたfunction/procedure
値
anonymous function(_params_)
関数
名前ありで定義されたfunction/procedure
値
function: _name_(_params_)
非同期関数
async宣言した関数
値
function: _name_(_params_)
組み込み関数
組み込み(ビルトイン)関数
値
builtin: _name_()
モジュール
module定義
参照
module: _name_
クラス定義
class定義
値
class: _name_
クラスインスタンス
クラス名() で得られる
参照
instance of _name_
NULL
def_dllにおけるNULL文字(chr(0))、UObjectのnull値
値
\0 (chr(0))
EMPTY
空の値、場合により空文字や0として扱われる
値
&quot;&quot; (空文字)
NOTHING
空オブジェクト
値
NOTHING
正規表現
正規表現パターン
値
regex: _pattern_
UObject
UObject
値
JSON文字列
列挙型
enum定義
値
Enum: _name_
タスク
非同期に行われる処理
参照
Task [_state_]
DLL関数
def_dll定義
値
_name_(_params_):_rtype_:_path_
構造体定義
struct-endstruct で得られる、関数のように呼ぶと構造体を得られる
値
_name_ {_member_: _type_}
構造体
構造体定義名() で得られる
参照
_name_(_address_)
COMオブジェクト
createoleobj/getactiveoleobj
参照
ComObject(_type_or_address_)
Unknownオブジェクト
COMオブジェクトのプロパティ・メソッドが返す場合がある
参照
IUnknown(_address_)
VARIANT
値
値
VARIANT(_vt_)
BrowserBuilderオブジェクト
起動するブラウザを構成するためのオブジェクト
参照
BrowserBuilder
Browserオブジェクト
操作対象のブラウザを示すオブジェクト
値
Browser: _id_
TabWindowオブジェクト
ブラウザのタブを示すオブジェクト
値
TabWindow: _id_
RemoteObjectオブジェクト
Webページ上のJavaScriptオブジェクト
値
RemoteObject(_id_)
WebRequestオブジェクト
HTTPリクエストを構成するためのオブジェクト
参照
WebRequest
WebResponseオブジェクト
HTTPレスポンスを示すオブジェクト
値
_responsebody_
HtmlNodeオブジェクト
パースされたHTMLドキュメントを示すオブジェクト
値
_html_
ファイルID
fopen()が返す
参照
_filepath_(_detail_)
バイト配列
encode()で得られる
値
[1,2,3] → [1, 2, 3]

---

# 特殊変数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/syntax/special_variables.html

## 特殊変数#
環境や実行方法により変化する特殊な変数です
GET_UWSC_PRO
EMPTY
Tip
実行環境の判定方法
select GET_UWSC_PRO
case EMPTY
print &quot;UWSCRです&quot;
case TRUE
print &quot;UWSC Pro版です&quot;
case FALSE
print &quot;UWSC Free版です&quot;
selend
GET_UWSC_VER
GET_UWSCR_VER
UWSCRのバージョンを返す
PARAM_STR
起動時パラメータを格納した配列
GET_UWSC_DIR
GET_UWSCR_DIR
uwscr.exeのあるフォルダ
GET_UWSC_NAME
GET_UWSCR_NAME
スクリプトファイルの名前
GET_FUNC_NAME
ユーザー定義関数の名前
関数内でのみ有効(関数外では未定義)
無名関数の場合はEMPTY
GET_WIN_DIR
windowsフォルダのパス
GET_SYS_DIR
systemフォルダのパス
GET_APPDATA_DIR
appdataのパス
GET_CUR_DIR
カレントディレクトリ
G_MOUSE_X
マウスポインタのX座標
G_MOUSE_Y
マウスポインタのY座標
G_SCREEN_W
画面全体の幅
G_SCREEN_H
画面全体の高さ
G_SCREEN_C
色数(１ピクセルのビット数)
THREAD_ID
現在のスレッドのスレッド識別子
THREAD_ID2
現在のスレッドのスレッド識別子

---

# 引数の型について - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/about_parameters.html

## 引数の型について#
ビルトイン関数の引数の型について
## 文字列#
文字列型の引数は主に文字列型の値を受けることを想定していますが、文字列型以外の値を渡すこともできます
文字列型以外の値を渡した場合、その値に対して 暗黙の文字列変換 を行います
どのような変換が行われるかは 値型 の型一覧を参照してください
真偽値について
UWSCでは真偽値型( TRUE , FALSE )を文字列型引数に渡すと &quot;1&quot; または &quot;0&quot; へと変換されていましたが、UWSCRでは &quot;True&quot; または &quot;False&quot; に変換されます
そのためUWSCにて &quot;1&quot; または &quot;0&quot; を想定しているスクリプトをUWSCRで実行すると問題になる可能性があります
fid = fopen ( &quot;hoge.txt&quot; , F_READ or F_WRITE )
fput ( fid , TRUE , 1 ) // 暗黙の文字列変換が行われて書き込みが行われる
// 書き込んだ値を読む
value = fget ( fid , 1 )
print value // UWSC: 1, UWSCR: True
## 真偽値#
真偽値を受ける引数は主に真偽値型の値を受けることを想定していますが、真偽値型以外の値を渡すこともできる場合があります
真偽値型以外の値が渡せる場合はその値の真偽性により TRUE または FALSE への暗黙の変換が行われます
真偽性の判定については 条件式の判定 を参照してください
logprint ( &quot;hoge&quot; ) // 文字列型は真偽性の判定結果がtruthyであるためTRUEとして扱われる

---

# 文字列操作関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/text.html

## 文字列操作関数#
## コピー#
copy ( 対象文字列 , 開始位置 [ , コピー文字数=EMPTY ] ) #
文字列をコピーします
パラメータ :
対象文字列 ( 文字列 ) -- コピー元の文字列
開始位置 ( 数値 ) -- コピー開始位置 (1から)
コピー文字数 ( 数値 省略可 ) -- 開始位置からコピーする文字数、省略時は末尾まで
戻り値 :
コピーした文字列
サンプルコード
moji = &quot;あいうえおかきくけこ&quot;
print copy ( moji , 6 ) // かきくけこ
print copy ( moji , 3 , 4 ) // うえおか
print copy ( moji , 11 ) // (範囲外のため空文字)
betweenstr ( 対象文字列 [ , 前文字=EMPTY , 後文字=EMPTY , n番目=1 , 数え方=FALSE ] ) #
対象文字列から前文字列と後文字に挟まれた部分の文字列をコピーします
パラメータ :
対象文字列 ( 文字列 ) -- コピー元の文字列
前文字 ( 文字列 省略可 ) -- コピーしたい文字列の前にある文字列、省略時は対象文字列の先頭から
後文字 ( 文字列 省略可 ) -- コピーしたい文字列の後にある文字列、省略時は対象文字列の末尾まで
n番目 ( 数値 省略可 ) --
n番目に一致する前後文字の組み合わせからコピーする、マイナスの場合後ろから探す
0指定時: 前後文字があれば該当文字列をすべて取得、それ以外は1として扱う
数え方 ( 真偽値 省略可 ) -- n番目の数え方を指定します
TRUEかつ正順: n番目の前文字を探し、その後に対となる後文字を探す
TRUEかつ逆順: 後ろからn番目の後文字を探し、その後に対となる前文字を探す
FALSEかつ正順: 前文字とその直後の後文字をペアとし、そのn番目を探す (ペア中に別の前文字があっても無視される)
FALSEかつ逆順: 後ろから見て後文字とその直前の前文字をペアとし、そのn番目を探す (ペア中に別の後ろ文字があっても無視される)
前文字または後文字省略時は無視されます
前文字を指定し後文字を省略
前文字を省略し後文字を指定
の場合この引数は無視されます
その為UWSCとは結果が異なります
戻り値の型 :
文字列またはEMPTYまたは配列
戻り値 :
前文字のみ指定: n番目の前文字以降の文字列を返す、該当なしならEMPTY
後文字のみ指定: n番目の後文字までの文字列を返す、該当なしならEMPTY
前後文字指定
n番目が0: 該当する文字列すべてを配列で返す、該当なしの場合空配列
n番目が0以外: 該当する文字列を返す、該当なしならEMPTY
サンプルコード
str = &quot;abc?def!ghi?jkl!mno&quot;
// 前後文字未指定の場合はそのままコピー
print betweenstr ( str )
// abc?def!ghi?jkl!mno
// 前文字のみの場合は前文字後から末尾まで
print betweenstr ( str , &#39;abc&#39; )
// ?def!ghi?jkl!mno
// 後文字のみの場合は先頭から後文字の前まで
print betweenstr ( str , , &#39;jkl&#39; )
// abc?def!ghi?
// 前文字と後文字を指定するとその間
print betweenstr ( str , &#39;abc&#39; , &#39;jkl&#39; )
// ?def!ghi?
// n番目の指定
print betweenstr ( str , &#39;?&#39; , &#39;!&#39; , 1 )
// def
print betweenstr ( str , &#39;?&#39; , &#39;!&#39; , 2 )
// jkl
str = &quot;?aaa?bbb!ccc?ddd!eee&quot;
// 数え方指定
print betweenstr ( str , &#39;?&#39; , &#39;!&#39; , 2 , TRUE )
// bbb
print betweenstr ( str , &#39;?&#39; , &#39;!&#39; , 2 , FALSE )
// ddd
token ( 区切り文字 , var 元文字列 [ , 区切り方法=FALSE , ダブルクォート=FALSE ] ) #
区切り文字から手前の文字を切り出します
もとの文字は切り出された状態になります
パラメータ :
区切り文字 ( 文字列 ) -- 区切りとなる文字、文字列の場合それぞれの文字が区切りとなる
元文字列 ( 文字列 参照渡し ) -- 切り出される文字列、関数実行後に切り出された残りの文字列が戻ります
区切り方法 ( 真偽値 省略可 ) -- 区切り文字が連続していた場合の処理方法を指定
TRUE
連続した区切り文字を一つの区切りとして扱う
FALSE
区切り文字が連続していてもそれぞれの文字を区切りとする
ダブルクォート ( 真偽値 省略可 ) -- ダブルクォートで括られた文字列内で区切るかどうか
TRUE
ダブルクォートで括られている文字列内の区切り文字を無視
FALSE
ダブルクォートがあっても区切る
戻り値 :
切り出した文字列
サンプルコード
moji = &quot;あ-い-う-え-お&quot;
print token ( &quot;-&quot; , moji ) // あ
print moji // い-う-え-お
print token ( &quot;-&quot; , moji ) // い
print moji // う-え-お
// 連続するトークン
// FALSEは個別に区切る
moji = &quot;あいうabcえお&quot;
// a で区切る
print token ( &quot;abc&quot; , moji , FALSE ) // あいう
print moji // bcえお
// b で区切る
print token ( &quot;abc&quot; , moji , FALSE ) // (空文字)
print moji // cえお
// TRUEならまとめて区切る
moji = &quot;あいうabcえお&quot;
print token ( &quot;abc&quot; , moji , TRUE ) // あいう
print moji // えお
// 該当する区切りがない場合文字列全体が切り出される
moji = &quot;あいうえお&quot;
print token ( &quot;abc&quot; , moji ) // あいうえお
print moji // (空文字)
// ダブルクォート内の区切り
csv = &quot;&lt;#DBL&gt;foo,bar&lt;#DBL&gt;,baz&quot;
print token ( &quot;,&quot; , csv ) // &quot;foo
print csv // bar&quot;,baz
csv = &quot;&lt;#DBL&gt;foo,bar&lt;#DBL&gt;,baz&quot;
print token ( &quot;,&quot; , csv , , TRUE ) // &quot;foo,bar&quot;
print csv // baz
## 置換#
replace ( 対象文字列 , 置換対象 , 置換文字列 [ , 正規表現モード=FALSE ] ) #
chgmoj ( 対象文字列 , 置換対象 , 置換文字列 [ , 正規表現モード=FALSE ] ) #
マッチした文字列を指定文字列で置換します
正規表現による置換も可能
パラメータ :
対象文字列 ( 文字列 ) -- 対象となる文字列
置換対象 ( 正規表現 ) -- 置換する文字列、正規表現モードの場合は正規表現を示す文字列
置換対象 -- 正規表現オブジェクト (これを指定した場合必ず正規表現モードになる)
置換文字列 ( 文字列 ) -- 置換後の文字列
マッチ文字列に置換
正規表現モードでは以下が使用可能
$0 がマッチした文字列そのものに置換される
$1 以降はサブマッチ
正規表現モード ( 真偽値 省略可 ) --
正規表現による置換を行う場合は TRUE
置換対象に正規表現オブジェクトを渡した場合はこの値は無視される
正規表現モードの場合は大文字小文字が区別されます
正規表現モードでない場合は大文字小文字は区別されません
戻り値 :
置換された文字列
置換対象が対象文字列にマッチしなかった場合は対象文字列がそのまま返る
サンプルコード
// 正規表現モードの場合は大文字小文字が区別される
print replace ( &quot;aA&quot; , &quot;A&quot; , &quot;B&quot; ) // BB
print replace ( &quot;aA&quot; , &quot;A&quot; , &quot;B&quot; , TRUE ) // aB
// マッチ文字列を使った置換
print replace ( &quot;aa11bb22cc33&quot; , &quot;([a-z]+)(\d+)&quot; , &quot;$1 = $2, &quot; , TRUE )
// aa = 11, bb = 22, cc = 33,
## サイズ#
length ( 値 ) #
文字列の文字数、配列や構造体のサイズを返します
長さを返せない値が渡された場合はエラー
パラメータ :
値 ( 文字列他 ) -- 文字数を得たい文字列
戻り値 :
文字数やサイズを示す数値
対応する値型
文字列: 文字数を返す
数値: 文字列とみなし文字数を返す
真偽値: 文字列とみなし文字数を返す
EMPTY: 0
NULL: 1 (UWSC互換、0でないことに注意)
配列: 配列長を返す
連想配列: 配列長を返す
バイト配列: 配列長を返す
構造体定義: 構造体のサイズを返す
構造体: 構造体サイズを返す
RemoteObject: lengthプロパティを持つならその値
UObject (配列): 配列長を返す
UObject (オブジェクト): 子要素の数を返す
サンプルコード
print length ( &quot;あいうえお&quot; ) // 5
print length ([ 1 , 2 , 3 ]) // 3
// 構造体定義
struct Point
x : long // 4
y : long // 4
endstruct
print length ( Point ) // 8
p = Point () // 構造体インスタンスにも対応
print length ( p ) // 8
// UObject
obj = @{
&quot;foo&quot; : [ 1 , 2 , 3 ] ,
&quot;bar&quot; : 123
}@
print length ( obj . foo ) // 3 (配列の場合は配列要素数
print length ( obj ) // 2 (オブジェクトの場合は子オブジェクトの数
lengthb ( 文字列 ) #
文字列のバイト数(ANSI)を得ます
パラメータ :
文字列 ( 文字列 ) -- 長さを得たい文字列
戻り値 :
ANSIバイト数
lengthu ( 文字列 ) #
文字列のバイト数(UTF-8)を得ます
パラメータ :
文字列 ( 文字列 ) -- 長さを得たい文字列
戻り値 :
UTF8バイト数
lengths ( 文字列 ) #
サロゲートペアの文字を2文字分としてカウントします
パラメータ :
文字列 ( 文字列 ) -- 長さを得たい文字列
戻り値 :
サロゲートペアを2文字とした文字数
サンプルコード
str = &quot;森鷗外𠮟る&quot;
print length ( str ) // 5
print lengths ( str ) // 6
lengthw ( 文字列 ) #
NULL終端Unicode文字列としての長さを得ます
パラメータ :
文字列 ( 文字列 ) -- 長さを得たい文字列
戻り値 :
符号なし16ビット整数の配列長
## 正規表現#
NewRE ( 正規表現 [ , 大小文字=FALSE , 複数行=FALSE , 改行=FALSE ] ) #
正規表現オブジェクトを返します
パラメータ :
正規表現 ( 文字列 ) -- 正規表現を表す文字列
大小文字 ( 真偽値 省略可 ) -- 大文字小文字を区別するなら TRUE
複数行 ( 真偽値 省略可 ) --
複数行を対象とするなら TRUE
その場合 ^ が行頭、 $ が行末と一致する
改行 ( 真偽値 省略可 ) -- TRUE であれば . が \n にマッチするようになる
戻り値 :
正規表現オブジェクト
サンプルコード
print NewRe ( &quot;hoge&quot; , FALSE , TRUE , TRUE ) // regex: (?ima)hoge
regex ( 文字列 , 正規表現 [ , 操作方法=REGEX_TEST ] ) #
正規表現による様々な文字列操作を行います
TestRE , Match 及び replace の一部の機能を持ちます
パラメータ :
文字列 ( 文字列 ) -- 対象となる文字列
正規表現 ( 文字列または正規表現オブジェクト ) -- 正規表現を示す文字列またはオブジェクト
操作方法 ( 定数または文字列 省略可 ) -- 指定方法により結果が異なる
REGEX_TEST (定数)
文字列に正規表現がマッチするかを調べる、 詳しくは TestRE を参照
結果は真偽値で返る
REGEX_MATCH (定数)
正規表現にマッチした文字列を得る、 詳しくは Match を参照
結果は文字列の配列で返る
文字列
文字列の置換を行う
置換後の文字列を返す
戻り値 :
操作方法による
サンプルコード
target = &quot;abc123def&quot;
re = &quot;\d+&quot;
print regex ( target , re ) // True
print regex ( target , re , REGEX_TEST ) // True
print regex ( target , re , REGEX_MATCH ) // [123]
print regex ( target , re , &quot;456&quot; ) // abc456def
TestRE ( 文字列 , 正規表現 ) #
文字列に対し正規表現がマッチするかを調べます
RegEx(文字列, 正規表現, REGEX_TEST) と同等です
パラメータ :
文字列 ( 文字列 ) -- 対象となる文字列
正規表現 ( 正規表現 ) -- 正規表現文字列またはオブジェクト
戻り値 :
真偽値
Match ( 文字列 , 正規表現 ) #
正規表現にマッチした文字列を列挙します
RegEx(文字列, 正規表現, REGEX_MATCH) と同等です
パラメータ :
文字列 ( 文字列 ) -- 対象となる文字列
正規表現 ( 正規表現 ) -- 正規表現文字列またはオブジェクト
戻り値 :
配列
グループマッチをしない場合: 文字列の配列
各要素がマッチした文字列
グループマッチした場合: 文字列の二次元配列
各要素の1番目がマッチした全体の文字列、2番目以降はサブマッチした文字列
サンプルコード
// グループマッチなし
for m in match ( &quot;aa11bb22cc33&quot; , &quot;\d+&quot; )
print &quot;found: &quot; + m
next
// found: 11
// found: 22
// found: 33
// グループマッチなし
for matches in match ( &quot;aa11bb22cc33&quot; , &quot;([a-z]+)(\d+)&quot; )
print &quot;found: &quot; + matches [ 0 ]
if length ( matches ) &gt; 1 then
print &quot; submatches:&quot;
for i = 1 to length ( matches ) - 1
print &quot; &quot; + matches [ i ]
next
endif
next
// found: aa11
// submatches:
// aa
// 11
// found: bb22
// submatches:
// bb
// 22
// found: cc33
// submatches:
// cc
// 33
## 利用可能な正規表現#
こちら を参照してください
## JSON#
FromJson ( json ) #
json文字列をUObjectにします
パラメータ :
json ( 文字列 ) -- json文字列
戻り値 :
変換に成功した場合は UObject 、失敗時は EMPTY
サンプルコード
textblock json
{
&quot;foo&quot;: 1,
&quot;bar&quot;: 2
}
endtextblock
obj = fromjson ( json )
print obj . foo // 1
ToJson ( UObject [ , 整形=FALSE ] ) #
UObjectをjson文字列にします
パラメータ :
UObject ( UObject ) -- json文字列にしたいUObject
整形 ( 真偽値 省略可 ) -- TRUEならjsonを見やすい形式にする
戻り値 :
json文字列
サンプルコード
obj = @{
&quot;foo&quot; : 1 ,
&quot;bar&quot; : {
&quot;baz&quot; : 2
}
}@
print tojson ( obj )
// {&quot;bar&quot;:{&quot;baz&quot;:2},&quot;foo&quot;:1}
// 整形する
print tojson ( obj , TRUE )
// {
// &quot;bar&quot;: {
// &quot;baz&quot;: 2
// },
// &quot;foo&quot;: 1
// }
// 子オブジェクトも変換可能
print tojson ( obj . bar )
// {&quot;baz&quot;: 2}
// yaml由来のUObjectも変換可能
textblock yaml
hoge: abc
piyo:
foo: 1
bar: 1
endtextblock
obj = FromYaml ( yaml )
print ToJson ( obj )
// {&quot;hoge&quot;:&quot;abc&quot;,&quot;piyo&quot;:{&quot;foo&quot;:1.0,&quot;bar&quot;:1.0}}
yaml由来のUObject
## YAML#
FromYaml ( yaml ) #
yaml文字列をUObjectにします
パラメータ :
yaml ( 文字列 ) -- yaml文字列
戻り値 :
変換に成功した場合は UObject 、失敗時は EMPTY
サンプルコード
textblock yaml
hoge: abc
piyo:
foo: 1
bar: 1
endtextblock
obj = FromYaml ( yaml )
print obj . hoge // abc
print obj . piyo . foo // 1
ToYaml ( UObject ) #
UObjectをyaml文字列にします
パラメータ :
UObject ( UObject ) -- yaml文字列にしたいUObject
戻り値 :
yaml文字列
サンプルコード
textblock yaml
hoge: abc
piyo:
foo: 1
bar: 1
endtextblock
obj = FromYaml ( yaml )
print ToYaml ( obj )
// hoge: abc
// piyo:
// foo: 1
// bar: 1
// json由来のUObjectも変換可能
obj = @{
&quot;foo&quot; : 1 ,
&quot;bar&quot; : {
&quot;baz&quot; : 2
}
}@
print toyaml ( obj )
// foo: 1.0
// bar:
// baz: 2.0
## 検索#
pos ( 検索文字列 , 対象文字列 [ , n番目=1 ] ) #
対象文字列の何文字目に検索文字列があるかを得ます
パラメータ :
検索文字列 ( 文字列 ) -- 探す文字列
対象文字列 ( 文字列 ) -- 探される文字列
n番目 ( 数値 省略可 ) -- n番目に一致する位置を得る、マイナスの場合後ろから探す
戻り値 :
見つかった位置、見つからなかった場合0
サンプルコード
moji = &quot;ももほげもももほげももももほげもも&quot;
print pos ( &#39;ほげ&#39; , moji ) // 3
print pos ( &#39;ほげ&#39; , moji , 2 ) // 8
print pos ( &#39;ほげ&#39; , moji , 3 ) // 14
print pos ( &#39;ほげ&#39; , moji , - 1 ) // 14 後ろから
// 見つからない場合は0
print pos ( &#39;ほげ&#39; , moji , 4 ) // 0
print pos ( &#39;ふが&#39; , moji ) // 0
## 変換系#
chknum ( 値 ) #
与えられた値が数値に変換可能かどうかを調べる
パラメータ :
値 ( 値 ) -- 調べたい値
戻り値 :
数値に変換可能かどうかを示す真偽値
サンプルコード
for v in [ &quot;1&quot; , 2 , &quot;３&quot; , &quot;四&quot; , &quot;Ⅴ&quot; , TRUE , &quot;FALSE&quot; ]
print v + &quot;: &quot; + chknum ( v )
next
// 1: True
// 2: True
// ３: False
// 四: False
// Ⅴ: False
// True: True
// FALSE: False
val ( 文字列 [ , エラー値=-999999 ] ) #
文字列を数値に変換します
パラメータ :
文字列 ( 文字列 ) -- 数値に変換したい文字列
エラー値 ( 数値 省略可 ) -- 変換できなかった場合に返す数値
戻り値 :
成功時は変換された数値、失敗時はエラー値
サンプルコード
print val ( 1 ) // 1
print val ( &quot;2&quot; ) // 2
print val ( &quot;３&quot; ) // -999999
print val ( TRUE ) // 1
print val ( &quot;ほげ&quot; , 0 ) // 0
trim ( 対象文字列 [ , 全角空白=FALSE ] ) #
trim ( 対象文字列 , 除去文字列 )
対象文字列の両端にあるホワイトスペースおよび制御文字を除去します
パラメータ :
対象文字列 ( 文字列 ) -- トリム対象文字列
全角空白 ( 真偽値 省略可 ) -- TRUEにした場合は全角の空白もトリム対象になります
除去文字列 ( 文字列 ) -- ホワイトスペース・制御文字ではなく指定文字を除去します
戻り値 :
トリム後の文字列
サンプルコード
print trim ( &quot; abc &quot; )
// abc
// 改行なども含む
print trim ( &quot; &lt;#CR&gt; abc&lt;#TAB&gt; &quot; )
// abc
// 制御文字
print trim ( NULL * 3 + &#39;abc&#39; + NULL * 3 )
// abc
// 全角スペース
print trim ( &quot; 　abc　 &quot; )
// 第2引数省略時は全角空白=FALSEとなる
// 　abc
print trim ( &quot; 　abc　 &quot; , FALSE )
// 　abc
print trim ( &quot; 　abc　 &quot; , TRUE )
// abc
// 指定文字
// この場合 e, d, f のいずれかが連続していれば除去する
print trim ( &quot;edeffededdabcedfffedeeddedf&quot; , &quot;edf&quot; )
// abc
chr ( コードポイント ) #
Unicodeコードポイントから文字を得ます
パラメータ :
コードポイント ( 数値 ) -- Unicodeコードポイント
戻り値 :
該当する文字、なければ空文字
サンプルコード
print chr ( 128021 ) // 🐕
chrb ( バイトコード ) #
バイトコードからASCII文字を得ます
パラメータ :
バイトコード ( 数値 ) -- 0～255
戻り値 :
該当する文字、なければ空文字
asc ( 文字 ) #
文字からUnicodeコードポイントを得ます
パラメータ :
文字 ( 文字列 ) -- コードポイントを得たい文字 (文字列の場合最初の文字のみ)
戻り値 :
該当するUnicodeコードポイント、なければ0
サンプルコード
print asc ( &quot;🐕&quot; ) // 128021
ascb ( 文字 ) #
ASCII文字からバイトコードを得ます
パラメータ :
文字 ( 文字列 ) -- バイトコードを得たい文字 (文字列の場合最初の文字のみ)
戻り値 :
該当するバイトコード、なければ0
isunicode ( 対象文字列 ) #
文字列中にUnicode専用文字(ANSIにない文字)が含まれるかどうかを調べる
パラメータ :
対象文字列 ( 文字列 ) -- 調べたい文字列
戻り値 :
Unicode専用文字が含まれていればTRUE
サンプルコード
print isunicode ( &quot;森鴎外叱る&quot; ) // FALSE
print isunicode ( &quot;森鷗外𠮟る&quot; ) // TRUE
strconv ( 対象文字列 , 変換方法 ) #
文字列を変換します (大文字↔小文字、ひらがな↔カタカナ、全角↔半角)
指定方法で変換できない文字列はそのまま出力されます
パラメータ :
対象文字列 ( 文字列 ) -- 変換したい文字列
変換方法 ( 定数 ) -- 変換方法を以下の定数で指定
SC_LOWERCASE
小文字に変換
SC_UPPERCASE
大文字に変換
SC_HIRAGANA
ひらがなに変換
SC_KATAKANA
カタカナに変換
SC_HALFWIDTH
半角文字に変換
SC_FULLWIDTH
全角文字に変換
戻り値 :
変換された文字列
サンプルコード
print strconv ( &#39;あいうえお&#39; , SC_KATAKANA ) // アイウエオ
print strconv ( &#39;あいうえお&#39; , SC_HALFWIDTH ) // あいうえお
print strconv ( &#39;あいうえお&#39; , SC_KATAKANA or SC_HALFWIDTH ) // ｱｲｳｴｵ
print strconv ( &#39;カキクケコ&#39; , SC_HIRAGANA ) // かきくけこ
print strconv ( &#39;カキクケコ&#39; , SC_HALFWIDTH ) // ｶｷｸｹｺ
print strconv ( &#39;ｻｼｽｾｿ&#39; , SC_FULLWIDTH ) // サシスセソ
print strconv ( &#39;ｻｼｽｾｿ&#39; , SC_FULLWIDTH or SC_HIRAGANA ) // さしすせそ
print strconv ( &#39;abcde&#39; , SC_UPPERCASE ) // ABCDE
print strconv ( &#39;abcde&#39; , SC_UPPERCASE or SC_FULLWIDTH ) // ＡＢＣＤＥ
format ( 数値 , 幅 [ , 桁数=0 , 埋め方法=FMT_DEFAULT ] ) #
数値を指定方法でフォーマットした文字列を返します
パラメータ :
数値 ( 数値 ) -- フォーマットしたい数値
幅 ( 数値 ) -- フォーマット後の文字列幅
幅が入力値の桁を越えている場合、埋め方法に従い不足分を埋めます
桁数 ( 数値 省略可 ) -- 小数点以下の桁数、または変換方法を指定
1以上の数値
小数点以下を指定桁数に丸める
0
変換しない
-1
16進数に変換 (アルファベット大文字)
-2
16進数に変換 (アルファベット小文字)
-3
2進数に変換
埋め方法 ( 定数 省略可 ) -- 幅に対する不足分を埋める方法
FMT_DEFAULT
半角スペースで左埋め
FMT_ZERO
0で左埋め
FMT_RIGHT
半角スペースで右埋め
FMT_ZEROR
0で右埋め
戻り値 :
フォーマットされた文字列
サンプルコード
// 幅指定
print format ( 1 , 8 ) // &#39; 1&#39;
// 小数点
print format ( 1 , 8 , 2 ) // &#39; 1.00&#39;
// 丸め
print format ( 1 . 234 , 0 , 2 ) // 1.23
print format ( 1 . 235 , 0 , 2 ) // 1.24
// 16進数
print format ( 42 , 0 , - 1 ) // 2A
// 16進数 (小文字)
print format ( 42 , 0 , - 2 ) // 2a
// 2進数
print format ( 42 , 0 , - 3 ) // 101010
// 0埋め
print format ( 42 , 4 , - 1 , FMT_ZERO ) // 002A
// 右埋め
print format ( 1 , 8 , 0 , FMT_RIGHT ) // &#39;1 &#39;
// 右0埋め
print format ( 1 , 8 , 0 , FMT_ZEROR ) // &#39;10000000&#39;
format ( 文字列 , 幅 )
パラメータ :
文字列 ( 文字列 ) -- フォーマットしたい文字列
幅 ( 数値 ) -- フォーマット後の文字列幅
幅が元の文字列長を越えた場合、指定幅まで元の文字を繰り返します
戻り値 :
フォーマットされた文字列
サンプルコード
// 文字列をフォーマット
print format ( &quot;abc&quot; , 8 ) // abcabcab
print format ( &quot;1&quot; , 8 ) // 11111111
format ( 秒数 , 日時フォーマット文字列 [ , ミリ秒=FALSE , ロケール文字列=EMPTY ] )
パラメータ :
秒数 ( 数値 ) -- 2000/01/01からの秒数またはミリ秒数
日時フォーマット文字列 ( 文字列 ) --
日時形式を示すフォーマット文字列
変換される日時はローカルタイムゾーン準拠
時刻フォーマットの書式
2023/01/23 13:24:56を基準に書式の例を以下に記します
書式
出力
備考
%Y
2023
年(4桁)
%y
23
年(下4桁)
%m
01
月(左0埋め)
%d
23
日(左0埋め)
%F
2023-01-23
年-月-日
%H
13
時(左0埋め、24時間)
%I
01
時(左0埋め、12時間)
%M
24
分(左0埋め)
%S
56
秒(左0埋め)
%R
13:24
hh:mm
%R
13:24
時:分
%T
13:24:56
時:分:秒
%X
13時24分56秒
ローカル時刻表示(日本の場合)
%+
2023-01-23T13:24:56+09:00
ISO8601/RFC3339形式
詳細な書式一覧は このリンク から確認できます
ミリ秒 ( 真偽値 省略可 ) -- TRUEなら秒数をミリ秒として扱う
ロケール文字列 ( 文字列 省略可 ) -- ロケールを示す文字列、省略時は現在のロケール設定 (日本または日本以外) に従う
ローカル日時フォーマットを行う場合はロケール設定が影響します
ロケール文字列の指定方法
利用可能なロケールは Locale を参照してください
表記は原則
言語 かつ XX
言語_国 かつ xx_XX
言語_国&#64;修飾子 かつ xx_XX&#64;xxxx
ですが区切り文字と大文字小文字の表記揺れを許容します
区切り文字として利用可能なのは _ , - , &#64; です
POSIX : 原則通り
Posix : 大文字小文字の揺れを許容
ja_JP : 原則通り
ja-JP : 区切り文字にハイフンを許容
JA_jp : 大文字小文字の揺れを許容
aa_ER&#64;saaho : 原則通り
aa-ER-saaho : 区切り文字の揺れを許容
aa&#64;ER&#64;saaho : 区切り文字の揺れを許容
戻り値 :
フォーマットされた文字列
サンプルコード
// 日時フォーマット
timestamp = gettime ( , &quot;2023/04/01 10:10:10&quot; )
print format ( timestamp , &quot;%c&quot; ) // 2023年04月01日 10時10分10秒
print format ( timestamp , &quot;%c&quot; , , &#39;en_US&#39; ) // Sat 01 Apr 2023 10:10:10 AM +09:00
print format ( timestamp , &quot;%c&quot; , , &#39;POSIX&#39; ) // Sat Apr 1 10:10:10 2023
encode ( 元文字列 , 変換方式 ) #
文字列をエンコードします
パラメータ :
元文字列 ( 文字列 ) -- エンコードしたい文字列
変換方式 ( 定数 ) -- 変換方式を示す定数
CODE_URL
URLエンコードを行う
CODE_HTML
一部の記号等を文字実態参照にする ( &lt; → &amp;lt; )
CODE_BYTEARRAY
バイト配列(ANSI)にする
CODE_BYTEARRAYW
バイト配列(Unicode)にする
CODE_BYTEARRAYU
バイト配列(UTF8)にする
CODE_ANSI
CODE_UTF8
互換性のため定数は存在していますが、無視されます
上記以外
元の文字列が返されます
戻り値 :
変換方式による
decode ( 文字列 , 変換方式 ) #
decode ( バイト配列 , 変換方式 )
文字列またはバイト配列をデコードします
パラメータ :
文字列 ( 文字列 ) -- デコードする文字列
バイト配列 ( バイト配列 ) -- デコードするバイト配列
変換方式 ( 定数 ) -- 変換方式を示す定数
CODE_URL
URLエンコードされた文字列を元の文字列に戻す
CODE_HTML
文字参照を文字に戻す ( &amp;lt; → &lt; )
CODE_BYTEARRAY
バイト配列(ANSI)を文字列に戻す
CODE_BYTEARRAYW
バイト配列(Unicode)を文字列に戻す
CODE_BYTEARRAYU
バイト配列(UTF8)を文字列に戻す
CODE_UTF8
互換性のため定数は存在していますが、無視されます
上記以外
EMPTYが返されます
戻り値 :
デコードされた文字列、変換できない場合は元文字列またはEMPTYを返す

---

# ウィンドウ操作関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/window.html

## ウィンドウ操作関数#
## ID取得#
getid ( タイトル [ , クラス名=EMPTY , 待ち時間=1 ] ) #
ウィンドウを検索し、該当するウィンドウを示すIDを返します
見つからない場合やタイムアウトした場合-1を返します
パラメータ :
タイトル ( 文字列 ) -- 検索するウィンドウのタイトル (部分一致)
クラス名 ( 文字列 省略可 ) -- 検索するウィンドウのクラス名 (部分一致)
待ち時間 ( 数値 省略可 ) -- ウィンドウが見つからない場合のタイムアウト時間
戻り値の型 :
数値
戻り値 :
ウィンドウID、失敗時は -1
getid ( 定数 )
パラメータ :
定数 ( 定数 ) -- 以下の定数を指定
GET_ACTIVE_WIN
アクティブウィンドウ
GET_FROMPOINT_WIN
マウスカーソル下のウィンドウ
GET_FROMPOINT_OBJ
マウスカーソル下の子ウィンドウ
GET_LOGPRINT_WIN
Printウィンドウ
GET_BALLOON_WIN
GET_FUKIDASI_WIN
吹き出し
GET_THISUWSC_WIN
GET_CONSOLE_WIN
UWSCRを実行しているコンソールウィンドウのIDを返します
GET_FORM_WIN
GET_FORM_WIN2
未実装 (-1を返す)
戻り値の型 :
数値
戻り値 :
ウィンドウID
getallwin ( [ ID=EMPTY ] ) #
すべてのウィンドウのIDを得ます
特定のウィンドウIDを指定した場合、そのウィンドウの子要素を得ます
パラメータ :
ID ( 数値 省略可 ) -- 子要素を取得したいウィンドウのID
戻り値 :
ウィンドウIDの配列
特殊変数の廃止
UWSCとは異なり見つかったウィンドウの個数ではなくウィンドウIDの配列が返るようになりました
それに伴い特殊変数 ALL_WIN_ID は廃止されました
idtohnd ( ID ) #
ウィンドウIDからウィンドウハンドル値を得ます
パラメータ :
ID ( 数値 ) -- ウィンドウID
戻り値 :
ウィンドウハンドル値、該当ウィンドウがない場合は 0
hndtoid ( hwnd ) #
ウィンドウハンドル値からウィンドウIDを得ます
パラメータ :
hwnd ( 数値 ) -- ウィンドウハンドル値
戻り値 :
ウィンドウID、該当ウィンドウがない場合 -1
getctlhnd ( ID , アイテム名 [ , n番目=1 ] ) #
getctlhnd ( ID , メニュー定数 )
子ウィンドウ(ボタン等)のウィンドウハンドル値、またはメニューハンドルを得ます
パラメータ :
ID ( 数値 ) -- ウィンドウID
アイテム名 ( 文字列 ) -- 子ウィンドウのタイトルまたはクラス名 (部分一致)
メニュー定数 ( 定数 ) -- 以下のいずれかを指定
GET_MENU_HND
メニューハンドルを返す
GET_SYSMENU_HND
システムメニューハンドルを返す
n番目 ( 数値 省略可 ) -- n番目に該当するアイテムを探す
戻り値 :
ハンドル値
サンプルコード
id = getid ( &quot;ファイル名を指定して実行&quot; )
h1 = getctlhnd ( id , &quot;実行するプログラム名、&quot; ) // タイトルを部分一致
h2 = getctlhnd ( id , &quot;static&quot; , 2 ) // クラス名指定、2番目
assert_equal ( h1 , h2 ) // 一致
## ID0について#
ウィンドウIDを使う一部の関数が実行されると、その関数の対象となったウィンドウが ID0 に記憶されます
次に同様の関数が実行されると ID0 は上書きされます
サンプルコード
ctrlwin ( getid ( &quot;TEST&quot; ) , HIDE )
// getid(&quot;TEST&quot;)のウィンドウがID0に記憶される
ctrlwin ( 0 , SHOW ) // 同じウィンドウに対して実行される
## ウィンドウ操作#
clkitem ( ID , アイテム名 [ , CLK定数=0 , チェック指定=TRUE , n番目=1 ] ) #
ボタン等をクリックします
パラメータ :
ID ( 数値 ) -- 対象のウィンドウID
アイテム名 ( 文字列 ) -- クリックしたいボタンや項目の名前
CLK定数 ( 定数 省略可 ) -- クリックしたいアイテムの種類やクリックの方法を指定します
これらの定数は OR で連結することにより複数指定が可能
アイテム種別
アイテム種別が未指定の場合はすべての種別を検索します
( CLK_BTN or CLK_LIST or CLK_TAB or CLK_MENU or CLK_TREEVIEW or CLK_LISTVIEW or CLK_TOOLBAR or CLK_LINK と同等)
複数指定時の検索順は以下の通り
CLK_BTN
CLK_LIST
CLK_TAB
CLK_MENU
CLK_TREEVIEW
CLK_LISTVIEW
CLK_TOOLBAR
CLK_LINK
CLK_BTN
ボタン、チェックボックス、ラジオボタン、その他
CLK_LIST
リストボックス、コンボボックス
ヒント
複数選択可能なリストボックスでの複数項目指定
アイテム名をタブ文字 ( &lt;#TAB&gt; ) で区切るか、配列指定で複数選択できます
// foo, bar, bazを選択状態にする
clkitem ( id , &quot;foo&lt;#TAB&gt;bar&lt;#TAB&gt;baz&quot; , CLK_LIST ) // タブ文字区切り
clkitem ( id , [ &quot;foo&quot; , &quot;bar&quot; , &quot;baz&quot; ] , CLK_LIST ) // タブ文字区切り
CLK_TAB
タブ
CLK_MENU
メニュー
ヒント
アイテム名のパス指定
ファイル\保存 のように階層構造をパス表記することもできます
注意
CLK_API でのみ使用可能です
CLK_TREEVIEW
CLK_TREEVEW
ツリービュー
制限事項
UWSCR x86版では CLK_TREEVIEW or CLK_API によるクリック操作に制限があり、
x64のウィンドウに対するクリックが行えません
CLK_API 以外の方式を指定してください
ヒント
アイテム名は root\branch\leaf のように階層構造を表すパス形式も指定できます
CLK_UIA で未展開のツリーを展開してクリックするためにはパス形式を指定する必要があります
CLK_UIA で枝要素を指定した場合、枝が閉じていれば開き、開いていれば閉じます
CLK_LISTVIEW
CLK_LSTVEW
リストビュー、ヘッダ
ヒント
UWSCからの機能拡張
リストビュー行の一番左だけでなく、どの列のアイテム名でも指定できるようになりました ( CLK_API/CLK_UIA )
ヘッダ名を指定することでヘッダをクリックできるようになりました ( CLK_API/CLK_ACC/CLK_UIA )
複数行を選択できるようになりました ( CLK_UIA )
CLK_TOOLBAR
ツールバー
CLK_LINK
リンク
注意
CLK_APIによるリンククリックは未対応です
CLK_ACCをご利用ください
マウスボタン指定
マウスボタン指定があった場合はクリック方式に関わらずメッセージ送信(PostMessage)による疑似クリック処理が行われます
未指定の場合はクリック方式別の処理を行います
CLK_RIGHTCLK
右クリック
CLK_LEFTCLK
左クリック (CLK_RIGHTCLKと同時指定ならこちらが優先)
CLK_DBLCLK
ダブルクリック (CLK_LEFTCLKと同時指定で2回目のクリック)
クリック方式(API)
クリック方式が未指定の場合はすべての方式で検索を行います
( CLK_API or CLK_UIA or CLK_ACC と同等)
クリック方式が複数指定された場合の適用順は以下の通り
CLK_API
CLK_UIA
CLK_ACC
CLK_API
Win32 APIによる検索およびクリック
クリックは対象アイテムに応じたメッセージ処理を行います
CLK_ACC
アクセシビリティコントロールによる検索およびクリック
クリックはACCオブジェクトのデフォルトアクションを実行、または選択を行います
CLK_UIA
UI Automationによる検索およびクリック
オプション
CLK_BACK
バックグラウンド処理 (ウィンドウをアクティブにしない)
CLK_MOUSEMOVE
CLK_MUSMOVE
クリック位置にマウスを移動
CLK_SHORT
アイテム名の部分一致
未指定の場合は完全一致する必要があります
CLK_FROMLAST
逆順サーチ (CLK_ACC指定時のみ有効)
CLK_HWND
戻り値を対象アイテムのHWNDにする (0は対象不明)
チェック指定 ( 真偽値 省略可 ) --
チェックボックスやメニューの場合、チェックのオンオフを指定 (TRUEならチェックを入れる、FALSEならはずす)
3状態チェックボックスの場合、 2 を指定することでグレー状態にできます
それ以外のアイテムの場合FALSEだとクリック動作を行いません (対象が存在していればTRUEを返す)
3状態チェックボックスのサポート
CLK_APIとCLK_UIAのみ
CLK_ACCは3状態チェックボックスをサポートしません
CLK_UIA指定時の2の動作
2状態チェックボックスに対してCLK_UIAで2を指定した場合は、クリック操作が複数回行われますが元々の状態に戻ります
n番目 ( 数値 省略可 ) -- 同名アイテムの場合何番目をクリックするか
UWSCとは順序が異なる場合があります
実装の違いによりUWSCとは別の番号を指定しなければならない可能性があります
ご注意ください
戻り値 :
成功時TRUE、 CLK_HWND 指定時は対象のウィンドウハンドル値を返す
アイテム名の一致について
CLK_SHORT を指定しない場合アイテム名は完全一致する必要がありますが、ニーモニックがある場合はそれを無視することができます
&amp; の有無は問わない
(&amp;A) のように括弧で括られたニーモニックは括弧ごと無視できる
括弧以降にある文字も無視できる
// &amp;Button
clkitem ( id , &quot;&amp;Button&quot; ) // ok, &quot;&amp;&quot;を含めても一致する
clkitem ( id , &quot;Button&quot; ) // ok, &quot;&amp;&quot;がなくても一致
// ボタン(&amp;B)
clkitem ( id , &quot;ボタン(&amp;B)&quot; ) // ok
clkitem ( id , &quot;ボタン(B)&quot; ) // ok, &quot;&amp;&quot;は無視できる
clkitem ( id , &quot;ボタン&quot; ) // ok, 括弧ごと無視できる
// ボタン (&amp;B)
clkitem ( id , &quot;ボタン&quot; ) // ok, 括弧の前に半角スペースがあった場合それも無視できる
// 選択 (&amp;S)...
clkitem ( id , &quot;選択&quot; ) // ok, 括弧以降も無視できる
ctrlwin ( ID , コマンド定数 ) #
対象ウィンドウに命令コマンドを送信します
ID0 を更新します
パラメータ :
ID ( 数値 ) -- 対象ウィンドウ
コマンド定数 ( 定数 ) -- 実行したいコマンドを示す定数
CLOSE
ウィンドウを閉じる
CLOSE2
ウィンドウを強制的に閉じる
ACTIVATE
ウィンドウをアクティブにする
HIDE
ウィンドウを非表示にする
SHOW
ウィンドウの非表示を解除する
MIN
ウィンドウを最小化する
MAX
ウィンドウを最大化する
NORMAL
ウィンドウを通常サイズに戻す
TOPMOST
ウィンドウを最前面に固定する
NOTOPMOST
ウィンドウの最前面固定を解除
TOPNOACTV
ウィンドウを最前面に移動するがアクティブにはしない
戻り値 :
なし
sckey ( ID , キー [ , キー , ... ] ) #
ショートカットキーを送信します
パラメータ :
ID ( ウィンドウID ) -- アクティブにするウィンドウのID、0指定でどのウィンドウもアクティブにしない
キー ( 定数または文字列 ) -- 仮想キーコード一覧 のいずれかまたはアルファベット一文字、35個まで
修飾子キー指定について
VK_SHIFT , VK_CTRL , VK_ALT , VK_WIN は押し下げられた状態になります (Rも含む)
これらのキーはすべてのキー入力が終了したあとにキーアップ状態に戻ります
戻り値 :
なし
setslider ( ID , 値 [ , n番目=1 , スクロール=FALSE ] ) #
スライダー(スクロールバー、トラックバー)の値を設定します
パラメータ :
ID ( ウィンドウID ) -- 対象ウィンドウのID
値 ( 数値 ) -- スライダーに設定する値
範囲外指定時の動作
最大値を上回る値だった場合は最大値に、最小値を下回る値だった場合は最小値に変更されます
n番目 ( 数値 省略可 ) -- n番目のスライダーを設定する
スクロール ( 真偽値 省略可 ) -- TRUEならスクロールバーを少しずつ動かす
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗または操作不能時はFALSE
sendstr ( ID , 文字列 [ , n番目=0 , 送信モード=FALSE , ACC指定=FALSE ] ) #
エディットボックスに文字列を送信します
パラメータ :
ID ( 数値 ) -- 対象ウィンドウのID
0ならクリップボードに送信 (その場合n番目、送信モード、ACC指定は無視されます)
文字列 ( 文字列 ) -- 送信する文字列
n番目 ( 数値 ) -- n番目のエディットボックスに送信
0ならフォーカスされたエディットボックス (対象ウィンドウは必ずアクティブになる)
UWSCとは順序が異なる場合があります
実装の違いによりUWSCとは別の番号を指定しなければならない可能性があります
ご注意ください
送信モード ( 真偽値または数値 ) --
FALSE または 0
追記
TRUE または 1
置き換え
2
一文字ずつ送信
ACC時は無視されます (TRUE扱い)
ACC指定 ( 真偽値または定数 ) --
FALSE または 0
APIまたはUIAを使用
ヒント
APIで検索を行い該当するものがなかった場合はUIAでの検索を試みます
UIA使用時は送信モードは無視され、常に置き換えられます
TRUE または 1
ACCを使用
STR_ACC_CELL (5)
DataGridView内のCell値の変更 (ACCを使用)
STR_UIA (6)
UIAを使用
送信モードは無視され、常に置き換えられます
UWSCとの違い
TRUEでも対象ウィンドウをアクティブにしないため、2は廃止されました
戻り値 :
なし
mouseorg ( ID [ , 基準=MORG_WINDOW , 画面取得=MORG_FORE , HWND=FALSE ] ) #
以下の関数にて座標の始点(0, 0)を特定のウィンドウ基準とする
mmv
btn
ChkImg (指定座標及び戻り値の座標)
chkclr (指定座標及び戻り値の座標)
peekcolor
MORG_DIRECT を指定した場合は以下も対象となる
kbd
パラメータ :
ID ( 数値 ) -- ウィンドウID または HWND
該当するIDが存在しない場合は失敗となるが、基準に MORG_DIRECT が指定されている場合はこの値をHWNDとして扱う
IDまたはHWNDに該当する有効なウィンドウが存在しない場合は失敗となる
0 が指定された場合はスクリーン座標基準に戻す (この場合以下の引数は無視される)
基準 ( 定数 省略可 ) -- 座標の始点を指定する
MORG_WINDOW (0)
対象ウィンドウのウィンドウ領域左上を基準とする
MORG_CLIENT
対象ウィンドウのクライアント領域左上を基準とする
MORG_DIRECT
対象ウィンドウのクライアント領域左上を基準とする
また mmv , btn 及び kbd 関数のマウス・キー操作をウィンドウに直接送信( SendMessage )する
送信するメッセージは以下 (対象ウィンドウがこれらのメッセージを処理しない場合操作は無効となる)
関数
操作
メッセージ
mmv
カーソル移動
WM_MOUSEMOVE
btn
左ボタン下げ (LEFT, DOWN)
WM_LBUTTONDOWN
btn
左ボタン上げ (LEFT, UP)
WM_LBUTTONUP
btn
右ボタン下げ (RIGHT, DOWN)
WM_RBUTTONDOWN
btn
右ボタン上げ (RIGHT, UP)
WM_RBUTTONUP
btn
中央ボタン下げ (MIDDLE, DOWN)
WM_MBUTTONDOWN
btn
中央ボタン上げ (MIDDLE, UP)
WM_MBUTTONUP
btn
マウスホイール回転(縦) (WHEEL)
WM_MOUSEWHEEL
btn
マウスホイール回転(横) (WHEEL2)
WM_MOUSEHWHEEL
kbd
キー下げ
WM_KEYDOWN
kbd
キー上げ
WM_KEYUP
kbd
文字送信(1文字ずつ)
WM_CHAR
btnについて
btn関数では各メッセージを送る前に以下のメッセージが送信されます
WM_MOUSEMOVE (x+1, y+1)
WM_MOUSEMOVE (x-1, y-1)
WM_MOUSEMOVE (x, y)
TOUCH非対応
btn関数でTOUCH指定時のMORG_DIRECTは無視されMORG_CLIENTとして動作します
画面取得 ( 定数 省略可 ) -- 画面取得方法を指定する
MORG_FORE
スクリーン上から画像を取得する ( ChkImg )、または色を得る ( peekcolor )
MORG_BACK
対象ウィンドウから直接画像の取得 ( ChkImg )、または色の取得 ( peekcolor ) を試みる
他のウィンドウに隠れている場合でも使用可能
動作しない場合
対象ウィンドウによっては正常に動作しない可能性があります
例: saveimgのIMG_BACKで画像が保存できないウィンドウ
CHKIMG_USE_WGCAPI指定時
chkimgでGraphicsCaptureAPI利用時にこれらのオプションは影響しません
ウィンドウの位置を問わずウィンドウ画像を取得します
HWND ( 真偽値またはEMPTY 省略可 ) -- MORG_DIRECT 指定時の第一引数の振る舞いを限定します ( MORG_DIRECT 以外の場合無視される)
FALSE
第一引数をIDとしますが、有効なIDが登録されていない場合はその値をHWNDとして扱います
例
30000 を指定
ID30000が登録済み→該当ウィンドウを対象とする
ID30000が未登録→HWNDが30000のウィンドウを対象とする
TRUE
第一引数をHWNDとして扱います
戻り値の型 :
真偽値
戻り値 :
成功した場合TRUE、失敗時はFALSE
サンプルコード
// MORG_DIRECTのHWND指定
id = getid ( hoge )
hnd = getctlhnd ( id , class_name )
// このとき hnd の値がいずれかの登録済みIDと一致してしまった場合は予期せぬ動作となる
mouseorg ( hnd , MORG_DIRECT )
// MORG_DIRECTかつ第四引数をTRUEにした場合hndはHWNDとして扱われる
mouseorg ( hnd , MORG_DIRECT , , TRUE )
chkmorg ( ) #
mouseorgで基準点となっているスクリーン座標を得る
戻り値の型 :
数値配列またはEMPTY
戻り値 :
基準点が変更されている場合は [x, y]、変更されていない場合はEMPTY
サンプルコード
mouseorg ( id )
print chkmorg () // [x, y]
mouseorg ( 0 )
print chkmorg () // EMPTY
## ウィンドウ情報取得#
status ( ID , ST定数 [ , ST定数... ] ) #
対象ウィンドウの各種状態を取得します
パラメータ :
ID ( 数値 ) -- ウィンドウID
ST定数 ( 定数 ) -- 取得したい状態を示す定数を指定
定数は最大21個指定できます
ST_TITLE
ウィンドウタイトル (文字列)
ST_CLASS
ウィンドウクラス名 (文字列)
ST_X
ウィンドウ左上のX座標 (数値)
ST_Y
ウィンドウ左上のY座標 (数値)
ST_WIDTH
ウィンドウの幅 (数値)
ST_HEIGHT
ウィンドウの高さ (数値)
ST_CLX
ウィンドウのクライアント領域左上のX座標 (数値)
ST_CLY
ウィンドウのクライアント領域左上のY座標 (数値)
ST_CLWIDTH
ウィンドウのクライアント領域の幅 (数値)
ST_CLHEIGHT
ウィンドウのクライアント領域の高さ (数値)
ST_PARENT
親ウィンドウのID (数値)
ST_ICON
最小化してればTRUE (真偽値)
ST_MAXIMIZED
最大化してればTRUE (真偽値)
ST_VISIBLE
ウィンドウが可視ならTRUE (真偽値)
ST_ACTIVE
ウィンドウがアクティブならTRUE (真偽値)
ST_BUSY
ウィンドウが応答なしならTRUE (真偽値)
ST_ISID
ウィンドウが有効ならTRUE (真偽値)
ST_WIN64
プロセスが64ビットかどうか (真偽値)
ST_PATH
プロセスの実行ファイルのパス (文字列)
ST_PROCESS
プロセスID (数値)
ST_MONITOR
ウィンドウが表示されているモニタ番号 ( monitor 関数に対応) (数値)
ST_WX
ウィンドウの補正なしX座標
ST_WY
ウィンドウの補正なしY座標
ST_WWIDTH
ウィンドウの補正なし幅
ST_WHEIGHT
ウィンドウの補正なし高さ
ST_ALL
すべての状態を取得
この定数を指定する場合ほかの定数は指定できません
戻り値 :
ST定数を一つだけ指定した場合は得られた値、複数指定時または ST_ALL 指定時は連想配列 (キーはST定数)
サンプルコード
id = getid ( &quot;uwsc&quot; , &quot;HH&quot; ) // uwscヘルプファイル
stat = status ( id , ST_TITLE , ST_CLASS , ST_HEIGHT , ST_WIDTH )
print stat [ ST_TITLE ] // uwsc
print stat [ ST_CLASS ] // HH Parent
print stat [ ST_HEIGHT ] // 778
print stat [ ST_WIDTH ] // 1251
monitor ( モニタ番号 [ , MON定数=MON_ALL ] ) #
モニタの情報を得ます
パラメータ :
モニタ番号 ( 数値 省略可 ) -- モニタを示す番号 (0から)
MON定数 ( 定数 省略可 ) -- 取得したい情報を示す定数
MON_X
モニタのX座標 (数値)
MON_Y
モニタのY座標 (数値)
MON_WIDTH
モニタの幅 (数値)
MON_HEIGHT
モニタの高さ (数値)
MON_PRIMARY
MON_ISMAIN
プライマリ(メイン)モニタならTRUE (真偽値)
MON_NAME
モニタ名 (文字列)
MON_WORK_X
作業エリアのX座標 (数値)
MON_WORK_Y
作業エリアのY座標 (数値)
MON_WORK_WIDTH
作業エリアの幅 (数値)
MON_WORK_HEIGHT
作業エリアの高さ (数値)
MON_DPI
画面のDPI
MON_SCALING
スケーリング倍率 (%)
MON_ALL
上記すべて (連想配列、キーはMON定数)
戻り値 :
定数指定 ( MON_ALL 以外): 得られた値
MON_ALL 指定: 連想配列 (キーはMON定数)
該当モニタなし: FALSE
monitor ( )
(引数なし) モニタの数を得ます
戻り値 :
モニタの数
サンプルコード
// すべてのモニタのサイズを表示
for i = 0 to monitor () - 1
m = monitor ( i , MON_ALL )
print &quot;モニタ&quot; + i + &quot;: &quot; + m [ MON_NAME ]
print m [ MON_X ] + &quot;, &quot; + m [ MON_Y ]
print m [ MON_WIDTH ] + &quot; x &quot; + m [ MON_HEIGHT ]
next
posacc ( ID , クライアントX座標 , クライアントY座標 [ , 種別=0 ] ) #
座標位置のアクセシビリティオブジェクトから情報を得ます
パラメータ :
ID ( ウィンドウID ) -- 対象ウィンドウのID
クライアントX座標 ( 数値 ) -- 対象ウィンドウのクライアント領域におけるX座標
クライアントY座標 ( 数値 ) -- 対象ウィンドウのクライアント領域におけるY座標
種別 ( 定数 省略可 ) -- 取得したい情報の種類を示す定数
0
ACC_ACC を実行し、取得できなければ ACC_API を実行 (デフォルト)
ACC_ACC
表示文字列の取得
ACC_API
DrawText, TextOut等のAPIで描画されたテキストを取得 (未実装)
ACC_NAME
オブジェクトの表示名
ACC_VALUE
オブジェクトの値 (エディットボックス等)
ACC_ROLE
オブジェクトの役割名
ACC_STATE
オブジェクトの状態
ACC_DESCRIPTION
オブジェクトの説明
ACC_LOCATION
オブジェクトの位置情報
[x, y, 幅, 高さ]
ACC_BACK (オプション)
他の定数とOR連結で指定
対象ウィンドウをアクティブにしない
戻り値の型 :
文字列または配列
戻り値 :
ACC_LOCATION 指定時は数値の配列を返します
ACC_STATE 指定時は文字列の配列を返します
それ以外は該当する値を文字列で返します
失敗時はEMPTYを返します
muscur ( ) #
マウスカーソルの種別を返します
戻り値の型 :
定数
戻り値 :
CUR_APPSTARTING (1)
砂時計付き矢印
CUR_ARROW (2)
標準矢印
CUR_CROSS (3)
十字
CUR_HAND (4)
ハンド
CUR_HELP (5)
クエスチョンマーク付き矢印
CUR_IBEAM (6)
アイビーム (テキスト上のカーソル)
CUR_NO (8)
禁止
CUR_SIZEALL (10)
４方向矢印
CUR_SIZENESW (11)
斜め左下がりの両方向矢印
CUR_SIZENS (12)
上下両方向矢印
CUR_SIZENWSE (13)
斜め右下がりの両方向矢印
CUR_SIZEWE (14)
左右両方向矢印
CUR_UPARROW (15)
垂直の矢印
CUR_WAIT (16)
砂時計
0
上記以外
peekcolor ( x , y [ , RGB指定=COL_BGR , クリップボード=FALSE ] ) #
指定位置の色を得ます
peekcolorを繰り返し実行する場合
peekcolor 関数を繰り返し実行した場合、パフォーマンスに問題が出る場合があります
これは、 peekcolor が GetPixel というWin32関数を利用していることが原因です
(実装自体はUWSCと同じであり、UWSCにも同様の問題があります)
特定範囲内で任意の色を探すという用途であれば chkclr() 関数の使用を検討してください
これによりパフォーマンスを改善できる可能性があります
パラメータ :
x ( 数値 ) -- X座標
y ( 数値 ) -- Y座標
RGB指定 ( 定数 省略可 ) -- 戻り値の指定
COL_BGR (0)
BGR値で返す
青は$FF0000、緑は$00FF00、赤は$0000FF
COL_RGB
RGB値で返す
赤は$FF0000、緑は$00FF00、青は$0000FF
COL_R
赤の成分のみ
COL_G
緑の成分のみ
COL_B
青の成分のみ
クリップボード ( 真偽値 省略可 ) --
FALSE
画面の指定座標から
TRUE
クリップボード画像の指定座標から
戻り値の型 :
数値
戻り値 :
指定座標の色を示す数値
失敗時は -1 (範囲外指定やクリップボード指定でクリップボード画像がない場合)
getslider ( ID [ , n番目=1 , パラメータ=SLD_POS ] ) #
スライダー(スクロールバー、トラックバー)の値を取得します
パラメータ :
ID ( ウィンドウID ) -- 対象ウィンドウのID
n番目 ( 数値 省略可 ) -- n番目のスライダーから値を得る
パラメータ ( 定数 省略可 ) -- 取得する値の種類を示す定数
SLD_POS
現在値
SLD_MIN
最小値
SLD_MAX
最大値
SLD_PAGE
1ページ移動量
SLD_BAR
表示方向 (横なら0、縦なら1を返す)
SLD_X
クライアントX座標
SLD_Y
クライアントY座標
戻り値の型 :
数値
戻り値 :
取得した値、該当するスライダーがない場合は -999999
chkbtn ( ID , アイテム名 [ , n番目=1 , ACC=FALSE ] ) #
ボタン(チェックボックス、ラジオボタン)やメニューのチェック状態を得る
パラメータ :
ID ( 数値 ) -- 対象ウィンドウのID
アイテム名 ( 文字列 ) -- ボタン名 (部分一致)
n番目 ( 数値 省略可 ) -- n番目に該当するボタンの状態を得る
UWSCとは順序が異なる場合があります
実装の違いによりUWSCとは別の番号を指定しなければならない可能性があります
ご注意ください
ACC ( 真偽値 省略可 ) --
FALSE
APIまたはUIAを使用
TRUE
ACCを使用
UWSCとの違い
TRUEでも対象ウィンドウをアクティブにしないため、2は廃止されました
戻り値の型 :
数値またはFALSE
戻り値 :
-1: 存在しない、または無効
0: チェックされていない
1: チェックされている
2: チェックボックスが灰色 (ACCでは判定不可)
FALSE: ウィンドウが存在しない
getstr ( ID [ , n番目=1 , 種別=STR_EDIT , マウス移動=FALSE ] ) #
ウィンドウ上の文字列を取得します
パラメータ :
ID ( 数値 ) -- 対象ウィンドウのID
0の場合クリップボードから取得します (その場合以降の引数は無視されます)
クリップボードへのアクセスができない場合
クリップボードアクセス時に何かしらのエラーが発生した場合はEMPTYを返します
n番目 ( 数値 省略可 ) -- n番目に該当するアイテム種別の文字列を得る
0の場合はフォーカスされたコントロール
-n の場合はDisableになっているものも含めたn番目
UWSCとは順序が異なる場合があります
実装の違いによりUWSCとは別の番号を指定しなければならない可能性があります
ご注意ください
種別 ( 定数 省略可 ) -- 文字列を取得するアイテム種別
STR_EDIT
エディットコントロール
STR_STATIC
スタティックコントロール
STR_STATUS
ステータスバー
STR_ACC_EDIT
エディットコントロール等 (ACCで取得)
STR_ACC_STATIC
スタティックコントロール (ACCで取得)
STR_ACC_CELL
DataGridView内のセルの値
マウス移動 ( 真偽値 省略可 ) -- TRUEなら該当アイテムまでマウス移動
戻り値の型 :
文字列またはEMPTY
戻り値 :
取得した文字列、対象がない場合はEMPTY
getitem ( ID , 種別 [ , n番目=1 , 列=1 , ディセーブル無視=FALSE , ACC最大取得数=0 ] ) #
ウィンドウ上の文字情報をアイテム種類別に取得する
パラメータ :
ID ( 数値 ) -- 対象ウィンドウのID
種別 ( 定数 ) -- 種類を示す定数、OR連結で複数指定可
ITM_BTN
ボタン、チェックボックス、ラジオボタン
ITM_LIST
リストボックス、コンボボックス
ITM_TAB
タブコントロール
ITM_MENU
メニュー
ITM_TREEVIEW (ITM_TREEVEW)
ツリービュー
ITM_LISTVIEW (ITM_LSTVEW)
リストビュー
ITM_EDIT
エディットボックス
ITM_STATIC
スタティックコントロール
ITM_STATUSBAR
ステータスバー
ITM_TOOLBAR
ツールバー
ITM_LINK
リンク
ITM_ACCCLK
ACCによりクリック可能なもの
ITM_ACCCLK2
ACCによりクリック可能なもの、選択可能テキスト
ITM_ACCTXT
ACCスタティックテキスト
ITM_ACCEDIT
ACCエディット可能テキスト
ITM_ACC_TREE
ACCツリー構造
この定数を指定した場合は他の定数および以降の引数は無視されます
ITM_FROMLAST
ACCで検索順序を逆にする
ITM_BACK
ACCでウィンドウをアクティブにしない
n番目 ( 数値 省略可 ) -- ITM_LIST、ITM_TREEVIEW、ITM_LISTVIEW指定時かつ対象が複数あった場合にいずれを取得するか指定、-1ならすべて取得
複数種別同時指定時の処理について
ITM_LIST、ITM_TREEVIEW、ITM_LISTVIEWのうち複数を同時に指定した場合、それぞれのn番目を検索します
// この場合リストまたはコンボボックスの2番目、及びツリービューの2番目をそれぞれ取得します
getitem ( id , ITM_LIST or ITM_TREEVIEW , 2 )
UWSCとは順序が異なる場合があります
実装の違いによりUWSCとは別の番号を指定しなければならない可能性があります
ご注意ください
列 ( 数値 省略可 ) -- ITM_LISTVIEW指定時にどの列から取得するかを指定(1から)、0ならすべての列、-1ならカラム名を取得
ディセーブル無視 ( 真偽値 省略可 ) -- FALSEならディセーブル状態でも取得する、TRUEなら取得しない
ACC最大取得数 ( 数値 省略可 ) -- ACC指定時に取得するアイテム数の上限を指定、0なら無制限、マイナス指定時は逆順(ITM_FROMLASTと同じ)
戻り値の型 :
文字列の配列またはUObject
戻り値 :
取得されたアイテム名の配列
ITM_ACC_TREE 指定時はACCのツリー構造を示すUObject (失敗時はNULL)
// 対象ウィンドウのACCツリー構造を取得
id = getid ( &quot;ファイル名を指定して実行&quot; )
uo = getitem ( id , ITM_ACC_TREE )
json = tojson ( uo , TRUE )
path = &quot;.\acc.json&quot;
fput ( fopen ( path , F_WRITE8 or F_AUTOCLOSE ) , json )
shexec ( path )
UWSCとの違い
戻り値が配列になったため ALL_ITEM_LIST は廃止されました
items = getitem ( id , ITM_BTN )
// 個数を得る
print length ( items )
// アイテム名の表示
for item in items
print item
next
また、空の文字列は結果に含まれなくなりました
// UWSCでは空文字を1つ目のアイテムとして出力していましたが、UWSCRでは空文字はスキップされます
i = 0
for item in getitem ( getid ( &#39;ファイル名を指定して実行&#39; ) , ITM_STATIC )
i += 1
print &quot;&lt;#i&gt;: &lt;#item&gt;&quot;
next
// 結果
// 1: 実行するプログラム名、または開くフォルダーやドキュメント名、インターネット リソース名を入力してください。
// 2: 名前(&amp;O):
ACC全般は以下の条件で取得します
- 条件に一致するロール
- 可視またはフォーカス可能
- ステータスが0ではない
- ディセーブル無視がTRUEの場合enabledのもののみ
ITM_ACCCLK
- リンクの場合親要素を含めないようにしました
ITM_ACCCLK2
ITM_ACCTXT
- 親要素を含めないようにしました
ITM_EDIT
- 読み取り専用も取得するようにしました
getslctlst ( ID [ , n番目=1 , 列=1 ] ) #
表示されているコンボボックス、リストボックス、ツリービュー、リストビューから選択されている項目を取得
パラメータ :
数値 ( ID ) -- 対象ウィンドウのID
n番目 ( 数値 省略可 ) -- n番目の該当コントロールから値を得る (1から)
列 ( 数値 省略可 ) -- リストビューの場合取得する列を指定 (1から)
戻り値の型 :
文字列、または文字列の配列
戻り値 :
選択項目、複数選択されている場合は配列で返る
UWSCとの違い
リストやリストビューが複数選択されていた場合にタブ連結された文字列ではなく、
それぞれの要素を持つ配列として返すようになりました
chkclr ( 探索色 [ , 閾値=0 , 範囲=[] , キャプチャ方法=-1 ] ) #
範囲内に探索色があればその位置を返します
mouseorg が実行されている場合は探索対象がそのウィンドウとなります
パラメータ :
探索色 ( 数値または配列 ) -- 探す色を指定します
数値: BGR値
配列: [B値, G値, R値]
閾値 ( 数値または配列 省略可 ) -- 探索する色の幅を指定します
数値: BGRそれぞれに対する閾値
配列: 個別指定 [対B, 対G, 対R]
閾値指定による色の幅について
探索色のB値が30でBに対する閾値が5の場合25～35であればヒットする
255 ($FF) を指定すると元の値に関わらずその色要素に対して必ずヒットします
chkclr ([ 0 , 100 , 0 ] , [ 255 , 5 , 255 ])
// 下限: [ 0, 95, 255]
// 上限: [255, 105, 255]
// が探索色となるB要素とR要素はすべてを対象とし、Gのみ95-105を対象とする
範囲 ( 配列 省略可 ) -- 探索範囲を [左上x, 左上y, 右下x, 右下y] で指定、省略時はモニタまたはウィンドウのサイズ
部分的な省略について
配列サイズが4より小さい場合、不足分は省略扱いとなります
null を記述することで省略であることを明示できます
[100] 左上xのみ指定、残りは省略
[100, 100] 左上xyを指定、右下xyは省略
[100, null, 100] 左上xと右下xを指定、左上yと右下yは省略
キャプチャ方法 ( 数値 省略可 ) -- キャプチャ方法及びキャプチャ対象モニタを指定する
mouseorg未使用時
- -1: スクリーン全体をGDIでキャプチャする
- 0以上の数値: 値をモニタ番号とし、Graphic Capture APIでキャプチャする
mouseorg使用時
- -1: ウィンドウをGDIでキャプチャする
- 0以上の数値: ウィンドウをGraphic Capture APIでキャプチャする
戻り値の型 :
二次元配列
戻り値 :
該当色のある座標および見つかった色([x, y, [b, g, r]])の配列
サンプルコード
function bgr_array_to_int ( arr : array )
result = arr [ 0 ] * $10000 + arr [ 1 ] * $100 + arr [ 2 ]
fend
mouseorg ( id )
offset_x = status ( id , ST_X )
offset_y = status ( id , ST_Y )
color = [ 0 , 100 , 0 ]
bgcolor = bgr_array_to_int ( color )
threshold = [ 0 , 5 , 0 ]
// [0, 95, 0] から [0, 105, 0] を探索範囲とする
for found in chkclr ( color , threshold )
x = found [ 0 ] + offset_x
y = found [ 1 ] + offset_y
color = found [ 2 ]
msg = &quot;座標: &lt;#x&gt;, &lt;#y&gt; 色: &lt;#color&gt;&quot;
balloon ( msg , x , y , FUKI_DOWN or FUKI_POINT , , , 0 , bgcolor )
if msgbox ( &quot;次へ&quot; ) = BTN_CANCEL then
break
endif
next
「Explorerが停止しているかもしれません、その場合Explorerを再起動してください」エラーについて
Explorerが停止状態のため画面のキャプチャに失敗している可能性があります
このエラーが出力された場合はタスクマネージャでExplorerの状態を確認してください
もし停止していたら再起動してください
Ctrl + Shift + Esc キーを同時に押し、タスクマネージャを起動します
エクスプローラー を右クリックします
メニューの 再起動 を押します
## 画像検索#
ChkImg ( [ ファイル名={clipboard} , 探索方式=0 , x1=EMPTY , y1=EMPTY , x2=EMPTY , y2=EMPTY , n番目=1 , 色幅=0 ] ) #
スクリーン上の指定画像と一致する位置の情報を返す
パラメータ :
ファイル名 ( 文字列 省略可 ) -- 探す画像ファイルのパス、省略時はクリップボード画像
探索方式 ( 数値 省略可 ) -- 一致判定の方式を指定
0: すべてのピクセルで一致を判定する
1: 指定画像の左上を透過色とし、指定画像の透過色に当たる部分は常に一致、それ以外は通常の一致判定を行う
2: 指定画像の右上を透過色とする
3: 指定画像の左下を透過色とする
4: 指定画像の右下を透過色とする
-1: 色ではなく形での一致を判定する、色幅は無視される
x1 ( 数値 省略可 ) -- 探索範囲の左上x座標、省略時はスクリーン全体の左上
y1 ( 数値 省略可 ) -- 探索範囲の左上y座標、省略時はスクリーン全体の左上
x2 ( 数値 省略可 ) -- 探索範囲の右下x座標、省略時はスクリーン全体の右下
y2 ( 数値 省略可 ) -- 探索範囲の右下y座標、省略時はスクリーン全体の右下
n番目 ( 数値 省略可 ) -- スクリーン左上から見てn番目の一致座標を返す、-1ならスクリーン上で一致するすべての座標を返す
色幅 ( 定数 省略可 ) -- 各ピクセルについて許容する色範囲を指定 (OR連結可)、省略時は完全一致
IMG_MSK_R1: RGBのうちR (赤) に対して R -2 &lt; 対象R &lt; R + 2 の範囲で許容する
IMG_MSK_R2: RGBのうちR (赤) に対して R -4 &lt; 対象R &lt; R + 4 の範囲で許容する
IMG_MSK_R3: RGBのうちR (赤) に対して R -8 &lt; 対象R &lt; R + 8 の範囲で許容する
IMG_MSK_R4: RGBのうちR (赤) に対して R -16 &lt; 対象R &lt; R + 16 の範囲で許容する
IMG_MSK_G1: RGBのうちG (緑) に対して G -2 &lt; 対象G &lt; G + 2 の範囲で許容する
IMG_MSK_G2: RGBのうちG (緑) に対して G -4 &lt; 対象G &lt; G + 4 の範囲で許容する
IMG_MSK_G3: RGBのうちG (緑) に対して G -8 &lt; 対象G &lt; G + 8 の範囲で許容する
IMG_MSK_G4: RGBのうちG (緑) に対して G -16 &lt; 対象G &lt; G + 16 の範囲で許容する
IMG_MSK_B1: RGBのうちB (青) に対して B -2 &lt; 対象B &lt; B + 2 の範囲で許容する
IMG_MSK_B2: RGBのうちB (青) に対して B -4 &lt; 対象B &lt; B + 4 の範囲で許容する
IMG_MSK_B3: RGBのうちB (青) に対して B -8 &lt; 対象B &lt; B + 8 の範囲で許容する
IMG_MSK_B4: RGBのうちB (青) に対して B -16 &lt; 対象B &lt; B + 16 の範囲で許容する
IMG_MSK_BGR1: RGBそれぞれに対して n -2 &lt; 対象色 &lt; n + 2 の範囲で許容する
IMG_MSK_BGR2: RGBそれぞれに対して n -4 &lt; 対象色 &lt; n + 4 の範囲で許容する
IMG_MSK_BGR3: RGBそれぞれに対して n -8 &lt; 対象色 &lt; n + 8 の範囲で許容する
IMG_MSK_BGR4: RGBそれぞれに対して n -16 &lt; 対象色 &lt; n + 16 の範囲で許容する
戻り値の型 :
配列
戻り値 :
n番目で1以上を指定した場合は該当する位置の [x, y]
-1を指定した場合は [x, y] の配列
見つからなかった場合は空の配列
クリップボード指定でクリップボードに画像がない場合も空の配列
サンプルコード
// n番目を-2にすることで探索が並列処理になる
for xy in chkimg ( image , 0 ,,,,, - 2 , IMG_MSK_BGR3 )
print xy
next
SearchImage ( 画像ファイルパス [ , スコア=95 , 最大検索数=5 , left=EMPTY , top=EMPTY , right=EMPTY , bottom=EMPTY , オプション=0 , モニタ番号=0 ] ) #
指定画像をスクリーン上から探してその座標を返します
UWSCとは互換性がありません
特殊変数 G_IMG_X , G_IMG_Y , ALL_IMG_X , ALL_IMG_Y は廃止
戻り値が変更されています
パラメータ :
画像ファイルパス ( 文字列 ) -- 検索する画像のパス (jpg, bmp, png)
スコア ( 数値 省略可 ) -- 画像に対する一致率を指定 (80.0-100.0)
一致率が指定値以上であれば結果を返します
小数も有効です (例: 99.75)
生スコア値指定
スコアは実際の処理では0.0から1.0の範囲の値として扱われます
例: 95 → 0.95
スコア値を0.0から1.0の範囲で指定した場合はそのままの値が使われます
この場合はスコアの下限がないため80未満を指定することも可能です
例: 0.75 (スコア75相当)
最大検索数 ( 数値 省略可 ) -- 検索の試行回数を指定
left ( 数値 省略可 ) -- 検索範囲指定: 左上X座標、省略時は画面左上X座標
top ( 数値 省略可 ) -- 検索範囲指定: 左上Y座標、省略時は画面左上Y座標
right ( 数値 省略可 ) -- 検索範囲指定: 右下X座標、省略時は画面右下X座標
bottom ( 数値 省略可 ) -- 検索範囲指定: 右下X座標、省略時は画面右下Y座標
オプション ( 定数 省略可 ) -- 実行時オプションを指定、OR連結可
SCHIMG_NO_GRAY
画像をグレースケール化せず探索を行う
SCHIMG_USE_WGCAPI
デスクトップまたはウィンドウの画像取得にGraphicsCaptureAPIを使う
デスクトップの場合は対象とするモニタを次の引数で指定
mouseorgを利用している場合はウィンドウを対象とする
このオプションを指定した場合mouseorgの MOUSE_FORE および MOUSE_BACK は無視されます (指定に関わらずフォア・バックをキャプチャ可能)
ヒント
このオプションにより通常ではキャプチャできないウィンドウがキャプチャできる可能性があります
キャプチャできないウィンドウ状態について
対象ウィンドウが最小化されている、または非表示になっている場合はキャプチャを行わず関数を終了します
このオプションでウィンドウをキャプチャする場合は対象ウィンドウが表示状態になっていることを確認してください
SCHIMG_METHOD_SQDIFF
類似度の計算にTM_SQDIFFを使用する、他の計算方法と併用不可
SCHIMG_METHOD_SQDIFF_NORMED
類似度の計算にTM_SQDIFF_NORMEDを使用する、他の計算方法と併用不可
SCHIMG_METHOD_CCORR
類似度の計算にTM_CCORRを使用する、他の計算方法と併用不可
SCHIMG_METHOD_CCORR_NORMED
類似度の計算にTM_CCORR_NORMEDを使用する、他の計算方法と併用不可
SCHIMG_METHOD_CCOEFF
類似度の計算にTM_CCOEFFを使用する、他の計算方法と併用不可
SCHIMG_METHOD_CCOEFF_NORMED
類似度の計算にTM_CCOEFF_NORMEDを使用する、他の計算方法と併用不可
計算方法未指定時はこれが適用される
モニタ番号 ( 定数 省略可 ) --
SCHIMG_USE_WGCAPI 時に検索するモニタ番号を0から指定、デフォルトは0 (プライマリモニタ)
mousemorg使用時はウィンドウを対象とするためこの引数指定は不要
戻り値の型 :
二次元配列
戻り値 :
該当する部分の座標とスコアを格納した二次元配列 [[X座標, Y座標, スコア], ...]
サンプルコード
for found in chkimg ( &quot;hoge.png&quot; )
print found // [x, y, スコア]
next
SearchImage ( 画像ファイルパス [ , スコア=95 , 最大検索数=5 , 範囲 , オプション=0 ] )
配列による範囲指定
パラメータ :
範囲 ( 配列 省略可 ) -- [left, top, right, bottom] で指定
サンプルコード
found = chkimg ( &quot;hoge.png&quot; , 95 , 1 , [ 100 , 100 , 400 , 400 ])
saveimg ( [ ファイル名=EMPTY , ID=0 , x=EMPTY , y=EMPTY , 幅=EMPTY , 高さ=EMPTY , クライアント領域=FALSE , 圧縮率=EMPTY , 取得方法=IMG_AUTO , WGCAPI=false , モニタ番号=0 ] ) #
ウィンドウの画像を保存します
パラメータ :
ファイル名 ( 文字列 省略可 ) -- 保存するファイル名 (対応する拡張子は jpg , bmp , png )、EMPTYの場合はクリップボードにコピー
拡張子が有効ではない場合
pngファイルとして保存されます
saveimg ( &quot;hoge&quot; ) // hoge.pngが保存される
ID ( 数値 省略可 ) -- ウィンドウID、0の場合スクリーン全体
x ( 数値 省略可 ) -- 取得範囲の起点となるx座標、EMPTYの場合は左上
y ( 数値 省略可 ) -- 取得範囲の起点となるy座標、EMPTYの場合は左上
幅 ( 数値 省略可 ) -- 取得範囲の幅、EMPTYの場合は ウィンドウ幅 - x
高さ ( 数値 省略可 ) -- 取得範囲の高さ、EMPTYの場合は ウィンドウ高さ - y
クライアント領域 ( 真偽値 省略可 ) -- FALSEならウィンドウ全体、TRUEならクライアント領域のみ
圧縮率 ( 数値 省略可 ) --
指定したファイル拡張子により指定値が異なります
ファイル名を省略した(クリップボードにコピーされる)場合この値は無視されます
jpg
JPEG画像の画質を0-100で指定します (高いほど高画質)
EMPTY指定時、または値が範囲外の場合は95になります
png
PNG画像の圧縮度合いを0-9で指定します (高いほどサイズが小さくなるが、遅くなる)
EMPTY指定時、または値が範囲外の場合は1になります
bmp
この値は無視されます
UWSCとの違い
UWSCでは1-100指定ならJPEG、0ならBMPで保存されていましたが、UWSCRではファイル名の拡張子で保存形式を指定します
取得方法 ( 定数 省略可 ) -- 画面の取得方法
IMG_FORE
スクリーン全体から対象ウィンドウの座標を元に画像を切り出す
IMG_BACK
対象ウィンドウから画像を取得
注意
他のウィンドウに隠れていても取得可能ですが、見た目が完全に一致しない場合があります
IMG_AUTO (0)
ウィンドウ全体が可視かどうかで取得方法を自動的に切り替えます
ウィンドウが見えていれば IMG_FORE を使用する (アクティブかどうかは問わない)
一部でも他のウィンドウに隠れていれば IMG_BACK を使用する
WGCAPI ( 真偽値 省略可 ) -- TRUEならGraphicsCaptureAPIにより画面またはウィンドウをキャプチャします
モニタ番号 ( 数値 省略可 ) -- IDに0を指定して、かつWGCAPIをTRUEにした場合にキャプチャするモニタ番号を0から指定
xy座標は0から
xy座標はモニタごとの座標を0から指定してください
0未満が指定された場合は0になります
戻り値 :
なし
## 低レベル関数#
mmv ( x , y [ , ms=0 ] ) #
マウスカーソルを移動します
パラメータ :
x ( 数値 ) -- 移動先のX座標
y ( 数値 ) -- 移動先のY座標
ms ( 数値 省略可 ) -- マウス移動を行うまでの待機時間 (ミリ秒)
戻り値 :
なし
btn ( ボタン定数 [ , 状態=CLICK , x=EMPTY , y=EMPTY , ms=0 ] ) #
指定座標にマウスボタン操作を送信します
パラメータ :
ボタン定数 ( 定数 ) -- 操作するマウスボタンを指定
LEFT
左クリック
RIGHT
右クリック
MIDDLE
ホイルクリック
WHEEL
ホイル回転 (上下方向)
WHEEL2
ホイル回転 (左右方向)
TOUCH
タッチ操作を行う
状態をCLICKにした場合指定座標をタッチして離す
状態をDOWNにした場合指定座標でタッチ
その後状態をUPで再実行した場合同一座標ならそのまま離し、座標が異なるならその座標までスワイプ操作を行う
msを指定した場合はスワイプ速度に影響する (移動区間の一区切り毎の移動速度を変更する)
重要
タッチできるのは一点のみ (複数箇所タッチは不可)
状態 ( 定数 省略可 ) -- マウスボタンに対してどのような操作を行うかを指定
LEFT , RIGHT , MIDDLE の場合以下のいずれかを指定
CLICK
ボタンクリック (デフォルト)
DOWN→UPを連続で行います
待機時間について
DOWNとUPの間と、UP後にわずかに待機時間が入ります
DOWN
待機
UP
待機(小)
DOWN
ボタン押し下げ
UP
ボタン開放
WHEEL : ノッチ数を指定 (正なら上方向、負なら下方向に回転)
WHEEL2 : ノッチ数を指定 (正なら右方向、負なら左方向に回転)
x ( 数値 省略可 ) -- ボタン操作を行う位置のX座標、省略時は現在のマウスのX座標
y ( 数値 省略可 ) -- ボタン操作を行う位置のY座標、省略時は現在のマウスのY座標
ms ( 数値 省略可 ) --
ボタン操作を行うまでの待機時間 (ミリ秒)
またはTOUCHのDOWN後のUPで別座標を指定した場合のスワイプ速度、0 (速)～10 (遅)
戻り値 :
なし
サンプルコード
btn ( TOUCH , DOWN , 100 , 100 )
btn ( TOUCH , UP , 200 , 200 ) // 別座標でUPした場合はスワイプ操作になる
btn ( TOUCH , DOWN , 150 , 150 )
btn ( TOUCH , UP , 250 , 250 , 0 ) // msが0なら最速
btn ( TOUCH , DOWN , 300 , 300 )
btn ( TOUCH , UP , 150 , 150 , 10 ) // 10ならとても遅い
kbd ( 仮想キーまたは文字コード [ , 状態=CLICK , ms=0 ] ) #
kbd ( 送信文字列 [ , 状態=CLICK , ms=0 ] )
キーボード入力を送信します
パラメータ :
仮想キーまたは文字コード ( 数値 ) -- 仮想キーコード一覧 のいずれか、または文字コード
送信文字列 ( 文字列 ) -- キー入力として送信される文字列
状態 ( 定数 省略可 ) -- キーの入力状態を指定、文字列送信時は無視される
CLICK
キークリック (デフォルト)
DOWN→UPを連続で行います
待機時間について
DOWNとUPの間と、UP後にわずかに待機時間が入ります
DOWN
待機
UP
待機(小)
DOWN
キー押し下げ
UP
キー開放
ms ( 数値 省略可 ) -- キーボード入力を行うまでの待機時間 (ミリ秒)
戻り値 :
なし
サンプルコード
// キーコード入力
// a が入力される
kbd ( VK_A )
// A が入力される
kbd ( VK_SHIFT , DOWN )
kbd ( VK_A , CLICK , 100 )
kbd ( VK_SHIFT , UP , 100 )
// 文字コード入力
// あ が入力される
kbd ( asc ( &quot;あ&quot; ))
// A が入力される
kbd ( &quot;A&quot; )
// あ が入力される
kbd ( &quot;あ&quot; )
// abcde が入力される
kbd ( &quot;abcde&quot; )
acw ( ID [ , x=EMPTY , y=EMPTY , w=EMPTY , h=EMPTY , ms=0 ] ) #
ウィンドウの位置やサイズを変更します
ID0 を更新します
パラメータ :
ID ( 数値 ) -- ウィンドウID
x ( 数値 省略可 ) -- 移動先のX座標、省略時は対象ウィンドウの現在のX座標
y ( 数値 省略可 ) -- 移動先のY座標、省略時は対象ウィンドウの現在のY座標
w ( 数値 省略可 ) -- 変更するウィンドウの幅、省略時は対象ウィンドウの現在の幅
h ( 数値 省略可 ) -- 変更するウィンドウの高さ、省略時は対象ウィンドウの現在の高さ
ms ( 数値 省略可 ) -- ウィンドウに変更を加えるまでの待機時間 (ミリ秒)
戻り値 :
なし
サンプルコード
acw ( getid ( GET_ACTIVE_WIN ) , 100 , 100 ) // ID0を更新
sleep ( 1 )
acw ( 0 , 200 , 200 )
## 仮想キーコード一覧#
VK_A
VK_B
VK_C
VK_D
VK_E
VK_F
VK_G
VK_H
VK_I
VK_J
VK_K
VK_L
VK_M
VK_N
VK_O
VK_P
VK_Q
VK_R
VK_S
VK_T
VK_U
VK_V
VK_W
VK_X
VK_Y
VK_Z
VK_0
VK_1
VK_2
VK_3
VK_4
VK_5
VK_6
VK_7
VK_8
VK_9
VK_START
VK_BACK
VK_TAB
VK_CLEAR
VK_ESC
VK_ESCAPE
VK_RETURN
VK_ENTER
VK_RRETURN
VK_SHIFT
VK_RSHIFT
VK_WIN
VK_RWIN
VK_ALT
VK_MENU
VK_RALT
VK_CTRL
VK_CONTROL
VK_RCTRL
VK_PAUSE
VK_CAPITAL
VK_KANA
VK_FINAL
VK_KANJI
VK_CONVERT
VK_NONCONVERT
VK_ACCEPT
VK_MODECHANGE
VK_SPACE
VK_PRIOR
VK_NEXT
VK_END
VK_HOME
VK_LEFT
VK_UP
VK_RIGHT
VK_DOWN
VK_SELECT
VK_PRINT
VK_EXECUTE
VK_SNAPSHOT
VK_INSERT
VK_DELETE
VK_HELP
VK_APPS
VK_MULTIPLY
VK_ADD
VK_SEPARATOR
VK_SUBTRACT
VK_DECIMAL
VK_DIVIDE
VK_NUMPAD0
VK_NUMPAD1
VK_NUMPAD2
VK_NUMPAD3
VK_NUMPAD4
VK_NUMPAD5
VK_NUMPAD6
VK_NUMPAD7
VK_NUMPAD8
VK_NUMPAD9
VK_F1
VK_F2
VK_F3
VK_F4
VK_F5
VK_F6
VK_F7
VK_F8
VK_F9
VK_F10
VK_F11
VK_F12
VK_NUMLOCK
VK_SCROLL
VK_PLAY
VK_ZOOM
VK_SLEEP
VK_BROWSER_BACK
VK_BROWSER_FORWARD
VK_BROWSER_REFRESH
VK_BROWSER_STOP
VK_BROWSER_SEARCH
VK_BROWSER_FAVORITES
VK_BROWSER_HOME
VK_VOLUME_MUTE
VK_VOLUME_DOWN
VK_VOLUME_UP
VK_MEDIA_NEXT_TRACK
VK_MEDIA_PREV_TRACK
VK_MEDIA_STOP
VK_MEDIA_PLAY_PAUSE
VK_LAUNCH_MEDIA_SELECT
VK_LAUNCH_MAIL
VK_LAUNCH_APP1
VK_LAUNCH_APP2
VK_OEM_PLUS
VK_OEM_COMMA
VK_OEM_MINUS
VK_OEM_PERIOD
VK_OEM_1
VK_OEM_2
VK_OEM_3
VK_OEM_4
VK_OEM_5
VK_OEM_6
VK_OEM_7
VK_OEM_8
VK_OEM_RESET
VK_OEM_JUMP
VK_OEM_PA1
VK_OEM_PA2
VK_OEM_PA3

---

# スクリプト制御 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/scriptcontrol.html

## スクリプト制御#
## 実行をブロック#
sleep ( 秒数 ) #
指定した秒数の間スクリプトの実行をブロックします
パラメータ :
秒数 ( 数値 ) -- スクリプトの実行を停止する秒数
sleep ( 関数 )
関数がTRUEを返す限り実行をブロックします
パラメータ :
関数 ( ユーザー定義関数 ) -- 条件判定を行う関数
## 動的評価#
eval ( 構文 ) #
渡された文字列をUWSCRの構文として評価します
式として評価された場合はその結果の値を返します
パラメータ :
構文 ( 文字列 ) -- UWSCRの構文を表す文字列
戻り値 :
式が評価された場合はその実行結果の 値 、文が評価された場合は EMPTY
サンプルコード
a = 1
eval ( &quot;a = 5&quot; ) // = で代入できる
for i = 0 to 3
print a
eval ( &quot;if a &gt; 3 then a -= 1 else a += 1&quot; ) // 単行IF
next
print a
## エラー発生#
raise ( エラーメッセージ , タイトル = 規定のタイトル ) #
実行時エラーを故意に発生させます
パラメータ :
エラーメッセージ ( 文字列 ) -- エラー内容を示す文字列
タイトル ( 文字列 省略可 ) -- エラーのタイトル
戻り値 :
なし
サンプルコード
try
print 1
raise ( &quot;エラーが発生しました&quot; )
print 2
except
print TRY_ERRMSG
endtry
# 結果
1
[ ユーザー定義エラー ] エラーが発生しました
assert_equal ( 値1 , 値2 ) #
2つの値を比較し、一致していない場合は実行時エラーになります
パラメータ :
値1 ( 値 ) -- 任意の値
値2 ( 値 ) -- 比較する値
戻り値 :
なし
サンプルコード
dim a = 5 , b = a , c = a * 2
assert_equal ( a , b ) // 一致するので何も起こらない
assert_equal ( b , c ) // [assert_equalエラー] left: 5; right: 10
## タスク#
Task ( func [ , args , ... ] ) #
関数を非同期に実行し、実行中の状態をタスクとして返します
await実行した場合
Task関数自体をawaitで実行した場合は関数の終了を待ちその戻り値を返します
パラメータ :
func ( 関数 ) -- 非同期実行させるユーザー定義関数
args ( 値 省略可 ) -- 関数に渡す引数 (最大20個まで指定可能)
戻り値の型 :
タスク
戻り値 :
実行中の タスク
WaitTask ( task ) #
タスク の完了を待ち、関数の戻り値を得ます
Promiseに相当する RemoteObject を受けた場合はそのPromiseの完了を待ち RemoteObject を返します
Promise以外はエラー
RemoteObject がPromiseではない場合エラーで終了します
パラメータ :
task ( タスク ) -- 未完了の タスク , または RemoteObject
戻り値 :
タスク として実行していた関数の戻り値、または RemoteObject
サンプルコード
function MyTask ( wait : number )
for i = 1 to wait
sleep ( 1 )
print &quot;タスク実行中: &quot; + ( wait - i )
next
result = &quot;タスク実行完了: &lt;#wait&gt;秒待ちました&quot;
fend
t = Task ( MyTask , 5 )
print &quot;タスクを開始しました&quot;
print &quot;タスクは非同期で実行されるため、その間別の処理を行えます&quot;
print &quot;WaitTaskを呼ぶと処理をブロックし、タスクの完了を待ちます&quot;
print &quot;タスクが完了すると関数のresult値を得られます&quot;
print WaitTask ( t ) // タスク実行完了: 5秒待ちました
## 型チェック#
type_of ( 値 ) #
値の型を返します
パラメータ :
値 ( すべて ) -- 型を調べたい値や変数
戻り値の型 :
定数
戻り値 :
型を示す文字定数
0.15.0時点での型定数
TYPE_NUMBER
TYPE_STRING
TYPE_BOOL
TYPE_ARRAY
TYPE_HASHTBL
TYPE_ANONYMOUS_FUNCTION
TYPE_FUNCTION
TYPE_BUILTIN_FUNCTION
TYPE_ASYNC_FUNCTION
TYPE_MODULE
TYPE_CLASS
TYPE_CLASS_INSTANCE
TYPE_NULL
TYPE_EMPTY
TYPE_NOTHING
TYPE_HWND
TYPE_REGEX
TYPE_UOBJECT
TYPE_VERSION
TYPE_THIS
TYPE_GLOBAL
TYPE_ENUM
TYPE_TASK
TYPE_DLL_FUNCTION
TYPE_STRUCT_DEFINITION
TYPE_STRUCT_INSTANCE
TYPE_COM_OBJECT
TYPE_IUNKNOWN
TYPE_VARIANT
TYPE_SAFEARRAY
TYPE_BROWSERBUILDER_OBJECT
TYPE_BROWSER_OBJECT
TYPE_TABWINDOW_OBJECT
TYPE_REMOTE_OBJECT
TYPE_FILE_ID
TYPE_BYTE_ARRAY
TYPE_REFERENCE
TYPE_WEB_REQUEST
TYPE_WEB_RESPONSE
TYPE_HTML_NODE
TYPE_WEBVIEW_FORM
TYPE_WEBVIEW_REMOTEOBJECT
TYPE_MEMBER_CALLER
TYPE_NOT_VALUE_TYPE

---

# システム関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/system.html

## システム関数#
## システム情報#
kindofos ( データ種別 = FALSE ) #
OS種別、またはアーキテクチャを判定します
パラメータ :
データ種別 ( 真偽値または定数 省略可 ) -- 以下のいずれかを指定
IS_64BIT_OS, TRUE
OSが64ビットかどうかを真偽値で返す
KIND_OF_OS, FALSE
OS種別をOS定数で返す
OS_WIN2000 (12)
OS_WINXP (13)
OS_WINSRV2003 (14)
OS_WINSRV2003R2 (15)
OS_WINVISTA (20)
OS_WINSRV2008 (21)
OS_WIN7 (22)
OS_WINSRV2008R2 (27)
OS_WIN8 (23)
OS_WINSRV2012 (24)
OS_WIN81 (25)
OS_WINSRV2012R2 (26)
OS_WIN10 (30)
OS_WINSRV2016 (31)
OS_WIN11 (32)
OSVER_MAJOR
OSのメジャーバージョンを数値で返す
OSVER_MINOR
OSのマイナーバージョンを数値で返す
OSVER_BUILD
OSのビルド番号を数値で返す
OSVER_PLATFORM
OSのプラットフォームIDを数値で返す
戻り値 :
データ種別による
env ( 環境変数 ) #
環境変数を展開します
パラメータ :
環境変数 ( 文字列 ) -- 環境変数を示す文字列
戻り値 :
展開された環境変数( 文字列 )
サンプルコード
print env ( &#39;programfiles&#39; ) // C:\Program Files
setenv ( 環境変数 , 設定値 ) #
プロセス環境変数を設定します
ヒント
この環境変数は実行中のuwscrプロセス及びその子プロセスに対してのみ有効です
パラメータ :
環境変数 ( 文字列 ) -- 環境変数名
設定値 ( 文字列 ) -- 環境変数にセットする値
サンプルコード
print env ( &#39;NO_PROXY&#39; ) // 空文字
setenv ( &#39;NO_PROXY&#39; , &#39;localhost&#39; )
print env ( &#39;NO_PROXY&#39; ) // localhost
wmi ( WQL , 名前空間 = 'root/cimv2' ) #
WQLを発行しWMIから情報を得ます
パラメータ :
WQL ( 文字列 ) -- WMIに対するクエリ文
名前空間 ( 文字列 省略可 ) -- 名前空間のパス
戻り値 :
クエリ結果( UObject配列 )
サンプルコード
res = wmi ( &#39;select name, processid, commandline from Win32_Process where name = &quot;uwscr.exe&quot;&#39; )
for obj in res
print obj . name
print obj . processid
print obj . commandline
next
cpuuserate ( ) #
システム全体での1秒間のCPU使用率を得る
戻り値の型 :
数値
戻り値 :
CPU使用率
sensor ( 種別 ) #
各種センサーから情報を得る (Sensor APIを使用)
パラメータ :
種別 ( 定数 ) -- センサー種別を指定する定数
SNSR_Biometric_HumanPresense
人が存在した場合に True
SNSR_Biometric_HumanProximity
人との距離(メートル)
SNSR_Electrical_Capacitance
静電容量(ファラド)
SNSR_Electrical_Resistance
電気抵抗(オーム)
SNSR_Electrical_Inductance
誘導係数(ヘンリー)
SNSR_Electrical_Current
電流(アンペア)
SNSR_Electrical_Voltage
電圧(ボルト)
SNSR_Electrical_Power
電力(ワット)
SNSR_Environmental_Temperature
気温(セ氏)
SNSR_Environmental_Pressure
気圧(バール)
SNSR_Environmental_Humidity
湿度(パーセンテージ)
SNSR_Environmental_WindDirection
風向(度数)
SNSR_Environmental_WindSpeed
風速(メートル毎秒)
SNSR_Light_Lux
照度(ルクス)
SNSR_Light_Temperature
光色温度(ケルビン)
SNSR_Mechanical_Force
力(ニュートン)
SNSR_Mechanical_AbsPressure
絶対圧(パスカル)
SNSR_Mechanical_GaugePressure
ゲージ圧(パスカル)
SNSR_Mechanical_Weight
重量(キログラム)
SNSR_Motion_AccelerationX
SNSR_Motion_AccelerationY
SNSR_Motion_AccelerationZ
X/Y/Z軸 加速度(ガル)
SNSR_Motion_AngleAccelX
SNSR_Motion_AngleAccelY
SNSR_Motion_AngleAccelZ
X/Y/Z軸 角加速度(度毎秒毎秒)
SNSR_Motion_Speed
速度(メートル毎秒)
SNSR_Scanner_RFIDTag
RFIDタグの40ビット値
SNSR_Scanner_BarcodeData
バーコードデータを表す文字列
注意
UWSCRではサポートされません (必ずEMPTYを返します)
SNSR_Orientation_TiltX
SNSR_Orientation_TiltY
SNSR_Orientation_TiltZ
X/Y/Z 軸角(度)
SNSR_Orientation_DistanceX
SNSR_Orientation_DistanceY
SNSR_Orientation_DistanceZ
X/Y/Z 距離(メートル)
SNSR_Orientation_MagHeading
磁北基準未補正コンパス方位
SNSR_Orientation_TrueHeading
真北基準未補正コンパス方位
SNSR_Orientation_CompMagHeading
磁北基準補正済みコンパス方位
SNSR_Orientation_CompTrueHeading
真北基準補正済みコンパス方位
SNSR_Location_Altitude
海抜(メートル)
SNSR_Location_Latitude
緯度(度数)
SNSR_Location_Longitude
経度(度数)
SNSR_Location_Speed
スピード(ノット)
戻り値の型 :
真偽値、数値、文字列
戻り値 :
種別に応じた値、値が取得できない場合はEMPTY
戻り値NaNの廃止
一部のエラーで値が取得できない場合にUWSCはNaNを返していましたが、UWSCRではEMPTYが返ります
## プロセス実行#
exec ( ファイル名 , 同期フラグ = FALSE , x = EMPTY , y = EMPTY , width = EMPTY , height = EMPTY ) #
プロセスを起動します
IDの取得に成功した場合はID0を更新します
パラメータ :
ファイル名 ( 文字列 ) -- 実行するexeのパス
同期フラグ ( 真偽値 省略可 ) --
TRUE: プロセス終了までブロックする
FALSE: プロセス終了を待たずに続行する
x ( 数値 省略可 ) -- ウィンドウ表示位置(X座標)、省略時はウィンドウのデフォルト
y ( 数値 省略可 ) -- ウィンドウ表示位置(Y座標)、省略時はウィンドウのデフォルト
width ( 数値 省略可 ) -- ウィンドウの幅、省略時はウィンドウのデフォルト
height ( 数値 省略可 ) -- ウィンドウの高さ、省略時はウィンドウのデフォルト
戻り値 :
同期フラグTRUE: プロセスの終了コード( 数値 )
同期フラグFALSE: ウィンドウID (取得できなければ-1)
失敗時: -1
shexec ( ファイル , パラメータ = EMPTY ) #
対象ファイルに対してシェルにより指定された動作で実行させます
(「ファイル名を指定して実行」とほぼ同じ)
パラメータ :
ファイル ( 文字列 ) -- 実行するファイルのパス
パラメータ ( 文字列 省略可 ) -- 実行時に付与するパラメータ
戻り値 :
真偽値 正常に実行されれば TRUE
サンプルコード
shexec ( &quot;cmd&quot; , &quot;/k ipconfig&quot; )
## CUIシェル#
doscmd ( コマンド , 非同期 = FALSE , 画面表示 = FALSE , Unicode = FALSE ) #
コマンドプロンプトを実行します
パラメータ :
コマンド ( 文字列 ) -- 実行するコマンド
非同期 ( 真偽値 省略可 ) -- FALSEなら終了するまで待つ
画面表示 ( 真偽値 省略可 ) -- TRUEならコマンドプロンプトを表示する
Unicode ( 真偽値 省略可 ) -- TRUEならUnicode出力
戻り値 :
非同期 と 画面表示 がいずれもFALSEであれば標準出力または標準エラー( 文字列 )を返す、それ以外は EMPTY
サンプルコード
// Unicode出力で文字化けを解消する
cmd = &quot;echo 森鷗外𠮟る 🐶&quot;
print doscmd ( cmd , FALSE , FALSE , FALSE ) // 森?外??る ??
print doscmd ( cmd , FALSE , FALSE , TRUE ) // 森鷗外𠮟る 🐶
powershell ( コマンド , 非同期 = FALSE , 画面表示 = FALSE , プロファイル無視 = FALSE ) #
Windows PowerShell (バージョン6未満)を実行します
pwsh ( コマンド , 非同期 = FALSE , 画面表示 = FALSE , プロファイル無視 = FALSE ) #
PowerShell (バージョン6以降)を実行します
パラメータ :
コマンド ( 文字列 必須 ) -- 実行するコマンド
非同期 ( 真偽値 省略可 ) -- FALSEなら終了するまで待つ
画面表示 ( 真偽値または2 省略可 ) -- TRUEならPowerShellを表示する、2なら表示して最小化
プロファイル無視 ( 真偽値 省略可 ) -- TRUEなら$PROFILEを読み込まない
戻り値 :
非同期と画面表示がいずれもFALSEであれば標準出力( 文字列 )を返す、それ以外は EMPTY
## 入力制御#
lockhard ( フラグ ) #
マウス、キーボードの入力を禁止する
要管理者特権
UWSCRを 管理者として実行 する必要があります
ロックの強制解除
Ctrl+Alt+Delete でロック状態を強制解除できます
ロックしたままでもUWSCRのプロセスが終了すればロックは解除されます
パラメータ :
フラグ ( 真偽値 省略可 ) -- TRUEで入力禁止、FALSEで解除
戻り値の型 :
真偽値
戻り値 :
関数が成功した場合TRUE
lockhardex ( [ ID=EMPTY , モード=LOCK_ALL ] ) #
ウィンドウに対するマウス、キーボードの入力を禁止する
ロックの強制解除
Ctrl+Alt+Delete でロック状態を強制解除できます
ロックしたままでもUWSCRのプロセスが終了すればロックは解除されます
ロック対象について
ロック可能な対象は常に一つです
ロック中に別のウィンドウに対してロックを行った場合元のウィンドウは開放されます
パラメータ :
ID ( 数値 省略可 ) -- 入力を禁止するウィンドウのID、0の場合はデスクトップ全体、EMPTYならロックを解除
モード ( 定数 省略可 ) -- 禁止内容を指定
LOCK_ALL (0)
マウス、キーボードの入力を禁止
LOCK_KEYBOARD
キーボードの入力のみ禁止
LOCK_MOUSE
マウスの入力のみ禁止
戻り値の型 :
真偽値
戻り値 :
関数が成功した場合TRUE
## 音声出力#
sound ( [ 名前=EMPTY , 同期フラグ=FALSE ] ) #
ファイル名、またはサウンドイベント名を指定しそれを再生する
UWSCとの違い
&quot;BEEP&quot; 指定のビープ音再生は廃止されました
代わりに beep 関数を使用してください
注意
wavの再生デバイス選択には対応していません
パラメータ :
名前 ( 文字列 省略可 ) --
ファイル名
再生したいwavファイルのパスを指定
サウンドイベント名
システム上で定義されているサウンドイベント名を指定
ヒント
サウンドイベント名について
環境により登録されているイベント名が異なる可能性があります
以下はWin32のドキュメントに記載されていたイベント名です
SystemAsterisk
SystemExclamation
SystemExit
SystemHand
SystemQuestion
SystemStar
EMPTY
再生を停止します
同期フラグ ( 真偽値 省略可 ) -- TRUEなら再生終了を待つ
戻り値 :
なし
beep ( [ 長さ=300 , 周波数=2000 , 繰り返し=1 ] ) #
ビープ音を再生します
パラメータ :
長さ ( 数値 省略可 ) -- ビープ音を再生する長さをミリ秒で指定
周波数 ( 数値 省略可 ) -- ビープ音の周波数(ヘルツ)を37-32767で指定
繰り返し ( 数値 省略可 ) -- 同じ長さと周波数のビープ音を繰り返し再生する回数
戻り値 :
なし
## キー入力#
getkeystate ( キーコード [ , ID=0 ] ) #
マウスやキーボードがクリックされたかどうか、または特定のキーのトグル状態を得る
パラメータ :
キーコード ( 定数 ) -- VK定数(クリック判定)またはTGL定数(トグル判定)
VK定数
仮想キーコード一覧 を参照
TGL_NUMLOCK
Num Lock
TGL_CAPSLOCK
Caps Lock
TGL_SCROLLLOCK
Scroll Lock
TGL_KANALOCK
カタカナ入力 (要ID指定)
TGL_IME
IME (要ID指定)
ID ( 数値 省略可 ) --
TGL_KANALOCK , TGL_IME にて入力方式を確認したいウィンドウのID
0ならアクティブウィンドウ
戻り値の型 :
真偽値
戻り値 :
クリックされていたらTRUE、またはトグル状態がオンならTRUE
sethotkey ( キーコード [ , 修飾子キー=0 , 関数=EMPTY ] ) #
関数をホットキーに登録
パラメータ :
キーコード ( 定数 ) -- 登録するキーコードをVK定数で指定 (VK定数は 仮想キーコード一覧 を参照)
修飾子キー ( 定数 省略可 ) -- 同時に押す修飾子キーを指定、OR連結で複数指定、0なら修飾子キーなし
MOD_ALT
Altキー
MOD_CONTROL
Controlキー
MOD_SHIFT
Shiftキー
MOD_WIN
Winキー
関数 ( 文字列またはユーザー定義関数 省略可 ) --
ホットキー入力時に実行するユーザー定義関数、またはその名前を文字列で指定
省略時、またはEMPTYや空文字が入力された場合はホットキーを解除する
指定関数の注意点
引数を受ける関数の場合、引数は無視されます (引数0の関数として扱われる)
関数内で引数へのアクセスを行う場合はエラーになります
関数内でエラーが発生した場合はスクリプトが強制終了されます
HOTKEY特殊変数
ホットキーで呼ばれる関数内では以下の変数が使えます
HOTKEY_VK
ホットキーのキーコード
HOTKEY_MOD
ホットキーの修飾子キー
戻り値 :
なし
## 仮想キーコード一覧#
VK_A
VK_B
VK_C
VK_D
VK_E
VK_F
VK_G
VK_H
VK_I
VK_J
VK_K
VK_L
VK_M
VK_N
VK_O
VK_P
VK_Q
VK_R
VK_S
VK_T
VK_U
VK_V
VK_W
VK_X
VK_Y
VK_Z
VK_0
VK_1
VK_2
VK_3
VK_4
VK_5
VK_6
VK_7
VK_8
VK_9
VK_BACK
VK_TAB
VK_CLEAR
VK_ESCAPE
VK_ESC
VK_ENTER
VK_RETURN
VK_RRETURN
VK_SHIFT
VK_RSHIFT
VK_WIN
VK_RWIN
VK_START
VK_MENU
VK_ALT
VK_RALT
VK_CONTROL
VK_CTRL
VK_RCTRL
VK_PAUSE
VK_CAPITAL
VK_KANA
VK_FINAL
VK_KANJI
VK_CONVERT
VK_NONCONVERT
VK_ACCEPT
VK_MODECHANGE
VK_SPACE
VK_PRIOR
VK_NEXT
VK_END
VK_HOME
VK_LEFT
VK_UP
VK_RIGHT
VK_DOWN
VK_SELECT
VK_PRINT
VK_EXECUTE
VK_SNAPSHOT
VK_INSERT
VK_DELETE
VK_HELP
VK_APPS
VK_MULTIPLY
VK_ADD
VK_SEPARATOR
VK_SUBTRACT
VK_DECIMAL
VK_DIVIDE
VK_NUMPAD0
VK_NUMPAD1
VK_NUMPAD2
VK_NUMPAD3
VK_NUMPAD4
VK_NUMPAD5
VK_NUMPAD6
VK_NUMPAD7
VK_NUMPAD8
VK_NUMPAD9
VK_F1
VK_F2
VK_F3
VK_F4
VK_F5
VK_F6
VK_F7
VK_F8
VK_F9
VK_F10
VK_F11
VK_F12
VK_NUMLOCK
VK_SCROLL
VK_PLAY
VK_ZOOM
VK_SLEEP
VK_BROWSER_BACK
VK_BROWSER_FORWARD
VK_BROWSER_REFRESH
VK_BROWSER_STOP
VK_BROWSER_SEARCH
VK_BROWSER_FAVORITES
VK_BROWSER_HOME
VK_VOLUME_MUTE
VK_VOLUME_DOWN
VK_VOLUME_UP
VK_MEDIA_NEXT_TRACK
VK_MEDIA_PREV_TRACK
VK_MEDIA_STOP
VK_MEDIA_PLAY_PAUSE
VK_LAUNCH_MEDIA_SELECT
VK_LAUNCH_MAIL
VK_LAUNCH_APP1
VK_LAUNCH_APP2
VK_OEM_PLUS
VK_OEM_COMMA
VK_OEM_MINUS
VK_OEM_PERIOD
VK_OEM_1
VK_OEM_2
VK_OEM_3
VK_OEM_4
VK_OEM_5
VK_OEM_6
VK_OEM_7
VK_OEM_8
VK_OEM_RESET
VK_OEM_JUMP
VK_OEM_PA1
VK_OEM_PA2
VK_OEM_PA3
VK_LBUTTON
VK_RBUTTON
VK_MBUTTON
## システム制御#
poff ( コマンド [ , スクリプト再実行=TRUE ] ) #
電源等の制御
パラメータ :
コマンド ( 定数 ) -- 制御方法を示す定数
P_POWEROFF
PCの電源オフ
P_SHUTDOWN
PCの電源を切れる状態までOSをシャットダウンする
P_LOGOFF または P_SIGNOUT
現在のユーザーをサインアウトする
P_REBOOT
PCを再起動する
P_SUSPEND または P_HIBERNATE
PCを休止状態にする
システムが休止をサポートしている必要があります
P_SUSPEND2 または P_SLEEP
PCをスリープ状態にする
P_MONIPOWER または P_MONITOR_POWERSAVE
モニタを省電力モードにする
モニタが省電力機能をサポートしている必要があります
P_MONIPOWER2 または P_MONITOR_OFF
モニタの電源を切る
モニタが省電力機能をサポートしている必要があります
P_MONIPOWER3 または P_MONITOR_ON
モニタの電源を入れる
モニタが省電力機能をサポートしている必要があります
P_SCREENSAVE
スクリーンセーバーを起動
P_UWSC_REEXEC
UWSCRの再起動
第二引数がTRUEならスクリプトを再実行する
無限ループに注意
スクリプト再実行を行う場合はpoffの実行条件に注意してください
繰り返しスクリプトの再実行が行われるおそれがあります
コンソールモード中の場合
ウィンドウモードで再実行されます
P_FORCE
アプリケーションの終了を待たずにサインアウトしたい場合や、シャットダウンを強制したい場合に指定
P_POWEROFF , P_SHUTDOWN , P_LOGOFF , P_REBOOT のいずれかに OR で連結指定する
それ以外の場合は無視される
poff ( P_POWEROFF or P_FORCE ) // 強制電源断
スクリプト再実行 ( 真偽値 省略可 ) -- TRUEなら P_UWSC_REEXEC 指定時にスクリプトを再実行する
UWSCとの違い
デフォルト値がTRUEになりました
戻り値 :
なし
OPTFINALLY指定時の動作
自身のプロセス終了を伴う以下のコマンドが try 節で実行された場合
OPTFINALLY指定時に限り finally 節が実行されます
自身を終了する前にfinallyを実行
P_UWSC_REEXEC
コマンド呼び出し前にfinallyを実行 (finally節が終了するまでこれらの処理は行われない)
P_POWEROFF
P_SHUTDOWN
P_LOGOFF
P_REBOOT
OPTION OPTFINALLY
try
poff ( P_UWSC_REEXEC )
msgbox ( &quot;poff以降は実行されない&quot; )
finally
msgbox ( &quot;finallyが実行される&quot; )
endtry
// OPTFINALLYが無い場合
try
poff ( P_UWSC_REEXEC )
finally
msgbox ( &quot;OPTFINALLYがないので実行されない&quot; )
endtry
OPTION OPTFINALLY
// poffがtryの外にある場合
poff ( P_UWSC_REEXEC )
try
finally
msgbox ( &quot;OPTFINALLYがあってもtry外だと実行されない&quot; )
endtry
シャットダウンの理由
poffによるシャットダウンは以下の理由で行われます
SHTDN_REASON_MAJOR_OTHER
SHTDN_REASON_MINOR_OTHER
SHTDN_REASON_FLAG_PLANNED
## 日時#
gettime ( [ 補正値=0 , 基準日時=EMPTY , 補正値オプション=G_OFFSET_DAYS , ミリ秒=FALSE ] ) #
指定日時の2000年1月1日からの経過時間を得る
またその時間に該当する日時情報を G_TIME_* 特殊変数に格納する
パラメータ :
補正値 ( 数値 省略可 ) -- 基準日時を起点として指定値分ずらした日時を得る
基準日時 ( 文字列 省略可 ) --
基準となる日時を指定、EMPTYで現在時刻
以下の形式で指定
&quot;YYYYMMDD&quot;
&quot;YYYY/MM/DD&quot;
&quot;YYYY-MM-DD&quot;
&quot;YYYYMMDDHHNNSS&quot;
&quot;YYYY/MM/DD HH:NN:SS&quot;
&quot;YYYY-MM-DD HH:NN:SS&quot;
RFC 3339 形式
タイムゾーン情報を含めた場合ローカル時間に変換されます
補正値オプション ( 定数 省略可 ) -- 補正値の指定方法を指定
G_OFFSET_DAYS
補正値を日数として扱う
G_OFFSET_HOURS
補正値を時間として扱う
G_OFFSET_MINUTES
補正値を分として扱う
G_OFFSET_SECONDS
補正値を秒として扱う
G_OFFSET_MILLIS
補正値をミリ秒として扱う
ミリ秒 ( 真偽値 省略可 ) -- 戻り値を秒ではなくミリ秒で返す
戻り値の型 :
数値
戻り値 :
2000年1月1日からの秒数
実行後に変更される特殊変数
以下の変数がgettime関数の結果に応じて変更されます
文字列型の場合は桁数分左0埋めされます
これらの変数の変更が適用されるのはgettimeを呼び出したスコープ内に限定されます
変数
型
内容
G_TIME_YY
数値
年
G_TIME_MM
数値
月
G_TIME_DD
数値
日
G_TIME_HH
数値
時
G_TIME_NN
数値
分
G_TIME_SS
数値
秒
G_TIME_ZZ
数値
ミリ秒
G_TIME_WW
数値
曜日 (0:日,1:月,2:火,3:水,4:木,5:金,6:土)
G_TIME_YY2
文字列
年 (下2桁)
G_TIME_MM2
文字列
月 (2桁)
G_TIME_DD2
文字列
日 (2桁)
G_TIME_HH2
文字列
時 (2桁)
G_TIME_NN2
文字列
分 (2桁)
G_TIME_SS2
文字列
秒 (2桁)
G_TIME_ZZ2
文字列
ミリ秒 (3桁)
G_TIME_YY4
文字列
年 (4桁)
以下の定数で G_TIME_WW と比較ができます
G_WEEKDAY_SUN (0)
G_WEEKDAY_MON (1)
G_WEEKDAY_TUE (2)
G_WEEKDAY_WED (3)
G_WEEKDAY_THU (4)
G_WEEKDAY_FRI (5)
G_WEEKDAY_SAT (6)
サンプルコード
dt = &quot;2023/04/01 10:10:10&quot;
// 戻り値
print gettime ( , dt ) // 733659010
// ミリ秒で返す
print gettime ( , dt ,, TRUE ) // 733659010000
// 第一引数の単位変更
// 従来の書き方で6時間ずらす
ts1 = gettime ( 0 . 25 , dt )
// 第三引数指定で時間扱いになる
ts2 = gettime ( 6 , dt , G_OFFSET_HOURS )
assert_equal ( ts1 , ts2 )
// format関数による日時フォーマット
ts = gettime ( , dt )
print format ( ts , &quot;%c&quot; ) // 2023年04月01日 10時10分10秒
// RFC3339形式
ts = gettime ( , &quot;2023-10-10T00:00:00+0000&quot; )
print format ( ts , &quot;%c&quot; ) // 2023年10月10日 09時00分00秒
## 音声#
speak ( 発声文字 [ , 非同期=FALSE , 中断=FALSE ] ) #
指定文字列を音声として再生する
パラメータ :
発声文字 ( 文字列 ) -- 発声させたい文字列
非同期 ( 真偽値 省略可 ) -- TRUEなら非同期で発声、FALSEなら発声終了を待つ
中断 ( 真偽値 省略可 ) -- 別の音声が発生中の場合にTRUEなら中断し、FALSEなら終了を待ってから発声させる
終了待ちについて
音声の終了待ちは、speak関数を非同期TRUEで事前に実行していた場合のみ有効です
また、speak関数は同一スレッド上で実行されている必要があります
戻り値 :
なし
recostate ( 開始フラグ [ , 登録単語... ] ) #
任意の単語を登録し音声認識を開始、または終了する
有効範囲はスレッド単位
登録及び解除は同一スレッド上でのみ有効です
パラメータ :
開始フラグ ( 真偽値 ) -- TRUEで音声認識を開始、FALSEで解除
登録単語 ( 文字列または文字列の配列 ) -- 開始フラグがTRUEの場合に音声認識させたい言葉を指定、未指定の場合は認識エンジンの標準辞書を使用する
戻り値の型 :
文字列
戻り値 :
使用する認識エンジン名、登録失敗時はEMPTY
dictate ( [ 拾得待ち=TRUE , タイムアウト=10000 ] ) #
recostate関数で登録した単語が音声入力されたらその文字列を返す
単語登録について
recostate関数が開始フラグTRUEで実行されていない場合は即座に終了しEMPTYを返します
また、recostate関数は同一スレッド上で実行されている必要があります
パラメータ :
拾得待ち ( 真偽値 省略可 ) -- TRUEなら入力を待つ、FALSEなら直近の入力を返す(入力がなければEMPTYを返す)
タイムアウト ( 数値 省略可 ) -- 拾得待ちがTRUEだった場合に待機する時間(ミリ秒)、0なら無限に待つ、拾得待ちFALSEなら無視される
戻り値の型 :
文字列
戻り値 :
拾得した文字列、拾得待ちTRUEでタイムアウトまたは拾得待ちFALSEで入力がない場合はEMPTY
サンプルコード
// 単語を登録する
print recostate ( TRUE , &quot;りんご&quot; , &quot;みかん&quot; , &quot;いちご&quot; , &quot;おわり&quot; )
print &quot;「りんご」「みかん」「いちご」に反応します&quot;
print &quot;「おわり」で終了&quot;
while TRUE
select word := dictate ( TRUE )
case &quot;おわり&quot;
print &quot;終了します&quot;
break
case EMPTY
// デフォルトでは10000ミリ秒経過でタイムアウト
print &quot;タイムアウトしました&quot;
break
default
print word
selend
wend
// 登録を解除
recostate ( false )

---

# ファイル操作関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/file.html

## ファイル操作関数#
## テキストファイル#
fopen ( ファイルパス [ , モード=F_READ , 追記文字列=EMPTY ] ) #
テキストファイルを開きます
UWSCとの違い
戻り値のファイルIDが数値ではなくなりました
パラメータ :
ファイルパス ( 文字列 ) -- 開きたいテキストファイルのパス
モード ( 定数 省略可 ) -- どのようにファイルを開くかを指定、 OR 連結可
F_EXISTS
パスの存在確認 (ディレクトリの場合は末尾に \ を付ける)
戻り値が真偽値になる
F_READ
ファイルを読む (SJIS、UTF8、UTF16対応)
F_WRITE
F_READ と併記しない場合はファイルを上書きする (UTF-8)
F_READ と併記した場合は元ファイルのエンコーディングを維持する
F_WRITE1
SJISで書き込む
F_WRITE8
UTF-8で書き込む
F_WRITE8B
BOM付きUTF-8で書き込む
F_WRITE16
UTF-16LEで書き込む
F_APPEND
文末に追記し、即ファイルを閉じる
追記文字列 を必ず指定する
F_WRITE 系と併記で書き込む文字列のエンコーディングを指定できる
戻り値が書き込んだバイト数になる
F_NOCR
文末に改行を入れない
F_TAB
CSVセパレータをカンマではなくタブ文字にする
F_EXCLUSIVE
排他モードでファイルを開く
F_AUTOCLOSE
ファイルIDが破棄された際に自動でファイルをクローズする
自動クローズについて
ファイルIDオブジェクトの参照がすべて失われた場合に自動クローズ処理が実施されます
// fopenが返すファイルIDをfputが処理した直後にファイルIDが失われるため自動クローズされhogehogeが書き込まれる
fput ( fopen ( &quot;hgoe.txt&quot; , F_WRITE or F_AUTOCLOSE ) , &quot;hogehoge&quot; )
fid = fopen ( &quot;fuga.txt&quot; , F_WRITE or F_AUTOCLOSE )
fput ( fid , &quot;fugafuga&quot; )
fid = EMPTY // この時点でfuga.txtが閉じられる
fid = fopen ( &quot;piyo.txt&quot; , F_WRITE or F_AUTOCLOSE )
// ファイルIDのコピー
fid2 = fid
fput ( fid , &quot;fugafuga&quot; )
fid = EMPTY // この時点ではまだpiyo.txtは閉じられない
msgbox ( 1 )
fid2 = EMPTY // ここでpiyo.txtが閉じられる
追記文字列 ( 文字列 省略可 ) -- F_APPEND 指定時に追記する文字列
戻り値 :
モードによる
真偽値
F_EXISTS 指定時、ファイルまたはディレクトリが存在する場合はTRUE
数値
F_APPEND 指定時、書き込んだバイト数
ファイルID
F_EXISTS , F_APPEND 以外を指定した場合、開いたファイルを示すIDを返す
ファイルが開けない場合の動作について
UWSCでは-1を返していましたが、UWSCRでは実行時エラーとなりファイルが開けない理由を明確にします。
例として、以下のような状況でエラーとなります
F_READ のみを指定し存在しないファイルを開こうとした場合 (読み出すファイルが無いため)
F_WRITE が含まれていて、読み取り専用のファイルを開こうとした場合 (書き込めないため)
fget ( ファイルID , 行 [ , 列=0 , ダブルクォート無視=FALSE ] ) #
ファイルを読み取ります
使用条件
F_READ を指定してファイルを開く必要があります
パラメータ :
ファイルID ( ファイルID ) -- fopen で開いたファイルのID
行 ( 数値 ) -- 読み取る行の番号、または以下の定数を指定 (定数指定時は以降の引数は無視される)
F_LINECOUNT
ファイルの行数を返す
F_ALLTEXT
ファイル全体のテキストを返す
列 ( 数値 ) -- 読み取るcsv列の番号 (1から)、0の場合は行全体
ダブルクォート無視 ( 真偽値または2 省略可 ) -- 列が1以上 (csv読み取り) の場合に有効
FALSE
両端のダブルクォートを削除する
TRUE
両端にダブルクォートがあってもなにもしない
2
連続するダブルクォート ( &quot;&quot; ) を単一のダブルクォート ( &quot; ) にする
その後両端のダブルクォートを削除する (FALSEと同じ処理)
戻り値 :
読み取った文字列
該当行または列が存在しない場合は EMPTY
EMPTYについて
UWSCでは指定行および列が存在しない場合に空文字( &quot;&quot; )を返していましたが、UWSCRでは EMPTY を返すように変更しています
これにより空文字を読み取った場合と、不正な行や列を読み取った場合を区別できるようになりました
サンプルコード
test.csv
foo,bar,baz
foo , bar , baz
&quot;ダブルクォートありのカラム&quot;,&quot;ダブルクォートの&quot;&quot;エスケープ&quot;&quot;&quot;,&quot;&quot;
スクリプト
fid = fopen ( &quot;test.csv&quot; , F_READ )
print fget ( fid , 1 ) // foo,bar,baz
print fget ( fid , 1 , 1 ) // foo
// 前後のホワイトスペースはトリムされる
print fget ( fid , 2 , 1 ) // 「 foo 」にはならず「foo」が返る
// ダブルクォートで括られたカラム
print fget ( fid , 3 , 1 , FALSE ) // ダブルクォートありのカラム
print fget ( fid , 3 , 1 , TRUE ) // &quot;ダブルクォートありのカラム&quot;
// 第4引数FALSEはUWSCにおける 2 の動作が標準になりました
print fget ( fid , 3 , 2 , FALSE ) // ダブルクォートの&quot;エスケープ&quot;
print fget ( fid , 3 , 2 , TRUE ) // &quot;ダブルクォートの&quot;&quot;エスケープ&quot;&quot;&quot;
fclose ( fid )
fput ( ファイルID , 値 [ , 行=0 , 列=0 ] ) #
ファイルに書き込みます
使用条件
F_WRITE 系を指定してファイルを開く必要があります
パラメータ :
ファイルID ( ファイルID ) -- fopen で開いたファイルのID
値 ( 文字列 ) -- 書き込む文字列
行 ( 数値 省略可 ) -- 書き込む行を指定
0
文末に新たな行として書き加えます
1以上
指定行に書き込みます (上書き)
F_ALLTEXT (定数)
ファイル全体を書き込む値で上書きします
列 ( 数値 省略可 ) -- 書き込むCSV列を指定
0
行全体に書き込み
1以上
CSVカラムとして書き込み
F_INSERT (定数)
指定した行へ上書きではなく挿入します
F_READ が未指定の場合無視されます
戻り値 :
なし
fdelline ( ファイルID , 行 ) #
指定行を削除します
使用条件
F_READ および F_WRITE 系を指定してファイルを開く必要があります
パラメータ :
ファイルID ( ファイルID ) -- fopen で開いたファイルのID
行 ( 数値 ) -- 削除する行の番号 (1から)、該当行がない場合なにもしない
戻り値 :
なし
fclose ( ファイルID [ , エラー抑止=FALSE ] ) #
ファイルを閉じて変更を適用します
ファイルの更新について
ファイルを閉じない限り fput や fdelline による変更はファイルに反映されません
パラメータ :
ファイルID ( ファイルID ) -- fopen で開いたファイルのID
エラー抑止 ( 真偽値 省略可 ) -- TRUEにするとファイル書き込み時のエラーを無視する
戻り値 :
ファイルへの書き込みが行われ正常に閉じられた場合はTRUE
サンプルコード
// 読み取り
fid = fopen ( path ) // fopen(path, F_READ) と同等
print fget ( fid , 1 )
fclose ( fid )
// 書き込み
fid = fopen ( path , F_WRITE )
fput ( fid , text )
fclose ( fid ) // 上書きされる
// 読み書き
fid = fopen ( path , F_READ or F_WRITE )
print fget ( fid , 1 )
fput ( fid , text )
fclose ( fid ) // 編集して保存
// エンコーディングを変更して保存
fid = fopen ( path , F_WRITE1 ) // SJISでファイルを書き込み
fput ( fid , text1 )
fclose ( fid )
fid = fopen ( path , F_READ or F_WRITE16 )
fput ( fid , text2 )
fclose ( fid ) // 編集してUTF-16で保存
// 追記
fopen ( path , F_APPEND or F_WRITE16 , text ) // UTF-16で末尾に追記
fopen ( path , F_APPEND ) // エラー; F_APPEND指定時は第三引数が必須
// 自動ファイルクローズ
print fput ( fopen ( path , F_WRITE or F_AUTOCLOSE ) , &quot;auto close&quot; )
// F_AUTOCLOSEによりfput実行後にファイルが自動でクローズされる
## CSVファイル#
csvopen ( CSVパス [ , ヘッダ有無=FALSE , TSVモード=FALSE ] ) #
CSVファイルを開く
パラメータ :
CSVパス ( 文字列 ) -- CSVファイルのパス
ヘッダ有無 ( 真偽値 省略可 ) -- 対象CSVファイルにヘッダ行があるかどうか
TSVモード ( 真偽値または文字 省略可 ) -- FALSEの場合はカンマ区切り、TRUEにするとタブ文字区切り、または任意のASCII文字
戻り値の型 :
CSVオブジェクト
戻り値 :
CSVファイルを示すオブジェクト、各種csv関数で利用される
csvclose ( csv ) #
編集したCSVをファイルに書き出す
バッファに変更があった場合のみ対象ファイルに書き込みを行う
csvopen で指定したファイルが存在しない場合は新しいファイルが作成される
この関数呼び出し後のCSVオブジェクトに対して再度この関数を実行しても書き込みは行われない
自動クローズ
CSVオブジェクトが破棄された場合は自動でこの関数と同等の処理が行われます
( fopen の F_AUTOCLOSE 指定時と同様です)
クローズ後のCSVオブジェクトについて
バッファに対する読み書きはできますが、再度 csvclose で書き込みを行うことはできません
パラメータ :
csv ( CSVオブジェクト ) -- csvopen の戻り値
戻り値 :
なし
csvread ( csv [ , 行 , 列 ] ) #
CSVバッファから値を読み出します
パラメータ :
csv ( CSVオブジェクト ) -- csvopen の戻り値
行 ( 数値 省略可 ) -- CSVの行番号
列 ( 数値または文字列 省略可 ) -- CSVの列番号、またはヘッダのカラム名
戻り値の型 :
文字列または配列
戻り値 :
行列の指定方法により得られる値が変わります
行
列
値
省略
省略
CSV全体の文字列
省略
1以上
該当列の配列
0
省略
ヘッダ行の配列
1以上
省略
該当行の配列
1以上
1以上
該当行及び列の文字列
サンプルコード
test.csv #
項目1 , 項目2 , 項目3
1 , 2 , 3
10 , 20 , 30
100 , 200 , 300
// ヘッダ行を有効にして開く
csv = csvopen ( &quot;test.csv&quot; , true )
// CSV全体を得る
print csvread ( csv )
// 項目1,項目2,項目3
// 1,2,3
// 10,20,30
// 100,200,300
// ヘッダ行の配列を得る
print csvread ( csv , 0 ) // [項目1, 項目2, 項目3]
// 2行目の配列を得る
print csvread ( csv , 2 ) // [10, 20, 30]
// 2行目1列目の文字列を得る
print csvread ( csv , 2 , 1 ) // 10
// 1列目の配列を得る
print csvread ( csv , , 1 ) // [1, 10, 100]
// 列をカラム名で指定
print csvread ( csv , 3 , &quot;項目2&quot; ) // 200
print csvread ( csv , , &quot;項目3&quot; ) // [3, 30, 300]
csvwrite ( csv , 行 , 列 , 値 ) #
CSVバッファに書き込みを行う
パラメータ :
csv ( CSVオブジェクト ) -- csvopen の戻り値
行 ( 数値 ) -- CSVの行番号
列 ( 数値または文字列 ) -- CSVの列番号、またはヘッダのカラム名
値 ( 文字列または配列 ) -- 書き込む値
戻り値の型 :
真偽値
戻り値 :
書き込み時true
サンプルコード
new_csv = &quot;new.csv&quot;
deletefile ( new_csv )
// ファイルを新規作成
csv = csvopen ( new_csv , true )
// 0行目指定でヘッダを書き込む
csvwrite ( csv , 0 , 1 , &quot;項目1&quot; )
csvwrite ( csv , 0 , 2 , &quot;項目2&quot; )
csvwrite ( csv , 0 , 3 , &quot;項目3&quot; )
// 指定位置に書き込み
csvwrite ( csv , 1 , 1 , &quot;1-1&quot; )
// 配列指定で複数列書き込み
csvwrite ( csv , 2 , 1 , [ &quot;2-1&quot; , &quot;2-2&quot; , &quot;2-3&quot; ])
// 3行目を飛ばして4行目に書き込み
csvwrite ( csv , 4 , 1 , [ &quot;4-1&quot; , &quot;4-2&quot; , &quot;4-3&quot; ])
// 2列目から書き込み
csvwrite ( csv , 5 , 2 , [ &quot;5-2&quot; , &quot;5-3&quot; ])
// 列の数は可変
csvwrite ( csv , 6 , 1 , [ &quot;6-1&quot; , &quot;6-2&quot; , &quot;6-3&quot; , &quot;6-4&quot; , &quot;6-5&quot; ])
// 全体読み出し
print csvread ( csv )
// 項目1,項目2,項目3
// 1-1
// 2-1,2-2,2-3
// &quot;&quot;
// 4-1,4-2,4-3
// ,5-2,5-3
// 6-1,6-2,6-3,6-4,6-5
csvclose ( csv ) // 保存
// 書き出したファイルも確認
print fget ( fopen ( new_csv , F_READ or F_AUTOCLOSE ) , F_ALLTEXT )
// 項目1,項目2,項目3
// 1-1
// 2-1,2-2,2-3
// &quot;&quot;
// 4-1,4-2,4-3
// ,5-2,5-3
// 6-1,6-2,6-3,6-4,6-5
## iniファイル#
readini ( [ セクション=EMPTY , キー=EMPTY , ファイル=&quot;&lt;#GET_UWSC_NAME&gt;.ini&quot; ] ) #
iniファイルを読み込みます
パラメータ :
セクション ( 文字列 省略可 ) -- 読み出したいキーのあるセクション名を指定、省略時はセクション一覧を得る
キー ( 文字列 省略可 ) -- 値を読み出したいキーの名前を指定、省略時はキー一覧を得る
ファイル ( 文字列またはファイルID 省略可 ) -- 読み出すiniファイルのパス、またはファイルID
ファイルIDを利用する場合
F_READ を含めてfopenしている必要があります
戻り値 :
セクション省略時
iniファイルのセクション一覧を格納した配列
セクション省略時のキー指定は無視されます
キーを省略
指定セクションのキー一覧を格納した配列
セクションとキーを指定
該当キーの値
該当キーが存在しない場合EMPTY
サンプルコード
test.ini
[section]
key1 = &quot;あ&quot;
key2 = &quot;い&quot;
key3 = &quot;う&quot;
[foo]
name = &quot;foo&quot;
[bar]
name = &quot;bar&quot;
[baz]
name = &quot;baz&quot;
スクリプト
ini = &#39;test.ini&#39;
print readini ( &#39;foo&#39; , &#39;name&#39; , ini ) // foo
// セクションを省略(またはEMPTY指定)するとセクション一覧を取得
print readini ( , , ini ) // [ section, foo, bar, baz ]
print readini ( , &#39;name&#39; , ini ) // ↑と同じ結果 (セクション省略時のキーは無視される)
// セクションを指定してキーを省略(またはEMPTY指定)するとキー一覧を収録
print readini ( &#39;section&#39; , , ini ) // [ key1, key2, key3 ]
writeini ( セクション , キー , 値 [ , ファイル=&quot;&lt;#GET_UWSC_NAME&gt;.ini&quot; ] ) #
iniファイルに書き込みます
パラメータ :
セクション ( 文字列 ) -- 書き込みたいキーのあるセクション名、存在しない場合新規に作成されます
キー ( 文字列 ) -- 書き込みたいキーの名前、存在しない場合新規に作成されます
値 ( 文字列 ) -- 該当キーに書き込む値
ファイル ( 文字列またはファイルID 省略可 ) -- 書き込むiniファイルのパス、またはファイルID
ファイルIDを利用する場合
ファイルIDは F_READ 及び F_WRITE 系を含めてfopenしている必要があります
また、ファイルIDを渡した場合はfcloseを呼ぶまで変更が反映されません
戻り値 :
なし
deleteini ( セクション [ , キー=EMPTY , ファイル=&quot;&lt;#GET_UWSC_NAME&gt;.ini&quot; ] ) #
指定キーまたはセクションを削除します
パラメータ :
セクション ( 文字列 ) -- 削除したいキーのあるセクション名
キー ( 文字列 ) -- 削除したいキーの名前
ファイル ( 文字列またはファイルID 省略可 ) -- 書き込むiniファイルのパス、またはファイルID
ファイルIDを利用する場合
ファイルIDは F_READ 及び F_WRITE 系を含めてfopenしている必要があります
また、ファイルIDを渡した場合はfcloseを呼ぶまで変更が反映されません
戻り値 :
なし
## INI関数のファイルID利用について#
iniファイルをfopenで開き、そのファイルIDを各種ini関数に渡すことでiniファイルの読み書きができるようになりました
サンプルコード
fid = fopen ( &quot;hoge.ini&quot; , F_READ or F_WRITE )
// ファイルパスの代わりにファイルIDを指定
print readini ( &quot;hoge&quot; , &quot;fuga&quot; , fid ) // 読む場合はF_READが必要
writeini ( &quot;hoge&quot; , &quot;fuga&quot; , &quot;fugafuga&quot; , fid ) // 書き込みにはF_READ or F_WRITEが必要
deleteini ( &quot;hoge&quot; , &quot;fuga&quot; , fid ) // 削除にもF_READ or F_WRITEが必要
fclose ( fid ) // iniファイルへの書き込みが反映される
以下のような用途を想定しています
同一iniファイルへの複数回の読み書きを行う場合にファイルアクセスを減らしたい
iniファイル編集時に排他制御( F_EXCLUSIVE )したい
## その他のファイル操作#
deletefile ( ファイルパス ) #
ファイルを削除します
* , ? によるワイルドカード指定も可能
パラメータ :
ファイルパス ( 文字列 ) -- 削除したいファイルのパス
戻り値 :
該当ファイルすべてを削除できた場合TRUE、一つでも該当ファイルが削除できなかった場合は該当ファイルが存在しない場合はFALSE
ワイルドカード指定時の動作について
UWSCではワイルドカード指定時に削除できないファイルが含まれていたとしても別のファイルが一つでも削除できればTRUEを返していましたが、UWSCRでは一つでも削除できないファイルが含まれていればFALSEを返します
getdir ( ディレクトリパス [ , フィルタ=&quot;*&quot; , 非表示ファイル=FALSE , 取得順=ORDERBY_NAME ] ) #
対象ディレクトリに含まれるファイル、またはディレクトリの一覧を取得します
パラメータ :
ディレクトリパス ( 文字列 ) -- 対象ディレクトリのパス
フィルタ ( 文字列 省略可 ) --
ファイル名のフィルタ、ワイルドカード( * , ? )可
\ のみ、または \ から始まる文字列指定でファイルではなくディレクトリ一覧を返す
非表示ファイル ( 真偽値 省略可 ) -- 非表示ファイルを含めるかどうか
取得順 ( 定数 省略可 ) -- 取得順を示す定数
ORDERBY_NAME
ファイル名順
ORDERBY_SIZE
サイズ順
ORDERBY_CREATED
作成日時順
ORDERBY_MODIFIED
更新日時順
ORDERBY_ACCESSED
最終アクセス日時順
戻り値 :
該当するファイル名またはディレクトリ名の一覧を格納した配列
UWSCとの違い
該当ファイルの個数ではなく配列が返るようになりました
それに伴い特殊変数 GETDIR_FILES は廃止されました
サンプルコード
ファイル構成
C:\test\
├ foo1.txt
├ foo2.txt
├ bar.txt
├ baz.txt
├ hidden.txt (隠しファイル)
├ dir1\
├ dir2\
├ folder1\
└ folder2\
スクリプト
// ファイル一覧の表示
print getdir ( &#39;C:\test&#39; ) // [foo1.txt, foo2.txt, bar.txt, baz.txt]
// ファイル名のフィルタ
print getdir ( &#39;C:\test&#39; , &#39;foo*&#39; ) // [foo1.txt, foo2.txt]
// 隠しファイルも表示
print getdir ( &#39;C:\test&#39; , , TRUE ) // [foo1.txt, foo2.txt, bar.txt, baz.txt, hidden.txt]
// フォルダ一覧の表示
print getdir ( &#39;C:\test&#39; , &#39;\&#39; ) // [dir1, dir2, folder1, folder2]
// フォルダ一名のフィルタ
print getdir ( &#39;C:\test&#39; , &#39;\dir*&#39; ) // [dir1, dir2]
dropfile ( ID , ディレクトリ , ファイル名 [ , ファイル名... ] ) #
ファイルをウィンドウにドロップします
ドロップ位置はクライアント領域の中央です
パラメータ :
ID ( 数値 ) -- ファイルをドロップするウィンドウのID
ディレクトリ ( 数値 ) -- ドロップするファイルの存在するディレクトリパス
ファイル名 ( 文字列または配列 ) -- ファイル名を示す文字列、またはファイル名を示す文字列を含む配列変数
戻り値 :
なし
dropfile ( ID , x , y , ディレクトリ , ファイル名 [ , ファイル名... ] )
第二、第三引数が数値だった場合はファイルのドロップ座標を指定します
対象ウィンドウのクライアント座標を指定します
パラメータ :
x ( 数値 ) -- クライアントX座標
y ( 数値 ) -- クライアントY座標
ファイル名指定数の下限および上限
上限は座標未指定時は34、座標指定時は32個まで (すべての引数の個数上限が36)
ファイル数がそれより多い場合は配列変数を使ってください
下限は1です (最低1つ指定する必要がある)
マウス移動が行われます
ドロップ処理時に瞬間的にマウスカーソルを指定座標に移動しています
(UWSCと同様の処理)
実行要件
対象ウィンドウが WM_DROPFILES メッセージを処理できる必要があります
## ZIPファイル#
zip ( zipファイル , ファイル [ , ファイル , ... ] ) #
zipファイルを作成します
パラメータ :
zipファイル ( 文字列 ) -- 作成するzipファイルのパス
ファイル ( 文字列または配列 ) --
zipファイルに含めたいファイルのパス (10個まで)
パスの配列を渡すこともできる
格納されるファイルのパス構成について
指定したパスがファイルの場合はそのファイル名でzipに格納します
フォルダが指定された場合はそのフォルダ以下のすべてのファイルをフォルダからの相対パスでzipに格納します
戻り値 :
成功時TRUE
サンプルコード
files = [
&#39;foo.uws&#39; ,
&#39;bar.uws&#39; ,
&#39;baz.uws&#39; ,
&#39;modules&#39; , // フォルダ指定
]
zip ( &quot;test.zip&quot; , files )
unzip ( zipファイル , 展開先フォルダ ) #
zipファイルを指定フォルダに展開します
展開先フォルダが存在しない場合は新規に作成されます
すでに同名ファイルが存在する場合は上書きされます
パラメータ :
zipファイル ( 文字列 ) -- 展開したいzipファイルのパス
展開先フォルダ ( 文字列 ) -- 展開先フォルダのパス
戻り値 :
成功時TRUE
ヒント
失敗した場合でも一部のファイルが展開されることがあります
サンプルコード
unzip ( &quot;test.zip&quot; , &quot;out&quot; )
for file in getdir ( &quot;out&quot; )
print file
next
for dir in getdir ( &#39;out&#39; , &#39;\&#39; )
for file in getdir ( &quot;out\&lt;#dir&gt;&quot; )
print &quot;&lt;#dir&gt;/&lt;#file&gt;&quot;
next
next
// foo.uws
// bar.uws
// baz.uws
// modules\qux.uws
// modules\quux.uws
zipitems ( zipファイル ) #
zipファイルに含まれるファイル一覧を取得します
パラメータ :
zipファイル ( 文字列 ) -- zipファイルのパス
戻り値 :
ファイル名を格納した配列 (フォルダの区切りは / )
サンプルコード
for item in zipitems ( &quot;test.zip&quot; )
print item
next
// foo.uws
// bar.uws
// baz.uws
// modules\qux.uws
// modules\quux.uws

---

# GUI - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/gui.html

## GUI#
## ダイアログ#
msgbox ( メッセージ [ , ボタン種=BTN_OK , x=EMPTY , y=EMPTY , フォーカス=EMPTY , リンク表示=FALSE ] ) #
メッセージボックスを表示します
クラス名
メッセージボックスのクラス名は UWSCR.MsgBox です
パラメータ :
メッセージ ( 文字列 ) -- ダイアログに表示するメッセージ
ボタン種 ( ボタン定数 省略可 ) -- 表示するボタンを示す定数、 OR 連結で複数表示
BTN_YES
はい
BTN_NO
いいえ
BTN_OK
OK
BTN_CANCEL
キャンセル
BTN_ABORT
中止
BTN_RETRY
再試行
BTN_IGNORE
無視
x ( 数値 省略可 ) -- ダイアログの初期表示位置のX座標を指定、省略時(EMPTY)なら画面中央
y ( 数値 省略可 ) -- ダイアログの初期表示位置のY座標を指定、省略時(EMPTY)なら画面中央
前回表示位置に表示
x, yに-1を指定するとそれぞれ前回表示した位置になります
フォーカス ( ボタン定数 省略可 ) -- カーソルの初期位置をボタン定数で指定、省略時や該当ボタンがない場合は一番左のボタンがフォーカスされます
リンク表示 ( 真偽値 省略可 ) -- TRUEであればURLをクリック可能なリンクにする
戻り値 :
押されたボタンを示すボタン定数 (×ボタンで閉じられた場合は BTN_CANCEL )
input ( メッセージ [ , デフォルト値=EMPTY , マスク表示=FALSE , x=EMPTY , y=EMPTY ] ) #
インプットボックスを表示します
クラス名
インプットボックスのクラス名は UWSCR.Input です
パラメータ :
メッセージ ( 文字列または配列 ) --
文字列
メッセージ欄に表示されるメッセージ
文字列の配列
1番目がメッセージ欄に表示される
2番目以降はラベルとして表示され、ラベル毎に入力欄が追加される
ラベルは最大5つまで
デフォルト値 ( 文字列または配列 省略可 ) --
文字列
入力欄に予め入力しておく値
文字列の配列
入力欄毎のデフォルト入力値
マスク表示 ( 真偽値または配列 省略可 ) --
真偽値
入力欄をマスク表示するかどうか
真偽値の配列
入力欄毎のマスク設定
x ( 数値 省略可 ) -- ダイアログの初期表示位置のX座標を指定、省略時(EMPTY)なら画面中央
y ( 数値 省略可 ) -- ダイアログの初期表示位置のY座標を指定、省略時(EMPTY)なら画面中央
前回表示位置に表示
x, yに-1を指定するとそれぞれ前回表示した位置になります
戻り値 :
入力欄が一つの場合
入力された値、キャンセル時はEMPTY
入力欄が複数の場合
それぞれに入力された値の配列、キャンセル時は空配列
サンプルコード
// ラベルを2つ指定し入力欄を2つにする
labels = [ &#39;ログイン&#39; , &#39;ユーザー名&#39; , &#39;パスワード&#39; ]
// 1つ目の入力欄のみデフォルト値を入れる
default = [ &#39;UserA&#39; , EMPTY ]
// 2つ目の入力欄がマスクされるようにする
mask = [ FALSE , TRUE ]
// 入力値は配列で返る
user = input ( labels , default , mask )
print &#39;ユーザー名: &#39; + user [ 0 ]
print &#39;パスワード: &#39; + user [ 1 ]
ファイルのドラッグアンドドロップについて
インプットボックスに対してExplorerなどからファイルをドラッグアンドドロップすると、そのファイルのパスが入力されます
複数ファイルをドロップした場合はパスがタブ文字で連結されて入力されます
入力欄が複数ある場合は、ドロップした入力欄にパスが挿入されます
入力欄以外にドロップされた場合は一つ目の入力欄にパスが入ります
slctbox ( 表示方法 , タイムアウト秒 , メッセージ=EMPTY , 表示項目 [ , 表示項目2 , ... , 表示項目31 ] ) #
slctbox ( 表示方法 , タイムアウト秒 , x , y , メッセージ=EMPTY , 表示項目 [ , 表示項目2 , ... , 表示項目29 ] )
セレクトボックスを表示します
クラス名
セレクトボックスのクラス名は UWSCR.SlctBox です
引数x, yについて
第3、第4引数が数値であった場合はx, yが指定されたものとします
&quot;100&quot; など数値に変換できる文字列であってもここでは数値として扱われません
x, yの有無による表示項目として渡せる引数の数が変わります
パラメータ :
表示方法 ( SLCT定数 ) -- 項目の表示方法および戻り値の形式を示す定数
表示方法と戻り値の形式をそれぞれ一つずつ OR で連結できます
表示方法
SLCT_BTN
ボタン
SLCT_CHK
チェックボックス
SLCT_RDO
ラジオボタン
SLCT_CMB
コンボボックス
SLCT_LST
リストボックス
戻り値の形式
SLCT_STR
項目名を返す
SLCT_NUM
インデックス番号で返す
タイムアウト秒 ( 数値 ) -- 指定秒数経過で自動的にダイアログを閉じる (キャンセル扱い)、0ならタイムアウトなし
x ( 数値 省略可 ) -- ダイアログの初期表示位置のX座標を指定、省略時(EMPTY)なら画面中央
y ( 数値 省略可 ) -- ダイアログの初期表示位置のY座標を指定、省略時(EMPTY)なら画面中央
前回表示位置に表示
x, yに-1を指定するとそれぞれ前回表示した位置になります
メッセージ ( 文字列 省略可 ) -- メッセージ欄に表示されるメッセージ
表示項目 ( 文字列または配列 ) -- 表示される項目名、または項目名を格納した配列
表示項目2-31 ( 文字列または配列 ) -- 表示される項目名、または項目名を格納した配列
戻り値 :
SLCT_NUM および SLCT_STR 未指定時
選択項目に応じた定数が返る
n番目の項目が選ばれれば SLCT_n
SLCT_1 から SLCT_31 まで
SLCT_CHK, SLCT_LST 以外
選択項目を示す値が返る
SLCT_CHK, SLCT_LST 指定時
選択項目の値が合算される
例
3番目と5番目が選ばれた場合 SLCT_3 or SLCT_5 が返る
警告
表示項目の配列指定で項目数が31を超える場合に、32個目以上を選択するとエラーになります
SLCT_NUM 指定時
SLCT_CHK, SLCT_LST 以外
選択位置のインデックス値(0から)が返る
SLCT_CHK, SLCT_LST 指定時
選択位置のインデックス値を格納した配列
注釈
項目数が31を超えてもOK
SLCT_STR 指定時
SLCT_CHK, SLCT_LST 以外
選択した項目の表示名
SLCT_CHK, SLCT_LST 指定時
選択した項目の表示名を格納した配列
注釈
項目数が31を超えてもOK
キャンセル時
-1 を返す
UWSCとの違い
タイムアウト時の戻り値が0ではなく -2 になった
表示項目に連想配列を渡した場合、値でなはくキーが表示される
SLCT_CHK , SLCT_LST 指定時の戻り値がタブ文字連結された文字列ではなく配列になった
popupmenu ( メニュー項目 [ , x=EMPTY , y=EMPTY ] ) #
ポップアップメニューを表示します
パラメータ :
メニュー項目 ( 配列 ) -- 表示項目を示す配列、要素が配列の場合サブメニューになる
x ( 数値 省略可 ) -- メニュー表示位置のX座標を指定、省略時(EMPTY)はマウスカーソル位置
y ( 数値 省略可 ) -- メニュー表示位置のY座標を指定、省略時(EMPTY)はマウスカーソル位置
戻り値 :
選択した項目の表示名、メニューの外側を選んだ場合はEMPTY
サンプルコード
// サブメニュー表示方法
list = [ &quot;項目1&quot; , &quot;項目2&quot; , &quot;サブメニュー&quot; , [ &quot;サブ項目1&quot; , &quot;サブ項目2&quot; ] , &quot;項目3&quot; ]
// 要素を配列にすると直前の項目のサブメニューになる
selected = popupmenu ( list )
// 項目1
// 項目2
// サブメニュー &gt; サブ項目1
// サブ項目2
// 項目3
// ネストも可能
list = [ &quot;menu&quot; , [ &quot;branch1&quot; , &quot;branch2&quot; , [ &quot;leaf1&quot; , &quot;leaf2&quot; ]]]
popupmenu ( list )
UWSCとの違い
メニュー項目に連想配列を渡した場合、値ではなくキーが表示されます
メニュー項目を選んだ場合の戻り値が項目のインデックス値ではなく選択項目の表示名になりました
メニュー項目外を選んだ場合の戻り値が-1ではなくEMPTYになりました
## メッセージ表示#
balloon ( メッセージ [ , X=0 , Y=0 , 変形=FUKI_DEFAULT , フォントサイズ=EMPTY , 文字色=$000000 , 背景色=$00FFFF , 透過=0 ] ) #
fukidasi ( メッセージ [ , X=0 , Y=0 , 変形=FUKI_DEFAULT , フォントサイズ=EMPTY , 文字色=$000000 , 背景色=$00FFFF , 透過=0 ] ) #
吹き出しを表示します
パラメータ :
メッセージ ( 文字列 ) -- 表示するメッセージ
X ( 数値 省略可 ) -- 表示位置 (X座標)
Y ( 数値 省略可 ) -- 表示位置 (Y座標)
変形 ( 定数 省略可 ) -- 変形方法を示す定数を指定
FUKI_DEFAULT
変形しない
FUKI_UP
吹き出しに上向きの嘴を付ける
FUKI_DOWN
吹き出しに下向きの嘴を付ける
FUKI_LEFT
吹き出しに左向きの嘴を付ける
FUKI_RIGHT
吹き出しに右向きの嘴を付ける
FUKI_ROUND
吹き出しの角を丸くする
FUKI_POINT
嘴定数に加えることで、表示位置の基準を吹き出しの左上ではなく嘴の先にする
balloon ( &quot;マウスカーソル位置が吹き出しの左上&quot; , G_MOUSE_X , G_MOUSE_Y , FUKI_DOWN )
sleep ( 2 )
balloon ( &quot;マウスカーソル位置に嘴を向ける&quot; , G_MOUSE_X , G_MOUSE_Y , FUKI_DOWN or FUKI_POINT )
sleep ( 2 )
フォントサイズ ( 数値 省略可 ) -- 表示される文字のサイズ、EMPTY時はフォント設定に従う
フォント名 ( 数値 省略可 ) -- 表示される文字のフォント、EMPTY時はフォント設定に従う
文字色 ( 数値 省略可 ) -- 文字の色をBGR値で指定、省略時は黒
背景色 ( 数値 省略可 ) -- 背景の色をBGR値で指定、省略時は黄色
ヒント
BGRの例
青: $FF0000
緑: $00FF00
赤: $0000FF
白: $FFFFFF
黒: $000000
黄: $00FFFF
UWSCとの違い
色指定を0にした場合、黄色ではなく黒になります
透過 ( 数値 省略可 ) -- ウィンドウを透過させます
0: 透過させない
1～255: 数値が大きいほど透明度が高い
-1: 背景を透明にするが枠線は残る
-2: 背景と枠線を透明にする
戻り値 :
なし
吹き出しの表示は1スレッドにつき1つまで
吹き出し表示中にballoon()を呼ぶと、以前の吹き出しは削除され新たな吹き出しが表示されます
吹き出し表示中に別のスレッドでballoonを呼んだ場合はそれぞれ表示されます
logprint ( 表示フラグ [ , X=EMPTY , Y=EMPTY , 幅=EMPTY , 高さ=EMPTY ] ) #
printウィンドウの表示状態を変更します
print窓が無効の場合
print窓が無効の場合この関数は無視されます
以下のいずれかの場合のみこの関数は有効です
OPTION GUIPRINT がTRUEに指定されている
設定ファイルの options.gui_print がTRUEになっている
ウィンドウ強制モード( uwscr --window )で起動している
UWSCRがguiビルドの場合
パラメータ :
表示フラグ ( 真偽値 ) --
TRUE
print窓を表示する
FALSE
print窓を非表示にする
既に表示済みなら消す
X ( 数値 省略可 ) -- 表示位置 (X座標)、EMPTYなら現状維持
Y ( 数値 省略可 ) -- 表示位置 (Y座標)、EMPTYなら現状維持
幅 ( 数値 省略可 ) -- 表示サイズ (幅)、EMPTYなら現状維持
高さ ( 数値 省略可 ) -- 表示サイズ (高さ)、EMPTYなら現状維持
戻り値 :
なし
## HTMLフォーム#
createform ( HTMLファイル , タイトル [ , 非同期フラグ=FALSE , オプション=FOM_DEFAULT , 幅=EMPTY , 高さ=EMPTY , X=EMPTY , Y=EMPTY ] ) #
フォームウィンドウを表示します
WebView2 Runtimeが必要です
Microsoft Edge WebView2 Runtime がインストールされていない場合この関数はエラーになります
UWSCとは互換性がありません
UWSCではIEコンポーネントを利用していたのに対してUWSCRではWebView2を利用しています
そのためUWSCで実行していたコードが動作しない場合があります
パラメータ :
HTMLファイル ( 文字列 ) -- 表示したいHTMLファイルのパス
ファイルの配置について
HTMLファイルから別のファイルを参照する場合、もとのHTMLファイルを起点とした相対パスを指定します
C:\Test\
form.html
js\
form.js
css\
form.css
img\
form.png
&lt;!DOCTYPE html&gt;
&lt; html lang = &quot;ja&quot; &gt;
&lt; head &gt;
&lt; meta charset = &quot;UTF-8&quot; &gt;
&lt; title &gt; 別ファイル参照例 &lt;/ title &gt;
&lt; link rel = &quot;stylesheet&quot; href = &quot;css/form.css&quot; &gt;
&lt; script src = &quot;js/form.js&quot; &gt;&lt;/ script &gt;
&lt;/ head &gt;
&lt; body &gt;
&lt; img src = &quot;img/form.png&quot; &gt;
&lt; form &gt;
&lt; input type = &quot;submit&quot; value = &quot;OK&quot; name = &quot;OK&quot; &gt;
&lt;/ form &gt;
&lt;/ body &gt;
&lt;/ html &gt;
html = &quot;c:\test\form.html&quot;
r = createform ( html , &quot;test&quot; )
タイトル ( 文字列 ) -- ウィンドウタイトル
非同期フラグ ( 真偽値 省略可 ) -- 非同期で実行するかどうか
FALSE: submitボタンが押される、またはウィンドウが閉じられるまで待機する
TRUE: 関数実行後にウィンドウが表示されたら制御を返す
オプション ( 定数 省略可 ) -- 以下の定数の組み合わせ(OR連結)を指定
FOM_NOICON
閉じるボタンを非表示にする
FOM_MINIMIZE
最小化ボタンを表示する
FOM_MAXIMIZE
最大化ボタンを表示する
FOM_NOHIDE
submitボタンが押されてもウィンドウを閉じない
FOM_NOSUBMIT
submitボタンが押されてもsubmitに割り当てられた処理(action)を行わない
FOM_NORESIZE
ウィンドウのサイズ変更不可
FOM_BROWSER
互換性のために残されていますが使用できません (指定しても無視されます)
FOM_FORMHIDE
ウィンドウを非表示で起動する
FOM_TOPMOST
ウィンドウを最前面に固定
FOM_NOTASKBAR
タスクバーにアイコンを表示しない
FOM_FORM2
互換性のために残されていますが使用できません (指定しても無視されます)
FOM_DEFAULT
オプションなし (0)
幅 ( 数値 省略可 ) -- ウィンドウの幅
高さ ( 数値 省略可 ) -- ウィンドウの高さ
X ( 数値 省略可 ) -- ウィンドウのX座標
Y ( 数値 省略可 ) -- ウィンドウのY座標
戻り値の型 :
Form情報 または Formオブジェクト
戻り値 :
非同期フラグによる
FALSE: Form情報
TRUE: Formオブジェクト
&lt;!DOCTYPE html&gt;
&lt; html lang = &quot;ja&quot; &gt;
&lt; head &gt;
&lt; meta charset = &quot;UTF-8&quot; &gt;
&lt; title &gt; Sample.html &lt;/ title &gt;
&lt;/ head &gt;
&lt; body &gt;
&lt; form &gt;
&lt; div &gt;
&lt; span &gt; ユーザー名 &lt;/ span &gt;
&lt; input type = &quot;text&quot; name = &quot;user&quot; &gt;
&lt;/ div &gt;
&lt; div &gt;
&lt; span &gt; パスワード &lt;/ span &gt;
&lt; input type = &quot;password&quot; name = &quot;pwd&quot; &gt;
&lt;/ div &gt;
&lt; div &gt;
&lt; select name = &quot;slct&quot; &gt;
&lt; option value = &quot;foo&quot; &gt; foo &lt;/ option &gt;
&lt; option value = &quot;bar&quot; &gt; bar &lt;/ option &gt;
&lt; option value = &quot;baz&quot; &gt; baz &lt;/ option &gt;
&lt;/ select &gt;
&lt;/ div &gt;
&lt; div &gt;
&lt; textarea name = &quot;txt&quot; cols = &quot;30&quot; rows = &quot;10&quot; &gt;&lt;/ textarea &gt;
&lt;/ div &gt;
&lt; div &gt;
&lt; input type = &quot;submit&quot; value = &quot;OK&quot; name = &quot;OK&quot; &gt;
&lt; input type = &quot;submit&quot; value = &quot;Cancel&quot; name = &quot;Cancel&quot; &gt;
&lt;/ div &gt;
&lt;/ form &gt;
&lt;/ body &gt;
&lt;/ html &gt;
r = createform ( &quot;sample.html&quot; , &quot;Sample&quot; )
select r . submit
case &quot;OK&quot;
print &quot;OKが押されました&quot;
print &quot;formの値は以下です&quot;
for data in r . data
print data . name + &quot;: &quot; + data . value
next
case &quot;Cancel&quot;
print &#39;キャンセルされました&#39;
case NULL
print &#39;submitされずにウィンドウが閉じられました&#39;
selend
## Form情報#
submit時のform情報を示す UObject
// submit時
{
&quot;submit&quot; : $submit , // $submitには押されたsubmitボタンのnameが入る
&quot;data&quot; : [
// form内の各要素のnameおよびvalueが格納される
{ &quot;name&quot; : $name , &quot;value&quot; , $value },
]
}
// ウィンドウが閉じられた場合
{
&quot;submit&quot; : null , // NULLになる
&quot;data&quot; : [] // 空配列
}
## Formオブジェクト#
Formウィンドウを示すオブジェクト
COMオブジェクトではありません
UWSCとは異なりCOMオブジェクトではなくUWSCR独自のオブジェクトとなります
別スレッドからは呼び出せません
Formオブジェクトをpublic変数に代入して別のスレッドから呼ぶことはできません
class Form #
property Document #
フォームに表示されているページのdocumentオブジェクト
戻り値の型 :
WebViewRemoteObject
Wait ( ) #
ウィンドウが閉じられるのを待つ
戻り値の型 :
Form情報
戻り値 :
submit時のform情報を示す Form情報 オブジェクト
submitせず閉じた場合は submit がNULLになります
// test.htmlにはOKとCancelのsubmitボタンがあるものとする
f = createform ( &quot;test.html&quot; , &quot;Test&quot; , true )
result = f . wait ()
select result . submit
case &quot;OK&quot;
for data in result . data
print data . name + &quot;: &quot; + data . value
next
case &quot;Cancel&quot;
print &quot;キャンセルされました&quot;
case NULL
print &quot;ウィンドウが閉じられました&quot;
default
print &quot;なにかおかしいです&quot;
selend
SetVisible ( [ 表示フラグ=TRUE ] ) #
ウィンドウの表示状態を変更する
パラメータ :
表示フラグ ( 真偽値 省略可 ) -- TRUEで表示、FALSEで非表示
戻り値 :
なし
Close ( ) #
ウィンドウを閉じる
戻り値 :
なし
SetEventHandler ( エレメント , イベント , 関数 ) #
任意のイベント発生時に実行する関数を登録します
関数は引数を2つまで受けられます、内訳は以下の通りです
イベント発生エレメントのvalue値
イベント発生エレメントのname属性値
パラメータ :
エレメント ( WebViewRemoteObject ) -- イベント発生元のエレメントを示す WebViewRemoteObject
イベント ( 文字列 ) -- イベント名
関数 ( ユーザー定義関数 ) -- イベント発生時に実行される関数
戻り値 :
なし
f = createform ( &quot;test.html&quot; , &quot;Test&quot; , true )
select = f . document . querySelector ( &quot;select&quot; )
f . SetEventHandler ( select , &quot;change&quot; , on_select_change )
button = f . document . querySelector ( &quot;input[type=button]&quot; )
f . SetEventHandler ( button , &quot;click&quot; , on_button_click )
f . wait ()
// 1つ目の引数でイベント発生エレメントのvalue
// 2つ目の引数でnameを受ける
procedure on_select_change ( value , name )
print value
print name
fend
// 引数は必須ではない
procedure on_button_click ()
print &quot;クリックされました&quot;
fend
## WebViewRemoteObject#
フォームに表示されているページのJavaScriptオブジェクトを示します
利用方法は RemoteObject と同等です
## 組み込みウィンドウのクラス名一覧#
関数名
クラス名
定数
msgbox
UWSCR.MsgBox
CLASS_MSGBOX
input
UWSCR.Input
CLASS_INPUTBOX
slctbox
UWSCR.Slctbox
CLASS_SLCTBOX
popupmenu
UWSCR.Popup
CLASS_POPUPMENU
balloon
UWSCR.Balloon
CLASS_BALLOON
logprintwin
UWSCR.LogPrintWin
CLASS_LOGPRINTWIN
createform
UWSCR.Form
CLASS_FORM

---

# 配列操作関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/array.html

## 配列操作関数#
## 配列の変更#
qsort ( var キー配列 [ , ソート順=QSRT_A , var 連動配列 , ... ] ) #
配列内の要素を並び替えます
ソート時の値型について
それぞれの値を文字列として扱いソートを行います
パラメータ :
キー配列 ( 配列 参照渡し ) -- ソートする配列
ソート順 ( 定数 省略可 ) -- ソート順を示す定数
QSRT_A
昇順
QSRT_D
降順
QSRT_UNICODEA
UNICODE文字列順 昇順
QSRT_UNICODED
UNICODE文字列順 降順
QSRT_NATURALA
数値順 昇順
QSRT_NATURALD
数値順 降順
連動配列 ( 配列 省略可 参照渡し ) -- キー配列のソートに連動してソートされる配列
キー配列よりサイズの小さい配列はソート前にリサイズされEMPTYで埋められます
8つまで指定可能
戻り値 :
なし
サンプルコード
// 連動ソート
// キー配列を並び替え、それと同じように別の配列も並び替えます
key = [ 5 , 2 , 1 , 4 , 3 ]
arr1 = [ &quot;お&quot; , &quot;い&quot; , &quot;あ&quot; , &quot;え&quot; , &quot;う&quot; ]
arr2 = [ &quot;お&quot; , &quot;い&quot; , &quot;あ&quot; , &quot;え&quot; , &quot;う&quot; , &quot;か&quot; ] // 余分はソート対象外、この場合「か」は位置が変更されない
arr3 = [ &quot;お&quot; , &quot;い&quot; , &quot;あ&quot; , &quot;え&quot; ] // 不足の場合末尾にEMPTYが追加されてからソート
qsort ( key , QSRT_A , arr1 , arr2 , arr3 )
print key // [1, 2, 3, 4, 5]
print arr1 // [あ, い, う, え, お]
print arr2 // [あ, い, う, え, お, か]
print arr3 // [あ, い, , え, お]
qsort ( key , QSRT_D , arr1 , arr2 , arr3 )
print key // [5, 4, 3, 2, 1]
print arr1 // [お, え, う, い, あ]
print arr2 // [お, え, う, い, あ, か]
print arr3 // [お, え, , い, あ]
reverse ( var 配列 ) #
配列の順序を反転させます
パラメータ :
配列 ( 配列 参照渡し ) -- 順序を反転させたい配列
戻り値 :
なし
サンプルコード
arr = [ 1 , 2 , 3 ]
print arr // [1,2,3]
reverse ( arr )
print arr // [3,2,1]
resize ( var 配列 [ , インデックス値=EMPTY , 初期値=EMPTY ] ) #
配列サイズを変更します
パラメータ :
配列 ( 配列 参照渡し ) -- サイズを変更したい配列
インデックス値 ( 数値 省略可 ) --
指定値 + 1 のサイズに変更される
省略時は変更なし
マイナス指定時はサイズ0の配列になる
初期値 ( 値 省略可 ) -- 元のサイズより大きくなる場合、追加される要素の初期値
戻り値 :
配列サイズ - 1 (配列インデックスの最大値)
サンプルコード
arr = [ 1 , 2 , 3 ]
// サイズ指定なしの場合は配列に変更なし
print resize ( arr ) // 2
print length ( arr ) // 3
// サイズ指定
print resize ( arr , 3 ) // 3
print length ( arr ) // 4
// マイナス指定でサイズ0になる
print resize ( arr , - 1 ) // -1
print length ( arr ) // 0
// サイズ変更+初期値指定
arr = []
print resize ( arr , 2 , &quot;a&quot; ) // 2
print length ( arr ) // 3
print arr // [a, a, a]
setclear ( var 配列 [ , 値=EMPTY ] ) #
指定した値で配列を埋めます
パラメータ :
配列 ( 配列 参照渡し ) -- 値を埋めたい配列
値 ( 値 省略可 ) -- 埋める値
戻り値 :
なし
サンプルコード
arr = [ 1 , 2 , 3 , 4 , 5 ]
print arr // [1, 2, 3, 4, 5]
// 値省略時はEMPTYで埋められる
setclear ( arr )
print arr // [, , , , ]
setclear ( arr , 111 )
print arr // [111, 111, 111, 111, 111]
shiftarray ( var 配列 , シフト値 ) #
指定値分配列内の要素をずらします
パラメータ :
配列 ( 配列 参照渡し ) -- 対象の配列
シフト値 ( 数値 ) -- 正の数なら要素を後方にずらす、負の数なら前方へずらす (空いた場所はEMPTYで埋められる)
戻り値 :
なし
サンプルコード
arr = [ 1 , 2 , 3 , 4 , 5 ]
print arr // [1, 2, 3, 4, 5]
shiftarray ( arr , 2 )
print arr // [, , 1, 2, 3]
shiftarray ( arr , - 2 )
print arr // [1, 2, 3, , ]
## 配列長を得る#
Length ( ) #
文字列操作関数の length 関数を参照
## 配列要素を使う#
slice ( 配列 [ , 開始=0 , 終了=EMPTY ] ) #
配列の一部をコピーし新たな配列を得ます
パラメータ :
配列 ( 配列 ) -- コピー元の配列
開始 ( 数値 省略可 ) -- コピーする開始位置のインデックス値
終了 ( 数値 省略可 ) -- コピーする終了位置のインデックス値、省略時は最後まで
戻り値 :
コピーされた配列
サンプルコード
// 開始と終了が未指定の場合は配列がそのまま複製される
base = [ 1 , 2 , 3 , 4 , 5 ]
new = slice ( base )
print new // [1, 2, 3, 4, 5]
print slice ( base , 2 ) // [3, 4, 5]
print slice ( base , , 2 ) // [1, 2, 3]
print slice ( base , 1 , 3 ) // [2, 3, 4]
// 範囲外が指定されたら空配列が返る
print slice ( base , 5 ) // []
calcarray ( 配列 , 計算方法 [ , 開始=0 , 終了=EMPTY ] ) #
配列内の数値で計算を行います
パラメータ :
配列 ( 配列 ) -- 数値を含む配列 (数値以外は無視される)
計算方法 ( 定数 ) -- 計算方法を示す定数
CALC_ADD
合計値を得る
CALC_MIN
最小値を得る
CALC_MAX
最大値を得る
CALC_AVR
平均値を得る
戻り値 :
計算結果
サンプルコード
arr = [ 1 , 2 , 3 , 4 , 5 ]
print calcarray ( arr , CALC_ADD ) // 15
print calcarray ( arr , CALC_MIN ) // 1
print calcarray ( arr , CALC_MAX ) // 5
print calcarray ( arr , CALC_AVR ) // 3
// 範囲指定
print calcarray ( arr , CALC_ADD , 2 , 3 ) // 7
print calcarray ( arr , CALC_MIN , 2 , 3 ) // 3
print calcarray ( arr , CALC_MAX , 2 , 3 ) // 4
print calcarray ( arr , CALC_AVR , 2 , 3 ) // 3.5
// 数値以外は無視される
arr = [ 1 , 2 , &quot;foo&quot; , 4 , 5 ]
print calcarray ( arr , CALC_ADD ) // 12
print calcarray ( arr , CALC_MIN ) // 1
print calcarray ( arr , CALC_MAX ) // 5
print calcarray ( arr , CALC_AVR ) // 3 ※ 数値要素が4つなので (1+2+4+5) / 4
## 文字列との相互変換#
join ( 配列 [ , 区切り文字=&quot; &quot; , 空文字除外=FALSE , 開始=0 , 終了=(配列長-1) ] ) #
配列要素を区切り文字で結合します
パラメータ :
配列 ( 配列 ) -- 結合したい配列
区切り文字 ( 文字列 省略可 ) -- 結合時の区切り文字
空文字除外 ( 真偽値 省略可 ) -- FALSEなら配列要素が空文字でも結合する、TRUEなら除外
開始 ( 数値 省略可 ) -- 結合範囲の開始位置のインデックス値
終了 ( 数値 省略可 ) -- 結合範囲の終了位置のインデックス値
戻り値 :
結合後の文字列
サンプルコード
arr = [ &quot;foo&quot; , &quot;bar&quot; , &quot;baz&quot; , &quot;qux&quot; ]
print join ( arr ) // foo bar baz qux
print join ( arr , &quot;+&quot; ) // foo+bar+baz+qux
print join ( arr , &quot;+&quot; , FALSE , 1 , 2 ) // bar+baz
// 空文字除外
print join ([ &quot;hoge&quot; , &quot;&quot; , &quot;fuga&quot; ] , &quot;&amp;&quot; , FALSE ) // hoge&amp;&amp;fuga
print join ([ &quot;hoge&quot; , &quot;&quot; , &quot;fuga&quot; ] , &quot;&amp;&quot; , TRUE ) // hoge&amp;fuga
split ( 文字列 [ , 区切り文字=&quot; &quot; , 空文字除外=FALSE , 数値変換=FALSE , CSV分割=FALSE ] ) #
文字列を区切り文字で分割して配列にします
パラメータ :
文字列 ( 文字列 ) -- 分割したい文字列
区切り文字 ( 文字列 省略可 ) -- 分割するための区切り、CSV分割が有効の場合最初の一文字のみ使用される
一文字ずつ分割
区切り文字として空文字を指定すると文字列を一文字ずつ分割できます
空文字除外 ( 真偽値 省略可 ) -- FALSEなら分割後に空文字があっても配列要素とする、TRUEなら除外
数値変換 ( 真偽値 省略可 ) -- TRUEなら分割後の文字列を数値へ変換し、変換できない場合は空文字とする
CSV分割 ( 真偽値 省略可 ) -- TRUEならCSVとして分割する (空文字除外と数値変換は無視される)
戻り値 :
分割された配列
サンプルコード
print split ( &quot;a b c&quot; ) // [a, b, c]
// 空文字除外
print split ( &quot;a,,b,,c&quot; , &quot;,&quot; , FALSE ) // [a, , b, , c]
print split ( &quot;a,,b,,c&quot; , &quot;,&quot; , TRUE ) // [a, b, c]
// 数値変換
print split ( &quot;1,2,f,4,5&quot; , &quot;,&quot; , FALSE , FALSE ) // [1, 2, f, 4, 5]
print split ( &quot;1,2,f,4,5&quot; , &quot;,&quot; , FALSE , TRUE ) // [1, 2, , 4, 5]
// 空文字除外と組み合わせると数値以外を排除できる
print split ( &quot;1,2,f,4,5&quot; , &quot;,&quot; , TRUE , TRUE ) // [1, 2, 4, 5]
// 空文字で分割
print split ( &quot;12345&quot; , &quot;&quot; , FALSE ) // [, 1, 2, 3, 4, 5, ]
print split ( &quot;12345&quot; , &quot;&quot; , TRUE ) // [1, 2, 3, 4, 5]
// CSV分割
// , で区切られる
print split ( &#39;a,b,&quot;c,d&quot;,e&#39; , &quot;,&quot; , , , FALSE ) // [a, b, &quot;c, d&quot;, e]
// &quot;&quot; 内を文字列扱いとし中の , では区切らない
print split ( &#39;a,b,&quot;c,d&quot;,e&#39; , &quot;,&quot; , , , TRUE ) // [a, b, c,d, e]

---

# 数学関数 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/math.html

## 数学関数#
isnan ( 数値 ) #
値が NaN であるかどうかを調べる
パラメータ :
数値 ( 数値 ) -- 調べる値
戻り値 :
NaN であればTRUE
サンプルコード
print IsNan ( NaN ) // True
n = NaN
print IsNan ( n ) // True
print IsNan ( 1 ) // False
print IsNan ( &quot;あ&quot; ) // False
random ( n ) #
0以上n未満の整数をランダムに返す
注釈
指定可能な最大値は2147483646です
パラメータ :
n ( 数値 ) -- 得たい数値の範囲を示す値
戻り値の型 :
数値
戻り値 :
得られたランダム値
abs ( n ) #
絶対値を得る
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
絶対値
zcut ( n ) #
マイナス値は0にする
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
0以上の整数
int ( n ) #
小数点以下切り落とし
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
整数
ceil ( n ) #
小数点以下を正方向に切り上げ
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
整数
round ( n [ , 桁=0 ] ) #
指定桁数で入力値を丸める
パラメータ :
n ( 数値 ) -- 入力値
桁 ( 数値 省略可 ) -- 丸める桁、マイナスなら小数点以下の桁数
戻り値の型 :
数値
戻り値 :
整数
sqrt ( n ) #
平方根
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
入力値の平方根、入力値がマイナスの場合NaN
power ( n , e ) #
nをe乗する
パラメータ :
n ( 数値 ) -- 入力値
e ( 数値 ) -- 指数
戻り値の型 :
数値
戻り値 :
数値
exp ( n ) #
指数関数
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
数値
ln ( n ) #
自然対数
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
数値
logn ( base , n ) #
baseを底とするnの対数
パラメータ :
base ( 数値 ) -- 底
n ( 数値 ) -- 値
戻り値の型 :
対数
戻り値 :
数値
sin ( n ) #
サイン
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン
cos ( n ) #
コサイン
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン
tan ( n ) #
タンジェント
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン
arcsin ( n ) #
アークサイン
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン
arccos ( n ) #
アークコサイン
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン
arctan ( n ) #
アークタンジェント
パラメータ :
n ( 数値 ) -- 入力値
戻り値の型 :
数値
戻り値 :
ラジアン

---

# COMオブジェクト - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/comobject.html

## COMオブジェクト#
## COMオブジェクトの作成・取得#
createoleobj ( ProgID ) #
COMオブジェクトのインスタンスを得ます
パラメータ :
ProgID ( 文字列 ) -- COMオブジェクトのProgIDまたはCLSID
戻り値 :
COMオブジェクト
getactiveoleobj ( ProgID [ , タイトル=EMPTY , n番目=1 ] ) #
既に起動中のCOMオブジェクトを得ます
タイトルが未指定の場合は指定ProgIDに該当しアクティブなオブジェクトを返します
タイトルを指定した場合はウィンドウタイトルに部分一致するウィンドウからProgIDに該当するオブジェクトを返します
パラメータ :
ProgID ( 文字列 ) -- COMオブジェクトのProgIDまたはCLSID
タイトル ( 文字列 省略可 ) -- ExcelやWordなど、オブジェクトを取得したいウィンドウのタイトルを指定 (部分一致)
MDI非対応
MDIウィンドウは対象外です
n番目 ( 数値 省略可 ) -- タイトルに一致するウィンドウが複数ある場合、n番目を取得
戻り値 :
COMオブジェクト
CLSIDの入力
CLSIDは {XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX} の形式で入力します
// WScript.ShellのCLSIDを指定
ws = createoleobj ( &quot;{72C24DD5-D70A-438B-8A42-98424B88AFB8}&quot; )
print ws // ComObject(IWshShell3)
ws . Popup ( &quot;Hello!&quot; )
## コレクション#
getoleitem ( コレクション ) #
コレクションを配列に変換します
パラメータ :
コレクション ( COMオブジェクト ) -- コレクションを示すCOMオブジェクト
戻り値の型 :
配列
戻り値 :
コレクションの要素を格納した配列
UWSCとの違い
要素の数ではなく要素の配列を返すようになりました
それに伴い ALL_OLE_ITEM は廃止されました
サンプルコード
ws = createoleobj ( &quot;WScript.Shell&quot; )
col = getoleitem ( ws . SpecialFolders )
print col [ 0 ]
## VARIANT#
vartype ( 値 ) #
VARIANTがどのような型であるかを調べます
パラメータ :
値 ( VARIANT ) -- VARIANT型の値
戻り値の型 :
定数
戻り値 :
VARIANTのデータ型を示す VAR定数
VARIANT以外を指定した場合
UWSCRの大半の値はVARIANT型ではありません
VARIANTではない値の場合は VAR_UWSCR が返ります
UWSCRにおける値の型を調べるには type_of() 関数をご利用ください
vartype ( 値 , VAR定数 )
任意の値を指定した型のVARIANTに変換します
パラメータ :
値 ( すべて ) -- 任意の値
VAR定数 ( 定数 ) -- 変換する型を 定数 で指定
以下の定数は使用できません
VAR_ASTR
VAR_USTR
VAR_UWSCR
戻り値の型 :
VARIANT
戻り値 :
指定した型のVARIANT
// 開いているExcelを取得
excel = getactiveoleobj ( &quot;Excel.Application&quot; )
// 日付型のVARIANTに変換
date = vartype ( &quot;2023/07/15&quot; , VAR_DATE )
// Excelのアクティブセルに日付型の値を入力
excel . ActiveCell . value = date
vartype ( COMオブジェクト , プロパティ名 )
COMオブジェクトのプロパティが返す値のVARIANT型を得ます
パラメータ :
COMオブジェクト ( COMオブジェクト ) -- 型を調べたいプロパティを持つCOMオブジェクト
プロパティ名 ( 文字列 ) -- 型を調べたいプロパティの名前
戻り値の型 :
VAR定数またはEMPTY
戻り値 :
プロパティの型、COMオブジェクト以外が渡された場合やプロパティが存在しない場合はEMPTY
excel = getactiveoleobj ( &quot;Excel.Application&quot; )
// アクティブセルの型を調べる
vt = vartype ( excel . activecell , &quot;value&quot; )
// 得た値をVAR_定数名に変換
print const_as_string ( vt , &quot;VAR_&quot; )
## VAR定数#
VAR定数一覧 #
定数
値
詳細
VAR_EMPTY
0
EMPTY
VAR_NULL
1
NULL
VAR_SMALLINT
2
符号付き2バイト整数
VAR_INTEGER
3
符号付き4バイト整数
VAR_SINGLE
4
単精度浮動小数点数
VAR_DOUBLE
5
倍精度浮動小数点数
VAR_CURRENCY
6
通貨型
VAR_DATE
7
日付型
VAR_BSTR
8
文字列型
VAR_DISPATCH
9
IDispatch型 (COMオブジェクト)
VAR_ERROR
10
エラー
VAR_BOOLEAN
11
真偽値
VAR_VARIANT
12
VARIANT型
VAR_UNKNOWN
13
IUnknown型
VAR_SBYTE
16
符号付き1バイト整数
VAR_BYTE
17
符号なし1バイト整数
VAR_WORD
18
符号なし2バイト整数
VAR_DWORD
19
符号なし4バイト整数
VAR_INT64
20
符号付き8バイト整数
VAR_ASTR
256
互換性のために残していますが実際には使用できません
VAR_USTR
258
互換性のために残していますが実際には使用できません
VAR_UWSCR
512
UWSCRのデータ型
VAR_ARRAY
$2000 (8192)
配列
## 非推奨関数#
非推奨の理由
UWSCRにはSAFEARRAY型の値が存在しないため以下の関数は非推奨となりました
互換性のため関数は残していますが、UWSCとは結果が異なります
safearray ( [ 下限=0 , 上限=-1 , 二次元下限=EMPTY , 二次元上限=(二次元下限-1) ] ) #
EMPTYを返します
パラメータ :
下限 ( 数値 省略可 ) -- 無視されます
上限 ( 数値 省略可 ) -- 無視されます
二次元下限 ( 数値 省略可 ) -- 無視されます
二次元上限 ( 数値 省略可 ) -- 無視されます
戻り値 :
EMPTY

---

# Excel - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/excel.html

## Excel#
xlopen ( [ ファイル名=EMPTY , 起動フラグ=XL_DEFAULT , パラメータ... ] ) #
Excelを起動します
パラメータ :
ファイル名 ( 文字列 省略可 ) -- 開きたいファイル名、EMPTYならExcelを新規に起動
起動フラグ ( 定数 省略可 ) -- Excelの起動方法を指定
XL_DEFAULT (0)
起動済みのExcelがあればそれを使い、なければ新規起動します
XL_NEW (1)
常にExcelを新規に起動します
XL_BOOK (2)
applicationではなくWorkbookオブジェクトを返します
XL_OOOC (3)
使用できません
パラメータ ( 文字列 可変長 ) --
ファイルオープン時の追加パラメータを &quot;パラメータ名:=値&quot; 形式の文字列で指定する
書式が不正な場合は無視される (エラーにはなりません)
以下は有効なパラメータ例
UpdateLinks
リンク更新方法
0: 更新しない
1: 外部更新のみ
2: リモート更新のみ
3: 外部、リモート共に更新
ReadOnly
読み取り専用で開く場合にTrueを指定
Format
CSVファイル時を開く場合にその区切り文字
1: タブ
2: カンマ
3: スペース
4: セミコロン
Password
パスワード保護されたブックを開くためのパスワード
WriteResPassword
書き込み保護されたブックに書き込むためのパスワード
IgnoreReadOnly
「読み取り専用を推奨する」のダイアログを抑止したい場合にTrue
// パスワード付きファイルを読み取り専用で開く
excel = xlopen ( &quot;hoge.xlsx&quot; , XL_NEW , &quot;ReadOnly:=True&quot; , &quot;Password:=hogehoge&quot; )
// カンマ区切りcsvファイルを開く
excel = xlopen ( &quot;hoge.xlsx&quot; , XL_NEW , &quot;Format:=2&quot; )
戻り値の型 :
COMオブジェクト
戻り値 :
Excel自身、またはWorkbookを示すCOMオブジェクト
xlclose ( Excel [ , ファイル名 ] ) #
Excelを終了します
ファイル名指定の有無で保存方法が異なります
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
ファイル名 ( 文字列 省略可 ) -- 保存するファイル名を指定、省略時は上書き保存
xlclose ( Excel , TRUE )
変更内容を保存せずに終了します
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
TRUE ( 真偽値 ) -- TRUE を指定 (固定値)
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗時FALSE
サンプルコード
excel = xlopen ( &quot;foo.xlsx&quot; )
// ブックが編集される
xlclose ( excel , &quot;bar.xlsx&quot; ) // 別名で保存
excel = xlopen ( &quot;bar.xlsx&quot; )
// ブックが編集される
xlclose ( excel ) // 上書き保存
excel = xlopen ( &quot;foo.xlsx&quot; )
// ブックが編集される
xlclose ( excel , TRUE ) // 保存せず終了
xlactivate ( Excel , シート識別子 [ , ブック識別子=EMPTY ] ) #
指定したシートをアクティブにします
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
シート識別子 ( 文字列または数値 ) -- アクティブにするシート名またはインデックス番号(1から)
ブック識別子 ( 文字列または数値 省略可 ) -- アクティブにするブック名またはインデックス番号(1から)
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗時FALSE
シート・ブックの識別子について
シート名は各シートの表示名を完全一致で指定する必要があります
シートのインデックス番号は左から数えた順番です
ブック名は拡張子を含めたファイル名を完全一致で指定する必要があります
新規作成したブックの場合は Book1 のようになります(拡張子がありません)
ブックのインデックス番号はブックを開いた順番です
ブック識別子を省略した場合はアクティブなブックが対象となります
Workbookオブジェクトを指定した場合ブック識別子は無視され、そのWorkbook内のシートをアクティブにします
xlsheet ( Excel , シート識別子 [ , 削除=FALSE ] ) #
アクティブなブックへのシートの追加、または削除を行う
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
シート識別子 ( 文字列または数値 ) -- アクティブにするシート名、削除時のみインデックス番号(1から)も可
削除 ( 真偽値 省略可 ) -- FALSEなら指定名のシートを追加、TRUEなら該当シートを削除
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗時FALSE
インデックス指定について
シート追加時はインデックス番号を文字列として扱います
xlsheet ( excel , 1 , FALSE ) // &quot;1&quot; という名前のシートが追加される
シート削除時はインデックスとシート名を厳密に区別します
そのためUWSCとは一部動作が異なります
xlsheet ( excel , 1 , FALSE ) // &quot;1&quot; という名前のシートを追加しておく
xlsheet ( excel , 1 , TRUE ) // 1を指定して削除を試みた場合
// UWSCの場合: &quot;1&quot; という名前のシートがあればそれを削除、なければ1番目のシートを削除
// UWSCRの場合: 必ず1番目のシートを削除、2番目以降にある&quot;1&quot;という名前のシートは対象とならない
xlgetdata ( Excel [ , 範囲=EMPTY , シート識別子=EMPTY ] ) #
xlgetdata ( Excel [ , 範囲=EMPTY , &lt;EMPTYPARAM&gt; , シート識別子=EMPTY ] )
範囲をA1形式の文字列で指定し、その値を返します
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
範囲 ( 文字列 省略可 ) -- 単一セル指定なら&quot;A1&quot;、範囲なら&quot;A1:C3&quot;のように指定
シート識別子 ( 文字列または数値 省略可 ) -- 得たい値のあるシート名またはインデックス番号(1から)を指定、省略時はアクティブシート
第三引数について
互換性のために第三引数を省略し、第四引数にシート名を指定することもできます
戻り値の型 :
値または配列、値の型はセルによる
戻り値 :
範囲の指定方法により異なります
単一セル指定: セルの値を返す
範囲指定: 範囲内の値を順に格納した配列を返す
範囲指定時の注意
UWSCではインデックスが1から始まるSafeArrayが返っていましたが
UWSCRでは通常の配列が返るためインデックスが0からになります
xlgetdata ( Excel , 行番号 , 列番号 [ , シート識別子=EMPTY ] )
セルの行と列の番号を指定しその値を得ます
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
行番号 ( 数値 ) -- 値を得たいセルの行番号 (1から)
列番号 ( 数値 ) -- 値を得たいセルの列番号 (1から)
シート識別子 ( 文字列または数値 省略可 ) -- 得たい値のあるシート名またはインデックス番号(1から)を指定、省略時はアクティブシート
戻り値の型 :
セルによる
戻り値 :
指定セルの値
xlsetdata ( Excel , 値 [ , 範囲=EMPTY , シート識別子=EMPTY , 文字色=EMPTY , 背景色=EMPTY ] ) #
xlsetdata ( Excel , 値 [ , 範囲=EMPTY , &lt;EMPTYPARAM&gt; , シート識別子=EMPTY , 文字色=EMPTY , 背景色=EMPTY ] )
A1形式で指定したセルまたはセル範囲に値を入力します
入力したい値が配列で、かつ単一セルを指定した場合は指定セルを起点として配列の値を入力します
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
値 ( 値 ) -- 入力したい値 (配列可)
入力値ごとの入力パターン
入力値により入力方法が代わります
単一の値: 指定範囲すべてに同一の値が入力されます
一次元配列: 指定行の列ごとに配列要素がそれぞれ入力されます、範囲が複数行の場合それぞれの行に入力されます
二次元配列: 配列を行列とみなし各要素を該当するセルに入力します
配列サイズが指定範囲を超える場合、超過分は入力されません
指定範囲が配列サイズを超える場合、不足箇所には #N/A が入力されます
範囲 ( 文字列 省略可 ) -- A1形式でセルまたはセル範囲を指定
シート識別子 ( 文字列または数値 省略可 ) -- 得たい値のあるシート名またはインデックス番号(1から)を指定、省略時はアクティブシート
第三引数について
互換性のために第三引数を省略し、第四引数にシート名を指定することもできます
文字色 ( 数値 省略可 ) -- 該当セルの文字色を変更する場合にBGRで指定
背景色 ( 数値 省略可 ) -- 該当セルの背景色を変更する場合にBGRで指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗時FALSE
xlsetdata ( Excel , 値 , 行 , 列 [ , シート識別子=EMPTY , 文字色=EMPTY , 背景色=EMPTY ] )
行列番号で指定したセルに値をセットする
入力したい値が配列の場合は指定セルを起点に配列の値を入力します
パラメータ :
Excel ( COMオブジェクト ) -- Excel.ApplicationまたはWorkbookを示すCOMオブジェクト
値 ( 値 ) -- 入力したい値 (配列可)
行 ( 数値 ) -- 入力したいセルの行番号 (1から)
列 ( 数値 ) -- 入力したいセルの列番号 (1から)
シート識別子 ( 文字列または数値 省略可 ) -- 得たい値のあるシート名またはインデックス番号(1から)を指定、省略時はアクティブシート
文字色 ( 数値 省略可 ) -- 該当セルの文字色を変更する場合にBGRで指定
背景色 ( 数値 省略可 ) -- 該当セルの背景色を変更する場合にBGRで指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE、失敗時FALSE
// A1セルに100が入力される
xlsetdata ( excel , 100 , &quot;A1&quot; )
// A2,B2,C2に200が入力される
xlsetdata ( excel , 200 , &quot;A2:C2&quot; )
// A3に301, B3に302, C3に303が入力される
xlsetdata ( excel , [ 301 , 302 , 303 ] , &quot;A3:C3&quot; )
// 単一セル指定で配列を渡した場合はそのセルを起点に配列の値が入力される
// A4に401, B4に402, C4に403になる
xlsetdata ( excel , [ 401 , 402 , 403 ] , &quot;A4&quot; )
// 配列サイズが範囲より大きい場合
// C5に503は入力されない
xlsetdata ( excel , [ 501 , 502 , 503 ] , &quot;A5:B5&quot; )
// 配列サイズが範囲より小さい場合
// 配列を超えた部分であるD6は#N/A
xlsetdata ( excel , [ 601 , 602 , 603 ] , &quot;A6:D6&quot; )
// 二次元配列は行列になる
// | A | B |
// 7 | 701| 702|
// 8 | 801| 802|
xlsetdata ( excel , [[ 701 , 702 ] , [ 801 , 802 ]] , &quot;A7:B8&quot; )
// 二次元配列で単一セルを指定した場合でもそのセルを起点に入力される
xlsetdata ( excel , [[ 901 , 902 , 903 ] , [ 1001 , 1002 , 1003 ]] , &quot;A9&quot; )
// 行列番号指定
xlsetdata ( excel , [[ 1101 , 1102 ] , [ 1201 , 1202 ] , [ 1301 , 1302 ]] , 11 , 1 )

---

# ウェブ関連 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/web.html

## ウェブ関連#
## ブラウザ操作#
破壊的変更が行われました
バージョン 0.11.0 以降のブラウザ操作機能はバージョン 0.10.2 以前とは互換性がありません
ブラウザパスの指定方法
通常はレジストリ等からブラウザの実行ファイルのパスを取得しそれを実行します (パスの自動取得)
自動取得を行わずに任意のパスで実行させるには設定ファイルにパスを記述します
{
&quot;browser&quot; : {
&quot;chrome&quot; : &quot;C:\\path\\to\\chrome.exe&quot; ,
&quot;msedge&quot; : &quot;C:\\path\\to\\msedge.exe&quot;
},
}
自動取得に戻す場合は null にします
{
&quot;browser&quot; : {
&quot;chrome&quot; : null ,
&quot;msedge&quot; : null
},
}
パスは必ずchrome.exeおよびmsedge.exeのものにしてください
それ以外は動作保証外です
BrowserControl ( ブラウザ定数 [ , ポート=9222 ] ) #
Devtools Protocolを利用したブラウザ操作を行うための Browserオブジェクト を返します
デバッグポートを開いたブラウザを起動します
対応ブラウザは以下
Google Chrome
Microsoft Edge
パラメータ :
ブラウザ定数 ( 定数 ) -- 以下のいずれかを指定
BC_CHROME
Google Chromeを操作します
BC_MSEDGE
Microsoft Edgeを操作します
ポート ( 数値 省略可 ) -- デバッグポートを指定する
戻り値の型 :
Browserオブジェクト
戻り値 :
対象ブラウザの Browserオブジェクト
ブラウザへの再接続について
対象ブラウザが同じデバッグポートを開けて起動している場合はそのブラウザに再接続できます
異なるポートを開いている、またはポートが開かれていない場合は再接続できずエラーになります
// 起動.uws
chrome = BrowserControl ( BC_CHROME , 9999 ) // ポート9999でChromeを起動
chrome [ 0 ] . navigate ( &quot;https://example.com&quot; ) // 0番目のタブで任意のサイトを開く
// 再接続.uws
chrome = BrowserControl ( BC_CHROME , 9999 ) // 9999ポートのChromeに再接続される
url = chrome [ 0 ] . document . URL // 0番目のタブのURLを取得
print url // https://example.com
起動中のブラウザとは別に自動操作用のブラウザを起ち上げるには
起動中のブラウザとは異なるプロファイルで新たなブラウザを起動する必要があります
このような場合はBrowserControl関数ではなく Browserbuilder 関数 を使用してください
Browserbuilder 関数が返す BrowserBuilderオブジェクト でプロファイルフォルダを指定します
Browserbuilder ( ブラウザ定数 ) #
BrowserBuilderオブジェクト を返します
最低限の設定でブラウザを起動する BrowserControl 関数とは異なり BrowserBuilderオブジェクト を介して様々な設定が行なえます
パラメータ :
ブラウザ定数 ( 定数 ) -- 以下のいずれかを指定
BC_CHROME
Google Chromeを操作します
BC_MSEDGE
Microsoft Edgeを操作します
戻り値の型 :
BrowserBuilderオブジェクト
戻り値 :
対象ブラウザの BrowserBuilderオブジェクト
ブラウザの起動方法
BrowserBuilderオブジェクト の start() メソッドでブラウザを起動、または再接続します
// BrowserBuilderオブジェクトを作成し、startメソッドを呼ぶ
builder = BrowserBuilder ( BC_CHROME )
chrome = builder . start ()
// 以下のようにも書ける
chrome = BrowserBuilder ( BC_CHROME ) . start ()
// ポートの変更
chrome = BrowserBuilder ( BC_CHROME ) _
. port ( 9999 ) _
. start ()
// ヘッドレス起動
chrome = BrowserBuilder ( BC_CHROME ) _
. headless ( TRUE ) _
. start ()
// プロファイルフォルダの変更
chrome = BrowserBuilder ( BC_CHROME ) _
. profile ( &quot;C:\uwscr\chrome\profile1&quot; ) _
. start ()
// 複合設定
chrome = BrowserBuilder ( BC_CHROME ) _
. port ( 12345 ) _
. headless ( TRUE ) _
. start ()
対象ブラウザが指定ポートを開いていなかった場合の動作
対象ブラウザのプロセスがすでに存在している
そのプロセスが指定ポートを開いていない
の2点を満たす場合、再接続が行えないためエラーになります
ただし、起動中のブラウザとは異なるプロファイルフォルダを指定した場合は指定ポートで新たなブラウザプロセスを起動します
(同一プロファイルにつき一つのデバッグポート(またはポートなし)でしかブラウザを起動できないため)
RemoteObjectType ( remote ) #
RemoteObject の型を返します
型名の他に可能であれば以下を含みます
型の詳細
クラス名
パラメータ :
remote ( RemoteObject ) -- 型情報を得たい RemoteObject
戻り値の型 :
文字列
戻り値 :
型の情報を示す文字列
## IE関数互換#
## IEGETDATA互換#
BRGetData ( タブ , name [ , value=EMPTY , n番目=1 ] ) #
エレメントのnameとvalue属性をもとに値を取得する
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
name ( 文字列 ) -- 値を取得するエレメントのname属性
value ( 文字列 省略可 ) -- nameが同一の場合にvalue属性の値を指定
n番目 ( 数値 省略可 ) -- nameもvalueも一致する場合順番を1から指定
戻り値 :
取得された値、取得できない場合はEMPTY
BRGetData ( タブ , タグ指定 [ , n番目=1 ] )
エレメントのタグ名と順番を指定して値を取得する
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
タグ指定 ( 文字列 ) -- &quot;TAG=タグ名&quot; でタグ指定モードになる
n番目 ( 数値 省略可 ) -- 該当タグの順番を1から指定
戻り値 :
取得された値、取得できない場合はEMPTY
BRGetData ( タブ , タグ指定 , プロパティ指定 [ , n番目=1 ] )
エレメントのタグ名とプロパティを指定して値を取得する
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
タグ指定 ( 文字列 ) -- &quot;TAG=タグ名&quot; でタグ指定モードになる
プロパティ指定 ( 文字列 省略可 ) -- &quot;プロパティ名=値&quot; を指定可(&quot;id=hoge&quot; など)、プロパティ名のみ大文字小文字の一致が必須
n番目 ( 数値 省略可 ) -- タグもプロパティも一致する場合順番を1から指定
戻り値 :
取得された値、取得できない場合はEMPTY
プロパティ指定について
UWSCとは異なりID, className, innerText, innerHTML以外のプロパティも指定できます
ただし、プロパティ名は大文字小文字が一致する必要があります(case sensitive)
プロパティの値は大文字小文字を無視しますが、完全一致する必要があります
BRGetData ( タブ , &quot;TAG=TABLE&quot; [ , n番目=1 , 行=1 , 列=1 ] )
テーブルエレメントの座標を指定して値を取得する
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
&quot;TAG=TABLE&quot; ( 文字列 ) -- &quot;TAG=TABLE&quot; を指定(固定)
n番目 ( 数値 省略可 ) -- テーブルの順番を1から指定
行 ( 数値 省略可 ) -- テーブルの行番号を1から指定
列 ( 数値 省略可 ) -- テーブルの列番号を1から指定
戻り値 :
取得された値、取得できない場合はEMPTY
## IESETDATA互換#
BRSetData ( タブ , 値 , name [ , value=EMPTY , n番目=1 , 直接入力=FALSE ] ) #
テキストボックス等に文字列を入力する
キー入力をエミュレートします
input[type=&quot;file&quot;] 要素に対してはファイルパスを設定します
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
値 ( 文字列 ) -- 入力したい値、ファイルパス複数登録の場合は文字列配列も可
name ( 文字列 ) -- 値を変更するエレメントのname属性
value ( 文字列 省略可 ) -- 同一nameのエレメントがある場合にvalue値を指定
n番目 ( 数値 省略可 ) -- nameとvalueが一致する場合に順番を1から指定
直接入力 ( 真偽値 省略可 ) -- 直接valueプロパティを変更する場合はTRUE
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
BRSetData ( RemoteObject , 値 )
テキストボックス等に文字列を入力する
キー入力をエミュレートします
input[type=&quot;file&quot;] 要素に対してはファイルパスを設定します
パラメータ :
タブ ( RemoteObject ) -- 入力したいエレメントを示す RemoteObject
値 ( 文字列 ) -- 入力したい値、ファイルパス複数登録の場合は文字列配列も可
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
browser = BrowserControl ( BC_CHROME )
tab = browser [ 0 ]
file = tab . querySelector ( &quot;input[type=file]&quot; )
files = [ &#39;C:\test\hoge.txt&#39; , &#39;C:\test\fuga.txt&#39; ]
print BRSetData ( file , files )
BRSetData ( タブ , TRUE , name [ , value=EMPTY , n番目=1 ] )
nameにより指定したエレメントをクリックします
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
TRUE ( 真偽値 ) -- TRUEを指定 (固定)
name ( 文字列 ) -- クリックするエレメントのname属性
value ( 文字列 省略可 ) -- 同一nameのエレメントがある場合にvalue値を指定
n番目 ( 数値 省略可 ) -- nameとvalueが一致する場合に順番を1から指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
BRSetData ( タブ , TRUE , タグ指定 [ , n番目=1 ] )
タグ名と順番により指定したエレメントをクリックします
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
TRUE ( 真偽値 ) -- TRUEを指定 (固定)
タグ指定 ( 文字列 ) -- &quot;TAG=タグ名&quot; でダグ指定モードになる
n番目 ( 数値 省略可 ) -- タグ名が一致する場合に順番を1から指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
BRSetData ( タブ , TRUE , タグ指定 , プロパティ指定 [ , n番目=1 ] )
タグ名とプロパティにより指定したエレメントをクリックします
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
TRUE ( 真偽値 ) -- TRUEを指定 (固定)
タグ指定 ( 文字列 ) -- &quot;TAG=タグ名&quot; でダグ指定モードになる
プロパティ指定 ( 文字列 ) -- &quot;プロパティ名=値&quot; を指定
n番目 ( 数値 省略可 ) -- タグ名とプロパティが一致する場合に順番を1から指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
プロパティ指定について
UWSCとは異なりID, className, innerText, innerHTML以外のプロパティも指定できます
ただし、プロパティ名は大文字小文字が一致する必要があります(case sensitive)
プロパティの値は大文字小文字を無視しますが、完全一致する必要があります
BRSetData ( タブ , TRUE , &quot;TAG=IMG&quot; [ , src=EMPTY , n番目=1 ] )
IMGエレメントをクリックします
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
TRUE ( 真偽値 ) -- TRUEを指定 (固定)
&quot;TAG=IMG&quot; ( 文字列 ) -- &quot;TAG=IMG&quot; を指定 (固定)
src ( 数値 省略可 ) -- 対象imgタグのsrcを指定
n番目 ( 数値 省略可 ) -- srcが一致する場合に順番を1から指定
戻り値の型 :
真偽値
戻り値 :
成功時TRUE
## IEGETSRC互換#
BRGetSrc ( タブ , タグ名 [ , n番目=1 ] ) #
指定タグのエレメントのouterHTMLを返します
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
タグ名 ( 文字列 ) -- HTMLを取得したいタグ名
n番目 ( 数値 省略可 ) -- タグの順番を1から指定
戻り値の型 :
文字列
戻り値 :
該当タグのHTMLソース、非該当ならEMPTY
## IESETSRC互換#
非推奨関数
ドキュメント全体の書き換えを非推奨としているため、互換関数は存在しません
## IELINK互換#
BRLink ( タブ , リンク文字 [ , n番目=1 , 完全一致=FALSE ] ) #
指定リンクをクリックします
パラメータ :
タブ ( TabWindowオブジェクト ) -- 値を取りたいページのタブを示す TabWindowオブジェクト
リンク文字 ( 文字列 ) -- リンクに表示されている文字列(デフォルトは部分一致)
n番目 ( 数値 省略可 ) -- リンク文字が同一の場合に順番を1から指定
完全一致 ( 真偽値 省略可 ) -- TRUEの場合完全一致するリンク文字を検索する
戻り値の型 :
真偽値
戻り値 :
該当するリンクが存在しクリックを実行した場合TRUE
## IEGETFRAME互換#
後日実装予定
TabWindowがフレーム対応し次第実装する予定です
## BrowserBuilderオブジェクト#
ブラウザの起動、再接続、起動時設定を行うオブジェクト
class BrowserBuilder #
port ( port ) #
ブラウザのデバッグポートを変更します、デフォルトは 9222
パラメータ :
port ( 数値 ) -- 変更するデバッグポート
戻り値の型 :
BrowserBuilder
戻り値 :
更新されたBrowserBuilder
headless ( 有効 = TRUE ) #
ブラウザをヘッドレスで起動するかどうかを設定します
この設定は再接続時には無視されます
パラメータ :
有効 ( 真偽値 ) -- TRUEの場合ブラウザをヘッドレスで起動
戻り値の型 :
BrowserBuilder
戻り値 :
更新されたBrowserBuilder
private ( 有効 = TRUE ) #
ブラウザをプライベートモードで起動するかどうかを設定します
この設定は再接続時には無視されます
パラメータ :
有効 ( 真偽値 ) -- TRUEの場合ブラウザをプライベートモードで起動
戻り値の型 :
BrowserBuilder
戻り値 :
更新されたBrowserBuilder
profile ( プロファイルパス ) #
プロファイルを保存するパスを指定します
この設定は再接続時には無視されます
パラメータ :
プロファイルパス ( 文字列 ) -- プロファイルを保存するパス
戻り値の型 :
BrowserBuilder
戻り値 :
更新されたBrowserBuilder
argument ( 起動時オプション ) #
ブラウザの起動時オプションを追加します
動作保証対象外の機能です
これはブラウザ起動時のオプションを任意に追加できる機能です
この機能を利用した際の動作は保証されません
ブラウザ等への影響を理解している場合のみご利用ください
この機能を利用することにより生じた不具合はUWSCRのバグとしては扱われません
パラメータ :
起動時オプション ( 文字列 ) -- 追加する起動時オプション
戻り値の型 :
BrowserBuilder
戻り値 :
更新されたBrowserBuilder
サンプルコード
// ブラウザの拡張機能を無効にする
builder = BrowserBuilder ( BC_CHROME )
builder . argument ( &quot;--disable-extensions&quot; )
chrome = builder . start ()
start ( ) #
ブラウザを起動し Browserオブジェクト を返します
戻り値の型 :
Browserオブジェクト
戻り値 :
対象ブラウザの Browserオブジェクト
## Browserオブジェクト#
操作対象となるタブを示すオブジェクト
Browserオブジェクトの取得に時間がかかる場合がある
Browserオブジェクト作成時に対象ブラウザに対してWebSocket接続を行います
WebSocket接続が確立されるまでにある程度の時間を要するのが原因です
class Browser #
property count #
ブラウザ上の操作可能なタブの数を返します
tabs[i]
インデックスを指定し TabWindowオブジェクト を返します
配列表記対応
Browserオブジェクトに直接インデックス指定することもできます
chrome = BrowserControl ( BC_CHROME )
// タブの取得
tab = chrome . tabs [ 0 ]
// 以下のようにも書ける
tab = chrome [ 0 ]
close ( ) #
ブラウザを閉じます
戻り値 :
なし
new ( url ) #
指定したURLを新しいタブを開きます
パラメータ :
url ( 文字列 ) -- 開きたいサイトのURL
戻り値の型 :
TabWindowオブジェクト
戻り値 :
新しく開いたタブの TabWindowオブジェクト
id ( ) #
ブラウザのウィンドウIDを返します
戻り値の型 :
数値
戻り値 :
ウィンドウID
## TabWindowオブジェクト#
タブごとのWindowオブジェクトを示すオブジェクト
一度目のプロパティ取得やメソッド実行に時間がかかる場合がある
タブ内のページ操作のためにWebSocketを使用していますが、初回のみWebSocketの接続処理が入ります
WebSocket接続が確立されるまでにある程度の時間を要するのが原因です
class TabWindow #
property document #
window.document に相当する RemoteObject を返します
ブラウザ操作の基本はdocument取得から
RemoteObject はブラウザ上のJavaScriptオブジェクトです
document を起点に querySelector 等でエレメントにアクセスできます
RemoteObject のプロパティやメソッドの実行結果は RemoteObject として返ります
そのためブラウザ上でJavaScriptを実行するかのようにブラウザ操作を行うことが可能です
詳しくは ブラウザ操作サンプル を参照してください
navigate ( url ) #
指定URLを開きます
ページの読み込み完了まで待機します (最大10秒)
読み込み時間が長い場合
読み込みに10秒以上かかるページに対しては navigate実行後に wait メソッドを呼んでください
パラメータ :
url ( 文字列 ) -- 開きたいサイトのURL
戻り値の型 :
真偽値
戻り値 :
タイムアウトした場合FALSE
reload ( [ キャッシュ無視=FALSE ] ) #
ページをリロードします
ページの読み込み完了まで待機します (最大10秒)
読み込み時間が長い場合
読み込みに10秒以上かかるページに対しては navigate実行後に wait メソッドを呼んでください
パラメータ :
キャッシュ無視 ( 真偽値 ) -- TRUEならキャッシュを無視してリロード ( Shift+refresh と同等)
戻り値の型 :
真偽値
戻り値 :
タイムアウトした場合FALSE
wait ( [ タイムアウト秒=10 ] ) #
ページの読み込みが完了するのを待ちます
リンクをクリックした後などに使用します
パラメータ :
タイムアウト秒 ( 数値 省略可 ) -- 読み込み完了まで待機する最大時間 (秒)
戻り値の型 :
真偽値
戻り値 :
タイムアウトした場合はFALSE
activate ( ) #
タブをアクティブにします
戻り値 :
なし
close ( ) #
タブを閉じます
戻り値 :
なし
dialog ( [ 許可=TRUE , プロンプト=EMPTY ] ) #
JavaScriptダイアログ(alert, confirm, prompt等)を処理します
パラメータ :
許可 ( 真偽値 省略可 ) -- ダイアログを閉じる方法を指定、TRUEでOK、FALSEでキャンセル
プロンプト ( 文字列 省略可 ) -- promptに入力する文字列
戻り値 :
なし
サンプルコード
select tab . dlgtype ()
case &quot;prompt&quot;
// プロンプトなら文字を入力
tab . dialog ( TRUE , &quot;hogehoge&quot; )
case &quot;confirm&quot;
if pos ( &quot;hoge&quot; , tab . dlgmsg ()) &gt; 0 then
// メッセージに hoge という文字列が含まれていればOKを押す
tab . dialog ( TRUE )
else
// hoge が含まれていないものはキャンセル
tab . dialog ( FALSE )
endif
case EMPTY
// ダイアログがなければなにもしない
default
// その他のダイアログであれば閉じる
tab . dialog ()
selend
dlgmsg ( ) #
JavaScriptダイアログに表示されているメッセージを取得します
戻り値 :
メッセージ文字列、ダイアログがない場合はEMPTY
dlgtype ( ) #
JavaScriptダイアログの種類を取得します
種類は以下のいずれかです
alert
confirm
prompt
beforeunload
戻り値 :
種類を示す文字列、ダイアログがない場合はEMPTY
leftClick ( x , y ) #
rightClick ( x , y ) #
middleClick ( x , y ) #
マウスクリックイベントを発生させます
それぞれ左クリック、右クリック、中央クリックを行います
パラメータ :
x ( 数値 ) -- ブラウザのビューポート上のX座標 (CSSピクセル単位、左上から)
y ( 数値 ) -- ブラウザのビューポート上のY座標 (CSSピクセル単位、左上から)
戻り値 :
なし
サンプルコード
// エレメントの取得
element = browser [ 0 ] . document . querySelector ( selector )
// getBoundingClientRectメソッドでエレメントの座標等の情報を得る
rect = element . getBoundingClientRect ()
// 座標を指定し右クリックする
tab . rightClick ( rect . x + 10 , rect . y + 10 )
eval ( JavaScript式 ) #
JavaScriptの式を評価し、オブジェクトの場合はRemoteObjectとして返します
パラメータ :
JavaScript式 ( 文字列 ) -- JavaScriptの式
戻り値の型 :
RemoteObject またはいずれかの値型
戻り値 :
評価結果がJavaScriptオブジェクトの場合は RemoteObject を返します
そうでない場合は該当するUWSCRの値型を返します
サンプルコード
chrome = BrowserControl ( BC_CHROME )
tab = chrome [ 0 ]
tab . navigate ( url )
func = tab . eval ( &quot;(a, b) =&gt; a + b&quot; ) // アロー関数を評価
print func ( 3 , 5 ) // 8 (関数として実行できる)
// コールバック用のJavaScript関数を作る
callback = tab . eval ( &quot;(event) =&gt; event.srcElement.style.backgroundColor = &#39;red&#39;&quot; )
slct = tab . document . querySelector ( &quot;select&quot; )
// イベントリスナをセット
slct . addEventListener ( &quot;change&quot; , callback )
## JavaScriptダイアログについて#
バージョン0.15.0まではUWSCRのスクリプトによりalert等のJavaScriptダイアログが表示された場合に動作がブロックされる問題がありました
0.16.0でこの問題が改善されましたがダイアログを TabWindow.dialog() で閉じない限り続くプロパティやメソッドがブロックされる場合があります
// スクリプトからalertを開く
tab . eval ( &quot;alert(&#39;hoge&#39;);&quot; ) // 0.16.0以降はブロックされない
tab . dialog () // ダイアログを閉じる
print tab . document // 正常動作
// ブロックされるパターン
tab . eval ( &quot;alert(&#39;hoge&#39;);&quot; )
// ダイアログを閉じずにプロパティにアクセスする
print tab . document // ダイアログが閉じられるまでブロックされる
## RemoteObject#
ブラウザ上に存在するJavaScriptオブジェクトを示すオブジェクト
## メソッドの実行#
RemoteObject.メソッド名(引数) でメソッドを実行します
メソッド名は大文字小文字を区別します
chrome = BrowserControl ( BC_CHROME )
foo = chrome [ 0 ] . document . querySelector ( &quot;#foo&quot; )
## プロパティの取得#
RemoteObject.プロパティ名 とすることでプロパティ値を取得します
配列要素であればインデックスを指定します RemoteObject.プロパティ名[i]
プロパティ名は大文字小文字を区別します
chrome = BrowserControl ( BC_CHROME )
url = chrome [ 0 ] . document . URL
## プロパティの変更#
RemoteObject.プロパティ名 = 値 とすることでプロパティ値を変更します
配列要素であればインデックスを指定します RemoteObject.プロパティ名[i] = 値
プロパティ名は大文字小文字を区別します
chrome = BrowserControl ( BC_CHROME )
foo = chrome [ 0 ] . document . querySelector ( &quot;#foo&quot; )
foo . value = &quot;ほげほげ&quot;
## インデックスによるアクセス#
RemoteObject 自身が配列であった場合は RemoteObject[i] とすることで要素を得られます
chrome = BrowserControl ( BC_CHROME )
links = chrome [ 0 ] . document . querySelectorAll ( &quot;a&quot; )
print links [ 0 ] . href
## 関数として実行#
RemoteObject 自身が関数である場合は RemoteObject(引数) として実行できます
## 非同期関数とPromise#
RemoteObject 自身、またはそのメソッドが非同期関数であった場合 await 構文でその終了を待ちます
RemoteObject がPromiseであった場合は WaitTask 関数でその終了を待ちます
いずれの場合も結果を返します
## 戻り値について#
RemoteObject のプロパティやメソッド、インデックスから得られる値の型は以下の通りです
JavaScript型
UWSCR型
string
文字列
number
数値
bool
真偽値
null
NULL
上記以外のオブジェクト
RemoteObject
オブジェクトでもプリミティブな値でもない場合 (undefinedなど)
EMPTY
## ブラウザ操作サンプル#
documentへのアクセス
// ブラウザを開く
chrome = BrowserControl ( BC_CHROME )
// ひとつめのタブを得る
tab1 = chrome . tabs [ 0 ]
// 以下のようにも書けます
// tab1 = chrome[0]
// 任意のサイトを開く
tab1 . navigate ( url )
// window.documentを得る
document = tab1 . document
// URLを得る
print document . URL
タブごとのURLを列挙
// タブの数を得る
print chrome . count
// URLを列挙
for tab in chrome . tabs
print tab . document . URL
next
// 以下のようにも書けます
// for tab in chrome
// print tab.document.URL
// next
自動操作用ブラウザを別途開く
// デバッグポートを開いていないブラウザがすでに開かれている場合
// 以下は再接続ができずエラーになる
// chrome = BrowserControl(BC_CHROME)
// プロファイルフォルダを指定して別のブラウザを起動する
chrome = BrowserBuilder ( BC_CHROME ) . profile ( &quot;C:\chrome\profile1&quot; ) . start ()
Seleniumテストページの操作
// ブラウザを開く
chrome = BrowserControl ( BC_CHROME )
// ブラウザをアクティブにする
ctrlwin ( chrome . id () , ACTIVATE )
// 新しいタブでSeleniumのテストページを開く
tab = chrome . new ( &#39;http://example.selenium.jp/reserveApp_Renewal/&#39; )
// ドキュメントを取得しておく
document = tab . document
// 宿泊日を入力
// 3日後の日付を得る
date = format ( gettime ( 3 , , G_OFFSET_DAYS ) , &#39;%Y/%m/%d&#39; )
document . querySelector ( &#39;#datePick&#39; ) . value = date
document . querySelector ( &#39;#reserve_year&#39; ) . value = G_TIME_YY4
document . querySelector ( &#39;#reserve_month&#39; ) . value = G_TIME_MM2
document . querySelector ( &#39;#reserve_day&#39; ) . value = G_TIME_DD2
// 宿泊日数を選択
reserve_term = 2
document . querySelector ( &quot;#reserve_term option[value=&#39;&lt;#reserve_term&gt;&#39;]&quot; ) . selected = TRUE
// 人数を選択
headcount = 5
document . querySelector ( &quot;#headcount option[value=&#39;&lt;#headcount&gt;&#39;]&quot; ) . selected = TRUE
// プラン選択
// お得な観光プランをチェック
document . querySelector ( &#39;#plan_b&#39; ) . checked = TRUE
// 名前入力
document . querySelector ( &#39;#guestname&#39; ) . value = &quot;おなまえ&quot;
// 利用規約に同意して次へ をクリック
document . querySelector ( &#39;#agree_and_goto_next&#39; ) . click ()
// 読み込み完了を待つ
tab . wait ()
// ページを移動したのでdocumentは取得しなおす
document = tab . document
// 合計金額を得る
price = document . querySelector ( &#39;#price&#39; ) . textContent
// 確定ボタンを押す
document . querySelector ( &#39;#commit&#39; ) . click ()
msgbox ( &quot;宿泊費用は&lt;#price&gt;円でした&quot; )
// タブを閉じる
tab . close ()
## ダウンロード先やその方法の制御について#
ダウンロードファイルの保存先フォルダの指定や、確認ダイアログの制御が現時点ではできません
ブラウザ操作にて特定のフォルダへのダウンロードを確認なしで行いたい場合は事前に以下の操作を行ってください
BrowserBuilderオブジェクト で専用のプロファイルフォルダを指定し、ブラウザを起動する
起動したブラウザの設定を手動で変更する
Chrome
設定画面の ダウンロード を開く
保存先 を任意のフォルダに変更する
ダウンロード前に各ファイルの保存場所を確認する をオフにする
MSEdge
設定画面の ダウンロード を開く
場所 を任意のフォルダに変更する
ダウンロード時の動作を毎回確認する をオフにする
変更を施したプロファイルを指定して改めてブラウザ操作を行う
ダウンロード開始と完了の検知
getdir関数で 未確認*.crdownload ファイルの数を確認し、1個以上であればダウンロードが開始されていると判定
ダウンロードするファイルの名前がわかっている場合、F_EXISTSがTRUEならダウンロード完了
あるいはgetdir関数で 未確認*.crdownload ファイルの数を確認し、0個であればダウンロード完了と判定
// ダウンロード開始検知
repeat
sleep ( 0 . 1 )
files = getdir ( download_path , &quot;未確認*.crdownload&quot; )
until length ( files ) &gt; 0
if filename != EMPTY then
// ファイル名が分かる場合
repeat
sleep ( 1 )
until fopen ( filename , F_EXISTS )
else
// ファイル名が分からない場合
repeat
sleep ( 1 )
files = getdir ( download_path , &quot;未確認*.crdownload&quot; )
until length ( files ) == 0
endif
## HTTPリクエスト#
プロキシ環境下における注意
OSのプロキシサーバー設定が有効な場合 Webrequest および WebRequestBuilder はその設定に従いプロキシサーバーを経由した通信を試みます
特定のドメインなどでプロキシサーバーを迂回すべく除外設定を行っていてもこれらの関数に反映されない場合があります
このような場合はプロセス環境変数 NO_PROXY で除外設定を行います
環境変数 NO_PROXY にはカンマ , 区切りで除外したいドメイン名等を指定します
// 環境変数名
const NO_PROXY = &#39;NO_PROXY&#39;
// 複数指定はカンマ区切り
const NO_PROXY_LIST = &#39;localhost,localserver.dev&#39;
// setenv関数でプロセス環境変数をセット
setenv ( NO_PROXY , NO_PROXY_LIST )
// プロセス環境変数の反映を確認
assert_equal ( NO_PROXY_LIST , env ( NO_PROXY ))
// プロキシサーバーを迂回したリクエストを送信
res = WebRequest ( &#39;http://localhost:8888/&#39; )
res = WebRequest ( &#39;http://localserver.dev/hoge/&#39; )
Webrequest ( url ) #
指定URLに対してGETリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
レスポンスを示す WebResponseオブジェクト
サンプルコード
res = WebRequest ( &quot;http://example.com&quot; )
print res . status
print res . body
WebRequestBuilder ( ) #
WebRequestオブジェクト を返します
WebRequest とは異なり詳細な設定を行い任意のメソッドでリクエストを送信できます
戻り値の型 :
WebRequestオブジェクト
戻り値 :
リクエストを行うための WebRequestオブジェクト
## WebRequestオブジェクト#
HTTPリクエストを行うためのオブジェクト
class WebRequest #
useragent ( UA ) #
UserAgent文字列をUser-Agentヘッダに設定します
未指定の場合設定されません
パラメータ :
UA ( 文字列 ) -- UserAgent文字列
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
header ( キー , 値 ) #
リクエストヘッダを追加します
パラメータ :
キー ( 文字列 ) -- ヘッダのキー
値 ( 文字列 ) -- ヘッダの値
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
timeout ( 秒 ) #
ヘッダを設定します
未指定の場合タイムアウトしません
パラメータ :
秒 ( 数値 ) -- タイムアウト秒
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
body ( 本文 ) #
リクエスト本文を設定します
未指定の場合は何も送信しません
パラメータ :
本文 ( 文字列またはUObject ) -- リクエスト本文、UObjectはjsonに変換されます
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
basic ( ユーザー名 [ , パスワード=EMPTY ] ) #
Basic認証のユーザー名とパスワードを設定したAuthorizationヘッダを追加します
未指定の場合は追加されません
パラメータ :
ユーザー名 ( 文字列 ) -- ユーザー名
パスワード ( 文字列 省略可 ) -- パスワード
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
bearer ( トークン ) #
Bearer認証のトークンを設定したAuthorizationヘッダを追加します
未指定の場合は追加されません
パラメータ :
トークン ( 文字列 ) -- 認証トークン
戻り値の型 :
WebRequestオブジェクト
戻り値 :
更新された WebRequestオブジェクト
get ( url ) #
GETリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
put ( url ) #
PUTリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
post ( url ) #
POSTリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
delete ( url ) #
DELETEリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
patch ( url ) #
PATCHリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
head ( url ) #
HEADリクエストを送信します
パラメータ :
url ( 文字列 ) -- リクエストを送るURL
戻り値の型 :
WebResponseオブジェクト
戻り値 :
WebResponseオブジェクト
サンプルコード
request = WebRequestBuilder ()
// ヘッダと認証情報を設定しておく
request . bearer ( MY_BEARER_TOKEN ) _
. header ( &#39;Content-Type&#39; , &#39;application/json&#39; )
// リクエストを送信
res1 = request . body ( json1 ) . post ( url1 )
res2 = request . body ( json2 ) . put ( url2 )
## WebResponseオブジェクト#
HTTPレスポンスを示すオブジェクト
class WebResponse #
property status #
レスポンスのステータスを数値で返します
property statusText #
レスポンスのステータスを示す文字列を返します
property succeed #
リクエストの成否を真偽値で返します
property header #
レスポンスヘッダを連想配列で返します
property body #
レスポンスボディを文字列で返します、返せない場合はEMPTY
property json #
レスポンスボディがjsonの場合UObjectを返します、返せない場合はEMPTY
## HTTPパーサー#
ParseHTML ( html ) #
HTMLをパースし HtmlNodeオブジェクト を返します
パラメータ :
html ( 文字列またはWebResponse ) -- HTMLドキュメントまたはその一部を示す文字列、またはHTMLドキュメントとして受けた WebResponseオブジェクト
戻り値の型 :
HtmlNodeオブジェクト
戻り値 :
パースされたHTMLドキュメントを示す HtmlNodeオブジェクト (ルートエレメント)
部分HTMLのパースについて
HTMLドキュメント全体ではなく一部をパースした場合に意図した結果が返らない場合があります
例: tbody``以下 (親の ``table がない)
&lt; tbody &gt;
&lt; tr &gt;
&lt; td &gt; りんご &lt;/ td &gt;
&lt; td &gt; バナナ &lt;/ td &gt;
&lt; td &gt; メロン &lt;/ td &gt;
&lt;/ tr &gt;
&lt;/ tbody &gt;
この動作は現状では仕様とし、このような部分パースは非推奨とします
サンプルコード
res = WebRequest ( url )
// WebResponseオブジェクトからHtmlNodeオブジェクトを得る
doc = ParseHTML ( res )
// ラジオボタンのvalue値を列挙
for radio in doc . find ( &#39;input[type=&quot;radio&quot;]&#39; )
print radio . attr ( &#39;value&#39; )
next
// 最初のselect要素内のoptionのテキストと値を列挙
slct = doc . first ( &#39;select&#39; )
for opt in slct . find ( &#39;option&#39; )
print opt . text
print opt . attr ( &#39;value&#39; )
next
## HtmlNodeオブジェクト#
以下のいずれかを示します
ルートエレメント
エレメントのコレクション
エレメント
class HtmlNode #
find ( selector ) #
cssセレクタに該当するエレメント郡を HtmlNodeオブジェクト の配列として返す
オブジェクトがコレクションの場合はEMPTYを返す
パラメータ :
selector ( 文字列 ) -- cssセレクタ
戻り値の型 :
HtmlNodeオブジェクト (コレクション)
戻り値 :
cssセレクタに該当するエレメントのコレクション
first ( selector ) #
findfirst ( selector ) #
cssセレクタに該当する最初のエレメントを HtmlNodeオブジェクト として返す
オブジェクトがコレクションの場合はEMPTYを返す
パラメータ :
selector ( 文字列 ) -- cssセレクタ
戻り値の型 :
HtmlNodeオブジェクト (エレメント)
戻り値 :
cssセレクタに該当する最初のエレメント
attr ( 属性名 ) #
attribute ( 属性名 ) #
該当するエレメント属性の値を返す
該当する属性がなかった場合はEMPTY
コレクションの場合はそれぞれの属性値の配列を返す
パラメータ :
属性名 ( 文字列 ) -- 属性の名前
戻り値の型 :
文字列またはEMPTY、またはそれらの配列
戻り値 :
該当する属性の値、属性がない場合EMPTY
textContent ( [ 結合文字=&quot;&quot; , トリム=TRUE ] )
エレメント内のテキストノードを結合してひとつの文字列として返す
コレクションの場合はそれぞれのエレメントについて結合処理を行い、それらを配列として返す
パラメータ :
結合文字 ( 文字列 省略可 ) -- この文字を挟んで結合する
トリム ( 真偽値 省略可 ) -- 結合前に各テキストノードの前後のホワイトスペースを除去する
戻り値の型 :
文字列、または文字列の配列
戻り値 :
結合した文字列またはその配列
count ( ) #
コレクションの場合は要素の数を返す
それ以外はEMPTY
戻り値の型 :
数値またはEMPTY
戻り値 :
コレクションの要素数
property outerhtml #
エレメントのHTMLの文字列を返す
コレクションの場合はそれぞれのHTMLの配列を返す
property innerhtml #
エレメント配下のHTMLの文字列を返す
コレクションの場合はそれぞれのHTMLの配列を返す
property text #
エレメント内のテキストノードの配列を返す
コレクションの場合はそれぞれのテキストノード配列の配列を返す
property textContent #
エレメント内のテキストノードを結合してひとつの文字列として返す
コレクションの場合はエレメント毎に結合された文字列の配列を返す
node.textContent() とした場合と同等
property isRoot #
ルートエレメントかどうか
property isElement #
エレメントかどうか
ルートエレメントの場合もTRUEを返す
property isCollection #
コレクションかどうか
## コレクションのインデックスアクセス#
コレクションは配列のようにインデックスアクセスによりエレメントを返します
root = ParseHTML ( html )
inputs = root . find ( &quot;form input&quot; )
print inputs [ 0 ] . attr ( &quot;value&quot; )
// for-in対応
for input in inputs
print input . attr ( &quot;value&quot; )
next

---

# ソケット通信 - UWSCR 1.1.9

> 参照元: https://stuncloud.github.io/UWSCR/builtins/socket.html

## ソケット通信#
UDP
TCP
WebSocket
を利用した通信を行うための関数群です
## 共通#
sclose ( ソケット ) #
ソケットを閉じます
注釈
自動クローズ
ソケットはソケットオブジェクトが破棄された時点で閉じられます
以下の方法のいずれでもソケットを閉じることができます
udp1 = udpclient ( addr , port )
sclose ( udp1 ) // sclose関数を使う
udp2 = udpclient ( addr , port )
udp2 = EMPTY // 別の値で上書きしソケットオブジェクトが失われれば自動クローズされる
パラメータ :
ソケット ( ソケットオブジェクト ) -- 以下のいずれかを指定
UDPClient
WebSocket
## UDP通信#
サンプルコード
const PORT_SEND = 50101
const PORT_RECV = 50303
function sender ()
// 呼び出しから3秒後にデータを送信する
const LOCALHOST = &quot;127.0.0.1&quot;
client = udpclient ( LOCALHOST , PORT_SEND )
sleep ( 3 )
udpsend ( client , LOCALHOST , PORT_RECV , &quot;UDP通信テスト&quot; )
fend
client = udpclient ( &quot;0.0.0.0&quot; , PORT_RECV )
// 送信スレッドを呼ぶ
thread sender ()
// 受信待機
r = udprecv ( client , 100 )
// 受信データを整形
data = decode ( r [ 0 ] , CODE_BYTEARRAYU )
addr = r [ 1 ]
port = r [ 2 ]
print &quot;&lt;#addr&gt;:&lt;#port&gt; からメッセージを受信しました: &lt;#data&gt;&quot;
// 127.0.0.1:50101 からメッセージを受信しました: UDP通信テスト
sclose ( client )
UdpClient ( IPアドレス , ポート ) #
任意のアドレスとポートで待ち受けるUDPクライアントオブジェクトを返す
パラメータ :
IPアドレス ( 文字列 ) -- 自身の待ち受けIPアドレス
ポート ( 数値 ) -- 自身の待ち受けポート
戻り値の型 :
UDPクライアント
戻り値 :
UDP送受信を行うためのオブジェクト
UdpSend ( udp , IPアドレス , ポート , 送信データ ) #
UDPによるデータ送信を行う
パラメータ :
udp ( UDPクライアント ) -- データを送信するUDPクライアント
IPアドレス ( 文字列 ) -- 送信先IPアドレス
ポート ( 数値 ) -- 送信先ポート
送信データ ( 値 ) --
以下のいずれかの型の値に対応
文字列: UTF8バイト配列に変換される
UObject: json文字列としてUTF8バイト配列に変換される
バイト配列: encode関数の戻り値等
数値配列: 数値 (0-255) の配列、数値以外や範囲外が含まれていたらエラーとなる
戻り値の型 :
真偽値
戻り値 :
送信成功時TRUE
UdpRecv ( バッファサイズ ) #
UDPによるデータ受信を行う
データを受信するまでブロックする
パラメータ :
バッファサイズ ( 数値 ) --
受信するデータ (バイト配列) のバッファサイズ
実際の受信データより小さいとデータが欠損する場合があります
戻り値の型 :
[バイト配列, 文字列, 数値]
戻り値 :
[受信データ, 送信元IPアドレス, 送信元ポート]
## TCP通信#
TcpSend ( IPアドレス , ポート , 送信データ ) #
TCPで接続先にデータを送信し、受け取ったレスポンスを返す
パラメータ :
IPアドレス ( 文字列 ) -- 対象サーバーのIPアドレス
ポート ( 数値 ) -- 対象サーバーのポート
送信データ ( 値 ) --
以下のいずれかの型の値に対応
文字列: UTF8バイト配列に変換される
UObject: json文字列としてUTF8バイト配列に変換される
バイト配列: encode関数の戻り値等
数値配列: 数値 (0-255) の配列、数値以外や範囲外が含まれていたらエラーとなる
戻り値の型 :
バイト配列
戻り値 :
レスポンスデータを示すバイト配列
サンプルコード
// example.comにGETリクエストを送る
// GETリクエストデータ
// 末尾に改行を2つ入れないとダメ
textblock request
GET /index.html HTTP/1.1
Host: example.com
Connection: close
endtextblock
// GETリクエストを送信
res = TcpSend ( &quot;23.192.228.80&quot; , 80 , request )
// レスポンスデータを文字列に変換してprint
print decode ( res , CODE_BYTEARRAYU )
TcpListener ( IPアドレス , ポート , ハンドラ [ , 終端文字=&quot;&lt;#CR&gt;&quot; , タイムアウト秒=10 ] ) #
指定アドレス及びポートでTCP接続の待ち受けを行う
パラメータ :
IPアドレス ( 文字列 ) -- 待ち受けIPアドレス
ポート ( 数値 ) -- 待ち受けポート
ハンドラ ( 関数 ) --
受信したデータをバイト配列として受け、クライアントに返信するデータを戻り値とする関数
返信に有効な型は以下
文字列: UTF8バイト配列に変換され返信される
UObject: json文字列がUTF8バイト配列に変換され返信される
バイト配列: encode関数の戻り値等
数値配列: バイト配列に変換可能であれば返信される
FALSE, NULL, EMPTY: 待ち受け状態を抜ける (クライアントには空データが返る)
終端文字 ( 文字 省略可 ) --
受信データの終端と判断するASCII文字 (chr(0)～chr(255))
この文字が送られてこないとデータ受信が終わらずレスポンスを返せない
省略時はCRLF ( &quot;#CR&quot; )
タイムアウト秒 ( 数値 ) -- 受信できない場合のタイムアウト秒 (終端文字が送られない場合などにタイムアウトする可能性がある)
戻り値 :
なし
サンプルコード
// 受信データハンドラ
// 受信内容により返信を変更する
function handler ( bytes )
received = decode ( bytes , CODE_BYTEARRAYU )
select received
case &quot;Ping&quot;
result = &quot;Pong&quot;
case &quot;さようなら&quot;
result = &quot;またね&quot;
default
result = &quot;こんにちは、&lt;#received&gt;さん&quot;
selend
fend
// 別スレッドでリッスン開始
thread TcpListener ( &quot;0.0.0.0&quot; , 9999 , handler )
// データ送信関数ラッパー
send = function ( data : string )
// デフォルトではTcpListenerのデータ終端が改行なので末尾に&lt;#CR&gt;を加える
res = TcpSend ( &quot;127.0.0.1&quot; , 9999 , &quot;&lt;#data&gt;&lt;#CR&gt;&quot; )
result = decode ( res , CODE_BYTEARRAYU )
fend
sleep ( 1 )
print send ( &quot;🐊&quot; )
// こんにちは、🐊さん
sleep ( 1 )
print send ( &quot;Ping&quot; )
// Pong
sleep ( 1 )
print send ( &quot;さようなら&quot; )
// またね
## WebSocket#
サンプルコード
// MSEdgeのデバッグポートを開いて起動
shexec ( &quot;msedge.exe&quot; , &quot;--remote-debugging-port=9515&quot; )
sleep ( 1 )
// WebSocket用のURLを得る
res = webrequest ( &quot;http://localhost:9515/json/version&quot; )
uri = res . json . webSocketDebuggerUrl
print &quot;webSocketDebuggerUrl: &lt;#uri&gt;&quot;
// WebSocketオブジェクトを作成
ws = WebSocket ( uri )
print ws
// リクエスト用jsonオブジェクトを作る
request = @{
&quot;id&quot; : 1 ,
&quot;method&quot; : &quot;Target.getTargets&quot; ,
&quot;params&quot; : {}
}@
// リクエストを送信
WsSend ( ws , request )
while TRUE
// データを受信
res = WsRecv ( ws )
obj = fromjson ( res )
if obj . id == request . id then
// idが一致したら抜ける
break
endif
wend
// Target.getTargetsメソッドの戻り値のうち、ページを示すものの情報を表示
for info in obj . result . targetInfos
if info . type == &quot;page&quot; then
print
print &quot;type : &quot; + info . type
print &quot;title: &quot; + info . title
print &quot;url : &quot; + info . url
endif
next
WebSocket ( wsuri ) #
WebSocketに接続する
パラメータ :
wsuri ( 文字列 ) -- ws:// から始まるURI
戻り値の型 :
WebSocket
戻り値 :
WebSocketオブジェクト
WsSend ( WebSocket , 送信データ ) #
WebSocketでデータを送信する
パラメータ :
WebSocket ( WebSocket ) -- WebSocketオブジェクト
送信データ ( 値 ) --
以下のいずれかの型の値に対応
文字列
UObject
バイト配列
定数
WS_PING : pingを送信する
WS_PONG : pongを送信する
戻り値の型 :
戻り値の型
戻り値 :
戻り値の説明
WsRecv ( WebSocket ) #
WebSocketでデータを受信する
パラメータ :
WebSocket ( WebSocket ) -- WebSocketオブジェクト
戻り値の型 :
文字列、バイト配列、定数、EMPTY
戻り値 :
受信データによる
受信データの型に注意
データの型が不明な場合は type_of 関数で型のチェックを行ってください
res = WsRecv ( ws )
select type_of ( res )
case TYPE_STRING
print &quot;received string: &lt;#res&gt;&quot;
case TYPE_BYTE_ARRAY
print &quot;received bytes: &lt;#res&gt;&quot;
case TYPE_NUMBER
select res
case WS_PING
print &quot;received ping&quot;
case WS_PONG
print &quot;received pong&quot;
selend
default
print &quot;received invalid data: &lt;#res&gt;&quot;
selend

---
