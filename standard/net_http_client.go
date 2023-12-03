package standard

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func NetHttpClient() {

	//GETメソッド
	res, _ := http.Get("https://example.com")

	fmt.Println(res.StatusCode)

	fmt.Println(res.Proto)

	// ヘッダーはキーを指定してバリューの配列が得られる
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	// レスポンスのcloseが必要
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Print(string(body))

	// Post
	// requestBody作成
	vs := url.Values{}

	// (key, value)
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	// エンコード
	fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ = io.ReadAll(res.Body)
	fmt.Print(string(body))

}
