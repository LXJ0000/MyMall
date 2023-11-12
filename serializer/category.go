package serializer

import "MyMall/repository/db/model"

type Category struct {
	Id           uint   `json:"id"`
	CreateAt     int64  `json:"create_at"`
	CategoryName string `json:"category_name"`
}

func BuildCategory(item *model.Category) *Category {
	return &Category{
		Id:           item.ID,
		CreateAt:     item.CreatedAt.Unix(),
		CategoryName: item.CategoryName,
	}
}

func BuildCategoryList(items []*model.Category) (categoryList []*Category) {
	for _, v := range items {
		categoryList = append(categoryList, BuildCategory(v))
	}
	return
}
