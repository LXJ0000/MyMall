package cache

import "fmt"

const (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%d", int(id))
}
