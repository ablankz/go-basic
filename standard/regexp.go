package standard

import (
	"fmt"
	"regexp"
)

func Regexp() {
	// (パターン, 対象)
	// 行うたびに正規表現をコンパイルするので大量にするには次のCompileを使う
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	// 正規表現パターンを予め生成
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)

	// コンパイルが失敗したときエラーを返すのではなく直接ランタイムエラーを発生させる
	// 動的にパターンを変えないならこちらが望ましい
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// 以下のようなエスケープが必要
	//regexp.MustCompile("\\d")
	// バッククォーテーションを使用するなら不要
	//regexp.MustCompile(`\d`)

	// 以下正規表現の例
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match)

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)
	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	/*
		re := regexp.MustCompile("ab")
		re.MatchStrings("abc")
	*/

	re5 := regexp.MustCompile(".")
	match = re5.MatchString("ABC")
	fmt.Println(match)
	match = re5.MatchString("\n")
	fmt.Println(match)

	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	re7 := regexp.MustCompile("A+?A+?X")
	fmt.Println(re7.MatchString("AAX"))
	fmt.Println(re7.MatchString("AAAAAAXXXX"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	re1 = regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re1.MatchString("abcxyz"))
	fmt.Println(re1.MatchString("ABCXYZ"))
	fmt.Println(re1.MatchString("ABCxyz"))
	fmt.Println(re1.MatchString("ABCabc"))

	// 正規表現にマッチした文字列の取得
	// マッチした部分の文字列を取得
	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1)
	// マッチした部分の文字列をまとめて取得
	// 第二引数は取得する文字列の最大数(-1ならすべて)
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)
	// 正規表現による文字列の分割
	// マッチした文字列で分割して配列作成
	// 第二引数は分割する最大数(-1ならすべて)
	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	// マッチしたスペース(\s+)をすべて , に置き換え
	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))
	// 、または。をマッチさせて空文字に置き換えることで削除
	re1 = regexp.MustCompile(`、|。`)
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

	// 正規表現のグループによるサブマッチ
	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	// サブマッチ文字列をすべて取得(二次元配列のような形)
	// ex. [[0123-456-7889 0123 456 7889] [111-222-333 111 222 333] [556-787-899 556 787 899]]
	// それぞれの1要素目はマッチ文字列すべて
	ms := re1.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)
	for _, v := range ms {
		fmt.Println(v)
	}

	// サブマッチした文字列を第2引数に渡して置き換え
	// $1などで受け取れる
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))

}
