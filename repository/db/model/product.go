package model

import (
	"MyMall/repository/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (p *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
func (p *Product) AddView() {
	//增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(p.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(p.ID)))
}
