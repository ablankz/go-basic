package standard

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Os() {
	// Exit 終了ステータス指定
	// deferも実行されない
	// os.Exit(1)

	f, err := os.Open("test.txt")
	if err != nil {
		// ログ出力後Exit(1)が呼ばれる
		log.Fatalln(err)
	}
	f.Close()

	fmt.Println(os.Args[0]) // コマンド名
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	// 存在する場合、新規に作り直す
	foo, _ := os.Create("foo.txt")
	foo.Write([]byte("Hello\n"))
	// 書き始めを指定
	foo.WriteAt([]byte("Golang"), 6)
	// seekを末尾にオフセットを移動する
	foo.Seek(0, io.SeekEnd)
	// オフセットの位置から書き始めるため追記になる
	foo.WriteString("Yeah")

	defer foo.Close()

	// エラーハンドリング省略
	bs := make([]byte, 128)
	// nには読み込んだバイト数, bsには文字列
	n, _ := foo.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	// オフセットの指定
	nn, _ := foo.ReadAt(bs, 10)
	fmt.Println(nn)
	fmt.Println(string(bs))

	// フラグ(複数ならパイプでつなぐ)
	// O_RDONLY
	// O_WRONLY
	// O_RDWR
	// O_APPEND 末尾に追記
	// O_CREATE なければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする
	// 第三引数 パーミッション
	ff, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
	}

	defer ff.Close()

	// エラーハンドリング省略
	bs = make([]byte, 128)
	nnn, _ := ff.Read(bs)
	fmt.Println(nnn)
	fmt.Println(string(bs))

	// ファイル情報
	fi, _ := ff.Stat()
	fmt.Println(fi.Name())
	fi.Size()
	fi.Mode()
	fi.ModTime()
	fi.IsDir()

	_ = os.Remove("foo.txt")

	_ = os.Rename("bar.txt", "foo.txt")
	_ = os.Rename("foo.txt", "bar/foo.txt")

	// カレントディレクトリ
	dir, _ := os.Getwd()
	fmt.Println(dir)

	// ディレクトリを移動する
	_ = os.Chdir("path/to/dir")

	// ディレクトリをパーミッションを付与して作成
	_ = os.Mkdir("test", 0775)
	// ネスト分すべて作成
	_ = os.MkdirAll("foo/fooo/foooo", 0775)

	// カレントディレクトリを開く
	d, err := os.Open(".")
	if err != nil {
		log.Println(err)
	}
	defer d.Close()

	// ディレクトリ内を見る
	fis, _ := d.Readdir(0)
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(fi)
		}
	}

	// ホスト名
	host, _ := os.Hostname()
	fmt.Println(host)

	// すべての環境変数
	for _, v := range os.Environ() {
		fmt.Println(v)
	}

	// 現在のプロセスid
	fmt.Println(os.Getpid())
	// 親プロセスid
	os.Getppid()
	// 呼び出し元の数値ユーザーID
	os.Getuid()
	// 実行ユーザーid root権限だと0
	os.Geteuid()
	// 呼び出し元の数値ユーザーグループID
	os.Getgid()
	// 実行ユーザーグループid
	os.Getegid()
}
