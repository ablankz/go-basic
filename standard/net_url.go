package standard

import (
	"fmt"
	"net/url"
)

func NetUrl() {
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	// http, httpsなど
	fmt.Println(u.Scheme)
	// ホスト名
	fmt.Println(u.Host)
	// パス
	fmt.Println(u.Path)
	// クエリ文字列
	fmt.Println(u.RawQuery)
	// クエリフラグ(#以降)
	fmt.Println(u.Fragment)

	// 絶対urlであるか(スキーマがあるか)
	fmt.Println(u.IsAbs())
	// クエリパラメータをmap形式で取得
	fmt.Println(u.Query())

	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	url.Path = "/api/test"
	q := url.Query()
	// (Key, Values)
	q.Set("q", "Golang")
	q.Set("r", "Python")
	// クエリ文字列に変換(エスケープ処理も行う)
	url.RawQuery = q.Encode()

	fmt.Println(url)

}
