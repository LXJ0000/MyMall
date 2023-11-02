package serializer

import "MyMall/repository/db/model"

type Carousel struct {
	Id        uint   `json:"id"`
	ImagePath string `json:"image_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) *Carousel {
	return &Carousel{
		Id:        item.ID,
		ImagePath: item.ImgPath,
		ProductId: item.ProductId,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

func BuildCarouselList(items []*model.Carousel) (carouselList []*Carousel) {
	for _, v := range items {
		carouselList = append(carouselList, BuildCarousel(v))
	}
	return
}
