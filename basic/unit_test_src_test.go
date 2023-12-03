package basic

import (
	"testing"
)

// テスト実行(指定ディレクトリ以下)
// go test -v ./basic/

// そのパッケージ内のすべてのテスト実行
// go test -v ./...

// カバレッジを検査(何％カバーできるか)
// go test -v -cover -coverprofile=coverage.out ./basic/

// カバレッジを見る
// go tool cover -html=coverage.out

func TestAddU(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2=3",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "2+3=5",
			args: args{x: 2, y: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddU(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("AddU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivideU(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1/2=0.5",
			args: args{x: 1, y: 2},
			want: 0.5,
		},
		{
			name: "2/0=0",
			args: args{x: 2, y: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivideU(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("DivideU() = %v, want %v", got, tt.want)
			}
		})
	}
}
