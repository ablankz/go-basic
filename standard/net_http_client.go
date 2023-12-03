package standard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Account struct {
	ID       string `json:"id"`
	PassWord string `json:"password"`
}

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
	// _body, _ := io.ReadAll(res.Body)
	// fmt.Print(string(body))

	// Post
	// requestBody作成
	vs := url.Values{}

	// クエリパラメータ(key, value)
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	// エンコード
	fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// body2, _ := io.ReadAll(res.Body)
	// fmt.Print(string(body2))

	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//GET ver
	//リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	//まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)

	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	//URLのクエリを確認
	q := req.URL.Query()

	//クエリを追加
	q.Add("name", "test")

	//クエリを表示
	fmt.Println(q)

	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())

	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()

	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//読み込み
	// body3, _ := io.ReadAll(resp.Body)

	//出力
	// fmt.Println(string(body3))

	//POST ver
	base, _ = url.Parse("https://example.com/")

	reference, _ = url.Parse("index/lists?id=1")

	// bodyデータ作成
	AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
	data, _ := json.Marshal(AccountData)

	endpoint = base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ = http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

	req.Header.Add("Content-Type", `application/json"`)

	q = req.URL.Query()

	q.Add("name", "test")

	fmt.Println(q)

	fmt.Println(q.Encode())

	req.URL.RawQuery = q.Encode()

	resp, _ = client.Do(req)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
