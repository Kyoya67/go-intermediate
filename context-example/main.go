package main

import (
	"context"
	"fmt"
	"time"
)

// work は、キャンセルされるまで処理を続けます。
func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("終了:", ctx.Err())
			return
		default:
			fmt.Println("作業中")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// exampleCancel は、キャンセル可能なcontextを使う例です。
func exampleCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go work(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

// exampleTimeout は、処理に制限時間を設定する例です。
func exampleTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("処理完了")
	case <-ctx.Done():
		fmt.Println("タイムアウト:", ctx.Err())
	}
}

type userNameKey struct{}

// exampleValue は、contextに値を入れて取り出す例です。
// WithValueはリクエストに紐づく値に限定して使います。
func exampleValue() {
	ctx := context.WithValue(context.Background(), userNameKey{}, "gopher")

	userName, ok := ctx.Value(userNameKey{}).(string)
	if ok {
		fmt.Println("ユーザー名:", userName)
	}
}

func main() {
	exampleCancel()
	exampleTimeout()
	exampleValue()
}
