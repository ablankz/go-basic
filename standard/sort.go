package standard

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここをカスタムすることでカスタムソート
// ソートの条件
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func Sort() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s)

	el := []Entry{
		{"A", 20},
		{"y", 30},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"w", 30},
		{"t", 15},
		{"c", 30},
	}

	fmt.Println(el)

	// 並び替えをカスタマイズ
	// 比較対象i, j i<jの全ての要素i,jにおいて第二引数に関数がtrueになるように並び替える(=は省く)
	// 例だと昇順
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	// Sliceの安定ソート
	// Sliceだとソートする前のオーダーを保持することを保証しない
	// SliceStableでは保証する
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	// カスタムソート
	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	sort.Sort(lt)

	fmt.Println(lt)

	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
