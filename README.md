# go-basic

## コマンド
- go
  goのヘルプ表示

- go help {コマンド}
  コマンドのヘルプ、ドキュメントを参照
  ex. go help version

- go version
  バージョン確認

- go env
  環境変数確認

- go fmt
  ソースコードの整形
  インデントなどの乱れがあるディレクトリで行う
  opt.
    -n 実行されるコマンドの表示(ファイルは書き換えない)　対象のファイルを確認する時に使う
    -x 実行されるコマンドの表示

- go doc パッケージ名
  パッケージのドキュメントを参照する
  ex. go doc math/rand

- go doc パッケージ名.識別子,メソッド名の指定
  識別子（関数や定数、変数など）を指定
  ex. go doc math/rand.Intn
  ex. go doc time.Time.Unix
  opt.
    -c 識別子のマッチングで大文字小文字を厳密に判定する
    -u 非公開な識別子やメソッドについてもドキュメントを出力する

- go build
  プログラムのビルド
  ファイルやパッケージ名を指定しないビルドはカレントディレクトリ内の*.goファイルを対象にコンパイルする。
  ビルド結果として、カレントディレクトリの名前を持つ実行ファイルを生成する。
  １つのディレクトリ内には１つのパッケージのみ定義可能
  通常のgo buildでは１つのディレクトリ内に複数パッケージを含む状態ではエラーを発生させる。個別にファイルを指定すればビルド可能
  ex. go build main.go config.go
  この場合生成される実行ファイル名は最初に指定したファイル名になる。(この場合main)
  実行ファイル名の指定の場合
  ex. go build -o appname main.go config.go
  opt.
    -x 内部処理の出力 go build -x
    -o 生成する実行ファイル名の指定 go build -o ファイル名

- go install パッケージ名
  パッケージや実行ファイルをビルドした結果をGOPATH内にインストールする
  $GOPATHな/srcに置かれたパッケージのビルドした結果が実行ファイルであれば$GOPATH/binへ、それ以外であれば$GOPATH/pkgへインストールされる。
  $GOPATH - bin (実行ファイルのインストール先)
              - pkg (ビルドしたパッケージのインストール先)
      - src (パッケージのソースコードのインストール先)

- go get パッケージ
  外部パッケージのダウンロードとインストールをまとめて実行する
  GitHubなどのサービス上で開発されているGoのパッケージなどに使用することができる。
  ダウンロード対象のパッケージがさらに別のパッケージに依存している場合でも、自動的に依存関係を抽出しつつ必要なパッケージのダウンロードとインストールを実行してくれる。
  GitHubでは数多くのGoのパッケージが開発されていて、これらの資産を生かすためにもGitの導入は必須である。
  ex. go get golang.org/x/net/http2
  HTTPプロトコルの最新版であるHTTP2に対応した拡張パッケージをインストールする
  このパッケージはgolang.orgによってホストされているため、パッケージにはgolang.org/x/net/http2を指定する。
  何も出力されなければ成功。内部の処理を確認したい場合は-xオプションを付与すれば確認できる。
  go install同様に$GOPATH内に保存される。
  このパッケージを使用する場合は以下のように指定する
  import "golang.org/x/net/http2"
  GitHubでホストされているパッケージも同様の手順で導入できる。
  opt.
    -d 対象パッケージのダウンロードのみ実行して停止する
    -f 対象パッケージのパスから推測されるリポジトリへの検証をスキップする (-u指定時のみ）
    -fix 対象パッケージの依存関係を解決する前にgo fixツールを適用する。
    -insecure カスタムリポジトリを使用する場合に非セキュアなプロトコルの使用を許可する。（例：HTTP)
    -t 対象パッケージに付属するテストが依存するパッケージも合わせてダウンロードする。
    -u 対象パッケージの更新と依存パッケージの更新を検出して再ダウンロードとインストールを実行する。

- go test パッケージ名
  テストの実行

## VSCode拡張機能
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - 拡張機能「Go」を入れるだけで、汎用的で優れたGo言語のサポート機能を使うことができます。
  - コードスニペット（予測変換機能 ）を活用して、効率よくコーディングしていきましょう。
- [Go Outliner](https://marketplace.visualstudio.com/items?itemName=766b.go-outliner)
  - Go言語における、アウトライン機能を強化してくれる拡張機能です。
- [Go Test Exlorer](https://marketplace.visualstudio.com/items?itemName=premparihar.gotestexplorer)
  - Go言語における、テスト機能を強化してくれる拡張機能です。
  - テストの実行結果が正しいか正しくないかがアイコンから直感的にわかります。
- [ indent-rainbow](https://marketplace.visualstudio.com/items?itemName=oderwat.indent-rainbow)
  - ソースコードのインデントの深さを色分けして表示します。
- [ZenKaku](https://marketplace.visualstudio.com/items?itemName=mosapride.zenkaku)
  - Zenkakuを入れることで、全角スペースがグレーで表示されるようになります。
- [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)
  - Better Commetsを入れることで、コメントを重要度と種類によって色分けすることができます。
  - コメント直後の半角記号か文字列によって色の指定ができます。
  ```
  !: 赤色　? : 青色　* : 黄緑　TODO: : オレンジ
  ```
- [Trailing Spaces](https://marketplace.visualstudio.com/items?itemName=shardulm94.trailing-spaces)
  - Trailing Spacesを入れることで、コード末尾の余計なスペースが赤字で表示され、ソースを綺麗に保つことができます。
