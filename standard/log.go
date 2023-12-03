package standard

import (
	"log"
	"os"
)

func Log() {
	// ログの出力先指定
	log.SetOutput(os.Stdout)

	// デフォルトで日時が記載される
	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	// // ログを書き込み終了
	// log.Fatal("Log\n")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)

	// // ログを書き込み終了
	// log.Panic("Log\n")
	// log.Panicln("Log2")
	// log.Panicf("Log%d\n", 3)

	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	// 出力先変更
	log.SetOutput(f)
	log.Println("ファイルに書き込む")

	// 以下ログフォーマット指定
	// デフォルト
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイル行番号(短縮形)
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	// 時刻とファイル行番号
	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	// デフォルトに戻す
	log.SetFlags(log.LstdFlags)

	// プレフィックスの指定
	log.SetPrefix("[LOG]")
	log.Println("E")

	// ロガーの生成, 一部だけ変えたいログ出力のためなどのユースケース
	logger := log.New(os.Stdout, "[PREFIX]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")

	_, err = os.Open("fdafdsafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

}
