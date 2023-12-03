package standard

import (
	"fmt"
	"log"
	"time"
)

func Time() {
	t := time.Now()
	fmt.Println(t)

	// args8: タイムゾーン, 以下の指定で実行環境のタイムゾーン
	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

	fmt.Printf("%T\nt", time.June)
	fmt.Printf("%T\n", time.June.String())

	// 時間間隔を表現
	// time.Durationを返す
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	// time.Durationを文字列から生成
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)

	t5 := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()

	// 時間の差分
	d2 := t5.Sub(t6)
	fmt.Println(d2)

	// t6がt5より前か
	fmt.Println(t6.Before(t5))
	// t6はt5より後か
	fmt.Println(t6.After(t5))
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))

	t7 := time.Date(2020, 10, 1, 9, 0, 0, 0, time.Local)
	t8 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC)
	// true
	fmt.Println(t7.Equal(t8))
	//fmt.Println(t7.Equal(t6))

	t9 := t5.AddDate(1, 0, 0)
	t10 := t5.AddDate(0, -1, 0)
	t11 := t5.AddDate(0, 0, 1)
	t12 := t5.AddDate(0, 2, -12)
	fmt.Println(t9)
	fmt.Println(t10)
	fmt.Println(t11)
	fmt.Println(t12)

	// 第一引数にフォーマット
	t13, err := time.Parse("2006/01/02", "2020/06/10")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t13)

	// 組み込みフォーマット
	t14, err := time.Parse(time.RFC822, "27 Nov 15 18:00 JST")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t14)

	// t13をレイアウト(引数の形でフォーマット)
	fmt.Printf("%T\n", t13.Format("2006/01/02"))

	// UTCのタイムゾーン
	utc := t14.UTC()
	fmt.Println(utc)

	// 環境変数のタイムゾーン
	jst := t.Local()
	fmt.Println(jst)

	// Unix 時刻を取得
	unix := t13.Unix()
	fmt.Println(unix)

	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")

	// 3秒おきにチャネルに時間を流す
	ticker := time.NewTicker(3 * time.Second)
	// 10秒後にチャネルに時間を流す
	ch2 := time.After(10 * time.Second)
ChanLoop:
	for {
		select {
		case t15 := <-ticker.C:
			fmt.Println(t15)
		case <-ch2:
			break ChanLoop
		}
	}
}
