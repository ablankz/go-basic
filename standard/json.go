package standard

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID int `json:"id"`
	// json変換後もプロパティ名でいい場合は`json:"id"`など省略する
	// ID int
	// 変換後、型を変えたい場合は以下のようにする
	// ID      int       `json:"id,string"`
	// 変換後、表示したくない場合は以下のようにする
	// ID      int       `json:"-"`
	// フィールドの型の初期値の場合(intなら0, stringなら""など)のみ変換したくないなら以下のようにする
	// ID      int       `json:"id,omitempty"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	// オブジェクトの場合は空オブジェクト({})が残るため、ポインタにしておくと良い
	A *A `json:"A,omitempty"`
}

// 以下のようにしてアンマーシャルをカスタマイズ
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name  string
		Email string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	u.Email = u2.Email
	return err
}

// 以下のようにしてマーシャルをカスタマイズ
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name  string
		Email string
	}{
		Name:  "Mr " + u.Name,
		Email: "custom." + u.Email,
	})
	return v, err
}

func Json() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// Marshal json変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	var u2 User

	// jsonを構造体に変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", u2)
}
