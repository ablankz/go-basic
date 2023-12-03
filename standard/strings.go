package standard

import (
	"fmt"
	"strings"
)

func Strings() {
	// 文字列結合
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 文字列検索
	// 最初に使われているインデックス、含まれていなければ-1
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	// 最後のインデクス番号
	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

	// 第二引数のAかBかCのどれが最初に現れるインデックス
	i4 := strings.IndexAny("ABC", "ABC")
	// 第二引数のAかBかCのどれが最後に現れるインデックス
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)

	// 文字列の始まり
	b1 := strings.HasPrefix("ABC", "A")
	// 文字列の終わり
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)

	// 含まれているか
	b3 := strings.Contains("ABC", "B")
	// 第二引数のBかDのどれかが含まれるか
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)

	// 何回出現するか
	i6 := strings.Count("ABCABC", "B")
	// 文字列の長さ+1
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7)

	// 繰り返す(負の値ならランタイムエラー)
	s3 := strings.Repeat("ABC", 4)
	// 空文字列
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列置換
	// 第四引数は置換する回数の最大数(-1ならすべて)
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 配列変換
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7)
	// セパレータを残して分割(A, という文字列など)
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8)
	//分割する最大数の指定
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9)
	// セパレータを残して分割(A, という文字列など)+最大数指定
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10)

	// 小文字変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")

	// 大文字変換
	s13 := strings.ToUpper("abc")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14)

	// 空白をtrim(前後のみで間は取り除かない)、全角や\t, \r, \nも対象
	s15 := strings.TrimSpace("    -    ABC    -    ")
	s16 := strings.TrimSpace("\tABC\r\n")
	s17 := strings.TrimSpace("　　　　ABC　　　　")
	fmt.Println(s15, s16, s17)

	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])

}
